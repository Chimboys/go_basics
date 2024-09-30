# Go Enums and iota

## Working with Enums in Go

### 1. Using `iota`
- `iota` is a predeclared identifier used to simplify constant definitions in Go. It starts at `0` and increments by `1` with each new constant in a constant block.

### 2. Using Functions with Enums
- You can define a custom type for enums and implement methods (e.g., `String()`) for enhanced functionality and readability.

## Implementing Nullable Properties

### 1. Using Pointers
- Set the property to a pointer type. If the pointer is `nil`, it indicates that the value is null.

### 2. Checking for Default Values
- Check if a value is equal to its default (zero) value to determine if it is effectively "null."

## `make` vs. `new`
* In Go, `make` is used for initializing **slices**, **maps**, and **channels**. It sets up the data structure and prepares it for use, including allocating memory and initializing values.
* `new` is a built-in function that allocates memory for a variable of a specified type but does not initialize it. **It returns a pointer to the variable with a zero value.**

## Capital first letter
-  In Go, to make a function, variable, type, or method exportable **(accessible from other packages)**, you need to capitalize the first letter of its name. For example, a function defined as `func myFunc()` will not be exported, but `func MyFunc()` will be.

## Omitting Types for Consecutive Named Parameters

* If two or more consecutive named function parameters share the same type, you can omit the type for all but the last parameter.
* **Equivalence Examples**:
  * `x int, y int`
  * `x, y int`

# Running Multiple Go Files

When working with Go applications that consist of multiple source files, it is essential to understand how to compile and run these files together. Here’s a breakdown of the process and the reasoning behind it.

## Command Structure

To run multiple Go source files, you can use the following command:

```bash
go run path/to/file1.go path/to/file2.go
```

## Files Seperation Logic in Go
**Driver Program:** Contains the main function and must be part of the main package.
**Other Files:** Should share the same package but should not include a main function.


functions/
- functions.go           *# Contains the main function*
- library.go             *# Contains additional functions and shares the same package*

``` bash
go run functions/functions.go functions/library.go
```

## Package-Level Variables vs. Function-Level Variables in Go

### Package-Level Variables
- **Scope**: 
  - Accessible throughout the package where defined.
  - Exported variables (starting with an uppercase letter) can be accessed from other packages.
  
- **Lifetime**: 
  - Exists for the entire duration of the program.
  - Initialized once when the package is imported or the program starts.

## Switch Statement in Go

- A `switch` statement in Go is a concise alternative to a series of `if-else` statements. It executes the first case that matches the condition expression.

### Key Differences from Other Languages:
- **Automatic Break**: Unlike C, C++, Java, JavaScript, and PHP, Go automatically breaks after executing a case, so there’s no need to include a `break` statement.
- **Flexibility**: Go's switch cases can evaluate any type of expression, not just constants or integers.

## Stacking Defers in Go
 **Deferred Function Calls**: In Go, deferred function calls are pushed onto a stack.
- **Execution Order**: When a function returns, its deferred calls are executed in a **last-in-first-out** (LIFO) order.

## Summary
- **Enums**: Use `iota` for efficient enum definitions.
- **Nullable Properties**: Use pointers for nullable fields or check for default values to manage absence of data.
- **Memory Management**: Use `make` for **slices**, **maps**, and **channels**; use `new` for allocating memory for other types.



[Learn more about Struct Literals](https://go.dev/tour/moretypes/5)
[Working with Slices](https://go.dev/tour/moretypes/15)