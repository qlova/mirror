package mirror_test

import (
	"testing"

	"qlova.org/mirror"
	"qlova.org/should"
)

func Test_Mirror(t *testing.T) {
	var get mirror.Type

	var StreetAddress struct {
		Number int
		Street string

		Postcode int
	}
	get.Reflect(&StreetAddress)

	should.Be("Number")(get.Field(StreetAddress.Number).Name).Test(t)
	should.Be("Postcode")(get.Field(StreetAddress.Postcode).Name).Test(t)
	should.Be("Street")(get.Field(StreetAddress.Street).Name).Test(t)

	var NestedStructure struct {
		A int
		B struct {
			C int
		}
	}
	get.Reflect(&NestedStructure)

	should.Be("A")(get.Field(NestedStructure.A).Name).Test(t)
	should.Be("B")(get.Field(NestedStructure.B).Name).Test(t)
	should.Be("C")(get.Field(NestedStructure.B.C).Name).Test(t)

	should.Be(".A")(get.Path(NestedStructure.A)).Test(t)
	should.Be(".B")(get.Path(NestedStructure.B)).Test(t)
	should.Be(".B.C")(get.Path(NestedStructure.B.C)).Test(t)
}
