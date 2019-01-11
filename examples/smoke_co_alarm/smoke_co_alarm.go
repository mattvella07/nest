package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattvella07/nest"
)

func main() {
	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	alarms, err := n.GetSmokeCOAlarms()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(alarms), " Smoke/CO Alarm(s): ", alarms)
}
