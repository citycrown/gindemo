package main

import (
	"fmt"
	"time"
)

func main(){
	go TestMyMqConnection()

	go TestMyMqConnection()

	for {
		time.Sleep(10 * time.Second)
	}


}

func TestMyMqConnection() {
	fmt.Printf("info")
}