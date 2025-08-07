package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CarContract contract for managing CRUD for Car 
type CarContract struct {
	contractapi.Contract
}

type Car struct {
	AssetType        string `json:"AssetType"`
	CarId            string `json:"CarId"`
	Color            string `json:"Color"`
	DateOfManufacture string `json:"DateOfManufacture"`
	OwnedBy          string `json:"OwnedBy"`
	Make             string `json:"Make"`
	Model            string `json:"Model"`
	Status           string `json:"Status"`
}

// CarExists returns true when asset with given ID exists in world state
func (c *CarContract) CarExists(ctx contractapi.TransactionContextInterface, carID string) (bool, error) {
	data, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateCar creates a new instance of Car
func (c *CarContract) CreateCar(ctx contractapi.TransactionContextInterface, carID string, make string, model string, color string, manufacturerName string, dateOfManufacture string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "Org1MSP" {
		return "", fmt.Errorf("user under MSPID %v cannot perform this action", clientOrgID)
	}

	exists, err := c.CarExists(ctx, carID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state: %v", err)
	}
	if exists {
		return "", fmt.Errorf("the asset %s already exists", carID)
	}

	car := Car{
		AssetType:        "car",
		CarId:            carID,
		Color:            color,
		DateOfManufacture: dateOfManufacture,
		Make:             make,
		Model:            model,
		OwnedBy:          manufacturerName,
		Status:           "In Factory",
	}

	bytes, err := json.Marshal(car)
	if err != nil {
		return "", fmt.Errorf("failed to marshal car: %v", err)
	}

	err = ctx.GetStub().PutState(carID, bytes)
	if err != nil {
		return "", fmt.Errorf("failed to put car in world state: %v", err)
	}

	return fmt.Sprintf("successfully added car %v", carID), nil
}

// ReadCar retrieves an instance of Car from the world state
func (c *CarContract) ReadCar(ctx contractapi.TransactionContextInterface, carID string) (*Car, error) {
	exists, err := c.CarExists(ctx, carID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state: %v", err)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", carID)
	}

	bytes, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return nil, fmt.Errorf("failed to get car from world state: %v", err)
	}

	car := new(Car)
	err = json.Unmarshal(bytes, car)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Car")
	}

	return car, nil
}

