package main

import "fmt"

type FullNameGetter interface {
	GetFullName() string
}

type Coords struct {
	Lat float64
	Lon float64
}

type PhoneNumber string

type OrderStatus int16

const (
	New         OrderStatus = 1
	InProgress  OrderStatus = 2
	InProgress2 OrderStatus = 3
	InProgress3 OrderStatus = 4
)

type Courier struct {
	FirstName     string
	LastName      string
	PhoneNumber   PhoneNumber
	CurrentCoords Coords
}

func NewCourier(
	a int,
) *Courier {
	return &Courier{
		FirstName:     "",
		LastName:      "",
		PhoneNumber:   "",
		CurrentCoords: Coords{},
	}
}

func (c *Courier) GetFullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func main() {
	pn1 := PhoneNumber("+7999")
	_ = pn1
}
