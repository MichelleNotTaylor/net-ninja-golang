package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  "Michelle",
		items: map[string]float64{},
		tip:   0.00,
	}

	return b
}
