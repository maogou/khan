package main

import (
	"os"
	"smallBot/internal"
)

func main() {
	robot := internal.NewKhanCommand()

	if err := robot.Run(os.Args); err != nil {
		panic(err)
	}

}
