package main

import (
	"fmt"
	"os"

	"go.arsenm.dev/logger"
	"go.arsenm.dev/logger/log"
)

func init() {
	log.Logger = logger.NewPretty(os.Stdout)
}

func main() {
	fmt.Println("start")

	

}
