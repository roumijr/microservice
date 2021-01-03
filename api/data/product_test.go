package data

import "testing"


func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:	"Mike",
		Price:	1.00,
		SKU:	"abc-abc-dls",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}