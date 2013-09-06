package main 

import (
	"fmt"
	"time"
)

const (
	RAWDATA = 1<<12		// 4 KB (aprox. 200 fields per record)
	RECORDS = 1000000 	// 1 Million	
)

func main() {

	// init	
	fmt.Printf("Initializing...\t")
	m := make([]map[int64][RAWDATA]byte, 0)
	fmt.Println("Done.")	

	for j := 0; j < 5; j++ {

		// fill
		fmt.Printf("Filling...\t")
		for i := 0; i < RECORDS; i++ {
			m = append(m, map[int64][RAWDATA]byte{int64(i):[RAWDATA]byte{}})
		}
		fmt.Printf("Done. %v records\n\n", len(m))

		// delete	
		fmt.Printf("Deleting all...\t")
		for i := 0; i < len(m); {
			m[len(m)-1], m[i], m = nil, m[len(m)-1], m[:len(m)-1]
		}
		fmt.Printf("Done. %v records\n\n", len(m))

		// chilling
		fmt.Printf("Chilling for 3 mins...\t")
		time.Sleep(180*time.Second)
		fmt.Printf("Done.\n--------------\n")

	}	

	fmt.Printf("Total records: %v\n", len(m))

	// wait
	func(){ var n int; fmt.Scanln(&n) }()
}
