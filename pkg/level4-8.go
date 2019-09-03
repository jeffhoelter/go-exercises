package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken not stirred", "Martinis", "Cars"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice Cream", "Sunsets"},
	}

	m["evil_dr"] = []string{"pinkies", "millions", "sharks with frickin laser beams"}

	for key, mapItem := range m {
		fmt.Println("Map Key:", key)
		for index, value := range mapItem {
			fmt.Println("	Index:", index, "Value:", value)
		}
	}

	delete(m, "no_dr")
	fmt.Println("After deletion:")

	for key, mapItem := range m {
		fmt.Println("Map Key:", key)
		for index, value := range mapItem {
			fmt.Println("	Index:", index, "Value:", value)
		}
	}

}
