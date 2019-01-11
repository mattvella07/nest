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

	thermostats, err := n.GetThermostats()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(thermostats), " thermostat(s): ", thermostats)
}
