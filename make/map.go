package main

import "fmt"

func main() {

	//declare and initalise map
	m := map[string]int{
		"Udith":  32,
		"Iranga": 27,
	}
	fmt.Println(m)
	fmt.Println(m["Iranga"])
	fmt.Println(m["Udith"])

	v, ok := m["Ir"]
	fmt.Println(v, ok)

	// check if key is found in the map or not
	if v, ok := m["Ir"]; !ok {
		fmt.Println("no value found in map", v)
	}
	if v, ok := m["Udith"]; ok {
		fmt.Println("value found in map", v)
	}

	// add new key/value to map
	m["Akein"] = 10
	fmt.Println(m)

	//loop through a map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete a record from the map
	delete(m, "Iranga")
	fmt.Println(m)

}
