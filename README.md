# P2EPRO Chaincode - Shipment Tracking

This is a smart contract code written in Go for tracking shipments in a supply chain network using Hyperledger Fabric.

## Functionality
The smart contract provides the following functionalities:

1. Initiate Shipment: Allows the manufacturer to initiate a new shipment by providing a unique shipment ID. The shipment is initialized with the status "Shipped from Manufacturer".

2. Update Shipment Status: Allows authorized participants to update the status of a shipment at various points, such as when it is received by the wholesaler, dealer, or retailer.

## Prerequisites
- Hyperledger Fabric network with peers, channels, and orderers properly configured.
- Go programming language installed on the development machine.
## Installation
1. Clone this repository or download the code.
2. Install the necessary dependencies:

   ```shell
   go get github.com/hyperledger/fabric/core/chaincode/shim
   go get github.com/hyperledger/fabric/protos/peer
   ```

3. Build the chaincode:

   ```shell
   go build
   ```

## Usage

1. Start your Hyperledger Fabric network.
2. Deploy the chaincode to your network.
3. Invoke the chaincode functions using the appropriate arguments to initiate a shipment or update the shipment status.
## Chaincode Functions
### InitiateShipment
This function is used to initiate a new shipment.
Arguments:
- shipmentID (string): The unique ID of the shipment.
### UpdateShipmentStatus
This function is used to update the status of a shipment.
Arguments:
- shipmentID (string): The ID of the shipment to update.
- status (string): The new status of the shipment.
## Contributing
## License
This project is licensed under the [MIT License](LICENSE).
