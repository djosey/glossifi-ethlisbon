//setup kube dependency

const vueFile
const auditFile
const graphJSON
const indexDict

function vuePubIPFS (vueFile){
    //publish vue file to IPFS and return header
    var fueFileHeaderIPFS
}

function auditPubIPFS (auditFile){
    //publish audit file to IPFS and return header
    var auditFileHeaderIPFS
}

function graphPubIPFS (graphJSON){
    //publish the contract information referencing the IPFS vue header and audit vue header
    var graphJSONHeaderIPFS
}

function indexPubIPFS (indexDict){
    //publish an index referencing the contract metadata, audit, and vue app IPFS headers
    var indexDictHeaderIPFS
}

function indexPubEth (indexDictHeaderIPFS){
    var indexEthAddress
}
