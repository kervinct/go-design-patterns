package creational

import (
	"errors"
	"fmt"
)

// ShirtCloner 克隆接口
type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	// White 白色
	White = 1
	// Black 黑色
	Black = 2
	// Blue 蓝色
	Blue = 3
)

// ShirtsCache 库存
type ShirtsCache struct{}

// GetClone 实现接口
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype // 取值赋值，发生拷贝
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

// ItemInfoGetter 产品接口
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor 命名类型
type ShirtColor byte

// Shirt 具体产品
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo 实现接口
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n",
		s.SKU, s.Color, s.Price)
}

// GetShirtsCloner 获取克隆接口
func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// GetPrice 获取产品信息
func (s *Shirt) GetPrice() float32 {
	return s.Price
}
