package creational

import "fmt"

// Vehicle 接口
type Vehicle interface {
	GetWheels() int
	GetSeats() int
}

// Car 接口
type Car interface {
	GetDoors() int
}

// Motorbike 接口
type Motorbike interface {
	GetType() int
}

// VehicleFactory 接口
type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)
const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 4
}

func (l *LuxuryCar) GetWheels() int {
	return 4
}

func (l *LuxuryCar) GetSeats() int {
	return 5
}

type FamilyCar struct{}

func (l *FamilyCar) GetDoors() int {
	return 5
}

func (l *FamilyCar) GetWheels() int {
	return 4
}

func (l *FamilyCar) GetSeats() int {
	return 5
}

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}

func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

func (c *CruiseMotorbike) GetSeats() int {
	return 2
}

func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}
