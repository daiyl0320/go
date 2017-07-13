package example

import "fmt"
import "framework"

type v struct {
	obj framework.Object
}

var v1 = v{framework.Object{"Variables", "This is a pratice about variables"}}

// global static variables, can just declaration by using var
var gUninitialInt int
var gInitialInt int = 1
var gUninitialBool bool
var gInitialBool bool = true
var gUninitialString string
var gInitialString string = "test1"

// global exported variables
var GUninitialInt int
var GInitialInt int = 2
var GUninitialBool bool
var GInitialBool bool = true
var GUninitialString string
var GInitialString string = "test2"

func (v1 *v) Id() string {
	return v1.obj.Id
}

func (v1 *v) Description() string {
	return v1.obj.Description
}

func init() {
	framework.Register(&v1)
}

func (v1 *v) Run() {
	// local variables
	var i = 1   // int
	var j = 1.1 // float
	ii := 2     // int
	jj := 2.2   //float
	fmt.Printf("i = %v, j = %f, ii = %v, jj = %f\n", i, j, ii, jj)
	fmt.Printf("gUninitialInt = %v, gUninitialBool = %v, gUninitialString = %v\n",
		gUninitialInt, gUninitialBool, gUninitialString)
	fmt.Printf("GUninitialInt = %v, GUninitialBool = %v, GUninitialString = %v\n",
		GUninitialInt, GUninitialBool, GUninitialString)
	fmt.Printf("gInitialInt= %v, gInitialBool= %v, gInitialString= %v\n",
		gInitialInt, gInitialBool, gInitialString)
	fmt.Printf("GInitialInt= %v, GInitialBool= %v, GInitialString= %v\n",
		GInitialInt, GInitialBool, GInitialString)
	// declare variable in if statement can just be ":="
	if k := true; k {
		// k can just visiable in if-else statement
		fmt.Printf("k = %v\n", k)
	} else {
		fmt.Printf("k = %v\n", k)
	}
	// this block can not see k, so we can define another k again
	k := false
	fmt.Printf("k` = %v\n", k)
}
