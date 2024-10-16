Quick Golang tips
==============

# fmt prints placeholders

%d: Decimal integer
%b: Binary integer
%o: Octal integer
%x: Hexadecimal integer (lowercase)
%X: Hexadecimal integer (uppercase)
%f: Floating-point number
%e: Scientific notation (lowercase)
%E: Scientific notation (uppercase)
%s: String
%q: Double-quoted string
%v: Default format for the value
%T: Type of the value
%p: Pointer address


# Arrays vs Slices

	arrays -> fixed size
	slice -> dynamic lengthOfNumbers

	Therefore, when declaring an array, we must specify its size, for example:

		friends := [4]string{"Dan", "Diana", "Paul", "John"}

	Declaring a slice uses a shorter syntaz

		friends := []string{"Dan", "Diana", "Paul", "John"}

	In the background, go actually uses an array to declare a slice. Go implements a slice as a data structure called slice header, which is the
	runtime representation of the slice.
	It contains three fields:
		1. The address of the backing array (pointing to the first element of the backing array)
		2. The length of the slice. The built-in function len() returns it
		3. The capacity of the slice which is the size of the backing array after the first element of the slice. Its returned 
			by the cap() built-in function


# Pointers

A pointer is a variable type that stores the memory address of another variable.
A pointer value is the address of a variable, or nil if it hasn't been initialized yet.

