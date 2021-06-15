/*
Map literals continued

If the top-level type is just a type name, you can omit it from the elements of the literal.
*/

package main

import "fmt"

type Vertex struct {
		lat, long float64
}

func main() {
		var m = map[string]Vertex{
				"Bell Labs": {40, -74},
				"Google": {37, -122},
		}

		fmt.Println(m)
}
