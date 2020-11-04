package structual

// Athlete 运动员
type Athlete struct{}

// Train 直接组合
func (a *Athlete) Train() {
	println("Training")
}

// Swim 函数
func Swim() {
	println("Swimming!")
}

// CompositeSwimmerA 直接成员组合
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

// -------------------------------------------
// 推荐使用

// Trainer 接口
type Trainer interface {
	Train()
}

// Swimmer 接口
type Swimmer interface {
	Swim()
}

// SwimmerImplementor 结构
type SwimmerImplementor struct{}

// Swim 实现接口
func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

// CompositeSwimmerB 接口对象组合
type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// -------------------------------------------

// Animal 结构
type Animal struct{}

// Eat 方法
func (r *Animal) Eat() {
	println("Eating")
}

// Shark 嵌入
type Shark struct {
	Animal
	Swim func()
}

// --------------------------------------------

// Tree 二叉树
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

// -------------------------------------------

// Parent 结构
type Parent struct {
	SomeField int
}

// Son 结构
type Son struct {
	// Parent // 匿名时，下面方法不能编译
	P Parent // 非匿名，可以引用值
}

// GetParentField 不能用Son代替此处的Parent参数
func GetParentField(p *Parent) int {
	return p.SomeField
}
