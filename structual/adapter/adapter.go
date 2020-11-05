package structual

import "fmt"

// LegacyPrinter 过时接口
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter 结构
type MyLegacyPrinter struct{}

// Print 实现接口
func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return newMsg
}

// ---------------------------------------------------

// NewPrinter 新接口
type NewPrinter interface {
	PrintStored() string
}

// PrinterAdapter 适配器
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

// PrintStored 适配器接口实现
func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print(fmt.Sprintf("Adapter: %s", p.Msg))
	}
	return p.Msg
}
