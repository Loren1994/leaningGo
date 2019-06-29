package learn

import (
	"fmt"
	"reflect"
)

type Cat struct {
	name string
	// 带有结构体tag的字段
	Type int `json:"type" id:"100"`
}

func Learn23() {
	var aa = 0
	typeOfA := reflect.TypeOf(aa)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	var cat Cat
	typeOfCat := reflect.TypeOf(cat)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	var cat1 = &Cat{}
	typeOfCat1 := reflect.TypeOf(cat1)
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat1.Name(), typeOfCat1.Kind())
	// 取类型的元素
	typeOfCat1 = typeOfCat1.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat1.Name(), typeOfCat1.Kind())
	fmt.Println(typeOfCat1.Field(0), typeOfCat1.NumField())
	f, b := typeOfCat1.FieldByName("name")
	fmt.Println(f, b)
	// 声明整型变量a并赋初值
	var a = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	// 获取interface{}类型的值, 通过类型断言转换
	var getA = valueOfA.Interface().(int)
	// 获取64位的值, 强制类型转换为int类型
	var getA2 = int(valueOfA.Int())
	fmt.Println(getA, getA2, valueOfA.Int())
	valueOfA1 := reflect.ValueOf(&a)
	//Elem() 方法获取 a 地址的元素
	valueOfA1 = valueOfA1.Elem()
	valueOfA1.SetInt(999)
	fmt.Println(valueOfA1.Int())
	// 根据反射类型对象创建类型实例
	// 等效于：new(int)，因此返回的是 *int 类型的实例
	aIns := reflect.New(typeOfA)
	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	//构造函数参数, 传入两个整型值
	params := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	retList := funcValue.Call(params)
	fmt.Println(retList[0])

}

// 普通函数
func add(a, b int) int {
	return a + b
}
