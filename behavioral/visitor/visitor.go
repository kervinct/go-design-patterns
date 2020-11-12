package behavioral

import (
	"fmt"
	"io"
	"os"
)

// MessageA 数据结构
type MessageA struct {
	Msg    string
	Output io.Writer
}

// MessageB 数据结构
type MessageB struct {
	Msg    string
	Output io.Writer
}

// Visitor 接口
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

// Visitable 数据接口
type Visitable interface {
	Accept(Visitor)
}

// Accept 数据接口实现
func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

// Accept 数据接口实现
func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

// MessageVisitor Visitor接口
type MessageVisitor struct{}

// VisitA 数据接口实现
func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}

// VisitB 数据接口实现
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

// Print 方法
func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

// Print 方法
func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

// --------------------------------

// MsgFieldVisitorPrinter 另一个Visitor
type MsgFieldVisitorPrinter struct{}

// VisitA 数据接口实现
func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}

// VisitB 数据接口实现
func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}
