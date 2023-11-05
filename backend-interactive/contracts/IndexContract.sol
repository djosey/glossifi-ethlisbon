// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract IndexContract {
    string public ipfsCID;

    // Constructor to initialize the IPFS CID
    constructor() {
        ipfsCID = "QmU7auJCkqwEet6TtTxD97xn5q4JtZMjpJkxtExS1GjyUK";
    }

    // Function to get the IPFS CID
    function getIPFSCID() public view returns (string memory) {
        return ipfsCID;
    }
}