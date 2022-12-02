package interfaceecample

import (
	"fmt"
	"testing"
)

type PetrolEngine struct{}

func (PetrolEngine) Refill() {
	fmt.Println("Petrol ENGINE refilled")
}

type ElectricEngine struct{}

func (ElectricEngine) Refill() {
	fmt.Println("Electric ENGINE refilled")
}

type CarBody struct{}

func (CarBody) Load() {
	fmt.Println("Car BODY loaded")
}

type TruckBody struct{}

func (TruckBody) Load() {
	fmt.Println("Truck BODY loaded")
}

// -----------------------------------------------------------------------

type Engine interface {
	Refill()
}

type Body interface {
	Load()
}

type Vehicle struct {
	Engine
	Body
}

// NewVehicle is a Vehicle constructor.
func NewVehicle(engine Engine, body Body) *Vehicle {
	return &Vehicle{
		Engine: engine,
		Body:   body,
	}
}

// -----------------------------------------------------------------------

func RefillVehicleEngine(engine Engine) {
	engine.Refill()
}

func LoadVehicleBody(body Body) {
	body.Load()
}

func TestBuildFlexibleVehicle(t *testing.T) {
	engine := PetrolEngine{} // or ElectricEngine{}
	body := TruckBody{}      // or CarBody{}

	vehicle := NewVehicle(engine, body)

	vehicle.Load()
	vehicle.Refill()

	// OR later on:
	LoadVehicleBody(vehicle)
	RefillVehicleEngine(vehicle)
}
