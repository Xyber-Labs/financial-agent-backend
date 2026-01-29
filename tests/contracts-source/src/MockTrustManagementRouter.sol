// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract MockTrustManagementRouter {

    struct Transaction {
        address target;
        uint256 value;
        bytes   data;
    }

    struct Permit {
        uint256 amount;
        uint256 deadline;
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    function execute(Transaction[] calldata txs, bytes calldata signature, uint256 deadline) external {}

    function deposit(
        address receiver, 
        address token, 
        uint256 amount, 
        uint256 lockDuration
    ) external returns(address walletAddress) {}

    function depositWithPermit( 
        address receiver, 
        address token, 
        uint256 lockDuration, 
        Permit calldata permit
    ) external returns(address walletAddress) {}


    function withdraw(
        address wallet,
        address token,
        address receiver,
        uint256[] calldata depositIds,
        uint256[] calldata amountsWithYield,
        bytes calldata signature,
        uint256 deadline
    ) external {}

}