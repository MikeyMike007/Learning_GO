/*
Mutating Maps

Insert or update an element in map m:

m[key] = elem

Retrieve an element:

elem = m[key]

Delete an element:

delete(m, key)

Test that a key is present with a two-value assignment:

elem, ok = m[key]

If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]
*/

package main

import "fmt"

func main() {
		m := make(map[string]int)

		m["answer"] = 42
		fmt.Println("The value: ", m["answer"])

		m["answer"] = 48
		fmt.Println("The value: ", m["answer"])
		
		delete(m, "answer")
		fmt.Println("The value: ", m["answer"])

		v, ok := m["answer"]
		fmt.Println("The value: ", v, "Present? ", ok)
}
