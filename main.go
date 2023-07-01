package main

import "golessons/project"

func main() {
	//sd()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.Demo3()
	//variables.Demo1()
	//loops.Demo1()
	//loops.Demo2()
	//loops.Demo3()
	//games.TwoGames()
	//oops.Demo4()
	//arrays.Demo2()
	//arrays.Demo4()
	//slices.Demo1()
	//slices.Demo2()
	//functions.SayHi2("amine")

	//functions.SayHi()
	//functions.Add(2, 3, 6)
	//var result = functions.Add2(2, 3)
	//fmt.Println(result * 90)
	//functions.SayHi1("sd")
	//add, substract, multiply, divide := functions.Multiplecalculations(3, 6)
	//fmt.Println(add, substract, multiply, divide)
	//var add, substract, multiply, divide = functions.Multiplecalculations(2, 3)
	//fmt.Println("addition:", add)
	//fmt.Println("substitution:", substract)
	//fmt.Println("multiplication:", multiply)
	//fmt.Println("division:", divide)
	//newarray1, newarray2 := functions.AddVariadic(2, 3, 6, 8, 6)
	//fmt.Println(newarray1)
	//fmt.Println(newarray2)
	//numbers := []int{6, 9, 5, 9, 63, 2, 4, 8}
	//we can add array to variadic fuctions. to indicate that we should use (...) after calling array
	//fmt.Println(functions.AddVariadic(numbers...))
	//maps.Demo1()
	//for_range.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	//number := 20
	//fmt.Println("actual number:", number)
	//pointers.Demo1(&number)
	//fmt.Println("actual number:", number)

	//pointers.Test()

	// // //structs.Demo1()
	// // //structs.Demo2()
	// // //pointers.Demo1(&number)
	// // //structs.Demo2()
	// // //go goroutines.EvenNumbers()
	// // //go goroutines.OddNumbers()
	// // //time.Sleep(time.Hour)

	// evenNumberCn := make(chan int)
	// oddNumberCn := make(chan int)

	// go channels.EvenNumbers(evenNumberCn)
	// go channels.OddNumbers(oddNumberCn)

	// evenNumberCnAddition, oddNumberCnAddition := <-evenNumberCn, <-oddNumberCn

	// addition := evenNumberCnAddition * oddNumberCnAddition
	// fmt.Println("multiplication:", addition)
	//games.TwoGames()

	//interfaces.Demo1()
	//interfaces.Demo2()
	//deferstatement.B()
	//deferstatement.Test()
	//deferstatement.Demo3()
	//errorhandling.Demo1()

	//interfaces.Demo3()
	//errorhandling.Demo2()

	//fmt.Println(errorhandling.Guess2(150))
	//stringfunctions.Demo1()
	//stringfunctions.Demo2()
	// restful.Demo1()
	// restful.Demo2()

	project.AddProduct()
	project.GetAllProducts()

}
