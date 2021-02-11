package mirror_test

import (
	"testing"
	"time"

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

	var IgnoreUnsupported struct {
		A int
		B string

		C time.Time `mirror:"ignore"`
		d time.Time
	}
	get.Reflect(&IgnoreUnsupported)

	should.Be("A")(get.Field(IgnoreUnsupported.A).Name).Test(t)
	should.Be("B")(get.Field(IgnoreUnsupported.B).Name).Test(t)

	var PrimitiveTypes struct {
		A bool
		B int
		C int8
		D int16
		E int32
		F int64
		G uint
		H uint8
		I uint16
		J uint32
		K uint64
		L float32
		M float64
		N complex64
		O complex128
		P string
	}
	get.Reflect(&PrimitiveTypes)

	should.Be("A")(get.Field(PrimitiveTypes.A).Name).Test(t)
	should.Be("B")(get.Field(PrimitiveTypes.B).Name).Test(t)
	should.Be("C")(get.Field(PrimitiveTypes.C).Name).Test(t)
	should.Be("D")(get.Field(PrimitiveTypes.D).Name).Test(t)
	should.Be("E")(get.Field(PrimitiveTypes.E).Name).Test(t)
	should.Be("F")(get.Field(PrimitiveTypes.F).Name).Test(t)
	should.Be("G")(get.Field(PrimitiveTypes.G).Name).Test(t)
	should.Be("H")(get.Field(PrimitiveTypes.H).Name).Test(t)
	should.Be("I")(get.Field(PrimitiveTypes.I).Name).Test(t)
	should.Be("J")(get.Field(PrimitiveTypes.J).Name).Test(t)
	should.Be("K")(get.Field(PrimitiveTypes.K).Name).Test(t)
	should.Be("L")(get.Field(PrimitiveTypes.L).Name).Test(t)
	should.Be("M")(get.Field(PrimitiveTypes.M).Name).Test(t)
	should.Be("N")(get.Field(PrimitiveTypes.N).Name).Test(t)
	should.Be("O")(get.Field(PrimitiveTypes.O).Name).Test(t)
	should.Be("P")(get.Field(PrimitiveTypes.P).Name).Test(t)

}
