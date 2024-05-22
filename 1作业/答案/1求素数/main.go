package main

import (
	"fmt"
)

func main() {
	type Config struct {
		RealTimeEvents []string `json:"real_time_events,omitempty"`
	}

	config := Config{RealTimeEvents: []string{"test1", "test2", "test3"}}

	list := []string{"test1", "test2", "test4"}

	//for _, event := range list {
	//	if i, _ := Find(config.RealTimeEvents, event); i != -1 {
	//		fmt.Println("Value found in slice")
	//	}
	//}

	index := DeleteSliceIndex(list, 2)
	fmt.Printf("%v", index)

	//config.RealTimeEvents = index
	fmt.Printf("%v", config.RealTimeEvents)

}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func DeleteSliceIndex(slice []string, index int) []string {
	le := len(slice)
	slice[index] = slice[le-1]
	return slice[:le-1]
}
