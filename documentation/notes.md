Some notes about go...
==============

Go is a staticly type language. Everytime we create a variable, we assign a type and that type will not change.
For example, in javascript we can assign:

var test = 123
test = undefined
test = "123"

In go, once we define the type we msut respect it.

-- Basic Go Types --

bool -> true false
string -> "test"
int -> 0 -10000 99999
float64 -> 10.00001


We can define a variable like so:

- var card string = "Ace of spades"
- card := "Ace of spades"

We only use := when we are defining a new variable!!!! If we are reassigning a different value to the string, we should only use =

- card = "Five of Diamonds"

------------------------------------------------------

# Arrays & Slices

Array -> fixed length of records
Slice -> An array that can grow or shrink

In slice we can use append function. This function does not modify the slice, it will create a new slice with more capacity.

sliceName[startIndexIncluding:upToNotIncluding]

-- Type conversion in go --

This will transform the string "Hi there!" in an array of bytes.

[]byte("Hi there!")

or, assuming arr is a []byte

string(arr)


A slice is a structure with three items, a pointer to the head of the array, a capacity and a length
Basically, go creates both a slice and an array internaly.
When a slice is created and passed to a function, what will happen is that go will create a copy of the slice, but the slice is a pointer to the array in memory
therefore, we do not need pointers while using slices because, even though the copy will be made, the copy will have a pointer to the array

This happens with other types!! 

Value types (need pointers to change things in a function)
int, float, string, bool, structs

Reference Types (dont worry about pointers)
slices, maps, channels, pointers, functions


-----------------------------------------------------

type Person struct {
	firstName string
	lastName string
}

alex := Person{firstName: "Alex", lastName: "Anderson"}

fmt.printf("%+v", alex) will print the values inside alex