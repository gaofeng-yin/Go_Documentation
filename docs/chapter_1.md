# Introduction to Golang
Go is programming language of the future, it borrows ideas from existing languages like C and C++.Easy to code, efficient on compilation and execution, in a nutshell: easy to learn and fast to compile. Invented by genius from Google,  **Robert Griesemer**, **Rob Pike**, and **Ken Thompson** to handle Google massive infrastructure and to increase developers productivity and software scalability as the systems grow. 
<p>

# Getting Started
Nothing better than starting with traditional "Hello world!" program:

<code>
package main

import(

		"fmt"
)

func main(){
	
		fmt.Println("Hello wolrd!")
	
}
</code>


<b>Package main</b> tells the Go compiler that the package should compile as an executable program instead of a shared library.<p>
The <b>main function</b> is a entry point of the program and here you write all your code. In this case have function <b>Println</b> which will print out "Hello world!" on the screen. Println comes from package <b>fmt</b> and we can make use of it by importing from Go library.<p>
In Go there is no necessity of ';' at end of each sentence. It makes our code more readable and our life easier because often programmers forget to put semi colon in one sentence and then the code doesn't compile and it takes time and energy to fix these kind of error.

# Variables

Normally variables are used to store values(string, numbers, or booleans). To declare a variable use var keyword or := which is a short assignment statement. <p>
To declare integer variable num with value 10:
<code>

		var num int = 10
		
</code>

shorthand declaration:

<code>

		num := 10
		
</code>

To define a string:
<code>

		var str string = "Hello world!"
		
		st := "How are you?"
		
</code>

Booleans can be 0 or 1 or true or false, zero is false and one is true:
<code>

		var t bool = true
		var f bool = false
		
		off := 0
		on  := 1
		
</code>

You can also declare multiple variables at once:
<code>

		var a, b, c, d, e int
		
		var str, name, gender string
		
</code>

