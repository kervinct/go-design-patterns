package structual

import (
	"errors"
	"fmt"
)

// IngredientAdd 接口
type IngredientAdd interface {
	AddIngredient() (string, error)
}

// PizzaDecorator 装饰器
type PizzaDecorator struct {
	Ingredient IngredientAdd
}

// AddIngredient 实现接口
func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

// Meat 结构
type Meat struct {
	Ingredient IngredientAdd
}

// AddIngredient 实现接口
func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "meat"), nil
}

// Onion 结构
type Onion struct {
	Ingredient IngredientAdd
}

// AddIngredient 实现接口
func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s, %s", s, "onion"), nil
}
