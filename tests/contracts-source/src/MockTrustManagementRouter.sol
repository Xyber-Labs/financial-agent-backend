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

    struct Deposit {
        uint256 amount;
        uint256 lockedUntil;
    }

    struct ChunkedX509Cert {
        bytes bodyPartOne;
        bytes publicKey;
        bytes bodyPartTwo;
        bytes signature;
        // expiration date
    }

    struct ChunkedSGXQuote {
        bytes header;
        bytes isvReport;
        bytes isvReportSignature;
        bytes attestationKey;
        bytes qeReport;
        bytes qeReportSignature;
        bytes qeAuthenticationData;
    }

    event Deposited(address wallet, address indexed user, address indexed token, uint256 amount, uint256 lockDuration);


    /*
     * Below are function from the original TrustManagementRouter contract
     */

    function execute(Transaction[] calldata txs, bytes calldata signature, uint256 deadline) external {}

    function initSessionKey(
        ChunkedX509Cert calldata leaf,
        ChunkedX509Cert calldata intermediate,
        ChunkedSGXQuote calldata quote,
        address sessionKey
    ) external {}

    function deposit(
        address receiver, 
        address token, 
        uint256 amount, 
        uint256 lockDuration
    ) external returns(address walletAddress) {
        emit Deposited(receiver, receiver, token, amount, lockDuration);
        return walletAddress;
    }

    function depositWithPermit( 
        address receiver, 
        address token, 
        uint256 lockDuration, 
        Permit calldata permit
    ) external returns(address walletAddress) {}


    function withdraw(
        address user,
        address token,
        address receiver,
        uint256[] calldata depositIds,
        uint256 amountWithYield
    ) external {}

    function getDeposits(address user, address token) external view returns(Deposit[] memory) {
        return mockDeposits;
    }

    function getWalletAddress(address user) public view returns(address walletAddress, bool isDeployed) {
        return (mockWalletAddress, mockIsDeployed);
    }

    function WALLET_BEACON() public view returns(address) {
        return address(this);
    }


    /*
     * Below are internal mock-functions to help replicating needed functionality from the original TrustManagementRouter contract
     */

     address mockWalletAddress;
     bool mockIsDeployed;

     function mockSetWalletAddress(address walletAddress, bool isDeployed) external {
        mockWalletAddress = walletAddress;
        mockIsDeployed = isDeployed;
     }

     Deposit[] mockDeposits;
     function mockSetDeposits(Deposit[] calldata deposits) external {
        while (mockDeposits.length > 0) {
            mockDeposits.pop();
        }
        for (uint256 i = 0; i < deposits.length; i++) {
            mockDeposits.push(deposits[i]);
        }
     }

}