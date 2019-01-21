pragma solidity 0.5.2;

import "./SafeMath.sol";
import "./IERC20.sol";
import "./MarketToken.sol";
import "./Parameterizer.sol";
import "./Voting.sol";


contract Market {
  using SafeMath for uint256;

  struct Listing {
    uint256 index; // pointer to the location of this listing's hash in the unordered keys array
    bool listed; // a 'listing' if true
    address owner; // owns the listing
    uint256 supply; // Number of tokens in the listing, banked by the market, put there by the owner
    bytes32 dataHash; // A magical construct which flies across the sky, pooping unicorns which, in turn, poop rainbows.
    uint256 rewards; // Number of Market tokens that have been minted + any challenge winnings
  }

  // TODO in another iteration allow challenger to stake a listing's supply/reward if they own one
  struct Challenge {
    address challenger; // adress making the challenge
    uint256 fromChallengeeSupply; // how much of `stake` was taken from listing.supply
    uint256 fromChallengeeRewards; // how much, if any, of `stake` was taken from listing.rewards
  }

  struct Investor {
    uint256 index; // pointer to the location of this investor's address in the unordered keys array
    uint256 invested; // cumulative total of investment in this market
    uint256 minted; // cumulative total of minted market tokens from investing
  }

  // Maps listingHash to associated challenge data
  mapping(bytes32 => Challenge) private selfChallenges;

  // Maps listingHashes to associated listingHash data
  mapping(bytes32 => Listing) private selfListings;
  // dynamic array which houses the hashes of all current listings
  bytes32[] private selfListingKeys;

  // Maps listing owner's address to how many listings they own
  mapping(address => uint256) private selfListingOwners;

  // Maps investor's address to their data
  mapping(address => Investor) private selfInvestors;
  // house the addresses of all current investors
  address[] private selfInvestorKeys;
  // running total of invested ether token
  uint256 private selfInvested;

  // though the ether token exposes more than the erc20 interface, the market only needs the latter
  IERC20 private selfEtherToken;
  MarketToken private selfMarketToken;
  Parameterizer private selfParameterizer;
  Voting private selfVoting;
  // will be the market factory
  address private selfOwner;

  // voting candidate kinds are denoted by uint8s. Not using an enum for just 2
  uint8 constant APPLICATION = 1;
  uint8 constant CHALLENGE = 2;

  /**
    @param marketTokenAddr Address of the deployed market token contract
    @param votingAddr Address of the deployed voting contract
    @param parameterizerAddr Address of the deployed parameterizer contract
  */
  constructor(
    address marketTokenAddr,
    address etherTokenAddr,
    address parameterizerAddr,
    address votingAddr
  ) public
  {
    selfOwner = msg.sender;
    selfMarketToken = MarketToken(marketTokenAddr);
    selfEtherToken = IERC20(etherTokenAddr);
    selfParameterizer = Parameterizer(parameterizerAddr);
    selfVoting = Voting(votingAddr);
  }

  /**
    @dev Starts a challenge for a listing
    @param listingHash The listing being challenged.
  */
  function challenge(bytes32 listingHash) external {
    uint256 stake = selfParameterizer.getChallengeStake();
    Listing storage listing = selfListings[listingHash];

    require(listing.listed == true, "Error:Market.challenge - Must be a listing to be challenged");
    // Prevent multiple challenges
    require(selfVoting.isCandidate(listingHash) != true, "Error:Market.challenge - Already challenged");
    // TODO requirements on _who_ can challenge?

    // NOTE we do allow rewards to keep a listing "above water" in a challenge
    if (listing.supply.add(listing.rewards) < stake) {
      // Not enough funds, listing is auto-deleted, TODO we could pass some val to signal a challenge auto-delete
      removeListing(listingHash);
    } else {
      // Takes tokens from challenger, do early as we'll revert if they aint got the funds
      // NOTE: in the next iteration make it possible for the challenger to stake a listing (by using its rewards)
      selfMarketToken.transferFrom(msg.sender, address(this), stake);

      uint256 fromChallengeeSupply;
      uint256 fromChallengeeRewards;

      // we will 'lock' only tokens in the supply if possible
      if (listing.supply >= stake) {
        listing.supply = listing.supply.sub(stake);
        fromChallengeeSupply = stake;
      } else {
        // lock supply and rewards if necessary
        fromChallengeeSupply = listing.supply;
        uint256 remainder = stake.sub(listing.supply);
        listing.supply = 0;
        listing.rewards = listing.rewards.sub(remainder);
        fromChallengeeRewards = remainder;
      }

      selfChallenges[listingHash] = Challenge({
        challenger: msg.sender,
        fromChallengeeSupply: fromChallengeeSupply,
        fromChallengeeRewards: fromChallengeeRewards
      });

      // places the candidate, associated to challenge by listingHash
      selfVoting.addCandidate(
        listingHash,
        CHALLENGE,
        selfParameterizer.getVoteBy()
      );

      emit ChallengedEvent(
        listingHash,
        msg.sender
      );
    }
  }

  /**
    @dev Allows the owner of a listingHash to increase its supply.
    @param listingHash A listingHash msg.sender is the owner of
    @param amount The number of Market tokens to increase the listing's supply by
  */
  function depositToListing(bytes32 listingHash, uint256 amount) external {
    // Listing storage listing = selfListings[listingHash];
    require(isListing(listingHash), "Error:Market.depositToListing - Listing does not exist");

    // TODO any need to restrict deposit to owner? I don't think so atm
    // require(selfListings[listingHash].owner == msg.sender, "Error:Market.depositToListing - Must be listing owner");

    selfMarketToken.transferFrom(msg.sender, address(this), amount);

    selfListings[listingHash].supply = selfListings[listingHash].supply.add(amount);

    emit ListingDepositEvent(
      listingHash,
      msg.sender,
      amount,
      selfListings[listingHash].supply,
      selfListings[listingHash].rewards
    );
  }

  /**
    @dev Allows the owner of a listingHash to remove the listing
    Returns all tokens to the owner of the listingHash. Burns any minted and/or rewards tokens.
    @param listingHash A listingHash msg.sender is the owner of.
  */
  function exit(bytes32 listingHash) external {
    require(selfListings[listingHash].owner == msg.sender, "Error:Market.exit - Must be listing owner");
    require(isListed(listingHash), "Error:Market.exit - Must be listed");
    // Cannot exit during ongoing challenge
    require(selfVoting.isCandidate(listingHash) != true, "Error:Market.exit - Cannot exit during a challenge");

    // Delete listingHash & return tokens
    removeListing(listingHash);
  }

  function getChallenge(bytes32 listingHash) external view returns (address, uint256, uint256) {
    return (
      selfChallenges[listingHash].challenger,
      selfChallenges[listingHash].fromChallengeeSupply,
      selfChallenges[listingHash].fromChallengeeRewards
    );
  }

  function getInvested() external view returns (uint256) {
    return selfInvested;
  }

  /**
    @dev Return the current minimum amount of ether tokens required to invest in this market (rate + slope * reserve)
    @return uint256 Said number of ether tokens
    TODO audit...
  */
  function getInvestmentPrice() public view returns (uint256) {
    uint256 rate = selfParameterizer.getConversionRate();
    uint256 slopeD = selfParameterizer.getConversionSlopeDenominator();
    uint256 slopeN = selfParameterizer.getConversionSlopeNumerator();
    uint256 reserve = selfEtherToken.balanceOf(address(this));
    return rate.add(slopeN.mul(reserve).div(slopeD));
  }

  function getInvestor(address addr) external view returns (uint256, uint256) {
    return (
      selfInvestors[addr].invested,
      selfInvestors[addr].minted
    );
  }

  function getListing(bytes32 listingHash) external view returns (bool, address, uint256, bytes32, uint256) {
    // TODO likely a better pattern to _not_ possibly revert on getters like this - just let evm return falsy stuff...
    // require(isListing(listingHash) == true, "Error:Market.getListing - Listing does not exist");

    // NOTE: we don't return the index
    return (
      selfListings[listingHash].listed,
      selfListings[listingHash].owner,
      selfListings[listingHash].supply,
      selfListings[listingHash].dataHash,
      selfListings[listingHash].rewards
    );
  }

  function getListings() external view returns (bytes32[] memory) {
    return selfListingKeys;
  }

  /**
    @dev An externally callable method that will return a listingHash using the same algorithm
    as the one used to produce any stored listingHash
    @param listing A string to hash and return
    @return the hashed bytes32
  */
  function getListingHash(string calldata listing) external pure returns (bytes32) {
    return keccak256(bytes(listing));
  }

  /**
    @dev Given ether tokens as entry, mint and return market tokens to the
    caller as dictated by the pricing curve.
    @param offered Number of ether tokens offered for investment
  */
  function invest(uint256 offered) external {
    require(isListingOwner(msg.sender) != true, "Error:Market.invest - Cannot invest while owning an active listing");
    // amount must be >= the investment price
    uint256 price = getInvestmentPrice();
    require(offered >= price, "Error:Market.invest - Amount offered must be greater than or equal to current investment price");
    // move those used tokens from investor to the reserve. NOTE this also subtracts from ether token allowance[sender][market]
    selfEtherToken.transferFrom(msg.sender, address(this), offered);
    // we mint `(offered/price) * 1e18`
    uint256 minted = (offered.div(price)).mul(1e18);
    // NOTE this is, at origin, owned by the market
    selfMarketToken.mint(minted);
    // now we can transfer those minted market tokens to the investor
    selfMarketToken.transfer(msg.sender, minted);
    // sender is an investor now (if not already)
    if(isInvestor(msg.sender)) {
      selfInvestors[msg.sender].invested = selfInvestors[msg.sender].invested.add(offered);
      selfInvestors[msg.sender].minted = selfInvestors[msg.sender].minted.add(minted);
      // TODO check here for council membership when changing to threshold rules...
    } else {
      Investor storage i = selfInvestors[msg.sender];
      // record what we took and gave - NOTE no need to .add here
      i.invested = offered;
      i.minted = minted;
      // address onto the end of the investor keys array, index stored in the struct
      i.index = selfInvestorKeys.push(msg.sender) - 1;
      // currently any new investor is a council member, change when switching to threshold rules TODO
      selfVoting.addToCouncil(msg.sender);
    }

    // the total invested in the market is bumped
    selfInvested = selfInvested.add(offered);

    emit InvestedEvent(msg.sender, offered, minted);
  }

  function isInvestor(address addr) public view returns (bool) {
    // if the keys array is empty, obv not...
    if(selfInvestorKeys.length == 0) return false;
    // now just confirm that they match
    return(selfInvestorKeys[selfInvestors[addr].index] == addr);
  }

  /**
    @dev Returns true if the provided listingHash is a listing
    @param listingHash The listingHash whose status is to be examined
  */
  function isListed(bytes32 listingHash) public view returns (bool) {
    return selfListings[listingHash].listed == true;
  }

  function isListing(bytes32 listingHash) public view returns (bool) {
    // if the keys array is empty, obv not...
    if(selfListingKeys.length == 0) return false;
    // now just confirm that they match
    return(selfListingKeys[selfListings[listingHash].index] == listingHash);
  }

  /**
  @dev Returns true if the provided address is a listing owner
  @param addr The address to check if 'owns' one or more listings
  */
  function isListingOwner(address addr) public view returns (bool) {
    return selfListingOwners[addr] >= 1;
  }

  /**
    @dev Allows a user to apply for listing status.
    @notice We note the owner as a `listOwner` as of this step
    @param listing A string "identifier" of a potential listing a user is applying to add to the registry
    Note that the string will be hased via keccak256, and that result must be unique to the market.
    @param dataHash unique key for the backend to find this bespoke data
    @param amount The number of (market) tokens a user wants to bank in the listing (optional)
  */
  function list(string calldata listing, bytes32 dataHash, uint256 amount) external {
    bytes32 listingHash = keccak256(bytes(listing));

    require(selfVoting.isCandidate(listingHash) != true, "Error:Market.apply - Already a candidate");
    require(isListed(listingHash) != true, "Error:Market.list - Already listed");

    // first, create the actual listing object
    Listing storage listing = selfListings[listingHash];
    listing.owner = msg.sender;
    listing.dataHash = dataHash; // a unique identifier known to a backend used to locate the data (or something)
    listing.index = selfListingKeys.push(listingHash) - 1;

    // the owner is official by this point
    selfListingOwners[msg.sender] = selfListingOwners[msg.sender].add(1);

    // the amount is optional here, but if sent, we need to bank them
    if (amount > 0) {
      // Transfers tokens from user to Market contract (market acts as banker for all listing market tokens)
      selfMarketToken.transferFrom(listing.owner, address(this), amount);
      listing.supply = amount; // optional arg here, may be 0
    }

    // with that in place, create the candidate for this listing, NOTE: will trigger a CandidateCreatedEvent as well
    selfVoting.addCandidate(
      listingHash,
      APPLICATION,
      selfParameterizer.getVoteBy()
    );

    emit AppliedEvent(
      listingHash,
      msg.sender,
      dataHash,
      selfParameterizer.getVoteBy(),
      amount
    );
  }

  /**
    @dev Deletes a listing and transfers tokens back to owner (minus minted rewards)
    @param listingHash The listing hash to delete
  */
  function removeListing(bytes32 listingHash) private {
    // must be a listing, but not necessarily listed
    require(isListing(listingHash), "Error:Market.removeListing - Must be a listing");

    // Transfers any remaining balance back to the owner and burns any minted tokens
    if (selfListings[listingHash].supply > 0) {
      selfMarketToken.transfer(selfListings[listingHash].owner, selfListings[listingHash].supply);
    }

    if (selfListings[listingHash].rewards > 0) {
      // burn any rewards remainder, the market has this amount banked atm.
      selfMarketToken.burn(selfListings[listingHash].rewards);
    }

    // remove the keys entry first
    uint256 deleted = selfListings[listingHash].index; // being removed
    bytes32 moved = selfListingKeys[selfListingKeys.length - 1]; // moving to overwrite 'deleted'
    selfListingKeys[deleted] = moved; // delete target now overwritten
    selfListingKeys.length--;
    // update the index of moved
    selfListings[moved].index = deleted;
    // note that the owner has one less listing now
    selfListingOwners[selfListings[listingHash].owner] = selfListingOwners[selfListings[listingHash].owner].sub(1);
    // if the owner has no listings, take them out of the council. TODO revisit when we implement threshold rules
    if (selfListingOwners[selfListings[listingHash].owner] < 1) {
      selfVoting.removeFromCouncil(selfListings[listingHash].owner);
    }
    // finally we can kill it
    delete selfListings[listingHash];

    emit ListingDeletedEvent(listingHash);
  }

  /**
    @dev Determines if an applicant becomes a listing
    @param listingHash The listingHash representing the applicant in question
  */
  function resolveApplication(bytes32 listingHash) public {
    require(selfVoting.inCouncil(msg.sender), "Error:Market.resolveApplication - Sender must be council member");
    // we assure it's not a listing already, there is a candidate, and voting has concluded
    require(selfVoting.candidateIs(listingHash, APPLICATION) == true, "Error:Market.resolveApplication - Must be an application");
    // the above and below remove the need to use isListing
    require(selfListings[listingHash].listed != true, "Error:Market.resolveApplication - Already listed");
    require(selfVoting.pollClosed(listingHash), "Error:Market.resolveApplication - Polls for this candidate must be closed");

    // Case: listing accepted
    if (selfVoting.didPass(listingHash, selfParameterizer.getQuorum())) {
      selfListings[listingHash].listed = true;
      // NOTE the market acts as a bank for the minted tokens as well (same as deposits)
      uint256 amount = selfParameterizer.getListReward();
      // there is no `to` adress as all mint balances are banked by the market
      selfMarketToken.mint(amount);
      // no need to .add here
      selfListings[listingHash].rewards = amount;
      // currently any new owner is made a council member, change when switching to threshold rules TODO
      if (selfVoting.inCouncil(selfListings[listingHash].owner) != true) {
        selfVoting.addToCouncil(selfListings[listingHash].owner);
      }

      emit ListedEvent(
        listingHash,
        selfListings[listingHash].owner,
        amount,
        selfListings[listingHash].supply
      );
      // Case: listing not accepted
    } else {
      // simply remove the listing and assure it happened.
      removeListing(listingHash);
    }
    // in either case the candidate is pruned
    selfVoting.removeCandidate(listingHash);
  }

  /**
    @dev Determines the winner in a challenge. Rewards the winner and possibly deletes the listingHash
    @param listingHash A listingHash with a challenge that is to be resolved
    @return true if succeeded
  */
  function resolveChallenge(bytes32 listingHash) public {
    require(selfVoting.inCouncil(msg.sender), "Error:Market.resolveChallenge - Sender must be council member");
    // dont operate on non challenges. TODO save the bytes arrays for kinds as CONST?
    require(selfVoting.candidateIs(listingHash, CHALLENGE), "Error:Market.resolveChallenge - Must be a challenge");
    require(selfVoting.pollClosed(listingHash), "Error:Market.resolveChallenge - Polls for this candidate must be closed");
    // somebody gets a prize...
    uint256 stake = selfParameterizer.getChallengeStake();

    // Case: challenge won
    if (selfVoting.didPass(listingHash, selfParameterizer.getQuorum())) {
      removeListing(listingHash);
      // Both refund the challenger their stake and give them the listing's
      selfMarketToken.transfer(selfChallenges[listingHash].challenger, stake.mul(2));
      // TODO update when we allow a challenger to stake a listing...

      emit ChallengeSucceededEvent(
        listingHash,
        selfChallenges[listingHash].challenger,
        stake
      );
    } else { // Case: challenge lost
      // first refund any locked supply amounts from the listing
      selfListings[listingHash].supply = selfListings[listingHash].supply.add(selfChallenges[listingHash].fromChallengeeSupply);
      // now refund any locked rewards + the rewarded stake
      selfListings[listingHash].rewards = selfListings[listingHash].rewards.add(selfChallenges[listingHash].fromChallengeeRewards).add(stake);

      emit ChallengeFailedEvent(
        listingHash,
        selfChallenges[listingHash].challenger,
        stake
      );
    }

    // clean up the challenge and candidates now
    delete selfChallenges[listingHash];
    selfVoting.removeCandidate(listingHash);
  }

  /**
    @dev Allows the owner of a listingHash to withdraw from its supply (minted is restricted).
    @param listingHash A listingHash msg.sender is the owner of.
    @param amount The number of Market tokens to withdraw from the supply (must be <= (supply - minted)).
  */
  function withdrawFromListing(bytes32 listingHash, uint256 amount) external {
    require(isListing(listingHash), "Error:Market.withdrawFromListing - Listing does not exist");
    require(selfListings[listingHash].owner == msg.sender, "Error:Market.withdraw - Must be listing owner");
    // can't outright withdraw more than the supply
    require(amount <= selfListings[listingHash].supply, "Error:Market.withdraw - amount must be less than or equal to the listing supply");

    selfListings[listingHash].supply = selfListings[listingHash].supply.sub(amount);
    selfMarketToken.transfer(msg.sender, amount);

    emit ListingWithdrawEvent(
      listingHash,
      msg.sender,
      amount,
      selfListings[listingHash].supply,
      selfListings[listingHash].rewards
    );
  }

  event ApplicationFailedEvent(bytes32 indexed listingHash);
  event AppliedEvent(bytes32 indexed listingHash, address indexed applicant, bytes32 indexed dataHash, uint256 voteBy, uint256 amount);
  event ChallengedEvent(bytes32 indexed listingHash, address indexed challenger);
  event ChallengeFailedEvent(bytes32 indexed listingHash, address indexed challenger, uint256 reward);
  event ChallengeSucceededEvent(bytes32 indexed listingHash, address indexed challenger, uint256 reward);
  event InvestedEvent(address indexed investor, uint256 offered, uint256 minted);
  event ListedEvent(bytes32 indexed listingHash, address indexed owner, uint256 reward, uint256 supply);
  event ListingDepositEvent(bytes32 indexed listingHash, address indexed owner, uint256 added, uint256 supply, uint256 rewards);
  event ListingDeletedEvent(bytes32 indexed listingHash);
  event ListingWithdrawEvent(bytes32 indexed listingHash, address indexed owner, uint256 added, uint256 supply, uint256 rewards);
}
