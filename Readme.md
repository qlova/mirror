# Mirror

Mirror is a small Go package that allows passing struct fields by value with the use of a mirror.Type

```go
var get mirror.Type
var StreetAddress struct {
    Number int
    Street string

    Postcode int
}

//Reflect the value onto the mirror.
get.Reflect(&StreetAddress)

//Inspect fields by struct-value.
get.Field(StreetAddress.Number).Name    //"Number"
get.Field(StreetAddress.Postcode).Name  //"Postcode"
get.Field(StreetAddress.Street).Name    //"Street"
```