package ipfsConn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Connects to ipfs, send json payload, and return cid(hash)
// Makes a call to IPFS
// Sends CID, Returns ipfsBody
func IpfsCall(cid string) (ipfsBody map[string]interface{}) {
	// Where local node is running - Connect to the IPFS API
	fmt.Println(cid)
	ipfsGateway := "http://127.0.0.1:8080/ipfs/" + cid

	// Make an HTTP GET request to the IPFS gateway
	resp, err := http.Get(ipfsGateway)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed with status code:", resp.Status)
		return
	}

	// Read the JSON data from the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Now, 'data' contains the JSON content retrieved from IPFS
	//fmt.Println("JSON Data:", data)

	/*	IPFSConn := configs.GoDotEnvVariable("IPFSConnection")
		shell := ipfs.NewShell(IPFSConn)
		fmt.Println("Ipfs Call")
		data, err := shell.Cat(cid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ipfsCall: ", data)

		// ...so we convert it to a string by passing it through
		// a buffer first. A 'costly' but useful process.
		// https://golangcode.com/convert-io-readcloser-to-a-string/
		buf := new(bytes.Buffer)
		buf.ReadFrom(data)
		newStr := buf.String()

		///**** Need to figure out how to change it from string to interface
		//***** Maybe something wrong with the json
		fmt.Println("newStr: ", newStr)*/
	ipfsBody = map[string]interface{}{}
	json.Unmarshal([]byte(data), &ipfsBody)
	fmt.Println("Data: ", ipfsBody)
	return ipfsBody
}
