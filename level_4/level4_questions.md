# Exercises - Ninja Level 4

### Hands-on exercise #1

- Using a COMPOSITE LITERAL:
  - create an ARRAY which holds 5 VALUES of TYPE int
  - assign VALUES to each index position.
- Range over the array and print the values out.
- Using format printing
  - print out the TYPE of the array
    solution: https://play.golang.org/p/tD0SzV3hdf

### Hands-on exercise #2

- Using a COMPOSITE LITERAL:
  - create a SLICE of TYPE int
  - assign 10 VALUES
- Range over the slice and print the values out.
- Using format printing
  - print out the TYPE of the slice
    solution: https://play.golang.org/p/sAQeFB7DIs

### Hands-on exercise #3

Using the code from the previous example, use SLICING to create the following new slices which are then printed:

```
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
```

solution: https://play.golang.org/p/SGfiULXzAB

### Hands-on exercise #4

Follow these steps:

- start with this slice
  `x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}`
- append to that slice this value ○ 52
- print out the slice
- in ONE STATEMENT append to that slice these values
  - 53
  - 54
  - 55
- print out the slice
- append to the slice this slice
  - `y := []int{56, 57, 58, 59, 60}`
- print out the slice
  solution: https://play.golang.org/p/QUYhtSBaDS

### Hands-on exercise #5

To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:

- start with this slice
  - `x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}`
- use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
  - `[42, 43, 44, 48, 49, 50, 51]`

solution: https://play.golang.org/p/u8zpHLfgSE

### Hands-on exercise #6

Create a slice to store the names of all of the states in the United States of America. Use make and append to do this. Goal: do not have the array that underlies the slice created more than once. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position, without using the range clause. Here is a list of the 50 states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

solution:

- my first take on this project, which was an INCORRECT solution
  - https://go.dev/play/p/vE9BCWMA2ET
- instructive proof from a student which showed me the error of my ways :)
  - https://go.dev/play/p/szSlBwJG2S-
- HERE IS THE CORRECT SOLUTION
  - https://go.dev/play/p/ZcyFCyo7_XH

It is good to think about the LEN and CAP. It is also good to think about the underlying array. You want to think about this so that you're thinking of memory use. Also, it is good to think about what got stored in what position. And it is good to remember that slices are dynamic - they can grow in size.
**FROM THE SPEC:**
The length of a slice s can be discovered by the built-in function len; unlike with arrays it may change during execution. The elements can be addressed by integer indices 0 through len(s)-1. The slice index of a given element may be less than the index of the same element in the underlying array.
A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.
The array underlying a slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by slicing a new one from the original slice. The capacity of a slice a can be discovered using the built-in function cap(a).

- referrence:
  https://go.dev/ref/spec#Slice_types

### Hands-on exercise #7

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

```
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."
```

Range over the records, then range over the data in each record.
solution: https://play.golang.org/p/FMM4c2PhZg

### Hands-on exercise #8

Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

```
`bond_james`: [`Shaken, not stirred`, `Martinis`, `Women`]
`moneypenny_miss`: [`James Bond`, `Literature`, `Computer Science`]
`no_dr`: [`Being evil`, `Ice cream`, `Sunsets`]
```

solution: https://play.golang.org/p/nTzSlRa9_A

### Hands-on exercise #9

Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
solution: https://play.golang.org/p/_CkxAhRrDJ

### Hands-on exercise #10

Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
solution: https://play.golang.org/p/TYl5EbjoeC
