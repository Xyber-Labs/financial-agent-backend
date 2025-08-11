// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract MockTEEWallet {

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

    function initSessionKey(
        ChunkedX509Cert calldata leaf,
        ChunkedX509Cert calldata intermediate,
        ChunkedSGXQuote calldata quote,
        address sessionKey
    ) external {}
}
