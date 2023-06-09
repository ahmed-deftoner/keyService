package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ahmed-deftoner/keyService/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":8080")

}

func main() {
	exc := []string{"Binance", "Bitget", "OKEX"}
	p := "Btget"
	found := false
	for _, v := range exc {
		if strings.Compare(v, p) == 0 {
			found = true
			break
		}
	}
	if found == true {
		fmt.Println("found")
	}
	Run()
}
