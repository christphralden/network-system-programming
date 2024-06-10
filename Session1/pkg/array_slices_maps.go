package pkg

import "fmt"

func RunARM() {

	// ARRAY
	// var intArr [3]int32
	// intArr[0] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(&intArr[0])

	// var intArr [3]int32 = [3]int32{1,2,3}
	// intArr := [3]int32{1,2,3}
	// intArr := [...]int32{1,2,3}

	//MAKE FOR SLICE MAP CHANNEL
	//SLICE -> wrapped version of array with added functionality

	// var intSlice []int32 = []int32{3, 4, 5}
	// fmt.Println(intSlice)
	// fmt.Printf("Length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("Length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// intSlice2  := []int32{8,9}
	// intSlice = append(intSlice, intSlice2...)
	// fmt.Printf("Length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice)

	//type, len, cap
	// var intSlice3 []int32 =  make([]int32, 3,8)
	// fmt.Println(intSlice3)

	//MAP
	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	var myMap2 = map[string]uint8{"alden": 20, "verren": 20}
	// fmt.Print(myMap2["alden"])
	var age, res = myMap2["alden"]
	if !res {
		fmt.Println("invalid")
	} else {
		fmt.Println(age)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	intArr := []int32{1, 2, 3, 4}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// i :=0
	// for{
	// 	if i>=10{
	// 		break
	// 	}
	// 	i++
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
