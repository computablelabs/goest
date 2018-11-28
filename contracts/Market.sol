pragma solidity 0.5.0;

import "./SafeMath.sol";
import "./IERC20.sol";
import "./MarketToken.sol";
import "./Parameterizer.sol";
import "./Voting.sol";


contract Market {
  using SafeMath for uint;

  struct Listing {
    uint index; // pointer to the location of this listing's hash in the unordered keys array
    bool listed; // a 'listing' if true
    address owner; // owns the listing
    uint supply; // Number of tokens in the listing, banked by the market, put there by the owner
    bytes32 dataHash; // A magical construct which flies across the sky, pooping unicorns which, in turn, poop rainbows.
    uint rewards; // Number of Market tokens that have been minted + any challenge winnings
  }

  // TODO in another iteration allow challenger to stake a listing's supply/reward if they own one
  struct Challenge {
    address challenger; // adress making the challenge
    uint fromChallengeeSupply; // how much of `stake` was taken from listing.supply
    uint fromChallengeeRewards; // how much, if any, of `stake` was taken from listing.rewards
  }

  struct Investor {
    uint index; // pointer to the location of this investor's address in the unordered keys array
    uint invested; // cumulative total of investment in this market
    uint minted; // cumulative total of minted market tokens from investing
  }

  // Maps listingHash to associated challenge data
  mapping(bytes32 => Challenge) private selfChallenges;

  // Maps listingHashes to associated listingHash data
  mapping(bytes32 => Listing) private selfListings;
  // dynamic array which houses the hashes of all current listings
  bytes32[] private selfListingKeys;

  // Maps listing owner's address to how many listings they own
  mapping(address => uint) private selfListingOwners;

  // Maps investor's address to their data
  mapping(address => Investor) private selfInvestors;
  // house the addresses of all current investors
  address[] private selfInvestorKeys;
  // running total of invested network token
  uint private selfInvested;

  IERC20 private selfNetworkToken;
  MarketToken private selfMarketToken;
  Voting private selfVoting;
  Parameterizer private selfParameterizer;
  bytes private selfName;
  // will be the market factory
  address private selfOwner;

  /**
    @param marketTokenAddr Address of the deployed market token contract
    @param votingAddr Address of the deployed voting contract
    @param parameterizerAddr Address of the deployed parameterizer contract
  */
  constructor(
    string memory name,
    address networkTokenAddr,
    address marketTokenAddr,
    address parameterizerAddr,
    address votingAddr
  ) public
  {
    selfOwner = msg.sender;

    selfName = bytes(name);
    selfNetworkToken = IERC20(networkTokenAddr);
    selfMarketToken = MarketToken(marketTokenAddr);
    selfParameterizer = Parameterizer(parameterizerAddr);
    selfVoting = Voting(votingAddr);
  }

  /**
    @dev Starts a challenge for a listing
    @param listingHash The listing being challenged.
    @return true if successful
  */
  function challenge(bytes32 listingHash) external returns (bool) {
    uint stake = selfParameterizer.getChallengeStake();
    Listing storage listing = selfListings[listingHash];

    require(listing.listed == true, "Error:Market.challenge - Must be a listing to be challenged");
    // Prevent multiple challenges
    require(!selfVoting.isCandidate(listingHash), "Error:Market.challenge - Already challenged");

    // NOTE we do allow rewards to keep a listing "above water" in a challenge
    if (listing.supply.add(listing.rewards) < stake) {
      // Not enough funds, listing is auto-deleted // TODO we could pass some val to signal a challenge auto-delete
      if (removeListing(listingHash) != true) {
        revert("Error:Market.challenge - Could not auto remove listing");
      }
    } else {
      // Takes tokens from challenger, do early as we'll revert if they aint got the funds
      // NOTE: in the next iteration make it possible for the challenger to stake a listing (by using its rewards)
      require(selfMarketToken.transferFrom(msg.sender, address(this), stake), "Error:Market.challenge - Could not transfer Market tokens");

      uint fromChallengeeSupply;
      uint fromChallengeeRewards;

      // we will 'lock' only tokens in the supply if possible
      if (listing.supply >= stake) {
        listing.supply = listing.supply.sub(stake);
        fromChallengeeSupply = stake;
      } else {
        // lock supply and rewards if necessary
        fromChallengeeSupply = listing.supply;
        uint remainder = stake.sub(listing.supply);
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
      bool added = selfVoting.addCandidate(
        "challenge",
        listingHash,
        selfParameterizer.getVoteBy()
      );

      require(added, "Error:Market.challenge - Could not add challenge to voting candidates list");

      emit ChallengedEvent(
        listingHash,
        msg.sender
      );
    }

    return true;
  }

  /**
    @dev Allows the owner of a listingHash to increase its supply.
    @param listingHash A listingHash msg.sender is the owner of
    @param amount The number of Market tokens to increase the listing's supply by
    @return true if succeeded
  */
  function depositToListing(bytes32 listingHash, uint amount) external returns (bool) {
    Listing storage listing = selfListings[listingHash];

    require(listing.owner == msg.sender, "Error:Market.deposit - Must be listing owner");
    require(selfMarketToken.transferFrom(msg.sender, address(this), amount), "Error:Market.deposit - Could not transfer Market Tokens");
    listing.supply = listing.supply.add(amount);

    emit ListingDepositEvent(
      listingHash,
      msg.sender,
      amount,
      listing.supply,
      listing.rewards
    );

    return true;
  }

  /**
    @dev Allows the owner of a listingHash to remove the listing
    Returns all tokens to the owner of the listingHash. Burns any minted and/or rewards tokens.
    @param listingHash A listingHash msg.sender is the owner of.
    @return true if success
  */
  function exit(bytes32 listingHash) external returns (bool) {
    require(selfListings[listingHash].owner == msg.sender, "Error:Market.exit - Must be listing owner");
    require(isListed(listingHash), "Error:Market.exit - Must be listed");
    // Cannot exit during ongoing challenge
    require(!selfVoting.isCandidate(listingHash), "Error:Market.exit - Cannot exit during a challenge");

    // Delete listingHash & return tokens
    if (removeListing(listingHash) != true) {
      revert("Error:Market.exit - Could not remove listing");
    }
    return true;
  }

  /**
    @dev Return the current minimum amount of network tokens required to invest in this market (rate + slope * reserve)
    @return uint Said number of network tokens
    TODO audit...
  */
  function getInvestmentPrice() public view returns (uint) {
    uint rate = selfParameterizer.getConversionRate();
    uint slopeD = selfParameterizer.getConversionSlopeDenominator();
    uint slopeN = selfParameterizer.getConversionSlopeNumerator();
    uint reserve = selfNetworkToken.balanceOf(address(this));
    return rate.add(slopeN.mul(reserve).div(slopeD));
  }

  function getName() external view returns (string memory) {
    return string(selfName);
  }

  /**
    @dev Given network tokens as entry, mint and return market tokens to the
    caller as dictated by the pricing curve.
    @param offered Number of network tokens offered for investment
    @return uint number of tokens minted and given to investor
  */
  function invest(uint offered) external returns (uint) {
    // TODO add a check that msg.sender is NOT a listing holder?
    require(!isListingOwner(msg.sender), "Error:Market.invest - Cannot invest while owning an active listing"); // this can be a modifier if we do this elswhere too
    // amount must be >= the investment price
    uint price = getInvestmentPrice();
    require(offered >= price, "Error:Market.invest - Amount offered must be greater than or equal to current investment price");
    // sender could have invest more than the price, but we will only use evenly divisable numbers as we cannot mint fractional tokens
    uint remainder = offered % price;
    // regardless of price multiples, any remainder will be unused
    uint taken = remainder != 0 ? offered.sub(remainder) : remainder;
    // move those used tokens from investor to the reserve. NOTE this also subtracts from network token allowance[sender][market]
    selfNetworkToken.transferFrom(msg.sender, address(this), taken);
    // we mint amount taken / investment price. TODO audit
    uint minted = taken.div(price);
    // NOTE this is, at origin, owned by the market
    require(selfMarketToken.mint(minted), "Error:Market.invest - Could not mint tokens");
    // now we can transfer those minted market tokens to the investor
    require(selfMarketToken.transfer(msg.sender, minted), "Error:Market.invest - Could not transfer tokens");
    // sender is an investor now (if not already)
    if(isInvestor(msg.sender)) {
      selfInvestors[msg.sender].invested = selfInvestors[msg.sender].invested.add(taken);
      selfInvestors[msg.sender].minted = selfInvestors[msg.sender].minted.add(minted);
      // TODO check here for council membership when changing to threshold rules...
    } else {
      Investor storage i = selfInvestors[msg.sender];
      // record what we took and gave - NOTE no need to .add here
      i.invested = taken;
      i.minted = minted;
      // address onto the end of the investor keys array, index stored in the struct
      i.index = selfInvestorKeys.push(msg.sender) - 1;
      // currently any new investor is a council member, change when switching to threshold rules TODO
      require(selfVoting.addToCouncil(msg.sender), "Error:Market.invest - could not add investor to council");
    }

    // the total invested in the market is bumped
    selfInvested = selfInvested.add(taken);

    emit InvestedEvent(
      msg.sender,
      offered,
      taken,
      minted
    );

    return minted;
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
    return selfListings[listingHash].listed;
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
    @notice Changed from Market.apply as apply is now (0.5) a reserved keyword
    @param listingHash The hash of a potential listing a user is applying to add to the registry
    @param dataHash unique key for the backend to find this bespoke data
    @param amount The number of (market) tokens a user wants to bank in the listing (optional)
    @return true if successful
  */
  function list(bytes32 listingHash, bytes32 dataHash, uint amount) external returns (bool) {
    require(!selfVoting.isCandidate(listingHash), "Error:Market.apply - Already a candidate");
    require(!isListed(listingHash), "Error:Market.list - Already listed");

    // first, create the actual listing object
    Listing storage listing = selfListings[listingHash];
    listing.owner = msg.sender;
    listing.dataHash = dataHash; // a unique identifier known to a backend used to locate the data (or something)

    // the amount is optional here, but if sent, we need to bank them
    if (amount > 0) {
      listing.supply = amount; // optional arg here, may be 0
      // Transfers tokens from user to Market contract (market acts as banker for all listing market tokens)
      require(selfMarketToken.transferFrom(listing.owner, address(this), amount), "Error:Market.list - Could not transfer Market Tokens");
    }

    // with that in place, create the candidate for this listing, NOTE: will trigger a CandidateCreatedEvent as well
    bool added = selfVoting.addCandidate(
      "application",
      listingHash,
      selfParameterizer.getVoteBy()
    );

    require(added, "Error:Market.list - Could not add applicant to voting candidates list");

    emit AppliedEvent(
      listingHash,
      msg.sender,
      dataHash,
      selfParameterizer.getVoteBy(),
      amount
    );

    return true;
  }

  /**
    @dev Deletes a listing and transfers tokens back to owner (minus minted rewards)
    @param listingHash The listing hash to delete
  */
  function removeListing(bytes32 listingHash) private returns (bool) {
    require(selfListings[listingHash].listed == true, "Error:Market.removeListing - Must be a listing");

    // Transfers any remaining balance back to the owner and burns any minted tokens
    if (selfListings[listingHash].supply > 0) {
      require(selfMarketToken.transfer(selfListings[listingHash].owner, selfListings[listingHash].supply), "Error:Market.removeListing - Could not refund remaining supply");
    }

    if (selfListings[listingHash].rewards > 0) {
      // burn any rewards remainder, the market has this amount banked atm.
      selfMarketToken.burn(selfListings[listingHash].rewards);
    }

    // remove the keys entry first
    uint deleted = selfListings[listingHash].index; // being removed
    bytes32 moved = selfListingKeys[selfListingKeys.length - 1]; // moving to overwrite 'deleted'
    selfListingKeys[deleted] = moved; // delete target now overwritten
    selfListingKeys.length--;
    // update the index of moved
    selfListings[moved].index = deleted;
    // finally we can kill it
    delete selfListings[listingHash];
    emit ListingDeletedEvent(listingHash);
    return true;
  }

  /**
    @dev Determines if an applicant becomes a listing
    @param listingHash The listingHash representing the applicant in question
  */
  function resolveApplication(bytes32 listingHash) public returns(bool) {
    require(selfVoting.inCouncil(msg.sender), "Error:Market.resolveApplication - Sender must be council member");
    // we assure it's not a listing already, there is a candidate, and voting has concluded
    require(!selfListings[listingHash].listed, "Error:Market.resolveApplication - Already listed");
    require(selfVoting.candidateIs(listingHash, "application"), "Error:Market.resolveApplication - Must be an application");
    require(selfVoting.pollClosed(listingHash), "Error:Market.resolveApplication - Polls for this candidate must be closed");

    // Case: listing accepted
    if (selfVoting.didPass(listingHash, selfParameterizer.getQuorum())) {
      // Listing storage listing = selfListings[listingHash];
      selfListings[listingHash].listed = true;
      // NOTE the market acts as a bank for the minted tokens as well (same as deposits)
      uint amount = selfParameterizer.getListReward();
      // there is no `to` adress as all mint balances are banked by the market
      require(selfMarketToken.mint(amount), "Error:Market.resolveApplication - Could not mint rewards");
      // no need to .add here
      selfListings[listingHash].rewards = amount;
      emit ListedEvent(
        listingHash,
        selfListings[listingHash].owner,
        amount,
        selfListings[listingHash].supply
      );
    } else {
      // simply remove the listing and assure it happened.
      if (removeListing(listingHash) != true) {
        // using revert as a pattern on internal method calls
        revert("Error:Market.resolveApplication - Could not remove applicant");
      }
    }

    // in either case the candidate is pruned
    require(selfVoting.removeCandidate(listingHash), "");
    return true;
  }

  /**
    @dev Determines the winner in a challenge. Rewards the winner and possibly deletes the listingHash
    @param listingHash A listingHash with a challenge that is to be resolved
    @return true if succeeded
  */
  function resolveChallenge(bytes32 listingHash) public returns(bool) {
    require(selfVoting.inCouncil(msg.sender), "Error:Market.resolveChallenge - Sender must be council member");
    // dont operate on non challenges. TODO save the bytes arrays for kinds as CONST?
    require(selfVoting.candidateIs(listingHash, "challenge"), "Error:Market.resolveChallenge - Must be a challenge");
    require(selfVoting.pollClosed(listingHash), "Error:Market.resolveChallenge - Polls for this candidate must be closed");
    // somebody gets a prize...
    uint stake = selfParameterizer.getChallengeStake();

    // Case: challenge won
    if (selfVoting.didPass(listingHash, selfParameterizer.getQuorum())) {
      removeListing(listingHash);
      // Both refund the challenger their stake and give them the listing's
      require(selfMarketToken.transfer(selfChallenges[listingHash].challenger, stake.mul(2)), "Error:market.resolveChallenge - Could not transfer MarketToken");
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
    require(selfVoting.removeCandidate(listingHash), "Error:Market.resolveChallenge - Could not remove candidate");
    return true;
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

    listing.supply = listing.supply.sub(amount);
    require(selfMarketToken.transfer(msg.sender, amount), "Error:Market.withdraw - Could not transfer tokens to listing owner");

    emit ListingWithdrawEvent(
      listingHash,
      msg.sender,
      amount,
      listing.supply,
      listing.rewards
    );
  }

  event ApplicationFailedEvent(bytes32 indexed listingHash);
  event AppliedEvent(bytes32 indexed listingHash, address indexed applicant, bytes32 indexed dataHash, uint voteBy, uint amount);
  event ChallengedEvent(bytes32 indexed listingHash, address indexed challenger);
  event ChallengeFailedEvent(bytes32 indexed listingHash, address indexed challenger, uint reward);
  event ChallengeSucceededEvent(bytes32 indexed listingHash, address indexed challenger, uint reward);
  event InvestedEvent(address indexed investor, uint offered, uint taken, uint minted);
  event ListedEvent(bytes32 indexed listingHash, address indexed owner, uint reward, uint supply);
  event ListingDepositEvent(bytes32 indexed listingHash, address indexed owner, uint added, uint supply, uint rewards);
  event ListingDeletedEvent(bytes32 indexed listingHash);
  event ListingWithdrawEvent(bytes32 indexed listingHash, address indexed owner, uint added, uint supply, uint rewards);
}
