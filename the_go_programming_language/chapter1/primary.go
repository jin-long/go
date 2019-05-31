package main
import "fmt"


func main() {
	fmt.Println("hello worod !")
/* 变量申明 */
	// var 变量名 类型
	var intVal int
	// 这时候会产生编译错误
	//intVal :=1
	// 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	intVal,intVal1 := 1,2
	fmt.Printf(" intVal:%d, intVal1:%d", intVal, intVal1)
	
	//多变量申明
	var a1,b1,c1 int
	//多变量赋值
	a1,b1,c1=10,20,30
	//多变量申明与赋值
	a2,b2,c2:=100,200,300

	fmt.Println(" ****** for test ",a1,b1,c1,a2,b2,c2);
/* for test*/
	number := [6]int{1,2,38,9}
	for i,x:=range number {
		fmt.Printf("location i %d, value x %d \n", i, x)
	}
	
	var a int
	var b int = 15
	for a=1; a<10; a++ { 
		fmt.Printf(" i:%d \n", a)
	}

	for a<b {
		a++
		fmt.Printf("a=%d, b=%d \n", a, b)
	}

/* if test */
	fmt.Println("****** if test ");
	if a<b {
		fmt.Printf("a<b a:%d, b:%d\n", a, b);
	} else {
		fmt.Printf("a=>b a:%d, b:%d\n", a, b);
	}

/* switch test */
	/* 1.go中每个case执行完后，不需要break */
	fmt.Println("****** if test ");

	/* 2.支持多条件匹配 */
	var v int = 2
	switch(v){
		case 1,2,3: fmt.Printf("1,2,3 \n")
		default:
			fmt.Printf("unknow  \n")
	}

	/* 3.判断interface 变量中实际存储的变量类型 */
	var x interface{}
	switch x.(type) {
		case int8:
			fmt.Println(" int8")
		case int32:
			fmt.Println(" int32")
		case int64:
			fmt.Println(" int64")
		default:
			fmt.Println(" unknow")
	}

	/* 4.通过fallthrough不判断条件继续执行下一个case */
	var str string = "a"
	switch str {
		case "a": 
			fmt.Println(" a ")
			fallthrough
		case "b":
			fmt.Println(" b ")
			//fallthrough
		case "c":
			fmt.Println(" c ")
			fallthrough
		default:
			fmt.Println(" default char ")
	}
}
