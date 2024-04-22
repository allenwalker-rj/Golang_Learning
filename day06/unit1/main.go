package main

import "fmt"

func main() {
	fmt.Println("unit map start")
	/**
	 * 1.只声明map内存是没有分配空间的,必须进行初始化，才会分配空间
	 * 2.map是无序的，key是不可以重复的
	 */
	CreateMap1()
	CreateMap2()
	CreateMap3()
	KeyRepeat()
	ValueRepeat()
	KeyDelete()
	ValueFind()
	LengthMap()
	IterateMap()
	NestingMap()

}

func CreateMap1() {
	fmt.Println("CreateMap1 output")
	a := make(map[int]string)
	a[123] = "张三"
	fmt.Println(a)
}

func CreateMap2() {
	fmt.Println("CreateMap2 output")
	b := make(map[int]string, 10)
	b[345] = "李四"
	fmt.Println(b)
}

func CreateMap3() {
	fmt.Println("CreateMap3 output")
	c := map[int]string{678: "王五"}
	fmt.Println(c)
}

func KeyRepeat() {
	fmt.Println("KeyRepeat output")
	//key 不可以重复，重复的key，value会被覆盖
	var a map[int]string = make(map[int]string, 10)
	a[123] = "张三"
	a[345] = "李四"
	a[678] = "王五"
	a[123] = "朱六"
	fmt.Println(a)
}

func ValueRepeat() {
	fmt.Println("ValueRepeat output")
	//value可以重复
	var a map[int]string = make(map[int]string, 10)
	a[123] = "张三"
	a[345] = "李四"
	a[678] = "王五"
	a[999] = "张三"
	fmt.Println(a)
}

func KeyDelete() {
	fmt.Println("KeyDelete output")
	d := make(map[int]string)
	d[123] = "张三"
	d[345] = "李四"
	d[678] = "王五"
	delete(d, 123)
	fmt.Println(d)
}

func ValueFind() {
	fmt.Println("ValueFind output")
	e := make(map[int]string)
	e[123] = "张三"
	e[345] = "李四"
	e[678] = "王五"
	//v=value , b=boolean是否存在
	v, b := e[345]
	fmt.Println(v, b)
}

func LengthMap() {
	fmt.Println("LengthMap output")
	d := make(map[int]string)
	d[123] = "张三"
	d[345] = "李四"
	d[678] = "王五"
	fmt.Println(len(d))
}

func IterateMap() {
	//map的遍历
	fmt.Println("IterateMap output")
	f := make(map[int]string)
	f[123] = "张三"
	f[345] = "李四"
	f[678] = "王五"
	for k, v := range f {
		fmt.Printf("key=%v value=%v\n", k, v)
	}
}

func NestingMap() {
	//map嵌套
	fmt.Println("NestingMap output")
	g := make(map[string]map[int]string)
	g["a"] = map[int]string{123: "张三", 234: "李四"}
	g["b"] = map[int]string{345: "王五", 456: "朱六"}
	fmt.Println(g)

	for k1, v1 := range g {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("k2=%v v2=%v\n", k2, v2)
		}
	}
}
