package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	fmt.Println("in answer")
	return 42
}

func init() {
	fmt.Println("in init")
	WhatIsThe = 0
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}

// -----output-----
//in answer
//in init
//It's all a lie.

//AnswerToLife() is guaranteed to run before init() is called, and init() is guaranteed to run before main() is called.
//Keep in mind that init() is always called, regardless if there's main or not, so if you import a package that has an init function, it will be executed.
//Also, keep in mind that you can have multiple init() functions per package, they will be executed in the order they show up in the code (after all variables are initialized of course).