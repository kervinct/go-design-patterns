package creational

import (
	"fmt"
)

// PaymentMethod 支付方法接口
type PaymentMethod interface {
	Pay(amount float32) string
}

// 实现的支付方法
const (
	Cash         = 1
	DebitCard    = 2
	NewDebitCard = 3
)

// GetPaymentMethod 返回一个PaymentMethod对象指针或错误
func GetPaymentMethod(m int) (PaymentMethod, error) {
	// return nil, errors.New("Not implemented yet")
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case NewDebitCard:
		return new(NewDebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

// CashPM 现金支付方法
type CashPM struct{}

// DebitCardPM 信用卡支付方法
type DebitCardPM struct{}

// Pay 接口实现
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// Pay 接口实现
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

// NewDebitCardPM 新的实例
type NewDebitCardPM struct{}

// Pay 接口实现
func (d *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}
