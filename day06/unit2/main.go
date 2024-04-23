package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

type Student struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) AddAge() {
	fmt.Println("AddAge")
	p.Age = p.Age + 1
	fmt.Println(p)
}

func (p Person) updateName() {
	fmt.Println("updateName")
	p.Name = "黑色"
	fmt.Println(p)
}

func (p *Person) updateNameTrue() {
	fmt.Println("updateNameTrue")
	(*p).Name = "白色"
	fmt.Println(*p)
}

func main() {
	fmt.Println("unit struct start")
	CreatePerson1()
	CreatePerson2()
	CreatePerson3()
	CreatePerson4()
	ConvertPersonToStudent()
	var p Person = Person{"张三", 18, "男"}
	p.AddAge()
	ValuePassing()
	PointerPassing()
}

func CreatePerson1() {
	fmt.Println("CreatePerson1")
	var p Person
	p.Name = "张三"
	p.Age = 18
	p.Sex = "男"
	var p2 Person
	p2.Name = "李四"
	p2.Age = 19
	p2.Sex = "男"
	fmt.Println(p)
	fmt.Println(p2)
}

func CreatePerson2() {
	fmt.Println("CreatePerson2")
	var p Person = Person{"张三", 18, "男"}
	fmt.Println(p)
}

func CreatePerson3() {
	fmt.Println("CreatePerson3")
	//p是指针，p其实指向的就是地址，应该给这个地址的指向的对象的字段赋值
	var p *Person = new(Person)
	(*p).Name = "张三"
	(*p).Age = 18
	(*p).Sex = "男"
	fmt.Println(*p)
	//为了符合正常的编程习惯，go提供的了简化的赋值方式
	//go编译器底层对p.Name转换为(*p).Name
	p.Name = "李四"
	p.Age = 19
	p.Sex = "男"
	fmt.Println(*p)
}

func CreatePerson4() {
	fmt.Println("CreatePerson4")
	var p *Person = &Person{"张三", 18, "男"}
	fmt.Println(*p)
}

func ConvertPersonToStudent() {
	fmt.Println("ConvertPersonToStudent")
	var p Person = Person{"张三", 18, "男"}
	var s Student = Student{"李四", 19, "男"}
	s = Student(p)
	fmt.Println(s)
}

func ValuePassing() {
	//方法是值传递，而非引用的传递
	fmt.Println("ValuePassing")
	var p Person = Person{"李四", 19, "男"}
	p.updateName()
	fmt.Println(p)
}

func PointerPassing() {
	//指针的传递
	fmt.Println("PointerPassing")
	var p Person = Person{"李四", 19, "男"}
	(&p).updateNameTrue()
	fmt.Println(p)
}
