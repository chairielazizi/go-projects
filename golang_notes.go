package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"reflect"
	"strconv" // to cast number to string
)

// if we declare at package level, cannot assign using :=
// must use
// var i int = 3

// can also intitialze like this
var (
	actorName  string = "Kristen Stewart"
	actorAge   int32  = 28
	actorHobby string = "Eating"
)

// visibilty of variable in golang
// lowercase naming :	var i int	--> if package level, any file in the package can access it
// uppercase naming :	var I int	--> exported from the package and globally visible
// block scope		:	--> not visible to the outside block itself

// Constants
const (
	myConst1 = iota
	myConst2 = iota
	myConst3 = iota
)
const (
	// errorSpecialList = iota
	// _ = iota + 5
	_ = iota // dont care what it is
	catSpecialList
	dogSpecialList
	snakeSpecialList
)
const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // intialize
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	fmt.Println("First programe in GoLang")

	var num1 int = 55 //intialization
	num2 := 65        //declare and initialization(one time only)
	var result int = num1 + num2
	fmt.Println(result)

	//declare and assign value
	var number int //declaration
	number = 4
	fmt.Println("The number is", number)

	//type casting
	number4 := float32(number)
	fmt.Printf("%v is now %T\n", number4, number4)
	//cast int to string, use "strconv" package
	number4string := strconv.Itoa(number)
	fmt.Printf("%v is now %T\n", number4string, number4string)

	var i, j = 100, "hello"
	fmt.Printf("i is %d and j is %s\n", i, j)
	var aa1 float32 = 35
	fmt.Printf("aa1 is %v and type %T\n", aa1, aa1)
	//fmt.Println("i is " + i + "\nj is " + j)
	// will cause compile time error if variables were declared and not used in program

	// summary of variables in GoLang
	// - cant redeclare variables, but can shadow them
	// - all variables must be used
	// - visibilty: - lowercase first leeter for packgae scope
	// 			 - uppercase first letter to export
	// 			 - no private scope
	// - declaration:	- var foo int
	// 				- var foo int = 43
	// 				- foo := 43
	// - type conversions: - destinationType(variable)
	// 					- use strconv package for strings

	// 2nd chapter: Primitives

	// boolean
	var goku bool = true
	fmt.Printf("\n%v, %T\n", goku, goku)
	// check equality
	// initially declaration variable is zero value
	goku1 := 1 == 1
	goku2 := 1 == 2
	fmt.Printf("%v, %T\n", goku1, goku1)
	fmt.Printf("%v, %T\n", goku2, goku2)
	// integer Primitives
	// uint16, uint32
	goku3 := 10 //  1010
	goku4 := 3  //	0011
	// arithmetic operator
	fmt.Println("Arithmetic operator")
	fmt.Println(goku3 + goku4)
	fmt.Println(goku3 - goku4)
	fmt.Println(goku3 * goku4)
	fmt.Println(goku3 / goku4)
	fmt.Println(goku3 % goku4)
	// bitwise operator
	fmt.Println("Bitwise operator")
	fmt.Println(goku3 & goku4)  // 0010
	fmt.Println(goku3 | goku4)  // 1011
	fmt.Println(goku3 ^ goku4)  // 1001
	fmt.Println(goku3 &^ goku4) // 0100
	fmt.Println("Shifting")
	// shifting bitwise
	goku5 := 8              // 2^3
	fmt.Println(goku5 << 3) //	2^3 * 2*3 = 2^6
	fmt.Println(goku5 >> 3) //	2^3 / 2^3 = 2^0
	//	floating point number
	goku6 := 3.14
	goku6 = 13.e72
	goku6 = 2.1E14
	fmt.Printf("%v, %T\n", goku6, goku6)

	goku7 := 10.2
	goku8 := 3.7
	fmt.Println(goku7 + goku8)
	fmt.Println(goku7 - goku8)
	fmt.Println(goku7 * goku8)
	fmt.Println(goku7 / goku8)
	// fmt.Println(goku7 % goku8)	no % operator for float

	// complex number : complex 64, complex 128
	fmt.Println("\nComplex number")
	var goku9 complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", goku9, goku9)
	var goku10 complex64 = 2 + 5.2i
	fmt.Println(goku9 + goku10)
	fmt.Println(goku9 - goku10)
	fmt.Println(goku9 * goku10)
	fmt.Println(goku9 / goku10)
	fmt.Printf("%v, %T\n", real(goku9), real(goku9))
	fmt.Printf("%v, %T\n", imag(goku9), imag(goku9))
	// can also initialize like this
	var goku11 complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", goku11, goku11)

	// text data type
	s1 := "this is a string"
	fmt.Printf("%v, %T\n", s1, s1)
	fmt.Printf("%v, %T\n", s1[2], s1[2]) // letter i in s1
	// s1[2] = "u"
	// fmt.Printf("%v, %T\n", s1, s1)  cause error

	// string concat
	s2 := " just another string"
	fmt.Printf("%v, %T\n", s1+s2, s1+s2)

	// convert to slice of bytes
	// because many func in golang use slice of bytes
	// flexible than hardcoded string
	b1 := []byte(s1)
	fmt.Printf("%v, %T\n", b1, b1)

	// rune
	// rune is type alias for int32
	// r1 := 'a'
	var r1 rune = 'a'
	fmt.Printf("%v, %T\n", r1, r1)

	// summary of Primitive data type in golang
	// - boolean type: - true or false
	// 				- not an alias for other types
	// 				- zero value is false
	// - numeric types: - integers: signed int
	// 				 - unsigned integers
	// 				 - arithmetic operations
	// 				 - bitwise operations
	// 				 - zero value is 0
	// 				 - cannot mix types in same family (int32 + int 64 = error!)

	// 				 - floating point:
	// 				 - follow ieee-754 standard
	// 				 - 32 and 64 bit version
	// 				 - decimal(3.14) , exponential(1e5), mixed(1.1e5)

	// 				 - complex number:
	// 				 - zero value is (0 +0i)
	// 				 - 64 and 128 bit version
	// 				 - complex, real, imag built-in func

	// 				 - text types:-
	// 				 - strings:
	// 				 - UTF-8
	// 				 - immutable
	// 				 - can be concatenated with +

	// 				 - rune:
	// 				 - UTF-32
	// 				 - alias for int32
	// 				 - special methods normally required to process
	// 				 - strings.Reader.#ReadRune

	// 3rd chapter: Constants
	// naming convention
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
	var nonConst int16 = 27
	fmt.Printf("%v, %T\n", int16(myConst)+nonConst, int16(myConst)+nonConst)
	//for declaring constant
	const bA = 99
	fmt.Println("\n", bA)

	// iota is a counter that we can use when we create a numerator constant
	fmt.Printf("%v, %T\n", myConst1, myConst1)
	fmt.Printf("%v, %T\n", myConst2, myConst2)
	fmt.Printf("%v, %T\n", myConst3, myConst3)

	var specialListType int = dogSpecialList
	fmt.Printf("%v\n", specialListType == dogSpecialList)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is hq? %v\n", isHeadquarters&roles == isHeadquarters)

	// summry of Constant
	// - immutable, but can be shadowed
	// - replaced by the compiler at compile time
	// - value must be calculate at compile time
	// - PascalCase for experted Constants
	// - camelCase for internal Constants
	// - typed constants work like immutable var :- only same type
	// - untyped constand work like literal :- work with similar type
	// - enumerated const
	// - iota starts at 0 each const block and increments by one

	// 4th chapter : Array and Slices
	// grades := [3]int{4, 13, 666}
	grades := [...]int{4, 13, 666} // [...] tells that we want to intialize array just large enoght to hold the value, no need to specify the size
	fmt.Printf("\nGrades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))

	// var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("%v\n", identityMatrix)

	arr1 := [...]int{1, 2, 3}
	// arr2 := arr1
	arr2 := &arr1 //arr1 also change with arr2
	arr2[1] = 5
	fmt.Println(arr1)
	fmt.Println(arr2)

	// Slice
	slice1 := []int{1, 2, 3} // intialization
	// slice2 := &slice1 // different from array, slice is a reference type
	slice2 := slice1
	slice2[1] = 5
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("Length: %v\n", len(slice1))
	fmt.Printf("Capacity: %v\n", cap(slice1)) // capacity in the array

	// various way to initialize slice
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice4 := slice3[:]   // slice of all element
	slice5 := slice3[3:]  // slice from 4th element to end
	slice6 := slice3[:6]  // slice of first 6 elements
	slice7 := slice3[3:6] // slice the 4th,5th and 6th elements
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(slice7)

	// lastly using built in function 'make' to declare slice
	fmt.Println("\nusing make fuction")
	slice8 := make([]int, 3, 100) // ([]type,length,capacity)
	fmt.Println(slice8)
	fmt.Printf("Length: %v\n", len(slice8))
	fmt.Printf("Capacity: %v\n", cap(slice8))

	// append function
	slice9 := []int{}
	fmt.Println(slice9)
	fmt.Printf("Length: %v\n", len(slice9))
	fmt.Printf("Capacity: %v\n", cap(slice9))
	slice9 = append(slice9, 1)
	fmt.Println(slice9)
	fmt.Printf("Length: %v\n", len(slice9))
	fmt.Printf("Capacity: %v\n", cap(slice9))
	// concat
	// slice9 = append(slice9, 2, 3, 4, 5)
	slice9 = append(slice9, []int{2, 3, 4, 5}...) // can also do this to concat two slices
	fmt.Println(slice9)
	fmt.Printf("Length: %v\n", len(slice9))
	fmt.Printf("Capacity: %v\n", cap(slice9))

	// stack operation
	slice10 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice10)
	slice11 := slice10[:len(slice10)-1] // [1 2 3 4]
	fmt.Println(slice11)
	slice12 := append(slice10[:2], slice10[3:]...) // [1 2 4 5]
	fmt.Println(slice12)
	fmt.Println(slice10)

	//array in goland
	var array1 [5]string
	array1[0] = "My"
	array1[1] = "name"
	array1[2] = "is"
	array1[3] = "Chairiel"
	array1[4] = "Azizi"
	for k := 0; k < len(array1); k++ {
		fmt.Printf(array1[k] + " ")
	}

	// summary of array and slices
	// - array: - collections of item with the sme type
	// 		 - fixed size
	// 		 - a := [3]int{1,2,3}
	// 		 - a := [_]int{1,2,3}
	// 		 - var a = []int
	// 		 - len() func returns size of array
	// 		 - copies refer to different underlying data
	// - slices: - backed by array
	// 		  - creation styles: - slice existing array or slice
	// 							 - literal styles
	// 							 - via make fuction
	// 		  - len() returns length of slice
	// 		  - cap() returns length of underlying array
	// 		  - append() add elements to slice: may cause expensive copy operation if underlyingarray is to small
	// 		  - copies refer to same underlying array :- affected other slices

	// 5th chapter : Map and Structs

	// Map
	fmt.Println()
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennyslvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	// map1 := map[[]int]string{}
	// error because slice cannot be a key of map
	map1 := map[[3]int]string{} // but array can
	fmt.Println(statePopulations, map1)
	fmt.Println(statePopulations["Ohio"])

	// add data to map
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	//delete data in map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations) // no error, instead return zero

	// ways to check data in map
	pop, ok := statePopulations["Oho"]
	// comma ok to check it
	fmt.Println(pop, ok)

	//length of map
	fmt.Println(len(statePopulations))

	// map also affected like array and slices
	sp := statePopulations
	// delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// map using make function
	statePopulations1 := make(map[string]int)
	// statePopulations1 := make(map[string]int, 10)
	fmt.Println(statePopulations1)

	// Structs
	// structs is independent data type unlike array,slices and map
	type Doctor struct {
		number     int
		doctorName string
		companions []string
	}
	doctor1 := Doctor{
		number:     3,
		doctorName: "Rahman Jailani",
		companions: []string{
			"Peta Jensen",
			"Osman Kencan",
		},
	}
	fmt.Println(doctor1)

	// anonymous struct
	doctor2 := struct{ name string }{name: "John Labu"}
	fmt.Println(doctor2)
	doctor3 := doctor2
	doctor3.name = "Alharbib Anyong"
	fmt.Println(doctor2)
	fmt.Println(doctor3)
	doctor4 := &doctor2
	fmt.Println(doctor4)

	// GoLang has no inheritance
	// but it has the same components like it which called as compositions
	// through embedding
	type Animal struct {
		name   string
		origin string
	}
	type Bird struct {
		Animal // this is the composition by embedding the Animal struct into Bird struct
		speed  float32
		canFly bool
	}

	bird1 := Bird{}
	bird1.name = "Emu"
	bird1.origin = "Australia"
	bird1.speed = 48
	bird1.canFly = false
	fmt.Println(bird1)

	type Birdie struct {
		// Animal: Animal{name: "Hornbill",origin:"Borneo"},
		// speed  float32,
		// canFly bool,
	}
	//for validation framework
	type AnimalValidation struct {
		name   string `required max:"100"`
		origin string
	}
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

	// summary of Structs and Map
	// - Map: - collections of valuea that are accessed by keys
	// 	   - crated via literal or make func
	// 	   - check for presence with "value, ok"
	// 	   - multiple assignments refer to same underlying data
	// - Structs: - collections of disparate data types that describe as single concept
	// 		   - keyed by neme fields
	// 		   - normally created as types, but anonymous structs is allowed
	// 		   - structs are value types
	// 		   - no inheritance, but can use compostion via embedding
	// 		   - tags can be added to struct fields to describe fields

	// 6th chapter: If Else statement

	// must always have bracket although it is only one line code
	fmt.Println()
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop, ok)
	}
	// once it pass the first condition, it will stop and not evaluate the rest condition
	number1 := 50
	guess1 := 30
	if guess1 < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess1 > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess1 < number1 {
			fmt.Println("Too low")
		}
		if guess1 > number1 {
			fmt.Println("Too high")
		}
		if guess1 == number1 {
			fmt.Println("You're right")
		}
	}

	myNum1 := 0.123
	// if myNum1 == math.Pow(math.Sqrt(myNum1), 2) {
	if math.Abs(myNum1/math.Pow(math.Sqrt(myNum1), 2)-1) < 0.001 {
		fmt.Println("These are the same.")
	} else {
		fmt.Println("They are different.")
	}

	// func returnTrue(){
	// 	fmt.Println("returning true")
	// 	return true
	// }

	//if else statement
	var max int = 100
	if max < 10 {
		fmt.Println(max, "is less than 10")
	} else if max > 10 && max < 100 {
		fmt.Println(max, "is between 10 and 100")
	} else {
		fmt.Println(max, "is equal or bigger than 100")
	}

	//switch statement
	a, b := 1, 2
	switch a + b {
	case 1, 0:
		fmt.Println("The sum is 1")
	case 2:
		{
			fmt.Println("The sum is 2")
		}
	case 3:
		fmt.Println("The sum is 3")
	case 4:
		fmt.Println("The sum is 4")
	default:
		fmt.Println("Printing default")
	}

	// fallthrough, not really being used in golang
	num3 := 10
	switch {
	case num3 <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // it is logicless
	case num3 <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// var num4 interface{} = 1
	var num4 interface{} = [3]int{} //array
	switch num4.(type) {
	case int:
		fmt.Println("num4 is an int")
	case float64:
		fmt.Println("num4 is an float64")
	case string:
		fmt.Println("num4 is an string")
	case [3]int:
		fmt.Println("num4 is an [3]int")
		// break // use to skip below code
		fmt.Println("this will print two")
	default:
		fmt.Println("num4 is an another type")
	}

	// summary of if else statement
	// - if statement: - initializer
	// 				- comparison operators
	// 				- logical operators
	// 				- short circuiting // remaining are not to be executed
	// 				- if-else statement
	// 				- if-else if statement
	// 				- equality and floats // when workin with floating point number
	// - switch statement: - swtiching on a tags
	// 					- cases with multiple tests
	// 					- initializers
	// 					- switches with no tags
	// 					- fallthrough
	// 					- inplicit break in golang
	// 					- type switches
	// 					- breaking out early

	// 7th chapter : Looping
	fmt.Println()
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	count1 := 0
	for ; count1 < 5; count1++ {
		fmt.Println(count1)
	}

	// do-while loop alike
	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}

	// break statement
	for count1 < 10 {
		fmt.Println("count1\n", count1)
		count1++
		if count1 == 8 {
			break
		}
	}

	// continue statement
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// the label loop
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// for range loop(type 2 loop)
	fmt.Println("\nFor range loop")
	loop1 := []int{1, 2, 3}
	// using key and value
	for k, v := range loop1 {
		fmt.Println(k, v)
	}

	loop2 := "hello Go!"
	for k, v := range loop2 {
		fmt.Println(k, string(v))
		fmt.Println(k, v)
	}

	// to print out the values only from statePopulations
	// by using _
	for _, v := range statePopulations {
		fmt.Println(v)
	}

	// for loop type 1
	var k int
	fmt.Println("\nFor loop type 1")
	for k = 1; k <= 10; k += 2 {
		fmt.Println(k)
	}

	// summary of Looping
	// - for statement
	// - no while and do-while loop
	// - can use for loop for those

	// - simple loop:
	// - for initializer;test;incrementer {}
	// - for test {}
	// - for {}

	// - exiting early: - break
	// 				 - continue
	// 				 - labels // can help to break inner and outer loop

	// - looping over collections:
	// - arrays,slices,maps,strings,channels
	// - for k,v := range collections {}

	// 8th chapter: defer,panic recover

	// defer the execution
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")
	// defer works in LIFO order(last in first out)
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close() // close after we done read them
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// defer1 := "start"
	// defer fmt.Println(defer1)
	// defer1 = "end"

	// this cause panic
	// panic1, panic2 := 1, 0
	// ans1 := panic1 / panic2
	// fmt.Println(ans1)

	// fmt.Println("start")
	// panic("something bas happen")
	// fmt.Println("end")

	// start GoLang web app
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Ahoyy Go!"))
	// })
	// err1 := http.ListenAndServe(":8080",nil)
	// if err1 != nil{
	// 	panic(err.Error())
	// }

	// recover
	// fmt.Println("start")
	// defer fmt.Println("this was deferred")
	// recover will return nil if app is panicking
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panicker()
	// panic("something bas happen")
	// fmt.Println("end")

	// summary of defer,panic and recover
	// - defer: - used to delay execution of a statement until func exits
	// 		 - useful to group open and clise func together
	// 		 - be careful in loops
	// 		 - run in LIFO(last in,first out) order
	// 		 - arguments evaluated at time defer executed,not at time of called func execution
	// - panic: - occur when program cannot continue at all
	// 		 - dont use when file cant be open, unless it is critical
	// 		 - use for unrecoverable events - cannot obtain TCP port for web server
	// 		 - func will stop executing
	// 		 - deferred func will still fire
	// 		 - if nothing handles panic, program will exits
	// - recover: - used to recover from panics
	// 		   - only useful in deferred func
	// 		   - current func will not attemp to continue, but higher func in call stack will

	// 9th chapter: Pointers

	var n1 int = 42
	var p1 *int = &n1 // holding the address og n1 in memory location
	fmt.Println(n1, p1)
	fmt.Println(n1, *p1) // deferencing operator
	n1 = 27
	fmt.Println(n1, *p1)
	*p1 = 14
	fmt.Println(n1, *p1)

	n2 := [3]int{1, 2, 3}
	p2 := &n2[0]
	p3 := &n2[1]
	fmt.Printf("%v %p %p\n", n2, p2, p3)

	// pointer to struct
	type pointerStruct struct {
		foo int
	}
	var ps *pointerStruct
	fmt.Println(ps)
	// ps = &pointerStruct{foo: 42}
	// fmt.Println(ps)
	ps = new(pointerStruct)
	ps.foo = 42
	fmt.Println(ps.foo)
	// (*ps).foo = 42  、、same as the above
	// fmt.Println((*ps).foo)

	// map and slices have their built in like pointer behaviour
	// where the copy will not follow the original data
	// slice13 := []int{1, 2, 3}

	// summary of Pointers
	// - pointer types use an asterisk(*) as a prefix to type pointed to
	// - *int - a pointer to an integer
	// - use the address of operator(&) to get address of variable

	// - deferencing pointer
	// - deference a pointer by preceding with an asterisk
	// - complex types(eg. structs) are automatically deferenced

	// - create pointer to objects
	// - can use  the address of operator(&) if value type already exists
	// - *ms := myStruct{foo:42}
	// - p := &ms

	// - use address of operator(&) before initializer
	// - &mysStruct{foo:42}

	// - use the new keyword
	// - cant initialize fields at the same time

	// - types with internal pointer
	// - all assignment operations in GO are copy operations
	// - slices and maps contain internal pointers,so copies point to same underlying data

	// 10th chapter : Functions

	for i := 0; i < 5; i++ {
		sayMessage("Ahoyy Go!", i)
	}
	sayGreeting("Ahoyy", "Stacey")

	greeting1 := "Ahoyy"
	name1 := "Hannah"
	sayGreeting1(&greeting1, &name1)
	fmt.Println(name1)

	sum("The sum is", 1, 2, 3, 4, 5)

	sum1 := sumPointer(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *sum1)

	div1, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(div1)

	// anonymous function
	func() {
		fmt.Println("Ahoyy Go!")
	}() // 		() to invoke the function

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	// assign anonymous to variable
	var f1 func() = func() {
		fmt.Println("Ahoyy Go!")
	}
	f1()

	// var divide func(float64, float64) (float64, error)
	// divide = func(a, b float64) (float64, error) {
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("Cannot divide by zero")
	// 	} else {
	// 		return a / b, nil
	// 	}
	// }
	// div2, err := divide(5.0, 0.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(div2)
	fmt.Println("Go")

	//Method
	g1 := greeter{
		greeting: "Ahoyy",
		name:     "Go!",
	}
	g1.greet()
	g1.greetPointer()

	// summary of Functions
	// - function - func foo(){...}
	// - parameter - comma delimited list of var and types
	// 			- func foo(var string,baz int)

	// 			- parameter of same type
	// 			- func foo(bar,baz int)

	// 			- when pointer are passed inm the func can change the value in the caller
	// 			- always true for data of slices and maps

	// 			- variadic parameters
	// 			- send list of the same type
	// 			- must be last parameter
	// 			- received as a slice
	// 			- func foo(bar string, baz ...int)

	// - return values - single return value : func foo() int
	// 				- multiple return value : func foo() (int,error)
	// 				- (result type,error)

	// 				- can ues named return values
	// 				- can return addresses of local variables

	// - anonymous function - func that dont have names
	// 					 - immediately invoked
	// 					 - func(){
	// 						 ...
	// 					 }()
	// 					 - a := func(){
	// 						 ...
	// 					 }
	// 					 a()

	// - func as types - can assign func to variables or use as arguments and return values in funcs
	// 				- type signature is like func signature, with no parameter names
	// 				- var f func(string,string,int) (int,error)

	// - method - func that executes in context of a type
	// 		 - func (g greeter) greet() {
	// 			 ...
	// 		 }
	// 		 - receiver can be value or pointer
	// 		 - value receiver gets copy of type
	// 		 - pointer receiver gets pointer to type

	// 11th chapter : Interfaces
	var w1 Writer = ConsoleWriter{} // polymorhpic behaviour
	w1.Write([]byte("Ahoyya Go!"))

	myInt := IntCounter(0) // cast an integer to it
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Peta Jensen, this is a test"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	// bwc := wc.(io.Reader)
	fmt.Println(bwc)

	r2, ok := wc.(io.Reader)
	// r, ok := wc.(*BufferedWriterCloser)
	// r, ok := wc.(BufferedWriterCloser) // error because of pointer
	if ok {
		fmt.Println(r2)
	} else {
		fmt.Println("Conversion failed")
	}

	// empty interface
	var myObj1 interface{} = NewBufferedWriterCloser()
	if wc1, ok := myObj1.(WriterCloser); ok {
		wc1.Write([]byte("Hello Alex Blake, this is a test"))
		wc1.Close()
	}
	r3, ok := myObj1.(io.Reader)
	if ok {
		fmt.Println(r3)
	} else {
		fmt.Println("Conversion failed")
	}

	var inter1 interface{} = 0
	switch inter1.(type) {
	case int:
		fmt.Println("inter1 is an integer")
	case string:
		fmt.Println("inter1 is a string")
	default:
		fmt.Println("I dont know what it isS")
	}

	// summary of interface
	// - basic:
	// type Writer interface {
	// 	Write([]byte) (int, error) // accept a slice bytes and return int & error
	// }

	// // implements the interface implicitly
	// type ConsoleWriter struct{}

	// func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// 	n, err := fmt.Println(string(data))
	// 	return n, err
	// }

	// - composing interfaces
	// - type conversion: - var wc WriterCloser = NewBufferedWriterCloser
	// 				   - bwc := wc.(*BufferedWriterCloser)
	// - The empty interface and type switches
	// - implementing with values vs. pointers
	// - method set of value is all with value receivers
	// - method set of pointer is all methods, regardless of receiver type

	// Best Practices of interfaces in Go
	// - use many, small interfaces
	// - single method interfaces are some of the most powerful and flexible
	// - io.Writer, io.Reader, interface{}

	// - dont export interfaces for types that will be consumed
	// - do export interfaces for types that will be used by packages
	// - design functions and methods to receive interfaces whenever possible

}

//panicker function
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// panic(err)
		}
	}()
	panic("something bad happen")
	fmt.Println("done panicking")
}

// 10th chapter : Functions
// open curly braces must be on the same line as the name of func

func sayMessage(msg string, index int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", index)
}
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	// name = "Ted"
	// fmt.Println(name)
}
func sayGreeting1(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}
func sumPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// not very common in go
func sumValues(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		// panic("Cannot provide zero as second value")
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// Methods
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // this is a method
	fmt.Println(g.greeting, g.name)
}
func (g *greeter) greetPointer() { // this is a method
	fmt.Println(g.greeting, g.name) // implicit deferencing pointer
}

// 11th chapter : Interfaces
// interfaces just describe the data's behaviour
// method definition
// type Writer interface {
// 	Write([]byte) (int, error) // accept a slice bytes and return int & error
// }

// implements the interface implicitly
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// 2nd example
type Incrementer interface {
	Increment() int // this is the method
}
type IntCounter int

//implementation of Incrementer interface
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// 3rd example
// composing interfaces
type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}
type WriterCloser interface { // like embedding in structs
	Writer
	Closer
}
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 8 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
