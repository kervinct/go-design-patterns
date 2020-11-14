package behavioral

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculate 逆波兰标记计算
func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(o, " ")

	for _, operatorString := range operators {
		if isOperator(operatorString) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := getOperatorFunc(operatorString)
			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			num, err := strconv.Atoi(operatorString)
			if err != nil {
				return 0, err
			}
			stack.Push(num)
		}
	}
	return int(stack.Pop()), nil
}

const (
	// SUM 加法
	SUM = "sum"
	// SUB 减法
	SUB = "sub"
	// MUL 乘法
	MUL = "mul"
	// DIV 除法
	DIV = "div"
)

type polishNotationStack []int

// Push 压栈
func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

// Pop 弹栈
func (p *polishNotationStack) Pop() int {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return 0
}

func isOperator(o string) bool {
	switch o {
	case SUM, SUB, MUL, DIV:
		return true
	default:
		return false
	}
}

func getOperatorFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}

// ----------------------------------------

// Interpreter 解释器接口
type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubstract struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSubstract) Read() int {
	return a.Left.Read() - a.Right.Read()
}

type operationMultiply struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationMultiply) Read() int {
	return a.Left.Read() * a.Right.Read()
}

type operationDivide struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationDivide) Read() int {
	return a.Left.Read() / a.Right.Read()
}

func operationFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSubstract{
			Left:  left,
			Right: right,
		}
	case MUL:
		return &operationMultiply{
			Left:  left,
			Right: right,
		}
	case DIV:
		return &operationDivide{
			Left:  left,
			Right: right,
		}
	}

	return nil
}

// PolishNotationStack 栈
type PolishNotationStack []Interpreter

// Push 压栈
func (p *PolishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}

// Pop 弹栈
func (p *PolishNotationStack) Pop() Interpreter {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return nil
}

// InterfaceCalculate 计算
func InterfaceCalculate(s string) (int, error) {
	stack := PolishNotationStack{}
	operators := strings.Split(s, " ")

	for _, opeartorString := range operators {
		switch opeartorString {
		case SUM, SUB, MUL, DIV:
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operationFactory(opeartorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		default:
			val, err := strconv.Atoi(opeartorString)
			if err != nil {
				return 0, fmt.Errorf("Error number: %v", val)
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}
	return int(stack.Pop().Read()), nil
}
