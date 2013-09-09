package main

import (
	"github.com/scottcagno/memdb"
	"fmt"
)

func main() {	

	// controled data set
	u1 := map[string][]string{
		"id"	: []string{"1"},
		"name"	: []string{"scott", "cagno"},
		"email"	: []string{"sac@example.com"},
		"cmpny"	: []string{"nei"},
	}
	u2 := map[string][]string{
		"id"	: []string{"2"},
		"name"	: []string{"kayla", "cagno"},
		"email"	: []string{"ksc@example.com"},
		"cmpny"	: []string{"nei"},
	}
	u3 := map[string][]string{
		"id"	: []string{"3"},
		"name"	: []string{"greg", "pechiro"},
		"email"	: []string{"gfp@example.com"},
		"cmpny"	: []string{"nei"},
	}

	// get new set instance
	d := memdb.InitDB()
	s := d.GetSet("users", 0)

	fmt.Println("ADDING DATA...")
	
	// add some data
	s.Add(u1)
	s.Add(u2)
	s.Add(u3)
	
	// display current database set info
	fmt.Printf("Set Details\n---------\n")
	for k, v := range s.Records {
		fmt.Printf("ID: %v\nDATA: %v\n\n", k, v)
	}

	fmt.Println("PERFORMING DELETE QUERY... (users?name=scott&name=cagno)")
	
	// perform delete query
	s.Del(map[string][]string{"name": []string{"scott", "cagno"}})
	
	// display current database set info
	fmt.Printf("Set Details\n---------\n")
	for k, v := range s.Records {
		fmt.Printf("ID: %v\nDATA: %v\n\n", k, v)
	}

	// pause
	func(){var n int; fmt.Scanln(&n)}()
}
