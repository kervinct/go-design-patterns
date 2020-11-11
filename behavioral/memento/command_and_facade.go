package behavioral

import "fmt"

// Command 接口
type Command interface {
	GetValue() interface{}
}

// Volume 实现接口
type Volume byte

// GetValue 接口
func (v Volume) GetValue() interface{} {
	return v
}

// Mute 实现接口
type Mute bool

// GetValue 接口
func (m Mute) GetValue() interface{} {
	return m
}

// Memento 抽象包装结构
type Memento struct {
	memento Command
}

// Originator 备忘结构
type Originator struct {
	Command Command
}

// NewMemento 创建备忘
func (o *Originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

// ExtractAndStoreCommand 存储备忘
func (o *Originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

// CareTaker 备忘列表
type CareTaker struct {
	mementoList []Memento
}

// Add 添加备忘
func (c *CareTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

// Pop 取出备忘
func (c *CareTaker) Pop() Memento {
	if len(c.mementoList) > 0 {
		tempMemento := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[0 : len(c.mementoList)-1]
		return tempMemento
	}

	return Memento{}
}

// Memento 获取指定备忘
func (c *CareTaker) Memento(i int) (Memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return Memento{}, fmt.Errorf("Index not found")
	}
	return c.mementoList[i], nil
}

// MementoFacade 备忘外观
type MementoFacade struct {
	originator Originator
	careTaker  CareTaker
}

// SaveSettings 备忘存储
func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}

// RestoreSettings 获取备忘
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
