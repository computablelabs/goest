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

  MarketToken private selfMarketToken;
  IERC20 private selfNetworkToken;
  Parameterizer private selfParameterizer;
  Voting private selfVoting;
  // will be the market factory
  address private selfOwner;

  // voting candidate kinds are denoted by uint8s. Not using an enum for just 2
  uint8 constant APPLICATION  = 1;
  uint8 constant CHALLENGE  = 2;

  /**
    @param marketTokenAddr Address of the deployed market token contract
    @param votingAddr Address of the deployed voting contract
    @param parameterizerAddr Address of the deployed parameterizer contract
  */
  constructor(
    address marketTokenAddr,
    address networkTokenAddr,
    address parameterizerAddr,
    address votingAddr
  ) public
  {
    selfOwner = msg.sender;
    selfMarketToken = MarketToken(marketTokenAddr);
    selfNetworkToken = IERC20(networkTokenAddr);
    selfParameterizer = Parameterizer(parameterizerAddr);
    selfVoting = Voting(votingAddr);
  }

  /**
    @dev Starts a challenge for a listing
    @param listingHash The listing being challenged.
  */
  function challenge(bytes32 listingHash) external {
    uint stake = selfParameterizer.getChallengeStake();
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
  function depositToListing(bytes32 listingHash, uint amount) external {
    // Listing storage listing = selfListings[listingHash];
    require(isListing(listingHash) == true, "Error:Market.depositToListing - Listing does not exist");

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
    require(isListed(listingHash) == true, "Error:Market.exit - Must be listed");
    // Cannot exit during ongoing challenge
    require(selfVoting.isCandidate(listingHash) != true, "Error:Market.exit - Cannot exit during a challenge");

    // Delete listingHash & return tokens
    removeListing(listingHash);
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

  function getListing(bytes32 listingHash) external view returns (bool, address, uint, bytes32, uint) {
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
    @dev Given network tokens as entry, mint and return market tokens to the
    caller as dictated by the pricing curve.
    @param offered Number of network tokens offered for investment
  */
  function invest(uint offered) external {
    require(isListingOwner(msg.sender) != true, "Error:Market.invest - Cannot invest while owning an active listing");
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
    uint mint = taken.div(price);
    // NOTE this is, at origin, owned by the market
    selfMarketToken.mint(mint);
    // now we can transfer those minted market tokens to the investor
    selfMarketToken.transfer(msg.sender, mint);
    // sender is an investor now (if not already)
    if(isInvestor(msg.sender) == true) {
      selfInvestors[msg.sender].invested = selfInvestors[msg.sender].invested.add(taken);
      selfInvestors[msg.sender].minted = selfInvestors[msg.sender].minted.add(mint);
      // TODO check here for council membership when changing to threshold rules...
    } else {
      Investor storage i = selfInvestors[msg.sender];
      // record what we took and gave - NOTE no need to .add here
      i.invested = taken;
      i.minted = mint;
      // address onto the end of the investor keys array, index stored in the struct
      i.index = selfInvestorKeys.push(msg.sender) - 1;
      // currently any new investor is a council member, change when switching to threshold rules TODO
      selfVoting.addToCouncil(msg.sender);
    }

    // the total invested in the market is bumped
    selfInvested = selfInvested.add(taken);

    emit InvestedEvent(
      msg.sender,
      offered,
      taken,
      mint
    );
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
  function list(string calldata listing, bytes32 dataHash, uint amount) external {
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
      listing.supply = amount; // optional arg here, may be 0
      // Transfers tokens from user to Market contract (market acts as banker for all listing market tokens)
      selfMarketToken.transferFrom(listing.owner, address(this), amount);
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
    require(selfListings[listingHash].listed == true, "Error:Market.removeListing - Must be a listing");

    // Transfers any remaining balance back to the owner and burns any minted tokens
    if (selfListings[listingHash].supply > 0) {
      selfMarketToken.transfer(selfListings[listingHash].owner, selfListings[listingHash].supply);
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
    require(selfVoting.inCouncil(msg.sender) == true, "Error:Market.resolveApplication - Sender must be council member");
    // we assure it's not a listing already, there is a candidate, and voting has concluded
    require(selfVoting.candidateIs(listingHash, APPLICATION) == true, "Error:Market.resolveApplication - Must be an application");
    // the above and below remove the need to use isListing
    require(selfListings[listingHash].listed != true, "Error:Market.resolveApplication - Already listed");
    require(selfVoting.pollClosed(listingHash) == true, "Error:Market.resolveApplication - Polls for this candidate must be closed");

    // Case: listing accepted
    if (selfVoting.didPass(listingHash, selfParameterizer.getQuorum()) == true) {
      selfListings[listingHash].listed = true;
      // NOTE the market acts as a bank for the minted tokens as well (same as deposits)
      uint amount = selfParameterizer.getListReward();
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
    require(selfVoting.inCouncil(msg.sender) == true, "Error:Market.resolveChallenge - Sender must be council member");
    // dont operate on non challenges. TODO save the bytes arrays for kinds as CONST?
    require(selfVoting.candidateIs(listingHash, CHALLENGE) == true, "Error:Market.resolveChallenge - Must be a challenge");
    require(selfVoting.pollClosed(listingHash) == true, "Error:Market.resolveChallenge - Polls for this candidate must be closed");
    // somebody gets a prize...
    uint stake = selfParameterizer.getChallengeStake();

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
  function withdrawFromListing(bytes32 listingHash, uint amount) external {
    require(isListing(listingHash) == true, "Error:Market.withdrawFromListing - Listing does not exist");
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
