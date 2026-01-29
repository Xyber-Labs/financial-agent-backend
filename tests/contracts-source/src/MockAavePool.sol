// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract MockAavePool {

    function supply(
        address asset,
        uint256 amount,
        address onBehalfOf,
        uint16 referralCode
    ) external {}

    function getReserveAToken(address asset) external view returns (address) {
        return mockAToken;
    }

    /*
     * Below are function from the original TrustManagementRouter contract
     */

    address mockAToken;
    function mockSetReserveAToken(address aToken) external {
        mockAToken = aToken;
    }
}