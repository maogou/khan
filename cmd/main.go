package main

import (
	"os"
	"smallBot/internal"
)

func main() {
	robot := internal.NewRobotCommand()

	if err := robot.Run(os.Args); err != nil {
		panic(err)
	}

}
