// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract MockTrustManagementRouter {

    struct Transaction {
        address target;
        uint256 value;
        bytes   data;
    }

    function execute(address target, Transaction[] calldata _transcations) external {}
}