pragma solidity 0.4.25;

import "./Parameterizer.sol";
import "./PLCRVoting.sol";
import "./MarketToken.sol";
import "./IERC20.sol";
import "./SafeMath.sol";


contract Market {
  using SafeMath for uint;

  struct Listing {
    uint applicationExpiry; // Expiration date of apply stage
    bool listed; // a 'listing' if true
    address owner; // owns the listing
    uint supply; // Number of tokens in the listing (both deposited and minted).
    uint challenge; // corresponts to a poll id in Voting if present
    string data; // A pointer to the actual data this listing represents (or possibly some small data primitive)
    uint minted; // Number of Market tokens that have been minted for this listing.
  }

  struct Challenge {
    uint rewardPool; // (remaining) Pool of tokens to be distributed to winning voters
    address challenger; // Owner of Challenge
    bool resolved; // Indication of if challenge is resolved
    uint stake; // Number of tokens at stake for either party during challenge
    uint totalTokens; // (remaining) Number of tokens used in voting by the winning side
    mapping(address => bool) rewardsClaimed; // Indicates whether a voter has claimed a reward yet
  }

  // Maps ids to associated challenge data
  mapping(uint => Challenge) private selfChallenges;

  // Maps listingHashes to associated listingHash data
  mapping(bytes32 => Listing) private selfListings;

  // Global Variables
  // IERC20 private selfNetworkToken;
  MarketToken private selfMarketToken;
  PLCRVoting private selfVoting;
  Parameterizer private selfParameterizer;
  string private selfName;

  /**
  @dev Contructor
  @param networkTokenAddr
  @param marketTokenAddr
  @param votingAddr
  @param parameterizerAddr
  */
  constructor(
    // address networkTokenAddr, TODO
    address marketTokenAddr,
    address votingAddr,
    address parameterizerAddr,
    string name
  ) public
  {
    // selfNetworkToken = IERC20(networkTokenAddr);
    selfMarketToken = MarketToken(marketTokenAddr);
    selfVoting = PLCRVoting(votingAddr);
    selfParameterizer = Parameterizer(parameterizerAddr);
    selfName = name;
  }

  /**
  @dev Returns true if apply was called for this listingHash
  @param listingHash The listingHash whose status is to be examined
  */
  function applied(bytes32 listingHash) view public returns (bool exists) {
    return selfListings[listingHash].applicationExpiry > 0;
  }

  /**
  @dev Allows a user to start an application. Takes tokens from user and sets
  apply stage end time.
  @param listingHash The hash of a potential listing a user is applying to add to the registry
  @param amount The number of (market) tokens a user is willing to potentially stake
  @param data Either some simple, small data primitive as a string or, more likely, a pointer to a location for some
  bespoke data that this listing represents
  // TODO likely provide a way to update the `data` attribute of a listing
  */
  function apply(bytes32 listingHash, uint amount, string data) external {
    require(!isListed(listingHash), "Error:Market.apply - Already listed");
    require(!applied(listingHash), "Error:Market.apply - Already applied");
    require(amount >= selfParameterizer.get("minDeposit"), "Error:Market.apply - amount must be greater than, or equal to, the minimum deposit");

    // Sets owner
    Listing storage listing = selfListings[listingHash];
    listing.owner = msg.sender;

    // Sets apply stage end time, funding, and the data pointer
    listing.applicationExpiry = block.timestamp.add(selfParameterizer.get("applyStageLen"));
    listing.supply = amount; // we will add a listReward if actually listed
    listing.minted = 0; // same as the line above, not added yet
    listing.data = data; // points to an off-chain resource (or is a string data primitive)

    // Transfers tokens from user to Market contract (market acts as banker for all listing market tokens)
    require(selfMarketToken.transferFrom(listing.owner, this, amount), "Error:Market.apply - Could not transfer Market Tokens");

    emit AppliedEvent(
      listingHash,
      amount,
      listing.applicationExpiry,
      data,
      msg.sender
    );
  }

  /**
  @dev Determines whether the given listingHash be whitelisted.
  @param listingHash The listingHash whose status is to be examined
  */
  function canBeListed(bytes32 listingHash) view public returns (bool) {
    uint id = selfListings[listingHash].challenge;

    // Ensures that the application was made,
    // the application period has ended,
    // the listingHash can be listed,
    // and either: the id == 0, or the challenge has been resolved.
    if (
      /* solium-disable-next-line */
      applied(listingHash) &&
      /* solium-disable-next-line */
      selfListings[listingHash].applicationExpiry < now &&
      /* solium-disable-next-line */
      !isListed(listingHash) &&
      (id == 0 || selfChallenges[id].resolved == true)
    ) {return true;}

    return false;
  }

  /**
  @dev Starts a poll for a listingHash which is either in the apply stage or
  already listed. Tokens are taken from the challenger and the applicant's deposits are locked.
  @param listingHash The listingHash being challenged, whether listed or in application
  */
  function challenge(bytes32 listingHash) external returns (uint id) {
    Listing storage listing = selfListings[listingHash];
    uint minDeposit = selfParameterizer.get("minDeposit");

    require(applied(listingHash) || listing.listed, "Error:Market.challenge - Must be a listing or an application");
    // Prevent multiple challenges
    require(listing.challenge == 0 || selfChallenges[listing.challenge].resolved, "Error:Market.challenge - Already challenged");

    // NOTE we do allow minted tokens to keep a listing "above water" in a challenge
    if ((listing.supply + listing.minted) < minDeposit) {
      // Not enough tokens, listingHash auto-deleted // TODO we could pass the deposit to differentiate auto and manual
      deleteListing(listingHash);
      return 0;
    }

    // Starts poll, uint id declared as named return
    id = selfVoting.startPoll(
      selfParameterizer.get("voteQuorum"),
      selfParameterizer.get("commitStageLen"),
      selfParameterizer.get("revealStageLen")
    );

    selfChallenges[id] = Challenge({
      challenger: msg.sender,
      rewardPool: uint(100).sub(selfParameterizer.get("dispensationPct")).mul(minDeposit).div(100),
      stake: minDeposit,
      resolved: false,
      totalTokens: 0
    });

    // Updates listingHash to store most recent challenge
    listing.challenge = id;

    // Locks tokens for listingHash during challenge
    listing.supply = listing.supply.sub(minDeposit);

    // Takes tokens from challenger
    require(selfMarketToken.transferFrom(msg.sender, this, minDeposit), "Error:Market.challenge - Could not transfer Market tokens");

    uint commitExpiry = selfVoting.getPollCommitExpiry(id);
    uint revealExpiry = selfVoting.getPollRevealExpiry(id);

    emit ChallengeEvent(
      listingHash,
      id,
      commitExpiry,
      revealExpiry,
      msg.sender
    );

    return id;
  }

  /**
  @dev Determines whether voting has concluded in a challenge for a given
  listingHash. Throws if no challenge exists.
  @param listingHash A listingHash with an unresolved challenge
  */
  function challengeCanBeResolved(bytes32 listingHash) view public returns (bool) {
    uint id = selfListings[listingHash].challenge;
    require(challengeExists(listingHash), "Error:Market.challengeCanBeResolved - Challenge does not exist");
    return selfVoting.pollEnded(id);
  }

  /**
  @dev Returns true if the application/listingHash has an unresolved challenge
  @param listingHash The listingHash whose status is to be examined
  */
  function challengeExists(bytes32 listingHash) view public returns (bool) {
    uint id = selfListings[listingHash].challenge;
    return (id > 0 && !selfChallenges[id].resolved);
  }

  /**
  @dev Called by a voter to claim their reward for each completed vote. Someone must call updateStatus() before this can be called.
  @param id The PLCR poll ID of the challenge a reward is being claimed for
  @param salt The salt of a voter's commit hash in the given poll
  */
  function claimReward(uint id, uint salt) public {
    // Ensures the voter has not already claimed tokens and challenge results have been processed
    require(selfChallenges[id].rewardsClaimed[msg.sender] == false);
    require(selfChallenges[id].resolved == true);

    uint voterTokens = selfVoting.getNumPassingTokens(msg.sender, id, salt);
    uint reward = voterReward(msg.sender, id, salt);

    // Subtracts the voter's information to preserve the participation ratios
    // of other voters compared to the remaining pool of rewards
    selfChallenges[id].totalTokens = selfChallenges[id].totalTokens.sub(voterTokens);
    selfChallenges[id].rewardPool = selfChallenges[id].rewardPool.sub(reward);

    // Ensures a voter cannot claim tokens again
    selfChallenges[id].rewardsClaimed[msg.sender] = true;

    require(selfMarketToken.transfer(msg.sender, reward));

    emit RewardClaimedEvent(id, reward, msg.sender);
  }

  /**
  @dev Deletes a listingHash from the whitelist and transfers tokens back to owner (minus minted rewards)
  @param listingHash The listing hash to delete
  */
  function deleteListing(bytes32 listingHash) private {
    Listing storage listing = selfListings[listingHash];

    // Emit events before deleting
    if (listing.listed) {
      emit ListingDeletedEvent(listingHash);
    } else {
      emit ApplicationDeletedEvent(listingHash);
    }

    // Deleting listing to prevent reentry
    delete selfListings[listingHash];

    // Transfers any remaining balance back to the owner and burns any minted tokens
    if (listing.supply > 0) {
      require(selfMarketToken.transfer(listing.owner, listing.supply), "Error:Market.deleteListing - Could not refund remaining supply");
      // burn any minted remainder, the market has this minted amount banked atm.
      selfMarketToken.burn(listing.minted);
    }
  }

  /**
  @dev Allows the owner of a listingHash to increase its supply.
  @param listingHash A listingHash msg.sender is the owner of
  @param amount The number of Market tokens to increase the listing's supply by
  */
  function depositToListing(bytes32 listingHash, uint amount) external {
    Listing storage listing = selfListings[listingHash];

    require(listing.owner == msg.sender, "Error:Market.deposit - Must be listing owner");

    listing.supply = listing.supply.add(amount);
    require(selfMarketToken.transferFrom(msg.sender, this, amount), "Error:Market.deposit - Could not transfer Market Tokens");

    emit ListingDepositEvent(
      listingHash,
      amount,
      listing.supply,
      listing.minted,
      msg.sender
    );
  }

  /**
  @dev Determines the number of tokens awarded to the winning party in a challenge.
  @param id The challenge/pollID to determine a reward for
  */
  function determineReward(uint id) public view returns (uint) {
    require(!selfChallenges[id].resolved && selfVoting.pollEnded(id), "Error:Market.determineReward - Poll must have ended");

    // Edge case, nobody voted, give all tokens to the challenger.
    if (selfVoting.getTotalNumberOfTokensForWinningOption(id) == 0) {
      return uint(2).mul(selfChallenges[id].stake);
    }

    // TODO recheck this math for Market use case
    return uint(2).mul(selfChallenges[id].stake).sub(selfChallenges[id].rewardPool);
  }

  /**
  @dev Allows the owner of a listingHash to remove the listing
  Returns all tokens to the owner of the listingHash. Burns any minted tokens.
  @param listingHash A listingHash msg.sender is the owner of.
  */
  function exit(bytes32 listingHash) external {
    Listing storage listing = selfListings[listingHash];

    require(listing.owner == msg.sender, "Error:Market.exit - Must be listing owner");
    require(isListed(listingHash), "Error:Market.exit - Must be listed");
    // Cannot exit during ongoing challenge
    require(listing.challenge == 0 || selfChallenges[listing.challenge].resolved, "Error:Market.exit - Cannot exit during a challenge");

    // Delete listingHash & return tokens
    deleteListing(listingHash);
  }

  function getName() external view returns (string) {
    return selfName;
  }

  /**
  @dev Returns true if the provided listingHash is a listing
  @param listingHash The listingHash whose status is to be examined
  */
  function isListed(bytes32 listingHash) view public returns (bool listed) {
    return selfListings[listingHash].listed;
  }

  /**
  @dev Called by updateStatus() if the applicationExpiry date passed without a
  challenge being made. Called by resolveChallenge() if an
  application/listing beat a challenge.
  @param listingHash The listingHash of an application/listingHash to be whitelisted
  */
  function listApplication(bytes32 listingHash) private {
    if (!selfListings[listingHash].listed) {
      Listing storage listing = selfListings[listingHash];
      listing.listed = true;
      // NOTE the market acts as a bank for the minted tokens as well (same as deposits)
      uint amount = selfParameterizer.get("listReward");
      selfMarketToken.mint(this, amount);
      listing.minted = amount;
      emit ListedEvent(listingHash, listing.supply, amount); // TODO send tokensMinted with event vs a separate minted event
    }
  }

  /**
  @dev Determines the winner in a challenge. Rewards the winner tokens and
  either lists or deletes the listingHash
  @param listingHash A listingHash with a challenge that is to be resolved
  */
  function resolveChallenge(bytes32 listingHash) private {
    uint id = selfListings[listingHash].challenge;

    // Calculates the winner's reward,
    // which is: (winner's full stake) + (dispensationPct * loser's stake)
    // TODO is this the same in market?
    uint reward = determineReward(id);

    // Sets flag on challenge being processed
    selfChallenges[id].resolved = true;

    // Stores the total tokens used for voting by the winning side for reward purposes
    selfChallenges[id].totalTokens = selfVoting.getTotalNumberOfTokensForWinningOption(id);

    // Case: challenge failed
    if (selfVoting.isPassed(id)) {
      listApplication(listingHash);
      // Unlock stake so that it can be retrieved by the applicant
      selfListings[listingHash].supply = selfListings[listingHash].supply.add(reward);

      emit ChallengeFailedEvent(
        listingHash,
        id,
        selfChallenges[id].rewardPool,
        selfChallenges[id].totalTokens
      );

    } else { // Case: challenge succeeded or nobody voted
      deleteListing(listingHash);
      // Transfer the reward to the challenger
      require(selfMarketToken.transfer(selfChallenges[id].challenger, reward));

      emit ChallengeSucceededEvent(
        listingHash,
        id,
        selfChallenges[id].rewardPool,
        selfChallenges[id].totalTokens
      );
    }
  }

  /**
  @dev Updates a listingHash's status from 'application' to 'listing' or resolves a challenge if one exists.
  @param listingHash The listingHash whose status is being updated
  */
  function updateStatus(bytes32 listingHash) public {
    if (canBeListed(listingHash)) {
      listApplication(listingHash);
    } else if (challengeCanBeResolved(listingHash)) {
      resolveChallenge(listingHash);
    } else {
      revert(); // TODO message?
    }
  }

  /**
  @dev Calculates the provided voter's token reward for the given poll.
  @param voter The address of the voter whose reward balance is to be returned
  @param id The poll ID of the challenge a reward balance is being queried for
  @param salt The salt of the voter's commit hash in the given poll
  @return The uint indicating the voter's reward
  */
  function voterReward(address voter, uint id, uint salt)
  public view returns (uint)
  {
    uint totalTokens = selfChallenges[id].totalTokens;
    uint rewardPool = selfChallenges[id].rewardPool;
    uint voterTokens = selfVoting.getNumPassingTokens(voter, id, salt);
    return voterTokens.mul(rewardPool).div(totalTokens);
  }

  /**
  @dev Allows the owner of a listingHash to withdraw from its supply (minted is restricted).
  @param listingHash A listingHash msg.sender is the owner of.
  @param amount The number of Market tokens to withdraw from the supply (must be <= (supply - minted)).
  */
  function withdrawFromListing(bytes32 listingHash, uint amount) external {
    Listing storage listing = selfListings[listingHash];

    require(listing.owner == msg.sender, "Error:Market.withdraw - Must be listing owner");
    // can't outright withdraw more than the supply
    require(amount <= listing.supply, "Error:Market.withdraw - amount must be less than or equal to the listing supply");
    // we do allow the minted amount to keep supply above min deposit threshold
    require((listing.supply.add(listing.minted)).sub(amount) >= selfParameterizer.get("minDeposit"), "Error:Market.withdraw - Cannot withdraw listing supply below minimum deposit");

    listing.supply = listing.supply.sub(amount);
    require(selfMarketToken.transfer(msg.sender, amount), "Error:Market.withdraw - Could not transfer tokens to listing owner");

    emit ListingWithdrawEvent(
      listingHash,
      amount,
      listing.supply,
      listing.minted,
      msg.sender
    );
  }

  event ApplicationDeletedEvent(bytes32 indexed listingHash);
  event AppliedEvent(bytes32 indexed listingHash, uint deposit, uint applicationExpiry, string data, address indexed applicant);
  event ChallengeEvent(bytes32 indexed listingHash, uint id, uint commitExpiry, uint revealExpiry, address indexed challenger);
  event ChallengeFailedEvent(bytes32 indexed listingHash, uint indexed id, uint rewardPool, uint totalTokens);
  event ChallengeSucceededEvent(bytes32 indexed listingHash, uint indexed id, uint rewardPool, uint totalTokens);
  event ListedEvent(bytes32 indexed listingHash, uint supply, uint minted);
  event ListingDepositEvent(bytes32 indexed listingHash, uint added, uint supply, uint minted, address indexed owner);
  event ListingDeletedEvent(bytes32 indexed listingHash);
  event ListingWithdrawEvent(bytes32 indexed listingHash, uint subtracted, uint supply, uint minted, address indexed owner);
  event RewardClaimedEvent(uint indexed id, uint reward, address indexed voter);
}
