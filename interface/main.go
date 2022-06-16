package main

import "fmt"

// IF ...
type IF interface {
	getName() string
}

// Human ...
type Human struct {
	firstName, lastName string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

// Plane ...
type Plane struct {
	vendor string
	model  string
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

// Car ...
type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)

	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}

	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())
}
