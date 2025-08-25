// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract MockERC20 {

    function balanceOf(address user) external view returns (uint256) {
        return mockBalance;
    }

    uint256 mockBalance;
    function mockSetBalance(uint256 balance) external {
        mockBalance = balance;
    }
}