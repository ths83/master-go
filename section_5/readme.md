# Time To Practice: Reflection

The standard formatting of structs is not very nice to read. Remember our CelestialBody, Planet, and Star structs?
Filling them with data and printing them with

`fmt.Printf("%v\n", ....)`
yealds this uninspiring output:

```
{{Sun 1988000000 1391400 587h0m0s} 0 4.83 -9223372036854775808 0xc4200a2000}
{{Mercury 330 4879 1407h0m0s} 3.7 false true [] 0xc4200a2060 <nil>}
{{Venus 4870 12104 5833h0m0s} 8.9 true false [] 0xc4200a20c0 0xc4200a2000}
{{Earth 5970 12756 24h0m0s} 9.8 true true [Moon] 0xc4200a2120 0xc4200a2060}
{{Mars 642 6792 24h37m0s} 3.7 true false [Phobos Deimos] <nil> 0xc4200a20c0}
```

I think we can do better.

## Your tasks

> NOTE: Great explanation [here](https://stackoverflow.com/a/39866671) and
> playground [here](https://go.dev/play/p/3kwe7ag1i1C)

### Task #1: Pretty-print a struct

Write a library package pretty with a function

`Print(i interface{})`

that takes an interface{} parameter and prints this parameter according to the following rules:

If the type is a struct, then print

- the struct’s type
- an opening brace
- with indent: all fields, including field name, type, and value
- a closing brace
  Else just print the type and the value of the parameter.
  Consider that the passed-in value may be invalid. The reflect package has a function that tests the validity of a
  value.

Obviously, we need the reflect package here, as all this information (types, struct field names, unexported struct
fields) are only accessible through reflection.

The provided code includes a file main.go with the sun and a few planets already being set up for use. There is also a
package file “pretty.go” prepared for you to implement.

## Task #2 (optional): Follow pointers

Expand your solution of task #1 to follow pointers.

That is, if the current value is a pointer,

- print “*”
- print the name of the struct field IF the pointer is a struct field
- print the value that the pointer points to (with indenting)
  Consider these edge cases:

- The pointer may be nil
- The data may contain circular references. Address this by doing either or both of:
- Choosing a maximum depth for follwing pointers (say, 10)
  Maintaining a list of already visited nodes, and only follow pointers that do not point to an already visited node.

> A hint: In my sample solution, I used a helper function “print(v reflect.Value,…)” with some more parameters that I
> call recursively when evaluating fields and following pointers. The additional parameters control the indent level and
> pass field names down one level, so that they can be printed with the proper indent level.

I also made this print() function a method of a struct, and in this struct I keep a list of already visited pointers.

However, the provided solution is only one way of implementing the tasks. Feel free to find a different, maybe better,
solution.

The file prettySolution.go.txt implements tasks #1 and #2. Its output looks like this:

Output:

```
main.Star: {
main.CelestialBody: {
Name (string): Sun
Mass (int64): 1988000000
Diameter (int64): 1391400
RotationPeriod (time.Duration): 587h0m0s
}
Distance (float64): 0
Magnitude (float64): 4.83
Discovery (int64): -9223372036854775808
*FirstPlanet (main.Planet): main.Planet: {
main.CelestialBody: {
Name (string): Mercury
Mass (int64): 330
Diameter (int64): 4879
RotationPeriod (time.Duration): 1407h0m0s
}
Gravity (float64): 3.7
HasAtmosphere (bool): false
HasMagneticField (bool): true
Satellites ([]string): []
*next (main.Planet): main.Planet: {
main.CelestialBody: {
Name (string): Venus
Mass (int64): 4870
Diameter (int64): 12104
RotationPeriod (time.Duration): 20998800000000000
}
Gravity (float64): 8.9
HasAtmosphere (bool): true
HasMagneticField (bool): false
Satellites ([]string): []
*next (main.Planet): main.Planet: {
main.CelestialBody: {
Name (string): Earth
Mass (int64): 5970
Diameter (int64): 12756
RotationPeriod (time.Duration): 86400000000000
}
Gravity (float64): 9.8
HasAtmosphere (bool): true
HasMagneticField (bool): true
Satellites ([]string): [Moon]
*next (main.Planet): main.Planet: {
main.CelestialBody: {
Name (string): Mars
Mass (int64): 642
Diameter (int64): 6792
RotationPeriod (time.Duration): 88620000000000
}
Gravity (float64): 3.7
HasAtmosphere (bool): true
HasMagneticField (bool): false
Satellites ([]string): [Phobos Deimos]
next (*main.Planet): <nil>
previous (!*main.Planet ALREADY VISITED)
}
previous (!*main.Planet ALREADY VISITED)
}
previous (!*main.Planet ALREADY VISITED)
}
previous (*main.Planet): <nil>
}
}
```

## Task #3 (optional): Print struct tags

Also print struct field tags after the field values.

Happy coding!
