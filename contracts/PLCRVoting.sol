pragma solidity 0.4.25;

import "./SafeMath.sol";
import "./IERC20.sol";

contract PLCRVoting {
  // Attribute store state vars
  mapping(bytes32 => uint) private selfStore;

  // DLL state vars and struct
  uint8 constant private NULL_NODE_ID = 0;

  struct DllNode {
    uint next;
    uint prev;
  }

  mapping(address => mapping(uint => DllNode)) private selfDllMap; // maps user address to linked list

  // PLCRVoting state vars and structs
  using SafeMath for uint;

  struct Poll {
    uint commitExpiry; // expiration date of commit period for poll
    uint revealExpiry; // expiration date of reveal period for poll
    uint voteQuorum; // number of votes required for a proposal to pass
    uint votesFor; // tally of votes supporting proposal
    uint votesAgainst; // tally of votes countering proposal
    mapping(address => bool) didCommit;  /// indicates whether an address committed a vote for this poll
    mapping(address => bool) didReveal;   /// indicates whether an address revealed a vote for this poll
  }

  uint constant private INITIAL_POLL_NONCE = 0;
  uint private selfPollNonce;
  mapping(uint => Poll) private selfPollMap; // maps poll ID to Poll struct
  mapping(address => uint) private selfVoteTokenBalance; // maps user's address to voteToken balance

  // Our voting contract interacts with a MarketToken, but only cares that it adhere to an ERC20 api (which it does)
  IERC20 private selfToken;

  // ----------------------------
  // Attribute store functionality (private)
  // We have removed the "Data" struct and went directly to a local mapping
  // -----------------------------
  function getStoreAttribute(bytes32 uuid, string name) private view returns (uint) {
    bytes32 key = keccak256(abi.encodePacked(uuid, name));
    return selfStore[key];
  }

  function setStoreAttribute(bytes32 uuid, string name, uint val) private {
    bytes32 key = keccak256(abi.encodePacked(uuid, name));
    selfStore[key] = val;
  }

  // ------------------
  // DLL functionality (private)
  // -----------------
  function dllContains(address sender, uint curr) private view returns (bool) {
    if (dllIsEmpty(sender) || curr == NULL_NODE_ID) {
      return false;
    }

    bool isSingleNode = (getDllStart(sender) == curr) && (getDllEnd(sender) == curr);
    bool isNullNode = (getDllNext(sender, curr) == NULL_NODE_ID) && (getDllPrev(sender, curr) == NULL_NODE_ID);
    return isSingleNode || !isNullNode;
  }

  function dllIsEmpty(address sender) private view returns (bool) {
    return getDllStart(sender) == NULL_NODE_ID;
  }

  function getDllEnd(address sender) private view returns (uint) {
    return getDllPrev(sender, NULL_NODE_ID);
  }

  function getDllNext(address sender, uint curr) private view returns (uint) {
    return selfDllMap[sender][curr].next;
  }

  function getDllPrev(address sender, uint curr) private view returns (uint) {
    return selfDllMap[sender][curr].prev;
  }

  function getDllStart(address sender) private view returns (uint) {
    return getDllNext(sender, NULL_NODE_ID);
  }

  /**
   * @dev Inserts a new node between prev and next. When inserting a node already existing in
   * the list it will be automatically removed from the old position.
   * @param prev the node which new will be inserted after
   * @param curr the id of the new node being inserted
   * @param next the node which new will be inserted before
  */
  function dllInsert(
    address sender,
    uint prev,
    uint curr,
    uint next
  ) private {
    require(curr != NULL_NODE_ID, "Error:PLCRVoting.dllInsert - ID of a new node may not be 0");

    dllRemove(sender, curr);

    // TODO are there savings to storing a CONST here for the message?
    require(prev == NULL_NODE_ID || dllContains(sender, prev), "Error:PLCRVoting.dLLInsert - Connot insert into DLL");
    require(next == NULL_NODE_ID || dllContains(sender, next), "Error:PLCRVoting.dLLInsert - Connot insert into DLL");

    require(getDllNext(sender, prev) == next, "Error:PLCRVoting.dLLInsert - Connot insert into DLL");
    require(getDllPrev(sender, next) == prev, "Error:PLCRVoting.dLLInsert - Connot insert into DLL");

    selfDllMap[sender][curr].prev = prev;
    selfDllMap[sender][curr].next = next;

    selfDllMap[sender][prev].next = curr;
    selfDllMap[sender][next].prev = curr;
  }

  function dllRemove(address sender, uint curr) private {
    if (!dllContains(sender, curr)) {
      return;
    }

    uint next = getDllNext(sender, curr);
    uint prev = getDllPrev(sender, curr);

    selfDllMap[sender][next].prev = prev;
    selfDllMap[sender][prev].next = next;

    delete selfDllMap[sender][curr];
  }

  // ------------------------
  // PLCRVoting functionality
  // ------------------------

  /**
  @param tokenAddr The address where the Network/Market token contract is deployed
  */
  constructor(address tokenAddr) public {
    selfToken = IERC20(tokenAddr);
    selfPollNonce = INITIAL_POLL_NONCE;
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
    require(selfToken.balanceOf(msg.sender) >= numTokens, "Error:Voting.requestVotingRights - Sender does not have sufficient funds");
    selfVoteTokenBalance[msg.sender] = selfVoteTokenBalance[msg.sender].add(numTokens);
    require(selfToken.transferFrom(msg.sender, this, numTokens), "Error:Voting.requestVotingRights - Token.transferFrom failure");
    emit VotingRightsGrantedEvent(numTokens, msg.sender);
  }

  /**
  @notice Withdraw numTokens ERC20 tokens from the voting contract, revoking these voting rights
  @param numTokens The number of ERC20 tokens desired in exchange for voting rights
  */
  function withdrawVotingRights(uint numTokens) external {
    uint availableTokens = selfVoteTokenBalance[msg.sender].sub(getLockedTokens(msg.sender));
    require(availableTokens >= numTokens, "Error:Voting.withdrawVotingRights - Insufficient available tokens");
    selfVoteTokenBalance[msg.sender] = selfVoteTokenBalance[msg.sender].sub(numTokens);
    require(selfToken.transfer(msg.sender, numTokens), "Error:Voting.withdrawVotingRights - Token.transfer failure");
    emit VotingRightsWithdrawnEvent(numTokens, msg.sender);
  }

  /**
  @dev Unlocks tokens locked in unrevealed vote where poll has ended
  @param id Integer identifier associated with the target poll
  */
  function rescueTokens(uint id) external {
    require(isExpired(selfPollMap[id].revealExpiry), "Error:Voting.rescueTokens - Poll must have ended");
    require(dllContains(msg.sender, id), "Error:Voting.rescueTokens - Poll not found");

    dllRemove(msg.sender, id);
    emit TokensRescuedEvent(id, msg.sender);
  }

  // =================
  // VOTING INTERFACE:
  // =================

  /**
  @notice Commits vote using hash of choice and secret salt to conceal vote until reveal
  @param id Integer identifier associated with target poll
  @param secretHash Commit keccak256 hash of voter's choice and salt (tightly packed in this order)
  @param numTokens The number of tokens to be committed towards the target poll
  @param prevId The ID of the poll that the user has voted the maximum number of tokens in which is still less than or equal to numTokens
  */
  function commitVote(
    uint id,
    bytes32 secretHash,
    uint numTokens,
    uint prevId
  ) external
  {
    require(commitPeriodActive(id), "Error:Voting.commitVote - Target poll must be active");
    require(selfVoteTokenBalance[msg.sender] >= numTokens, "Error:Voting.commitVote - Sender cannot overspend"); // prevent user from overspending
    require(id != 0, "Error:Voting.commitVote - Poll ID cannot be 0"); // prevent user from committing to zero node placeholder

    // Check if prevId exists in the user's DLL or if prevId is 0
    require(prevId == 0 || dllContains(msg.sender, prevId), "Error:Voting.commitVote - Poll must exist");

    uint nextId = getDllNext(msg.sender, prevId);

    // if nextId is equal to id, id is being updated,
    if (nextId == id) {
      nextId = getDllNext(msg.sender, id);
    }

    require(validPosition(prevId, nextId, msg.sender, numTokens), "Error:Voting.commitVote - Poll position is not valid");

    dllInsert(msg.sender, prevId, id, nextId);

    bytes32 uuid = attrUuid(msg.sender, id);

    setStoreAttribute(uuid, "numTokens", numTokens);
    setStoreAttribute(uuid, "commitHash", uint(secretHash));

    selfPollMap[id].didCommit[msg.sender] = true;
    emit VoteCommittedEvent(id, numTokens, msg.sender);
  }

  /**
  @dev Compares previous and next poll's committed tokens for sorting purposes
  @param prevId Integer identifier associated with previous poll in sorted order
  @param nextId Integer identifier associated with next poll in sorted order
  @param voter Address of user to check DLL position for
  @param numTokens The number of tokens to be committed towards the poll (used for sorting)
  @return valid Boolean indication of if the specified position maintains the sort
  */
  function validPosition(
    uint prevId,
    uint nextId,
    address voter,
    uint numTokens
  ) public view returns (bool valid)
  {
    bool prevValid = (numTokens >= getNumTokens(voter, prevId));
    // if next is zero node, numTokens does not need to be greater
    bool nextValid = (numTokens <= getNumTokens(voter, nextId) || nextId == 0);
    return prevValid && nextValid;
  }

  /**
  @notice Reveals vote with choice and secret salt used in generating commitHash to attribute committed tokens
  @param id Integer identifier associated with target poll
  @param voteOption Vote choice used to generate commitHash for associated poll
  @param salt Secret number used to generate commitHash for associated poll
  */
  function revealVote(uint id, uint voteOption, uint salt) external {
    // Make sure the reveal period is active
    require(revealPeriodActive(id), "Error:Voting.revealVote - Reveal period must be active");
    require(selfPollMap[id].didCommit[msg.sender], "Error:Voting.revealVote - Sender must have committed their vote"); // make sure user has committed a vote for this poll
    require(!selfPollMap[id].didReveal[msg.sender], "Error:Voting.revealVote - Sender cannot reveal their vote multiple times"); // prevent user from revealing multiple times
    require(keccak256(abi.encodePacked(voteOption, salt)) == getCommitHash(msg.sender, id), "Error:Voting.revealVote - Hash mismatch"); // compare resultant hash from inputs to original commitHash

    uint numTokens = getNumTokens(msg.sender, id);

    if (voteOption == 1) {// apply numTokens to appropriate poll choice
      selfPollMap[id].votesFor = selfPollMap[id].votesFor.add(numTokens);
    } else {
      selfPollMap[id].votesAgainst = selfPollMap[id].votesAgainst.add(numTokens);
    }

    // remove the node referring to this vote upon reveal
    dllRemove(msg.sender, id);

    selfPollMap[id].didReveal[msg.sender] = true;

    emit VoteRevealedEvent(
      id,
      numTokens,
      selfPollMap[id].votesFor,
      selfPollMap[id].votesAgainst,
      voteOption,
      msg.sender
    );
  }

  /**
  @param id Integer identifier associated with target poll
  @param salt Arbitrarily chosen integer used to generate secretHash
  @return correctVotes Number of tokens voted for winning option
  */
  function getNumPassingTokens(address voter, uint id, uint salt) external view returns (uint correctVotes) {
    require(pollEnded(id), "Error:Voting.getNumPassingTokens - Poll must have ended");
    require(selfPollMap[id].didReveal[voter], "Error:Voting.getNumPassingTokens - Voter must have revealed their vote");

    uint winningChoice = isPassed(id) ? 1 : 0;
    bytes32 winnerHash = keccak256(abi.encodePacked(winningChoice, salt));
    bytes32 commitHash = getCommitHash(voter, id);

    require(winnerHash == commitHash, "Error:Voting.getNumPassingTokens - Hash mismatch");

    return getNumTokens(voter, id);
  }

  // ==================
  // POLLING INTERFACE:
  // ==================

  /**
  @dev Initiates a poll with canonical configured parameters at poll ID emitted by PollCreated event
  @param voteQuorum Type of majority (out of 100) that is necessary for poll to be successful
  @param commitDuration Length of desired commit period in seconds
  @param revealDuration Length of desired reveal period in seconds
  */
  function startPoll(uint voteQuorum, uint commitDuration, uint revealDuration) external returns (uint id) {
    selfPollNonce = selfPollNonce.add(1);

    uint commitEndDate = block.timestamp.add(commitDuration);
    uint revealEndDate = commitEndDate.add(revealDuration);

    selfPollMap[selfPollNonce] = Poll({
      voteQuorum: voteQuorum,
      commitExpiry: commitEndDate,
      revealExpiry: revealEndDate,
      votesFor: 0,
      votesAgainst: 0
    });

    emit PollCreatedEvent(
      voteQuorum,
      commitEndDate,
      revealEndDate,
      selfPollNonce,
      msg.sender
    );

    return selfPollNonce;
  }

  /**
  @notice Determines if proposal has passed
  @dev Check if votesFor out of totalVotes exceeds votesQuorum (requires pollEnded)
  @param id Integer identifier associated with target poll
  */
  function isPassed(uint id) public view returns (bool passed) {
    require(pollEnded(id), "Error:Voting.isPassed - Poll must have ended");

    Poll memory poll = selfPollMap[id];
    return (poll.votesFor.mul(100)) > (poll.voteQuorum.mul(poll.votesFor.add(poll.votesAgainst)));
  }

  // ----------------
  // POLLING HELPERS:
  // ----------------

  /**
  @dev Gets the total winning votes for reward distribution purposes
  @param id Integer identifier associated with target poll
  @return Total number of votes committed to the winning option for specified poll
  */
  function getTotalNumberOfTokensForWinningOption(uint id) external view returns (uint numTokens) {
    require(pollEnded(id), "Error:Voting.getTotalNumberOfTokensForWinningOption - Poll must have ended");

    if (isPassed(id)) {
      return selfPollMap[id].votesFor;
    } else {
      return selfPollMap[id].votesAgainst;
    }
  }

  /**
  @notice Determines if poll is over
  @dev Checks isExpired for specified poll's revealEndDate
  @param id Integer identifier associated with target poll
  @return Boolean indication of whether polling period is over
  */
  function pollEnded(uint id) public view returns (bool ended) {
    require(pollExists(id), "Error:Voting.pollEnded - Poll must exist");

    return isExpired(selfPollMap[id].revealExpiry);
  }

  /**
  @notice Checks if the commit period is still active for the specified poll
  @dev Checks isExpired for the specified poll's commitEndDate
  @param id Integer identifier associated with target poll
  @return Boolean indication of isCommitPeriodActive for target poll
  */
  function commitPeriodActive(uint id) public view returns (bool active) {
    require(pollExists(id), "Error:Voting.commitPeriodActive - Poll must exist");

    return !isExpired(selfPollMap[id].commitExpiry);
  }

  /**
  @notice Checks if the reveal period is still active for the specified poll
  @dev Checks isExpired for the specified poll's revealEndDate
  @param id Integer identifier associated with target poll
  */
  function revealPeriodActive(uint id) public view returns (bool active) {
    require(pollExists(id), "Error:Voting.revealPeriodActive - Poll must exist");

    return !isExpired(selfPollMap[id].revealExpiry) && !commitPeriodActive(id);
  }

  /**
  @dev Checks if user has committed for specified poll
  @param voter Address of user to check against
  @param id Integer identifier associated with target poll
  @return Boolean indication of whether user has committed
  */
  function didCommit(address voter, uint id) external view returns (bool committed) {
    require(pollExists(id), "Error:Parameterizer.didCommit - Poll must exist");

    return selfPollMap[id].didCommit[voter];
  }

  /**
  @dev Checks if user has revealed for specified poll
  @param voter Address of user to check against
  @param id Integer identifier associated with target poll
  @return Boolean indication of whether user has revealed
  */
  function didReveal(address voter, uint id) external view returns (bool revealed) {
    require(pollExists(id), "Error:Voting.didReveal - Poll must exist");

    return selfPollMap[id].didReveal[voter];
  }

  /**
  @dev Return the commit expiry for a given poll
  @param id Integer identifier associated with target poll
  @return the expiry as uint
  */
  function getPollCommitExpiry(uint id) external view returns (uint expiry) {
    require(pollExists(id), "Error:Voting.getPollCommitExpiry - Poll must exist");

    return selfPollMap[id].commitExpiry;
  }

  /**
  @dev Return the reveal expiry for a given poll
  @param id Integer identifier associated with target poll
  @return the expiry as uint
  */
  function getPollRevealExpiry(uint id) external view returns (uint expiry) {
    require(pollExists(id), "Error:Voting.getPollrevealExpiry - Poll must exist");

    return selfPollMap[id].revealExpiry;
  }

  /**
  @dev Checks if a poll exists
  @param id The poll ID whose existance is to be evaluated.
  @return Boolean Indicates whether a poll exists for the provided poll ID
  */
  function pollExists(uint id) public view returns (bool exists) {
    return (id != 0 && id <= selfPollNonce);
  }

  // ---------------------------
  // DOUBLE-LINKED-LIST HELPERS:
  // ---------------------------

  /**
  @dev Gets the bytes32 commitHash property of target poll
  @param voter Address of user to check against
  @param id Integer identifier associated with target poll
  @return Bytes32 hash property attached to target poll
  */
  function getCommitHash(address voter, uint id) public view returns (bytes32 commitHash) {
    return bytes32(getStoreAttribute(attrUuid(voter, id), "commitHash"));
  }

  /**
  @dev Wrapper for getAttribute with attrName="numTokens"
  @param voter Address of user to check against
  @param id Integer identifier associated with target poll
  @return Number of tokens committed to poll in sorted poll-linked-list
  */
  function getNumTokens(address voter, uint id) public view returns (uint numTokens) {
    return getStoreAttribute(attrUuid(voter, id), "numTokens");
  }

  /**
  @dev Gets top element of sorted poll-linked-list
  @param voter Address of user to check against
  @return Integer identifier to poll with maximum number of tokens committed to it
  */
  function getLastNode(address voter) public view returns (uint id) {
    return getDllPrev(voter, 0);
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
  is found, if the provided poll ID matches the found nodeID, this operation is an in-place
  update. In that case, return the previous node of the node being updated. Otherwise return the
  first node that was found with a value less than or equal to the provided numTokens.
  @param voter The voter whose DLL will be searched
  @param numTokens The value for the numTokens attribute in the node to be inserted
  @param id Integer identifier associated with target poll
  @return the node which the propoded node should be inserted after
  */
  function getInsertPointForNumTokens(address voter, uint numTokens, uint id)
  external view returns (uint prevNode)
  {
    // Get the last node in the list and the number of tokens in that node
    uint nodeId = getLastNode(voter);
    uint tokensInNode = getNumTokens(voter, nodeId);

    // Iterate backwards through the list until reaching the root node
    while (nodeId != 0) {
      // Get the number of tokens in the current node
      tokensInNode = getNumTokens(voter, nodeId);
      if (tokensInNode <= numTokens) { // We found the insert point!
        if (nodeId == id) {
          // This is an in-place update. Return the prev node of the node being updated
          nodeId = getDllPrev(voter, nodeId);
        }
        // Return the insert point
        return nodeId;
      }
      // We did not find the insert point. Continue iterating backwards through the list
      nodeId = getDllPrev(voter, nodeId);
    }

    // The list is empty, or a smaller value than anything else in the list is being inserted
    return nodeId;
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
  @param id Integer identifier associated with target poll
  @return UUID Hash which is deterministic from user and poll ID
  */
  function attrUuid(address user, uint id) private pure returns (bytes32 uuid) {
    return keccak256(abi.encodePacked(user, id));
  }

  event VoteCommittedEvent(uint indexed id, uint numTokens, address indexed voter);
  event VoteRevealedEvent(uint indexed id, uint numTokens, uint votesFor, uint votesAgainst, uint indexed choice, address indexed voter);
  event PollCreatedEvent(uint voteQuorum, uint commitExpiry, uint revealExpiry, uint indexed id, address indexed creator);
  event VotingRightsGrantedEvent(uint numTokens, address indexed voter);
  event VotingRightsWithdrawnEvent(uint numTokens, address indexed voter);
  event TokensRescuedEvent(uint indexed id, address indexed voter);
}
