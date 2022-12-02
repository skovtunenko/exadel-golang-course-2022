package interfaceecample_test

import (
	"fmt"
	"testing"
)

type PetrolEngine struct{}

func (PetrolEngine) Refill() {
	fmt.Println("Petrol ENGINE refilled")
}

type CarBody struct{}

func (CarBody) Load() {
	fmt.Println("Car BODY loaded")
}

type Vehicle struct {
	PetrolEngine
	CarBody
}

// NewVehicle is a Vehicle constructor.
func NewVehicle(petrolEngine PetrolEngine, carBody CarBody) *Vehicle {
	return &Vehicle{PetrolEngine: petrolEngine, CarBody: carBody}
}

// -----------------------------------------------------------------------

func RefillVehicleEngine(vehicle *Vehicle) {
	vehicle.Refill()
}

func LoadVehicleBody(vehicle *Vehicle) {
	vehicle.Load()
}

func TestBuildRigidVehicle(t *testing.T) {
	engine := PetrolEngine{}
	body := CarBody{}

	vehicle := NewVehicle(engine, body)

	vehicle.Load()
	vehicle.Refill()

	// OR later on:
	LoadVehicleBody(vehicle)
	RefillVehicleEngine(vehicle)
}
