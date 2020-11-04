package creational

// BuildProcess builder
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector container
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct run process
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder set builder
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct abstract object
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder concrete builder
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels interface implemented
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats interface implemented
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure interface implemented
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle interface implemented
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// BikeBuilder concrete builder
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels interface implemented
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats interface implemented
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure interface implemented
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "MotorBike"
	return b
}

// GetVehicle interface implemented
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
