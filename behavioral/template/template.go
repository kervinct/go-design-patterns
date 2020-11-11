package behavioral

import "strings"

// MessageRetriever 算法中一步的抽象接口
type MessageRetriever interface {
	Message() string
}

// Template 模版接口
type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

// TemplateImpl 模版实现
type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

// ExecuteAlgorithm 算法执行接口实现
func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

// AnonymousTemplate 匿名模版
type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

// ExecuteAlgorithm 算法执行接口实现
func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

// --------------------------------------
// 避免修改接口

// TemplateAdapter 模版适配器
type TemplateAdapter struct {
	myFunc func() string
}

// Message 算法实现
func (a *TemplateAdapter) Message() string {
	return ""
}

// MessageRetrieverAdapter 获取接口对象
func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}

type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}
