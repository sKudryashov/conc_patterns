package race_conditions

import "fmt"

func GenerateRaceCondition() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	fmt.Print("Race condition generator")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
