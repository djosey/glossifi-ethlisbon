1. User inputs smart contract address
2. Address gets sent thru endpoint to back-end
3. Back-end makes web3 call to smart contract
4. Brings the index payload forward to parse based on smart contract address and find ipfs CID
of ContractDetails.json 
 5. parse ContractDetails.json leads to audit.txt with vue website
7. Audit.txt is outputted to the user on front-end(Get Request)

Recursive