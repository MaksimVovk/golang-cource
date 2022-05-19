package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMails struct {
	amount int
}

func (s SpamMails) String() string {
	return "Spam mail"
}

func (s SpamMails) Ship() Shipping {
	return Air
}

func (s SpamMails) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMails{40000}
	automate(&mail)

	// waste := ToxicWaste{300}
	// automate(&waste)
}
