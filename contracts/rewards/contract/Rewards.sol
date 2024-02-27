// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

import "./libs/Initializable.sol";
import "./TreasuryInterface.sol";
import "./BlockSignerInterface.sol";

contract Rewards is Initializable {
    address public owner;

    address public treasuryAddress;
    address public blockSignerAddress;
    uint256 public rewardTransferEpoch;

    mapping(uint256 => bool) public epochsWithRewardsCalculated;
    mapping(uint256 => mapping(address => bool)) public epochsWithRewardsCalculatedForNode;

    mapping(address => bool) public whitelist;
    mapping(address => uint256) public standbyNodeHistory;
    mapping(address => uint256) public pendingRewardsTransaction;
    mapping(address => SlashedNode) public slashedNodes;

    address[] public pendingRewardTransferAddresses;

    mapping(uint256 => uint256) public standbyNodeBlocksConfirmedHistory;
    mapping(uint256 => uint256) public masterNodeBlocksConfirmedHistory;
    mapping(address => mapping(uint256 => uint256)) public blockHistory;

    uint256 public currentEpoch;

    struct SlashedNode {
        bool isSlashed;
        uint256 untilEpoch;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function");
        _;
    }

    modifier onlyWhitelisted() {
        require(whitelist[msg.sender], "Not whitelisted");
        _;
    }

    modifier treasuryAddressSet() {
        require(treasuryAddress != address(0), "Treasury not set" );
        _;
    }

    function initialize(
        address _blockSignerAddress,
        uint256 _rewardTransferEpoch,
        uint256 _initialEpoch
    ) external initializer {
        blockSignerAddress = _blockSignerAddress;
        rewardTransferEpoch = _rewardTransferEpoch;
        currentEpoch = _initialEpoch;

        owner = msg.sender;
        whitelist[msg.sender] = true; // Add owner to whitelist
    }

    function setTreasuryAddress(address _treasuryAddress) external {
        treasuryAddress = _treasuryAddress;
    }

    function setCurrentEpochByOwner(uint256 epoch) external onlyOwner {
        currentEpoch = epoch;
    }

    // Function to increment the current epoch (accessible only internally)
    function incrementCurrentEpoch() internal {
        currentEpoch++;
    }

    function addWhitelisted(address _address) external onlyOwner {
        whitelist[_address] = true;
    }

    function removeWhitelisted(address _address) external onlyOwner {
        whitelist[_address] = false;
    }

    function markEpochWithRewardsCalculated(uint256 epoch) internal {
        epochsWithRewardsCalculated[epoch] = true;
    }

    function requireRewardsNotCalculated(uint256 epoch) internal view {
        require(!epochsWithRewardsCalculated[epoch], "Rewards already calculated for this epoch");
    }

    function isRewardTransferEpoch(uint256 epoch) internal view returns (bool) {
        return epoch % rewardTransferEpoch == 0;
    }

    function setRewardTransferEpoch(uint256 _epoch) external onlyOwner {
        rewardTransferEpoch = _epoch;
    }

    function requireRewardsNotCalculatedForNode(uint256 epoch, address node) internal view {
        require(!epochsWithRewardsCalculatedForNode[epoch][node], "Rewards already calculated for this node in the current epoch");
    }

    function processStandbyNode(
        uint256 chainReward,
        address node,
        uint256 epoch
    ) internal {
        requireRewardsNotCalculatedForNode(epoch, node);

        uint256 verifiedBlocks = getVerifiedBlocks(node, epoch);

        if (shouldSlash(verifiedBlocks)) {
            slashNode(node, epoch+4);
        }

        bool slashed = isSlashed(node, epoch);

        if (!slashed) {
            uint256 rewards = calculateReward(chainReward, verifiedBlocks, slashed);
            setPendingRewards(node, rewards);

        }
        
        standbyNodeHistory[node] += verifiedBlocks;
        epochsWithRewardsCalculatedForNode[epoch][node] = true;
    }

    function countConfirmedBlocks(bytes32[] memory blockHashes, address[] memory standbyNodes) internal {
        uint256 masterNodeBlocksConfirmed = 0;
        uint256 standbyNodeBlocksConfirmed = 0;

        for (uint256 i = 0; i < blockHashes.length; i++) {
            bytes32 blockHash = blockHashes[i];

            // Get the signers for the current block hash
            address[] memory signers = BlockSignerInterface(blockSignerAddress).getSigners(blockHash);

            // Iterate over each signer
            for (uint256 j = 0; j < signers.length; j++) {
                address signer = signers[j];
                
                // Check if the signer is in the standbyNodes list
                bool isStandbyNode = false;
                for (uint256 k = 0; k < standbyNodes.length; k++) {
                    if (standbyNodes[k] == signer) {
                        isStandbyNode = true;
                        break;
                    }
                }

                if (!isStandbyNode) {
                    // Increment the number of blocks confirmed by master nodes
                    masterNodeBlocksConfirmed++;
                } else {
                    // Increment the number of blocks confirmed by standby nodes
                    standbyNodeBlocksConfirmed++;
                    blockHistory[signer][currentEpoch]++;
                }
            }
        }

        // Store the counts in the history mappings
        standbyNodeBlocksConfirmedHistory[currentEpoch] = standbyNodeBlocksConfirmed;
        masterNodeBlocksConfirmedHistory[currentEpoch] = masterNodeBlocksConfirmed;
    }


    function calculateRewards(uint256 chainReward, bytes32[] memory blockHashes, address[] memory standbyNodes, uint256 epoch) external onlyWhitelisted treasuryAddressSet {
        require(epoch == currentEpoch + 1, "Invalid epoch");
        incrementCurrentEpoch();

        requireRewardsNotCalculated(currentEpoch);

        countConfirmedBlocks(blockHashes, standbyNodes);

        for (uint256 i = 0; i < standbyNodes.length; i++) {
            address node = standbyNodes[i];
            processStandbyNode(chainReward, node, currentEpoch);
        }

        
        // Transfer rewards if it's a reward transfer epoch
        if (isRewardTransferEpoch(currentEpoch)) {
            transferRewards();
        }

        // Mark the epoch as calculated
        markEpochWithRewardsCalculated(currentEpoch);
    }


    function setPendingRewards(address node, uint256 rewards) internal {
        pendingRewardsTransaction[node] = pendingRewardsTransaction[node] + rewards;
        // Add the address to the list of pending reward transfer addresses if it's not already present
        if (!isAddressPendingRewardTransfer(node)) {
            pendingRewardTransferAddresses.push(node);
        }
    }

    function resetPendingRewards(address node) internal {
        delete pendingRewardsTransaction[node];
        // Remove the address from the list of pending reward transfer addresses
        removePendingRewardTransferAddress(node);
    }

    function removePendingRewardTransferAddress(address node) internal {
        for (uint256 i = 0; i < pendingRewardTransferAddresses.length; i++) {
            if (pendingRewardTransferAddresses[i] == node) {
                pendingRewardTransferAddresses[i] = pendingRewardTransferAddresses[pendingRewardTransferAddresses.length - 1];
                pendingRewardTransferAddresses.pop();
                break;
            }
        }
    }

    // Function to check if an address is in the list of pending reward transfer addresses
    function isAddressPendingRewardTransfer(address node) internal view returns (bool) {
        for (uint256 i = 0; i < pendingRewardTransferAddresses.length; i++) {
            if (pendingRewardTransferAddresses[i] == node) {
                return true;
            }
        }
        return false;
    }

    function transferRewards() internal {
        // Create an instance of the Treasury contract
        TreasuryInterface treasury = TreasuryInterface(treasuryAddress);

        // Iterate over pending reward transfer addresses and transfer rewards
        for (uint256 i = 0; i < pendingRewardTransferAddresses.length; i++) {
            address node = pendingRewardTransferAddresses[i];
            uint256 rewards = pendingRewardsTransaction[node];
            if (rewards > 0) {
                // Create a transaction in the Treasury contract to transfer rewards to the node
                uint256 transactionId = treasury.createTransaction(node, rewards);

                // Set isConfirmed for the transactionId and msg.sender to true
                treasury.confirmTransaction(transactionId);
            }
        }

        for (uint256 i = 0; i < pendingRewardTransferAddresses.length; i++) {
            // Clear pending rewards after creating the transaction
            address node = pendingRewardTransferAddresses[i];
            resetPendingRewards(node);
        }
    }

    function getVerifiedBlocks(address node, uint256 epoch) internal view returns (uint256) {
        // Retrieve the number of blocks verified by a node in a specific epoch
        return blockHistory[node][epoch];
    }

    function isSlashed(address node, uint256 epoch) internal view returns (bool) {
        return slashedNodes[node].isSlashed && slashedNodes[node].untilEpoch >= epoch;
    }

    function slashNode(address node, uint256 untilEpoch) internal {
        slashedNodes[node] = SlashedNode(true, untilEpoch);
    }

    function shouldSlash(uint256 verifiedBlocks) internal pure virtual returns (bool) {
        return verifiedBlocks == 0;
    }

    function calculateReward(uint256 chainReward, uint256 verifiedBlocks, bool slashed) internal view virtual returns (uint256) {
        if (slashed) {
            return 0;
        }

        uint256 totalSigner = standbyNodeBlocksConfirmedHistory[currentEpoch];
        if (totalSigner > 0) {
            uint256 calcReward = chainReward / totalSigner * verifiedBlocks;
            return calcReward;
        }

        return 0;
    }
}
