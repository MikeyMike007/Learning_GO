/*
Map literals

Map literals are like struct literals, but the keys are required.
*/

package main

import "fmt"

type Vertex struct {
		lat, long float64
}

func main() {
		var m = map[string]Vertex{
				"Bell Labs": Vertex{40, -74},
				"Google": Vertex{37, -122},
		}

		fmt.Println(m)
}
