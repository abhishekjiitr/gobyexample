// ## Sorting by Functions

// Sorting a slice by a function is a bit tricker in Go
// than you may be used to in other languages. Let's look
// at some examples to see how it works.

package main

import "fmt"
import "sort"

type Person struct {
    Name string
    Age  int
}

type ByName []Person

func (this ByName) Len() int {
    return len(this)
}
func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
    return len(this)
}
func (this ByAge) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func main() {
    kids := []Person{
        {"Jack", 10},
        {"Jill", 9},
        {"Bob", 12},
    }
    fmt.Println("Original:", kids)

    sort.Sort(ByName(kids))
    fmt.Println("ByName:  ", kids)

    sort.Sort(ByAge(kids))
    fmt.Println("ByAge:   ", kids)
}
