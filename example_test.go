package keyvalue_test

import (
	"fmt"

	"github.com/womat/keyvalue"
)

func Example() {
	r := keyvalue.NewRecord()
	r.Set("port", 8443)
	r.Set("debug", "yes")
	r.Set("factor", "1.5")

	fmt.Println(r.Int("port"))
	fmt.Println(r.Bool("debug"))
	fmt.Println(r.Float64("factor"))
	fmt.Printf("%q\n", r.String("host")) // missing: no error, no panic
	fmt.Println(r.Exists("host"))

	for _, k := range r.GetSortedKeys() {
		v, _ := r.Value(k)
		fmt.Printf("%s=%v (%T)\n", k, v, v)
	}

	// Output:
	// 8443
	// true
	// 1.5
	// ""
	// false
	// debug=yes (string)
	// factor=1.5 (string)
	// port=8443 (int)
}

// A string is parsed strictly per target type, so the same text reads
// differently depending on the accessor.
func ExampleRecord_Int() {
	r := keyvalue.Record{"text": "3.9", "number": 3.9}

	fmt.Println(r.Int("text"))   // strconv.Atoi rejects "3.9"
	fmt.Println(r.Int("number")) // a float is truncated
	fmt.Println(r.Float64("text"))

	// Output:
	// 0
	// 3
	// 3.9
}

func ExampleRecord_Bool() {
	r := keyvalue.Record{"on": "ON", "two": 2, "zero": "0", "word": "maybe"}

	fmt.Println(r.Bool("on"), r.Bool("two"), r.Bool("zero"), r.Bool("word"))

	// Output:
	// true true false false
}

// Value tells a missing key from one that holds a zero value.
func ExampleRecord_Value() {
	r := keyvalue.Record{"count": 0}

	for _, k := range []string{"count", "missing"} {
		v, ok := r.Value(k)
		fmt.Printf("%s: %v %v\n", k, v, ok)
	}

	// Output:
	// count: 0 true
	// missing: <nil> false
}
