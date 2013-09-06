package main 

import (
	"fmt"
	"time"
	//"runtime"
)

const (
	RECORDS = 1000000 	// 1 Million	
)

func main() {

	// init	
	fmt.Printf("Initializing...\t")
	m := map[int]map[string][]string{}
	fmt.Println("Done.")	

	for j := 0; j < 5; j++ {

		// fill
		fmt.Printf("Filling...\t")
		for i := 0; i < RECORDS; i++ {
			m[i]=map[string][]string{
				fmt.Sprintf("key_%v", i):[]string{
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
						"jfkdajfkldafjasklsfjkljasdfklsdjafkjsadkfasjdfkljdasklfjakdsjfkdasjfkljasdjfs",
					},
			}
		}
		fmt.Printf("Done. %v records\n\n", len(m))

		// delete	
		fmt.Printf("Deleting all...\t")
		for k1, m2 := range m {
			for k2, _ := range m2 {
				delete(m2, k2)
			}
			delete(m, k1)
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
