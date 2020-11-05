package structual

import (
	"errors"
	"fmt"
	"io"
)

// PrinterAPI 接口
type PrinterAPI interface {
	PrintMessage(string) error
}

// PrinterImpl1 结构
type PrinterImpl1 struct{}

// PrintMessage 实现接口
func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

// PrinterImpl2 另一个结构
type PrinterImpl2 struct {
	Writer io.Writer
}

// PrintMessage 实现接口
func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

//-----------------------------------------------------

// PrinterAbstraction 抽象接口
type PrinterAbstraction interface {
	Print() error
}

// NormalPrinter 组合
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print 接口实现
func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

// PacketPrinter 组合
type PacketPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print 接口实现
func (c *PacketPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packet: %s", c.Msg))
	return nil
}

// ----------------------------------------------------
