package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ShipmentContract struct{}

type Shipment struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (sc *ShipmentContract) InitiateShipment(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	shipmentID := args[0]

	existingShipmentBytes, err := stub.GetState(shipmentID)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to read from world state: %s", err.Error()))
	}

	if existingShipmentBytes != nil {
		return shim.Error(fmt.Sprintf("Shipment with ID %s already exists", shipmentID))
	}

	shipment := Shipment{
		ID:     shipmentID,
		Status: "Shipped from Manufacturer",
	}

	shipmentBytes, err := json.Marshal(shipment)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal shipment: %s", err.Error()))
	}

	err = stub.PutState(shipmentID, shipmentBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put shipment in world state: %s", err.Error()))
	}

	return shim.Success(nil)
}

func (sc *ShipmentContract) UpdateShipmentStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	shipmentID := args[0]
	status := args[1]

	shipmentBytes, err := stub.GetState(shipmentID)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to read from world state: %s", err.Error()))
	}

	if shipmentBytes == nil {
		return shim.Error(fmt.Sprintf("Shipment with ID %s does not exist", shipmentID))
	}

	shipment := new(Shipment)
	err = json.Unmarshal(shipmentBytes, shipment)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal shipment: %s", err.Error()))
	}

	shipment.Status = status

	updatedShipmentBytes, err := json.Marshal(shipment)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal updated shipment: %s", err.Error()))
	}

	err = stub.PutState(shipmentID, updatedShipmentBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put updated shipment in world state: %s", err.Error()))
	}

	return shim.Success(nil)
}

func (sc *ShipmentContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(ShipmentContract))
	if err != nil {
		fmt.Printf("Error starting shipment chaincode: %s", err.Error())
	}
}
