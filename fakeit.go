package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// Create structs with random injected data
type Foo struct {
	Str           string
	Friendname    string `fake:"{friendname}"`
	Int           int
	Pointer       *int
	Name          string         `fake:"{firstname}"`  // Any available function all lowercase
	Sentence      string         `fake:"{sentence:3}"` // Can call with parameters
	RandStr       string         `fake:"{randomstring:[hello,world]}"`
	Number        string         `fake:"{number:1,10}"`       // Comma separated for multiple values
	Regex         string         `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
	Map           map[string]int `fakesize:"2"`
	Array         []string       `fakesize:"2"`
	Bar           Bar
	Skip          *string   `fake:"skip"` // Set to "skip" to not generate data for
	Created       time.Time // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

type Bar struct {
	Name   string
	Number int
	Float  float32
}

func main() {
	fmt.Println(gofakeit.Name())

	gofakeit.AddFuncLookup("friendname", gofakeit.Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return gofakeit.RandomString([]string{"bill", "bob", "sally"}), nil
		},
	})

	// Pass your struct as a pointer
	var f Foo
	gofakeit.Struct(&f)

	fmt.Println(f.Str)              // hrukpttuezptneuvunh
	fmt.Println(f.Friendname)       // hrukpttuezptneuvunh
	fmt.Println(f.Int)              // -7825289004089916589
	fmt.Println(*f.Pointer)         // -343806609094473732
	fmt.Println(f.Name)             // fred
	fmt.Println(f.Sentence)         // Record river mind.
	fmt.Println(f.RandStr)          // world
	fmt.Println(f.Number)           // 4
	fmt.Println(f.Regex)            // cbdfc
	fmt.Println(f.Map)              // map[PxLIo:52 lxwnqhqc:846]
	fmt.Println(f.Array)            // cbdfc
	fmt.Printf("%+v", f.Bar)        // {Name:QFpZ Number:-2882647639396178786 Float:1.7636692e+37}
	fmt.Println(f.Skip)             // <nil>
	fmt.Println(f.Created.String()) // 1908-12-07 04:14:25.685339029 +0000 UTC
}
