package book

import "log"

func ItemSave() chan interface{}  {
	outChan := make(chan interface{})
	go func() {
		itemCount := 0
		for  {
			item := <- outChan
			log.Printf("count: %d, item: %s\n", itemCount, item)
			itemCount++
		}
	}()
	return outChan
}
