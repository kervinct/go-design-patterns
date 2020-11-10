package behavioral

import "fmt"

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	memento Command
}

type Originator struct {
	Command Command
}

func (o *Originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *Originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type CareTaker struct {
	mementoList []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) Pop() Memento {
	if len(c.mementoList) > 0 {
		tempMemento := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[0 : len(c.mementoList)-1]
		return tempMemento
	}

	return Memento{}
}

func (c *CareTaker) Memento(i int) (Memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return Memento{}, fmt.Errorf("Index not found")
	}
	return c.mementoList[i], nil
}

type MementoFacade struct {
	originator Originator
	careTaker  CareTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings(i int) Command {
	memento, _ := m.careTaker.Memento(i)
	m.originator.ExtractAndStoreCommand(memento)
	return m.originator.Command
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
func main() {
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}
