package main

import (
	//"encoding/json"

	"context"
	"encoding/json"
	configs "ethlisbon-2023/configs"
	index "ethlisbon-2023/index"
	ipfsConn "ethlisbon-2023/ipfs"
	"ethlisbon-2023/parsing"
	client "ethlisbon-2023/web3"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/audit-request/{address}", getAudit).Methods("GET")
	http.ListenAndServe(":12345", router)
}

func getAudit(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	contractAddress, _ := (params["address"])
	auditCid, frontendLink, indexCid, ContDeetsCid := recursive(contractAddress)
	res := map[string]interface{}{"AuditCID": auditCid, "FrontEndCID": frontendLink, "IndexCID": indexCid, "ContractMetadataCID": ContDeetsCid}
	//_ = json.Unmarshal(audit, &res)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(res); err != nil {
		panic(err)
	}

}

func recursive(contractAddress string) (audit string, frontendLink string, indexCid string, contDeetsCid string) {
	indexCid = web3Call() // Smart contract address sent to web3, returns indexCid
	fmt.Println("This is: ", indexCid)
	index := ipfsConn.IpfsCall(indexCid)
	indexCidTarget := contractAddress
	contDeetsCid = parsing.JsonParsing(indexCidTarget, index) // Finds cid of audit
	//frontendTarget := "glossifi-frontend"
	//frontendLink = parsing.JsonParsing(frontendTarget, index)
	contractDetails := ipfsConn.IpfsCall(contDeetsCid) // CID send to ipfs, pulls down contractDetails
	//res := map[string]interface{}{"data": contractDetails}
	cidTarget := "glossifi-audit"
	auditCid := parsing.JsonParsing(cidTarget, contractDetails) // Finds cid of audit
	frontendTarget := "glossifi-frontend"
	frontendLink = parsing.JsonParsing(frontendTarget, contractDetails) // Finds front-end link, how to display on front-end**
	//audit = ipfsConn.IpfsCall(auditCid)
	// Cid of audit is sent to ipfs, pulls down audit
	/*audit, err := os.ReadFile("audit.txt")
	if err != nil {
		log.Fatal(err)
	}*/

	//frontendLink = "k51qzi5uqu5di97q258amcztskk0tj05s6t5aa2h56o9vnrvo9rp59p6lb0zs9"
	return auditCid, frontendLink, indexCid, contDeetsCid
}

// Makes a call to smart contract index to find CID based on address
// Sends address, Returns CID
// Need to have a parser or send it thru the json parser
func web3Call() (ipfsCID string) {

	CONTRACT_ADDRESS := configs.GoDotEnvVariable("SMARTCONTRACTADDRESS")
	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	SENDER_ADDRESS := configs.GoDotEnvVariable("SENDERADDRESS")
	senderAddress := common.HexToAddress(SENDER_ADDRESS)
	client := client.Client()
	// Create a new instance of IndexCaller (replace with your contract and bind package)
	indexCaller, err := index.NewIndexCaller(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	ctx := context.Background()
	// Create CallOpts with your Ethereum client and other options
	callOpts := &bind.CallOpts{
		Context: ctx,
		//GasLimit: gasLimit,
		From: senderAddress,
		//Value: value,
		//BlockNumber: big.NewInt(5),
	}

	// Call the GetIPFSCID function to retrieve the IPFS CID
	ipfsCID, err = indexCaller.GetIPFSCID(callOpts)
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}

	// Print the retrieved IPFS CID
	fmt.Println("IPFS CID:", ipfsCID)
	return ipfsCID
}

// Parse thru contractDetails.json to find Audit CID and Vue App link
// Makes second call to ipfs(maybe have as one ipfs call function)
// Sends audit CID, returns Audit
// Audit is displayed on front end with link

// Audit ipfs cid: Qmbcgj1V5gTyknz8HfZvS6dSE9BabEAkjdVmGdNuotoogD
// curl "http://127.0.0.1:8080/ipfs/Qmbcgj1V5gTyknz8HfZvS6dSE9BabEAkjdVmGdNuotoogD"
// COntract details cid: QmcJ8rGnXwnLkMK5r3L5DsWoahuArEtQhGaiYWsAp4FPMa
// index cid: QmU7auJCkqwEet6TtTxD97xn5q4JtZMjpJkxtExS1GjyUK
