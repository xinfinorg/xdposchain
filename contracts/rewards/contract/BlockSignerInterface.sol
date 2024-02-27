// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

interface BlockSignerInterface {
    function getSigners(bytes32 _blockHash) external view returns(address[] memory);
}