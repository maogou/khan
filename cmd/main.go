package main

import (
	"maogou/khan/internal"
	"os"
)

func main() {
	robot := internal.NewKhanCommand()

	if err := robot.Run(os.Args); err != nil {
		panic(err)
	}

}
