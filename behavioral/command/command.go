package behavioral

import (
	"fmt"
	"time"
)

// type Command interface {
// 	Execute()
// }

// type ConsoleOutput struct {
// 	message string
// }

// func (c *ConsoleOutput) Execute() {
// 	fmt.Println(c.message)
// }

// func CreateCommand(s string) Command {
// 	fmt.Println("Creating command")

// 	return &ConsoleOutput{
// 		message: s,
// 	}
// }

// type CommandQueue struct {
// 	queue []Command
// }

// func (p *CommandQueue) AddCommand(c Command) {
// 	p.queue = append(p.queue, c)

// 	if len(p.queue) == 3 {
// 		for _, command := range p.queue {
// 			command.Execute()
// 		}

// 		p.queue = make([]Command, 0)
// 	}
// }

// func main() {
// 	queue := CommandQueue{}

// 	queue.AddCommand(CreateCommand("First message"))
// 	queue.AddCommand(CreateCommand("Second message"))
// 	queue.AddCommand(CreateCommand("Third message"))

// 	queue.AddCommand(CreateCommand("Fourth message"))
// 	queue.AddCommand(CreateCommand("Fifth message"))
// }

type Command interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}

func main() {
	var timeCommand Command
	timeCommand = &TimePassed{time.Now()}

	var helloCommand Command
	helloCommand = &HelloMessage{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
