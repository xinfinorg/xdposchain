pragma solidity 0.5.0;

///@notice SafeMath library is utilized for overflow/underflow safeguarding.
import "./libs/SafeMath.sol";


/// @author XinFin Organisation
/// @title XINFIN DPOS CONSENSUS: XDPOS
contract XDCValidator {
    using SafeMath for uint256;

    struct ValidatorState {
        address validator;
        bool hasBeenElected; 
        uint256 stake;
        mapping(address => uint256) nominators;
    }

    struct WithdrawState {
      mapping(uint256 => uint256) stakes;
      uint256[] blockNumbers;
    }
    
    address[] public validators;
    address[] public candidates;

    uint256 public validatorsTotal;
    uint256 public candidatesTotal;
    
    uint256 public minCandidateStake;
    uint256 public minNominatorStake;
    
    uint256 public maxRegisteredValidators; 
    
    uint256 public candidateWithdrawDelay; 
    uint256 public nominatorWithdrawDelay; 
    
    mapping(address => WithdrawState) internal withdrawsState;
    mapping(address => ValidatorState) internal validatorsState;

    mapping(address => address[]) internal nominators;

    mapping(address => string[]) public hashedKYCs;
    
    mapping(address => mapping(address => bool)) public invalidKYCReported;
    
    mapping(address => uint) public invalidKYCsTotal;
    
    mapping(address => address[]) public validatorCandidateProposals;
    
    event UploadedKYC(
        address _validator, 
        string hashedKYC
        );
        
    event Voted(
        address _nominator, 
        address _candidate, 
        uint256 _stake
        );
        
    event Unvoted(
        address _nominator, 
        address _candidate, 
        uint256 _stake
        );
        
    event Proposed(
        address _validator, 
        address _candidate, 
        uint256 _stake
        );
        
    event Resigned(
        address _validator, 
        address _candidate
        );
        
    event Withdrawn(
        address _validator, 
        uint256 _blockNumber, 
        uint256 _stake
        );
        
    event InvalidatedValidator(
        address _invalidValidator, 
        address[] _allValidators
        );

    modifier onlyKYCWhitelisted {
       require(hashedKYCs[msg.sender].length > 0 || validatorCandidateProposals[msg.sender].length > 0);
       _;
    }
    
    modifier onlyQualifiedCandidateStake {
        require(msg.value >= minCandidateStake);
        _;
    }
    
    modifier onlyQualifiedNominatorStake {
        require(msg.value >= minNominatorStake);
        _;
    }

    modifier onlyQualifiedValidators(address _candidate) {
        require(validatorsState[_candidate].validator == msg.sender);
        _;
    }

    modifier onlyQualifiedCandidates(address _candidate) {
        require(validatorsState[_candidate].hasBeenElected);
        _;
    }
    
    modifier onlyUnqualifiedCandidates (address  _candidate) {
        require(!validatorsState[_candidate].hasBeenElected);
        _;
    }
    
    modifier onlyQualifiedHashedKYC (string memory hashedKYC) {
        bytes memory c = bytes (hashedKYC);
        require (c.length > 6);
        for (uint i = 0; i < c.length; i++)
            require (0x7ffeffe07ff7dfe03fe000000000000 & (uint(1) << uint8 (c[i])) > 0);
        _;
    }

    modifier onlyValidVotes (address _candidate, uint256 _stake) {
        require(validatorsState[_candidate].nominators[msg.sender] >= _stake);
        
        if (validatorsState[_candidate].validator == msg.sender) {
            require(validatorsState[_candidate].nominators[msg.sender].sub(_stake) >= minCandidateStake);
        }
        _;
    }  

    modifier onlyValidWithdraw (uint256 _blockNumber, uint _index) {
        require(_blockNumber > 0);
        require(block.number >= _blockNumber);
        require(withdrawsState[msg.sender].stakes[_blockNumber] > 0);
        require(withdrawsState[msg.sender].blockNumbers[_index] == _blockNumber);
        _;
    }

    ///@notice Constructor which initializes the initial configuration values.
    ///@param _candidates the initial candidates proposed by the genesis validator.
    ///@param _stakes the initial candidates individual stakes.
    ///@param _genesisValidator the first validator to exist.
    ///@param _minCandidateStake the minimum staking requirement for a candidate.
    ///@param _minNominatorStake the minimum staking requirement for a nominator.
    ///@param _maxRegisteredValidators the maximum amount of allowed validator registration.
    ///@param _candidateWithdrawDelay a delay related to candidate withdrawals.
    ///@param _nominatorWithdrawDelay a delay related to nominator withdrawals.
    ///@dev _minCandidateStake, _minNominatorStake, _maxRegisteredValidators, _candidateWithdrawDelay, _nominatorWithdrawDelay
    ///should all be modified to constant variables if known beforehand as those values are not dynamic but only read (We should help the compiler, where we can).
    ///@dev You should also consider adding a candidate LIMIT, to avoid endless loops. If the array of candidates gets too big, it will break.
    constructor (
        address[] memory _candidates,
        uint256[] memory _stakes,
        address _genesisValidator,
        uint256 _minCandidateStake,
        uint256 _minNominatorStake,
        uint256 _maxRegisteredValidators,
        uint256 _candidateWithdrawDelay,
        uint256 _nominatorWithdrawDelay
    ) public {
        minCandidateStake = _minCandidateStake;
        minNominatorStake = _minNominatorStake;
        maxRegisteredValidators = _maxRegisteredValidators;
        candidateWithdrawDelay = _candidateWithdrawDelay;
        nominatorWithdrawDelay = _nominatorWithdrawDelay;
        candidatesTotal = _candidates.length;
        validators.push(_genesisValidator);
        validatorsTotal++;
        
        for (uint256 i = 0; i < _candidates.length; i++) {
            candidates.push(_candidates[i]);
            validatorsState[_candidates[i]] = ValidatorState({
                validator: _genesisValidator,
                hasBeenElected: true,
                stake: _stakes[i]
            });
            
            nominators[_candidates[i]].push(_genesisValidator);
            validatorCandidateProposals[_genesisValidator].push(_candidates[i]);
            validatorsState[_candidates[i]].nominators[_genesisValidator] = minCandidateStake;
        }
    }

    ///@notice uploadKYC implements the functionality which allows a validator to provide KYC information.
    ///@param _hashedKYC is string format of a IPFS Hash containing the KYC information.
    function uploadKYC(string calldata _hashedKYC) external onlyQualifiedHashedKYC(_hashedKYC) {
        hashedKYCs[msg.sender].push(_hashedKYC);
        
        emit UploadedKYC(msg.sender, _hashedKYC);
    }
    
    ///@notice vote implements the functionality which allows validators to vote on candidates. 
    ///@param _candidate is the address of target candidate of concern.
    function vote(address _candidate) external payable onlyQualifiedNominatorStake onlyQualifiedCandidates(_candidate) {
        validatorsState[_candidate].stake = validatorsState[_candidate].stake.add(msg.value);
        
        if (validatorsState[_candidate].nominators[msg.sender] == 0) {
            nominators[_candidate].push(msg.sender);
        }
        
        validatorsState[_candidate].nominators[msg.sender] = validatorsState[_candidate].nominators[msg.sender].add(msg.value);
        
        emit Voted(msg.sender, _candidate, msg.value);
    }

    ///@notice propose implements the functionality which turns network participants into validators through participation.
    ///@param _candidate is the address of the target candidate of concern.
    function propose(address _candidate) external payable onlyQualifiedCandidateStake onlyKYCWhitelisted onlyUnqualifiedCandidates(_candidate) {
        uint256 stake = validatorsState[_candidate].stake.add(msg.value);
        candidates.push(_candidate);
        validatorsState[_candidate] = ValidatorState({
            validator: msg.sender,
            hasBeenElected: true,
            stake: stake
        });
        
        validatorsState[_candidate].nominators[msg.sender] = validatorsState[_candidate].nominators[msg.sender].add(msg.value);
        candidatesTotal = candidatesTotal.add(1);
        
        if (validatorCandidateProposals[msg.sender].length == 0){
            validators.push(msg.sender);
            validatorsTotal++;
        }
        
        validatorCandidateProposals[msg.sender].push(_candidate);
        nominators[_candidate].push(msg.sender);
        
        emit Proposed(msg.sender, _candidate, msg.value);
    }
    
    ///@notice withdraw implements the functionality which allows to cleanup the state of a particuclar period of time.
    ///@param _blockNumber and _index, observes and modifies informations from a certain period of time. 
    function withdraw(uint256 _blockNumber, uint _index) external onlyValidWithdraw(_blockNumber, _index) {
        uint256 stake = withdrawsState[msg.sender].stakes[_blockNumber];
        delete withdrawsState[msg.sender].stakes[_blockNumber];
        delete withdrawsState[msg.sender].blockNumbers[_index];
        msg.sender.transfer(stake);
        
        emit Withdrawn(msg.sender, _blockNumber, stake);
    }
    
    ///@notice unvote implements the functionality which allows a nominator to withdraw their vote and stake.
    ///@param _candidate is the address of the target candidate of concern.
    ///@param _stake is the amount of value staked with the vote.
    function unvote(address _candidate, uint256 _stake) external onlyValidVotes(_candidate, _stake) {
        validatorsState[_candidate].stake = validatorsState[_candidate].stake.sub(_stake);
        validatorsState[_candidate].nominators[msg.sender] = validatorsState[_candidate].nominators[msg.sender].sub(_stake);

        // refund after delay X blocks
        uint256 withdrawBlockNumber = nominatorWithdrawDelay.add(block.number);
        withdrawsState[msg.sender].stakes[withdrawBlockNumber] = withdrawsState[msg.sender].stakes[withdrawBlockNumber].add(_stake);
        withdrawsState[msg.sender].blockNumbers.push(withdrawBlockNumber);
        
        emit Unvoted(msg.sender, _candidate, _stake);
    }

    ///@notice resign implements the functionality which allows a validator to resign from an election.
    ///@param _candidate is the address of the target candidate of concern.
    function resign(address _candidate) external onlyQualifiedValidators(_candidate) onlyQualifiedCandidates(_candidate) {
        validatorsState[_candidate].hasBeenElected = false;
        candidatesTotal = candidatesTotal.sub(1);
        
        for (uint256 i = 0; i < candidates.length; i++) {
            if (candidates[i] == _candidate) {
                delete candidates[i];
                break;
            }
        }
        uint256 stake = validatorsState[_candidate].nominators[msg.sender];
        validatorsState[_candidate].stake = validatorsState[_candidate].stake.sub(stake);
        validatorsState[_candidate].nominators[msg.sender] = 0;
        // refunding after resigning X blocks
        uint256 withdrawBlockNumber = candidateWithdrawDelay.add(block.number);
        withdrawsState[msg.sender].stakes[withdrawBlockNumber] = withdrawsState[msg.sender].stakes[withdrawBlockNumber].add(stake);
        withdrawsState[msg.sender].blockNumbers.push(withdrawBlockNumber);
        
        emit Resigned(msg.sender, _candidate);
    }
    
    ///@notice invalidKYCReport implements the functionality which allows a candidate to report malicious/wrong KYC information. 
    ///@param _invalidValidator is the address of the target validator of concern.
    function invalidKYCReport(address _invalidValidator) external onlyQualifiedCandidates(msg.sender) onlyQualifiedCandidates(_invalidValidator) {
        address reportingValidator = getCandidateValidator(msg.sender);
        address invalidValidator = getCandidateValidator(_invalidValidator);
        require(!invalidKYCReported[reportingValidator][invalidValidator]);
        invalidKYCReported[reportingValidator][invalidValidator] = true;
        invalidKYCsTotal[_invalidValidator] += 1;
        
        if( invalidKYCsTotal[_invalidValidator] * 100 / validatorsTotal >= 75) {
            // 75% owners say that the KYC is invalid
            address[] memory allValidators = new address[](candidates.length-1) ;
            uint count = 0;
            for (uint i= 0; i < candidates.length; i++) {
                if (getCandidateValidator(candidates[i]) == _invalidValidator) {
                    // logic to remove stake.
                    candidatesTotal = candidatesTotal.sub(1);
                    allValidators[count++] = candidates[i];
                    delete candidates[i];
                    delete validatorsState[candidates[i]];
                    delete hashedKYCs[_invalidValidator];
                    delete validatorCandidateProposals[_invalidValidator];
                    delete invalidKYCsTotal[_invalidValidator];
                }
            }
            for (uint k=0; k < validators.length; k++) {
                        if (validators[k] == _invalidValidator) {
                            delete validators[k];
                            validatorsTotal--;
                            break;
                } 
            }
            emit InvalidatedValidator(_invalidValidator, allValidators);
        }
    }
    
    ///@notice hasBeenElected implements the functionality which allows the public to verify that a particular candidate has been elected.
    ///@param _candidate is the address of the target validator of concern.
    ///@return a result showcasing whether the _candidate has been elected or not.
    function hasBeenElected(address _candidate) public view returns(bool) {
        return validatorsState[_candidate].hasBeenElected;
    }

    ///@notice getCandidates implements the functionality which allows the public to verify all candidates.
    ///@return a result showcasing all the candidates elected by the validators of the network.
    function getCandidates() public view returns(address[] memory) {
        return candidates;
    }
    
    ///@notice getCandidateStake implements the functionality which allows the public to observe the associated stake of a candidate.
    ///@param _candidate is the address of the target candidate of concern.
    ///@return a result showcasing the stake associated with a particular candidate.
    function getCandidateStake(address _candidate) public view returns(uint256) {
        return validatorsState[_candidate].stake;
    }

    ///@notice getCandidateValidator implements the functionality which allows the public to verify the validator of a candidate. 
    ///@param _candidate is the address of the target candidate of concern.
    ///@return a result showcasing the address of the validator of particular candidate.
    function getCandidateValidator(address _candidate) public view returns(address) {
        return validatorsState[_candidate].validator;
    }

    ///@notice getNominatorStake implements the functionality which allows the public to observe the associated stake of a nominator for a candidate.
    ///@param _candidate is the address of the target candidate of concern.
    ///@param _nominator is the address of the target nominator of concern.
    ///@return a result showcasing the stake of a particular nominator for a candidate.
    function getNominatorStake(address _candidate, address _nominator) public view returns(uint256) {
        return validatorsState[_candidate].nominators[_nominator];
    }

    ///@notice getNominators implements the functionality which allows the public to verify all nominators to a particular candidate.
    ///@param _candidate is the address of the target candidate of concern.
    ///@return a result showcasing all the nominators of a particular candidate.
    function getNominators(address _candidate) public view returns(address[] memory) {
        return nominators[_candidate];
    }

    ///@notice getWithdrawBlockNumbers implements the functionality which allows the public to observe the current withdraw state. 
    ///@return a result showcasing the withdraw state period.
    function getWithdrawBlockNumbers() public view returns(uint256[] memory) {
        return withdrawsState[msg.sender].blockNumbers;
    }

    ///@notice getWithdrawStakes implements the functionality which allows the public to observe the withdraw stakes associated with a particular period.
    ///@param _blockNumber the withdraw period of concern.
    ///@return a result showcasing the stake associated with the period.
    function getWithdrawStakes(uint256 _blockNumber) public view returns(uint256) {
        return withdrawsState[msg.sender].stakes[_blockNumber];
    }

    ///@notice invalidKYCPercent implements the functionality which allows the public to analyse the percentage of speceific invalid KYC.
    ///@param _invalidValidator is the address of the target invalid Validator of concern.
    ///@return a result showcasing the invalid KYC in percentage.
    function invalidKYCPercent(address _invalidValidator) public view onlyQualifiedCandidates(_invalidValidator) returns(uint){
        address invalidValidator = getCandidateValidator(_invalidValidator);
        return (invalidKYCsTotal[invalidValidator] * 100 / validatorsTotal);
    }
    
    ///@notice getHashedKYC implements the functionality which allows the public to observe the IPFS hash related to the KYC of specific validator.
    ///@param _validator is the address of the target associated with the related KYC lookup query.
    ///@return a result showcasing the IPFS hash related to the KYC of a specific validator.
    function getHashedKYC(address _validator) public view returns (string memory) {
        if(hasBeenElected(_validator)) {
            return hashedKYCs[getCandidateValidator(_validator)][hashedKYCs[getCandidateValidator(_validator)].length-1];
        } else {
            return hashedKYCs[_validator][hashedKYCs[_validator].length-1];
        }
    }
    
    ///@notice getHashedKYCsTotal implements the fucntionality which allows the public to observe an accumulated amount of IPFS hashes of a particular validator. 
    ///@param _validator is the address of the target associated with the related KYC lookup query.
    ///@return a result showcasing the accumulated amount of IPFS hashes of a specific validator.
    function getHashedKYCsTotal(address _validator) public view returns(uint){
        return hashedKYCs[_validator].length;
    }

    
