package testdemo

import "fmt"

// ArrayFunc ...
func ArrayFunc() {
	// 定义一个数组
	var hens [6]float64
	//给每隔数组元素赋值
	hens[0] = 1.1
	hens[1] = 2.2
	hens[2] = 3.3
	hens[3] = 4.4
	hens[4] = 5.5
	hens[5] = 6.6
	//遍历数组，求出总体和
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	//求出平均体重
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)
	fmt.Println(hens)
	fmt.Printf("数组首地址%p\n", &hens)
}

// ArrayFunc01 ...
func ArrayFunc01() {
	//从终端输入5个成绩
	var score [5]float64
	for i := 0; i < 5; i++ {
		fmt.Printf("请输入%d个元素的值\n", i+1)
		fmt.Scanln(&score[i])
	}
	fmt.Println(score)
}

// ArrayFunc02 ... 四种初始化数组方式
func ArrayFunc02() {
	var array01 [3]int = [3]int{11, 22, 33}
	fmt.Println("array01=", array01)

	var array02 = [3]int{11, 22, 33}
	fmt.Println("array02=", array02)

	var array03 = [...]int{11, 22, 33}
	fmt.Println("array03=", array03)
	//类型推导，指定下标赋值
	var array04 = [...]int{3: 11, 2: 22, 0: 33}
	fmt.Println("array04=", array04)
}

// ArrayFunc03 ... 遍历数组
func ArrayFunc03() {
	strArr := [...]string{"小昆虫", "小白菜", "小东风"}
	fmt.Println("strArr=", strArr)
	//下标遍历数组
	for i := 0; i < len(strArr); i++ {
		fmt.Println("i=", i, strArr[i])
	}
	fmt.Println("=========================")
	//range遍历数组
	for index, value := range strArr {
		fmt.Println("index=", index, "value=", value)
	}
	fmt.Println("=========================")
	for _, value := range strArr {
		fmt.Println("value=", value)
	}
}
