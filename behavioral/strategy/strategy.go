package behavioral

import (
	"errors"
	"flag"
	"log"
)

var (
	output = flag.String("output", "console", "The output to use between 'console' and 'image' file")
)

type PrintStrategy interface {
	Print() error
}

type ConsoleSquare struct{}

func (c *ConsoleSquare) Print() error {
	return errors.New("not implemented yet")
}

type ImageSquare struct {
	DestinationFilePath string
}

func (t *ImageSquare) Print() error {
	return errors.New("not implemented yet")
}

func main() {
	flag.Parse()

	var activeStrategy PrintStrategy

	switch *output {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{}
	default:
		activeStrategy = &ConsoleSquare{}
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
