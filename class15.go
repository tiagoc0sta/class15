/*package main

import (
 "fmt"
)

func main() {

 //anonymous function for filtering even numbers

 filterEven := func(numbers []int) []int {
  var result []int
  for _, number := range numbers {
   if number%2 == 0 {
    result = append(result, number)
   }

  }
  return result
 }

 // Calling the anonymous function
 numbers := []int{1, 2, 3, 4, 5, 6}
 evenNumbers := filterEven(numbers)
 fmt.Println("Even numbers: ", evenNumbers)

 applylistOperation := func(number []int, operation func([]int) int) int {

  return operation(numbers)

 }

 sum := applylistOperation(numbers, func(x []int) int {

  total := 0

  for _, n := range x {
   total += n
  }

  return total

 })

 fmt.Println("Sum:", sum)
}

package main

import (
	"fmt"
)

func main() {

 concat := func(s1, s2 string) string {

  return s1 + s2
 }

 result := concat("Hello ", "World")

 fmt.Println("Concatenation: ", result)

 applyStringOperation := func(s string, operation func(string) string) string {
  return operation(s)
 }

 result = applyStringOperation("Paradise", func(s string) (result string) {

  for _, v := range s {
   result = string(v) + result
  }
  return

 })

 fmt.Println("Reversed String: ", result)

}

package main

import (
 "fmt"
)

func main() {

 //anonymous function to check if a bool value is true
 isTrue := func(b bool) bool {
  return b == true
 }

 //calling the anonymous function

 result := isTrue(true)

 fmt.Println("Is true: ", result)

 //Create an anonymous function to negate a bool

 negate := func(b bool) bool {
  return !b
 }

 result = negate(false)

 fmt.Println("Negation: ", result)

}

package main

import (
 "fmt"
)

func main() {

 //anonymous function to append an element to a slice

 appendElement := func(slice []int, element int) []int {
  return append(slice, element)
 }

 //calling the anonymous function
 numbers := []int{1, 2, 3}
 fmt.Println("OG Slice: ", numbers)
 numbers = appendElement(numbers, 4)

 fmt.Println("Appended Slice: ", numbers)

 //Using anonymous function to filter elements greater than a threshold

 filterGreaterThan := func(slice []int, threshold int) []int {

  var result []int
  for _, num := range slice {
   if num > threshold {
    result = append(result, num)
   }
  }
  return result
 }

 filteredNumbers := filterGreaterThan(numbers, 2)

 fmt.Println("Filtered slice: ", filteredNumbers)

}*/

/*package main

import (
 "fmt"
)

func main() {

 //Anonymous function to check if a key exists in a map

 keyExists := func(m map[string]int, key string) bool {
  _, exists := m[key]
  return exists
 }

 //Calling the anonymous function

 ages := map[string]int{
  "Jake":   31,
  "Holly":  25,
  "Jordan": 28,
 }

 result := keyExists(ages, "Holly")
 fmt.Println("Key exists: ", result)

 deleteKey := func(m map[string]int, key string) {
  delete(m, key)
 }

 deleteKey(ages, "Holly")
 fmt.Println("After Deletion: ", ages)
}*/

//write an anonymous function in go that takes two integers as parameters and returns theis product
package main

import "fmt"

func main() {
    // Define and call the anonymous function
    product := func(a, b int) int {
        return a * b
    }(5, 3) // Passing 5 and 3 as arguments

    fmt.Println("Product:", product)
}

