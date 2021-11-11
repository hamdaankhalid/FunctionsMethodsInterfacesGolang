package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"math"
)
/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t^2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/
 func main(){

	reader := bufio.NewReader(os.Stdin)
	// input acceleration, initial velocity, and initial displacement.
	fmt.Println("Enter value for acceleration:")

	accelerationInput, _ := reader.ReadString('\n')

	fmt.Println("Enter value for initial velocity:")

	initialVelocityInput, _ := reader.ReadString('\n')
	

	fmt.Println("Enter value for initial displacemnet:")

	displacementInput, _ := reader.ReadString('\n')

	sanitizedAcceleration, _ := strconv.ParseFloat(strings.Replace(accelerationInput, "\n", "", -1), 64)

	sanitizedVelocity, _ := strconv.ParseFloat(strings.Replace(initialVelocityInput, "\n", "", -1), 64)

	sanitizedDisplacement, _ := strconv.ParseFloat(strings.Replace(displacementInput, "\n", "", -1), 64)


	// generate function
	displacementFunc := GenDisplaceFn(sanitizedAcceleration, sanitizedVelocity, sanitizedDisplacement)

	// ask for time
	fmt.Println("Please enter value for time:")

	time, _ := reader.ReadString('\n')

	sanitizedTime, _ := strconv.ParseFloat(time, 64)
	// use generated function to calc amd return
	result := displacementFunc(sanitizedTime)

	fmt.Println("Result: ", result)
 }

 func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64{
	generatedFunc := func(t float64) float64{
		return (a*(math.Pow(t, 2)*0.5) + (v0 * t) + (s0))
	}
	return generatedFunc
 }