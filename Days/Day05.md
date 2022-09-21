## Day 05

## Golang

Exercise items for Go will be in a separate folder.

**Basics**

Go is syntactically similar to C and often referred to as Golang because of its domain name. Basic concepts will introduce three major features: Packages, Functions, and Variables.

**Packages**

Go applications are organized in packages. A package is a collection of sources files located in the same directory. All source files in a directory must share the same package name. Conventional for the package name to be the last directory in the import path. When a package is imported, only entities (e.g. functions, types, variables, constants) whose names start with a captial letter can be used / accessed. Recommended style of naming in Go is that identifiers will be named using ```camelCase```, except for those meant to be accessible across packages which should be ```PascalCase```.

Ex: ```package lasagna```

## Variables

Go is statically-typed, which means all variables must have a defined type at compile-time.

Ex: ``` var explicit int // Explicitly typed```

Can also use an initializer, and the compiler will assign the variable type to match the type of the initializer.

Ex: ``` implicit := 10 // Implicitly typed as an int```

Once declared, variables can be assigned values using the ```=``` operator. Once declared, a variable's type can never change.

Ex:
```
count := 1 // Assigned initial value
count = 2 // Update value

count = false // This will give a compiler error due to the value not being an int type
```

Constants

They hold data like variables, but their value can't be changed during the execution of a program. They are defined using ```const``` and can be numbers, characters, strings, or booleans.

Ex: ```const Age = 21 // Defines numeric constant```

Functions

Go functions accept zero or more parameters. They must be explicitly typed as there is no type inference.

Values are returned from functions using ```return```

Functions are invoked by specifying the function name and passing arguments for each of the parameters.

*Note* Go supports two types of comments. Single line comments are preceded by ```//``` and multiline comments are inserted between ```/*``` and ```*/```

Ex:
```
package greeting

// Hello is public function
func Hello (name string) string {
    return hi(name)
}

// hi is private function
func hi (name string) string {
    return "hi " + name
}
```

Resource: https://exercism.org/tracks/go/concepts/basics
