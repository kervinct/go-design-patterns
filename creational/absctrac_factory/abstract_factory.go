package creational

import "fmt"

const (
	// CarFactoryType 汽车
	CarFactoryType = 1
	// MotorbikeFactoryType 摩托车
	MotorbikeFactoryType = 2
)

// GetVehicleFactory 获取工厂
func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
