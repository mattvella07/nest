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

	cameras, err := n.GetCameras()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(cameras), " Camera(s): ", cameras)
}
