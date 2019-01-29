pragma solidity ^0.4.21;

// This contract is under development.
// Refer to readme for further details.

library SafeMath {

  /**
  * @dev Multiplies two numbers, throws on overflow.
  */
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    assert(c / a == b);
    return c;
  }

  /**
  * @dev Integer division of two numbers, truncating the quotient.
  */
  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    // uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return a / b;
  }

  /**
  * @dev Subtracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  /**
  * @dev Adds two numbers, throws on overflow.
  */
  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}

contract TomoValidator {
    using SafeMath for uint256;

    event Vote(address _voter, address _candidate, uint256 _cap);
    event Unvote(address _voter, address _candidate, uint256 _cap);
    event Propose(address _owner, address _candidate, uint256 _cap);
    event Resign(address _owner, address _candidate);
    event Withdraw(address _owner, uint256 _blockNumber, uint256 _cap);

    struct ValidatorState {
        address owner;
        bool isCandidate;
        uint256 cap;
        mapping(address => uint256) voters;
    }

    struct WithdrawState {
      mapping(uint256 => uint256) caps;
      uint256[] blockNumbers;
    }

    mapping(address => WithdrawState) withdrawsState;

    mapping(address => ValidatorState) validatorsState;
    mapping(address => address[]) voters;
    
    // KYC mapping
    mapping(address => bytes32) KYCData;
    mapping(address => uint) public invalidKYCCount;
    mapping(address => mapping(address => bool)) public hasVotedInvalid;
    mapping(address => address[]) public ownerToCandidate;
    address[] public owners;
    
    address[] public candidates;

    uint256 public candidateCount = 0;
    uint256 public minCandidateCap;
    uint256 public minVoterCap;
    uint256 public maxValidatorNumber;
    uint256 public candidateWithdrawDelay;
    uint256 public voterWithdrawDelay;

    modifier onlyValidCandidateCap {
        // anyone can deposit X TOMO to become a candidate
        require(msg.value >= minCandidateCap);
        _;
    }

    modifier onlyValidVoterCap {
        
        require(msg.value >= minVoterCap);
        _;
    }
    
    modifier onlyKYCWhitelisted {
        if(KYCData[msg.sender] != "")
        {_;}
        else{
            address addr = msg.sender;
            while(validatorsState[addr].owner!=address(0)){
                addr = validatorsState[addr].owner;
            }
            if(KYCData[addr] != "")
            {_;}
        }
    }

    modifier onlyOwner(address _candidate) {
        require(validatorsState[_candidate].owner == msg.sender);
        _;
    }

    modifier onlyCandidate(address _candidate) {
        require(validatorsState[_candidate].isCandidate);
        _;
    }

    modifier onlyValidCandidate (address _candidate) {
        require(validatorsState[_candidate].isCandidate);
        _;
    }

    modifier onlyNotCandidate (address _candidate) {
        require(!validatorsState[_candidate].isCandidate);
        _;
    }

    modifier onlyValidVote (address _candidate, uint256 _cap) {
        require(validatorsState[_candidate].voters[msg.sender] >= _cap);
        if (validatorsState[_candidate].owner == msg.sender) {
            require(validatorsState[_candidate].voters[msg.sender].sub(_cap) >= minCandidateCap);
        }
        _;
    }

    modifier onlyValidWithdraw (uint256 _blockNumber, uint _index) {
        require(_blockNumber > 0);
        require(block.number >= _blockNumber);
        require(withdrawsState[msg.sender].caps[_blockNumber] > 0);
        require(withdrawsState[msg.sender].blockNumbers[_index] == _blockNumber);
        _;
    }

    function TomoValidator (
        address[] _candidates,
        uint256[] _caps,
        address _firstOwner,
        uint256 _minCandidateCap,
        uint256 _minVoterCap,
        uint256 _maxValidatorNumber,
        uint256 _candidateWithdrawDelay,
        uint256 _voterWithdrawDelay
    ) public {
        minCandidateCap = _minCandidateCap;
        minVoterCap = _minVoterCap;
        maxValidatorNumber = _maxValidatorNumber;
        candidateWithdrawDelay = _candidateWithdrawDelay;
        voterWithdrawDelay = _voterWithdrawDelay;
        candidateCount = _candidates.length;
        owners.push(_firstOwner);
        for (uint256 i = 0; i < _candidates.length; i++) {
            candidates.push(_candidates[i]);
            validatorsState[_candidates[i]] = ValidatorState({
                owner: _firstOwner,
                isCandidate: true,
                cap: _caps[i]
            });
            voters[_candidates[i]].push(_firstOwner);
            ownerToCandidate[_firstOwner].push(_candidates[i]);
            validatorsState[_candidates[i]].voters[_firstOwner] = minCandidateCap;
        }
    }
    
    
    // KYC : Upload KYC
    function uploadKYC(bytes32 _kycdata) external {
        require(KYCData[msg.sender]=="");
        KYCData[msg.sender]=_kycdata;
    }

    // KYC : Upload KYC
    function propose(address _candidate) external payable onlyValidCandidateCap onlyKYCWhitelisted onlyNotCandidate(_candidate) {
        uint256 cap = validatorsState[_candidate].cap.add(msg.value);
        candidates.push(_candidate);
        validatorsState[_candidate] = ValidatorState({
            owner: msg.sender,
            isCandidate: true,
            cap: cap
        });
        validatorsState[_candidate].voters[msg.sender] = validatorsState[_candidate].voters[msg.sender].add(msg.value);
        candidateCount = candidateCount.add(1);
        owners.push(msg.sender);
        ownerToCandidate[msg.sender].push(_candidate);
        voters[_candidate].push(msg.sender);
        emit Propose(msg.sender, _candidate, msg.value);
    }

    function vote(address _candidate) external payable onlyValidVoterCap onlyValidCandidate(_candidate) {
        validatorsState[_candidate].cap = validatorsState[_candidate].cap.add(msg.value);
        if (validatorsState[_candidate].voters[msg.sender] == 0) {
            voters[_candidate].push(msg.sender);
        }
        validatorsState[_candidate].voters[msg.sender] = validatorsState[_candidate].voters[msg.sender].add(msg.value);
        emit Vote(msg.sender, _candidate, msg.value);
    }

    function getCandidates() public view returns(address[]) {
        return candidates;
    }

    function getCandidateCap(address _candidate) public view returns(uint256) {
        return validatorsState[_candidate].cap;
    }

    function getCandidateOwner(address _candidate) public view returns(address) {
        return validatorsState[_candidate].owner;
    }

    function getVoterCap(address _candidate, address _voter) public view returns(uint256) {
        return validatorsState[_candidate].voters[_voter];
    }

    function getVoters(address _candidate) public view returns(address[]) {
        return voters[_candidate];
    }

    function isCandidate(address _candidate) public view returns(bool) {
        return validatorsState[_candidate].isCandidate;
    }

    function getWithdrawBlockNumbers() public view returns(uint256[]) {
        return withdrawsState[msg.sender].blockNumbers;
    }

    function getWithdrawCap(uint256 _blockNumber) public view returns(uint256) {
        return withdrawsState[msg.sender].caps[_blockNumber];
    }

    function unvote(address _candidate, uint256 _cap) public onlyValidVote(_candidate, _cap) {
        validatorsState[_candidate].cap = validatorsState[_candidate].cap.sub(_cap);
        validatorsState[_candidate].voters[msg.sender] = validatorsState[_candidate].voters[msg.sender].sub(_cap);

        // refund after delay X blocks
        uint256 withdrawBlockNumber = voterWithdrawDelay.add(block.number);
        withdrawsState[msg.sender].caps[withdrawBlockNumber] = withdrawsState[msg.sender].caps[withdrawBlockNumber].add(_cap);
        withdrawsState[msg.sender].blockNumbers.push(withdrawBlockNumber);

        emit Unvote(msg.sender, _candidate, _cap);
    }

    function resign(address _candidate) public onlyOwner(_candidate) onlyCandidate(_candidate) {
        validatorsState[_candidate].isCandidate = false;
        candidateCount = candidateCount.sub(1);
        for (uint256 i = 0; i < candidates.length; i++) {
            if (candidates[i] == _candidate) {
                delete candidates[i];
                break;
            }
        }
        uint256 cap = validatorsState[_candidate].voters[msg.sender];
        validatorsState[_candidate].cap = validatorsState[_candidate].cap.sub(cap);
        validatorsState[_candidate].voters[msg.sender] = 0;
        // refunding after resigning X blocks
        uint256 withdrawBlockNumber = candidateWithdrawDelay.add(block.number);
        withdrawsState[msg.sender].caps[withdrawBlockNumber] = withdrawsState[msg.sender].caps[withdrawBlockNumber].add(cap);
        withdrawsState[msg.sender].blockNumbers.push(withdrawBlockNumber);
        emit Resign(msg.sender, _candidate);
    }
    
    // KYC : Upload KYC
    function VoteInvalidKYC(address _invalidMasternode) onlyValidCandidate(msg.sender) public {
        address candidateOwner = getOwner(msg.sender);
        require(!hasVotedInvalid[candidateOwner][_invalidMasternode]);
        invalidKYCCount[_invalidMasternode] += 1;
        if( invalidKYCCount[_invalidMasternode]*100/getOwnerCount() >= 75 ){
            // 75% owners say that the KYC is invalid
            for (uint i=0;i<candidates.length;i++){
                if (getOwner(candidates[i])==_invalidMasternode){
                    // logic to rmove cap.
                    
                    candidateCount = candidateCount.sub(1);
                    delete candidates[i];
                    delete validatorsState[candidates[i]];
                    delete ownerToCandidate[_invalidMasternode];
                    for(uint k=0;k<owners.length;k++){
                        if (owners[k]==_invalidMasternode)
                        delete owners[k];
                    }
                }
            }
        }
    }
    
    function InvalidPercent(address _invalidMasternode) view public returns(uint){
        return (invalidKYCCount[_invalidMasternode]*100/getOwnerCount());
    }
    
    function getOwner(address _masternode) view public returns(address){
        // address addr = _masternode;
        // while (validatorsState[addr].owner != 0) {
        //     addr = validatorsState[addr].owner;
        // }
        return validatorsState[_masternode].owner;
    } 
    
    function getOwnerCount() view public returns (uint){
        return owners.length;
    }


    function withdraw(uint256 _blockNumber, uint _index) public onlyValidWithdraw(_blockNumber, _index) {
        uint256 cap = withdrawsState[msg.sender].caps[_blockNumber];
        delete withdrawsState[msg.sender].caps[_blockNumber];
        delete withdrawsState[msg.sender].blockNumbers[_index];
        msg.sender.transfer(cap);
        emit Withdraw(msg.sender, _blockNumber, cap);
    }
}
