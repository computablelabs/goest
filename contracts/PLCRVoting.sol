pragma solidity 0.4.24;

import "./ERC20.sol";
import "./DLL.sol";
import "./AttributeStore.sol";
import "./SafeMath.sol";


/**
@title Partial-Lock-Commit-Reveal Voting scheme with ERC20 tokens
Thanks to this team: Aspyn Palatnick, Cem Ozer, Yorke Rhodes
*/
contract PLCRVoting {

  // ============
  // EVENTS:
  // ============

  event VoteCommitted(uint indexed pollID, uint numTokens, address indexed voter);
  event VoteRevealed(uint indexed pollID, uint numTokens, uint votesFor, uint votesAgainst, uint indexed choice, address indexed voter);
  event PollCreated(uint voteQuorum, uint commitEndDate, uint revealEndDate, uint indexed pollID, address indexed creator);
  event VotingRightsGranted(uint numTokens, address indexed voter);
  event VotingRightsWithdrawn(uint numTokens, address indexed voter);
  event TokensRescued(uint indexed pollID, address indexed voter);

  // ============
  // DATA STRUCTURES:
  // ============

  using AttributeStore for AttributeStore.Data;
  using DLL for DLL.Data;
  using SafeMath for uint;

  struct Poll {
    uint commitEndDate;     /// expiration date of commit period for poll
    uint revealEndDate;     /// expiration date of reveal period for poll
    uint voteQuorum;    /// number of votes required for a proposal to pass
    uint votesFor;    /// tally of votes supporting proposal
    uint votesAgainst;      /// tally of votes countering proposal
    mapping(address => bool) didCommit;  /// indicates whether an address committed a vote for this poll
    mapping(address => bool) didReveal;   /// indicates whether an address revealed a vote for this poll
  }

  /* ============
   * STATE VARIABLES:
   * pollNonce
   * pollMap
   * voteTokenBalance
   * dllMap
   * token
  ============ */

  uint constant public INITIAL_POLL_NONCE = 0;
  uint public pollNonce;

  mapping(uint => Poll) public pollMap; // maps pollID to Poll struct
  mapping(address => uint) public voteTokenBalance; // maps user's address to voteToken balance

  mapping(address => DLL.Data) dllMap;
  AttributeStore.Data store;

  ERC20 public token;

  // ============
  // CONSTRUCTOR:
  // ============

  /**
  @dev Initializes voteQuorum, commitDuration, revealDuration, and pollNonce in addition to token contract and trusted mapping
  @param tokenAddr The address where the ERC20 token contract is deployed
  */
  constructor(address tokenAddr) public {
    token = ERC20(tokenAddr);
    pollNonce = INITIAL_POLL_NONCE;
  }

  // ================
  // TOKEN INTERFACE:
  // ================

  /**
  @notice Loads numTokens ERC20 tokens into the voting contract for one-to-one voting rights
  @dev Assumes that msg.sender has approved voting contract to spend on their behalf
  @param numTokens The number of votingTokens desired in exchange for ERC20 tokens
  */
  function requestVotingRights(uint numTokens) external {
    require(token.balanceOf(msg.sender) >= numTokens, "Error:Voting.requestVotingRights - Sender does not have sufficient funds");
    voteTokenBalance[msg.sender] = voteTokenBalance[msg.sender].add(numTokens);
    require(token.transferFrom(msg.sender, this, numTokens), "Error:Voting.requestVotingRights - token.transferFrom failure");
    emit VotingRightsGranted(numTokens, msg.sender);
  }

  /**
  @notice Withdraw numTokens ERC20 tokens from the voting contract, revoking these voting rights
  @param numTokens The number of ERC20 tokens desired in exchange for voting rights
  */
  function withdrawVotingRights(uint numTokens) external {
    uint availableTokens = voteTokenBalance[msg.sender].sub(getLockedTokens(msg.sender));
    require(availableTokens >= numTokens, "Error:Voting.withdrawVotingRights - Insufficient available tokens");
    voteTokenBalance[msg.sender] = voteTokenBalance[msg.sender].sub(numTokens);
    require(token.transfer(msg.sender, numTokens), "Error:Voting.withdrawVotingRights - token.transfer failure");
    emit VotingRightsWithdrawn(numTokens, msg.sender);
  }

  /**
  @dev Unlocks tokens locked in unrevealed vote where poll has ended
  @param pollID Integer identifier associated with the target poll
  */
  function rescueTokens(uint pollID) external {
    require(isExpired(pollMap[pollID].revealEndDate), "Error:Voting.rescueTokens - Poll must have ended");
    require(dllMap[msg.sender].contains(pollID), "Error:Voting.rescueTokens - Poll not found");

    dllMap[msg.sender].remove(pollID);
    emit TokensRescued(pollID, msg.sender);
  }

  // =================
  // VOTING INTERFACE:
  // =================

  /**
  @notice Commits vote using hash of choice and secret salt to conceal vote until reveal
  @param pollID Integer identifier associated with target poll
  @param secretHash Commit keccak256 hash of voter's choice and salt (tightly packed in this order)
  @param numTokens The number of tokens to be committed towards the target poll
  @param prevPollID The ID of the poll that the user has voted the maximum number of tokens in which is still less than or equal to numTokens
  */
  function commitVote(
    uint pollID,
    bytes32 secretHash,
    uint numTokens,
    uint prevPollID
  ) external
  {
    require(commitPeriodActive(pollID), "Error:Voting.commitVote - Target poll must be active");
    require(voteTokenBalance[msg.sender] >= numTokens, "Error:Voting.commitVote - Sender cannot overspend"); // prevent user from overspending
    require(pollID != 0, "Error:Voting.commitVote - Poll ID cannot be 0"); // prevent user from committing to zero node placeholder

    // Check if prevPollID exists in the user's DLL or if prevPollID is 0
    require(prevPollID == 0 || dllMap[msg.sender].contains(prevPollID), "Error:Voting.commitVote - Poll must exist");

    uint nextPollID = dllMap[msg.sender].getNext(prevPollID);

    // if nextPollID is equal to pollID, pollID is being updated,
    nextPollID = (nextPollID == pollID) ? dllMap[msg.sender].getNext(pollID) : nextPollID;

    require(validPosition(prevPollID, nextPollID, msg.sender, numTokens), "Error:Voting.commitVote - Poll position is not valid");

    dllMap[msg.sender].insert(prevPollID, pollID, nextPollID);

    bytes32 UUID = attrUUID(msg.sender, pollID);

    store.setAttribute(UUID, "numTokens", numTokens);
    store.setAttribute(UUID, "commitHash", uint(secretHash));

    pollMap[pollID].didCommit[msg.sender] = true;
    emit VoteCommitted(pollID, numTokens, msg.sender);
  }

  /**
  @dev Compares previous and next poll's committed tokens for sorting purposes
  @param prevID Integer identifier associated with previous poll in sorted order
  @param nextID Integer identifier associated with next poll in sorted order
  @param voter Address of user to check DLL position for
  @param numTokens The number of tokens to be committed towards the poll (used for sorting)
  @return valid Boolean indication of if the specified position maintains the sort
  */
  function validPosition(
    uint prevID,
    uint nextID,
    address voter,
    uint numTokens
  ) public view returns (bool valid)
  {
    bool prevValid = (numTokens >= getNumTokens(voter, prevID));
    // if next is zero node, numTokens does not need to be greater
    bool nextValid = (numTokens <= getNumTokens(voter, nextID) || nextID == 0);
    return prevValid && nextValid;
  }

  /**
  @notice Reveals vote with choice and secret salt used in generating commitHash to attribute committed tokens
  @param pollID Integer identifier associated with target poll
  @param voteOption Vote choice used to generate commitHash for associated poll
  @param salt Secret number used to generate commitHash for associated poll
  */
  function revealVote(uint pollID, uint voteOption, uint salt) external {
    // Make sure the reveal period is active
    require(revealPeriodActive(pollID), "Error:Voting.revealVote - Reveal period must be active");
    require(pollMap[pollID].didCommit[msg.sender], "Error:Voting.revealVote - Sender must have committed their vote"); // make sure user has committed a vote for this poll
    require(!pollMap[pollID].didReveal[msg.sender], "Error:Voting.revealVote - Sender cannot reveal their vote multiple times"); // prevent user from revealing multiple times
    require(keccak256(abi.encodePacked(voteOption, salt)) == getCommitHash(msg.sender, pollID), "Error:Voting.revealVote - Hash mismatch"); // compare resultant hash from inputs to original commitHash

    uint numTokens = getNumTokens(msg.sender, pollID);

    if (voteOption == 1) {// apply numTokens to appropriate poll choice
      pollMap[pollID].votesFor = pollMap[pollID].votesFor.add(numTokens);
    } else {
      pollMap[pollID].votesAgainst = pollMap[pollID].votesAgainst.add(numTokens);
    }

    dllMap[msg.sender].remove(pollID); // remove the node referring to this vote upon reveal
    pollMap[pollID].didReveal[msg.sender] = true;

    emit VoteRevealed(
      pollID,
      numTokens,
      pollMap[pollID].votesFor,
      pollMap[pollID].votesAgainst,
      voteOption,
      msg.sender
    );
  }

  /**
  @param pollID Integer identifier associated with target poll
  @param salt Arbitrarily chosen integer used to generate secretHash
  @return correctVotes Number of tokens voted for winning option
  */
  function getNumPassingTokens(address voter, uint pollID, uint salt) public view returns (uint correctVotes) {
    require(pollEnded(pollID), "Error:Voting.getNumPassingTokens - Poll must have ended");
    require(pollMap[pollID].didReveal[voter], "Error:Voting.getNumPassingTokens - Voter must have revealed their vote");

    uint winningChoice = isPassed(pollID) ? 1 : 0;
    bytes32 winnerHash = keccak256(abi.encodePacked(winningChoice, salt));
    bytes32 commitHash = getCommitHash(voter, pollID);

    require(winnerHash == commitHash, "Error:Voting.getNumPassingTokens - Hash mismatch");

    return getNumTokens(voter, pollID);
  }

  // ==================
  // POLLING INTERFACE:
  // ==================

  /**
  @dev Initiates a poll with canonical configured parameters at pollID emitted by PollCreated event
  @param voteQuorum Type of majority (out of 100) that is necessary for poll to be successful
  @param commitDuration Length of desired commit period in seconds
  @param revealDuration Length of desired reveal period in seconds
  */
  function startPoll(uint voteQuorum, uint commitDuration, uint revealDuration) public returns (uint pollID) {
    pollNonce = pollNonce.add(1);

    uint commitEndDate = block.timestamp.add(commitDuration);
    uint revealEndDate = commitEndDate.add(revealDuration);

    pollMap[pollNonce] = Poll({
      voteQuorum: voteQuorum,
      commitEndDate: commitEndDate,
      revealEndDate: revealEndDate,
      votesFor: 0,
      votesAgainst: 0
    });

    emit PollCreated(
      voteQuorum,
      commitEndDate,
      revealEndDate,
      pollNonce,
      msg.sender
    );

    return pollNonce;
  }

  /**
  @notice Determines if proposal has passed
  @dev Check if votesFor out of totalVotes exceeds votesQuorum (requires pollEnded)
  @param pollID Integer identifier associated with target poll
  */
  function isPassed(uint pollID) public view returns (bool passed) {
    require(pollEnded(pollID), "Error:Voting.isPassed - Poll must have ended");

    Poll memory poll = pollMap[pollID];
    return (poll.votesFor.mul(100)) > (poll.voteQuorum.mul(poll.votesFor.add(poll.votesAgainst)));
  }

  // ----------------
  // POLLING HELPERS:
  // ----------------

  /**
  @dev Gets the total winning votes for reward distribution purposes
  @param pollID Integer identifier associated with target poll
  @return Total number of votes committed to the winning option for specified poll
  */
  function getTotalNumberOfTokensForWinningOption(uint pollID) public view returns (uint numTokens) {
    require(pollEnded(pollID), "Error:Voting.getTotalNumberOfTokensForWinningOption - Poll must have ended");

    if (isPassed(pollID))
      return pollMap[pollID].votesFor;
    else
      return pollMap[pollID].votesAgainst;
  }

  /**
  @notice Determines if poll is over
  @dev Checks isExpired for specified poll's revealEndDate
  @param pollID Integer identifier associated with target poll
  @return Boolean indication of whether polling period is over
  */
  function pollEnded(uint pollID) public view returns (bool ended) {
    require(pollExists(pollID), "Error:Parameterizer.pollEnded - Poll must exist");

    return isExpired(pollMap[pollID].revealEndDate);
  }

  /**
  @notice Checks if the commit period is still active for the specified poll
  @dev Checks isExpired for the specified poll's commitEndDate
  @param pollID Integer identifier associated with target poll
  @return Boolean indication of isCommitPeriodActive for target poll
  */
  function commitPeriodActive(uint pollID) public view returns (bool active) {
    require(pollExists(pollID), "Error:Parameterizer.commitPeriodActive - Poll must exist");

    return !isExpired(pollMap[pollID].commitEndDate);
  }

  /**
  @notice Checks if the reveal period is still active for the specified poll
  @dev Checks isExpired for the specified poll's revealEndDate
  @param pollID Integer identifier associated with target poll
  */
  function revealPeriodActive(uint pollID) public view returns (bool active) {
    require(pollExists(pollID), "Error:Parameterizer.revealPeriodActive - Poll must exist");

    return !isExpired(pollMap[pollID].revealEndDate) && !commitPeriodActive(pollID);
  }

  /**
  @dev Checks if user has committed for specified poll
  @param voter Address of user to check against
  @param pollID Integer identifier associated with target poll
  @return Boolean indication of whether user has committed
  */
  function didCommit(address voter, uint pollID) public view returns (bool committed) {
    require(pollExists(pollID), "Error:Parameterizer.didCommit - Poll must exist");

    return pollMap[pollID].didCommit[voter];
  }

  /**
  @dev Checks if user has revealed for specified poll
  @param voter Address of user to check against
  @param pollID Integer identifier associated with target poll
  @return Boolean indication of whether user has revealed
  */
  function didReveal(address voter, uint pollID) public view returns (bool revealed) {
    require(pollExists(pollID), "Error:Parameterizer.didReveal - Poll must exist");

    return pollMap[pollID].didReveal[voter];
  }

  /**
  @dev Checks if a poll exists
  @param pollID The pollID whose existance is to be evaluated.
  @return Boolean Indicates whether a poll exists for the provided pollID
  */
  function pollExists(uint pollID) public view returns (bool exists) {
    return (pollID != 0 && pollID <= pollNonce);
  }

  // ---------------------------
  // DOUBLE-LINKED-LIST HELPERS:
  // ---------------------------

  /**
  @dev Gets the bytes32 commitHash property of target poll
  @param voter Address of user to check against
  @param pollID Integer identifier associated with target poll
  @return Bytes32 hash property attached to target poll
  */
  function getCommitHash(address voter, uint pollID) public view returns (bytes32 commitHash) {
    return bytes32(store.getAttribute(attrUUID(voter, pollID), "commitHash"));
  }

  /**
  @dev Wrapper for getAttribute with attrName="numTokens"
  @param voter Address of user to check against
  @param pollID Integer identifier associated with target poll
  @return Number of tokens committed to poll in sorted poll-linked-list
  */
  function getNumTokens(address voter, uint pollID) public view returns (uint numTokens) {
    return store.getAttribute(attrUUID(voter, pollID), "numTokens");
  }

  /**
  @dev Gets top element of sorted poll-linked-list
  @param voter Address of user to check against
  @return Integer identifier to poll with maximum number of tokens committed to it
  */
  function getLastNode(address voter) public view returns (uint pollID) {
    return dllMap[voter].getPrev(0);
  }

  /**
  @dev Gets the numTokens property of getLastNode
  @param voter Address of user to check against
  @return Maximum number of tokens committed in poll specified
  */
  function getLockedTokens(address voter) public view returns (uint numTokens) {
    return getNumTokens(voter, getLastNode(voter));
  }

  /*
  @dev Takes the last node in the user's DLL and iterates backwards through the list searching
  for a node with a value less than or equal to the provided numTokens value. When such a node
  is found, if the provided pollID matches the found nodeID, this operation is an in-place
  update. In that case, return the previous node of the node being updated. Otherwise return the
  first node that was found with a value less than or equal to the provided numTokens.
  @param voter The voter whose DLL will be searched
  @param numTokens The value for the numTokens attribute in the node to be inserted
  @param pollID Integer identifier associated with target poll
  @return the node which the propoded node should be inserted after
  */
  function getInsertPointForNumTokens(address voter, uint numTokens, uint pollID)
  public view returns (uint prevNode)
  {
    // Get the last node in the list and the number of tokens in that node
    uint nodeID = getLastNode(voter);
    uint tokensInNode = getNumTokens(voter, nodeID);

    // Iterate backwards through the list until reaching the root node
    while (nodeID != 0) {
      // Get the number of tokens in the current node
      tokensInNode = getNumTokens(voter, nodeID);
      if (tokensInNode <= numTokens) { // We found the insert point!
        if (nodeID == pollID) {
          // This is an in-place update. Return the prev node of the node being updated
          nodeID = dllMap[voter].getPrev(nodeID);
        }
        // Return the insert point
        return nodeID;
      }
      // We did not find the insert point. Continue iterating backwards through the list
      nodeID = dllMap[voter].getPrev(nodeID);
    }

    // The list is empty, or a smaller value than anything else in the list is being inserted
    return nodeID;
  }

  // ----------------
  // GENERAL HELPERS:
  // ----------------

  /**
  @dev Checks if an expiration date has been reached
  @param terminationDate Integer timestamp of date to compare current timestamp with
  @return expired Boolean indication of whether the terminationDate has passed
  */
  function isExpired(uint terminationDate) public view returns (bool expired) {
    return (block.timestamp > terminationDate);
  }

  /**
  @dev Generates an identifier which associates a user and a poll together
  @param user Address to associate with a poll
  @param pollID Integer identifier associated with target poll
  @return UUID Hash which is deterministic from user and pollID
  */
  function attrUUID(address user, uint pollID) public pure returns (bytes32 UUID) {
    return keccak256(abi.encodePacked(user, pollID));
  }
}
