package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they are %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'car' and was %s\n", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they are %d\n", car.Seats)
	}
}
