# Exercises - Ninja Level 1

### Hands-on exercise #1

##### 1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the

IDENTIFIERS “x” and “y” and “z”
a. 42
b. “James Bond”
c. true

##### 2. Now print the values stored in those variables using

a. a single print statement
b. multiple print statements
code: here’s the solution: https://play.golang.org/p/yYXnWXIQNa video: 017

### Hands-on exercise #2

##### 3. Use var to DECLARE three VARIABLES.

The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
a. identifier “x” type int
b. identifier “y” type string
c. identifier “z” type bool

##### 4. in func main

a. print out the values for each identifier
b. The compiler assigned values to the variables. What are these values called?
code: here’s the solution: https://play.golang.org/p/jzHwSlles9 video: 018

### Hands-on exercise #3

Using the code from the previous exercise,

##### 5. At the package level scope, assign the following values to the three variables

a. for x assign 42
b. for y assign “James Bond”
c. for z assign true

##### 6. in func main

a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
b. print out the value stored by variable “s”
code: here’s the solution: https://play.golang.org/p/QFctSQB_h3 video: 019

### Hands-on exercise #4

● FYI - nice documentation and new terminology “underlying type”
○ https://golang.org/ref/spec#Types For this exercise

##### 7. Create your own type. Have the underlying type be an int.

##### 8. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR”

keyword

##### 9. in func main

a. print out the value of the variable “x”
b. print out the type of the variable “x”
c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
d. print out the value of the variable “x”
code: here’s the solution: https://play.golang.org/p/snm4WuuYmG video: 020

### Hands-on exercise #5

Building on the code from the previous example

##### 10. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

```go
  type hotdog int
  var x hotdog
  var y int
```

##### 11. in func main

a. this should already be done
i. print out the value of the variable “x”
ii. print out the type of the variable “x”
iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
iv. print out the value of the variable “x”

b.now do this
i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE

1. then use the “=” operator to ASSIGN that value to “y”
2. print out the value stored in “y”
3. print out the type of “y”
   code: here’s the solution: https://play.golang.org/p/cj8RrYgBOD video: 021

## Quiz for "variables, values & type"

### Hands-on exercise #6

An explanation of the quiz questions and their answers is provided. video: 022
