package myfakeit

import (
	"github.com/brianvoe/gofakeit/v6"
)

// Create global variable to deal with global function call.
var globalFaker *gofakeit.Faker = gofakeit.New(0)

func init() {
	gofakeit.SetGlobalFaker(globalFaker)

	addSystemLookup()
	addSoftwareLookup()
}
