# 🔗 Hyperledger Fabric Blockchain Project

An enterprise-grade permissioned blockchain system using **Hyperledger Fabric** with **Go chaincode** for managing digital assets securely in a modular architecture.

---

## 📚 Overview

Hyperledger Fabric is a modular and extensible permissioned blockchain platform designed for enterprise use cases. It features:

- ✅ Pluggable consensus mechanisms  
- 🔐 Fine-grained access control  
- 🔒 Private data via channels  
- ⚡ High-throughput transaction processing  

---

## 🧩 Core Components

| Component     | Description                                                                 |
|---------------|-----------------------------------------------------------------------------|
| **Peer**      | Maintains the ledger and executes chaincode (smart contracts)              |
| **Orderer**   | Orders endorsed transactions and packages them into blocks                  |
| **Chaincode** | Encapsulates business logic (written in Go, JavaScript, or Java)            |
| **Channel**   | Enables private communication between a group of participants               |
| **LevelDB / CouchDB** | State databases used by peers to store world state                |
| **MSP**       | Membership Service Provider: manages identity and access policies           |

---

## 💡 Chaincode Functions

Chaincode represents the business logic layer. Typical functions include:

- `InitLedger()` – Pre-load data into the ledger  
- `CreateAsset()` – Add a new asset  
- `ReadAsset()` – Fetch asset by key  
- `UpdateAsset()` – Modify asset details  
- `DeleteAsset()` – Remove asset from the ledger  

---

## 🔄 Transaction Lifecycle

1. 🧾 **Client application** submits a transaction proposal  
2. 🧪 **Endorsing Peers** simulate and sign the transaction  
3. 📬 **Client** collects endorsements and sends to **Orderer**  
4. 📦 **Orderer** assembles transactions into blocks  
5. 📚 **Committing Peers** validate and append blocks to the ledger  

---

## 🗃 Ledger Structure

### 1. Blockchain (Immutable)

- Chain of blocks  
- Each block contains:
  - List of transactions  
  - Hash of the previous block  

### 2. World State (Mutable)

- Key-value store  
- Reflects the latest state of each asset  
- Stored in **LevelDB** (default) or **CouchDB** (for rich queries)  

---

## 📦 Useful Go Libraries

```go
import (
    "encoding/json"  // JSON serialization
    "fmt"            // Formatting
    "errors"         // Error handling

    "github.com/hyperledger/fabric-contract-api-go/contractapi"  // Contract API
    "github.com/hyperledger/fabric-chaincode-go/shim"            // Chaincode shim interface
)
