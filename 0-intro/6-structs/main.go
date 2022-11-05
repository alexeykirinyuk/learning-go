package main

import "fmt"

type Coords struct {
	Lat float64
	Lon float64
}

type PhoneNumber string

type Courier struct {
	FirstName     string
	LastName      string
	PhoneNumber   PhoneNumber
	CurrentCoords Coords
}

func (c *Courier) GoToSeller() {
	fmt.Println("Пошел к продавцу...")
}

func (c *Courier) GetFullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func main() {
	c := Courier{
		FirstName:   "Алексей",
		LastName:    "Алексеев",
		PhoneNumber: "+79998887766",
		CurrentCoords: Coords{
			Lat: 10.11,
			Lon: 11.11,
		},
	}
	fmt.Printf("%v\n", c)

	c.GoToSeller()

	r := &c

	changeCourier(*r)
	fmt.Printf("changeCourier(c): %v\n", c)

	changeCourierLinked(r)
	fmt.Printf("changeCourierLinked(&c): %v\n", c)
}

func changeCourier(c Courier) {
	c.LastName = "Антонов"
}

func changeCourierLinked(c *Courier) {
	c.LastName = "Антонов"
}
