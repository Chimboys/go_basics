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



## Summary
- **Enums**: Use `iota` for efficient enum definitions.
- **Nullable Properties**: Use pointers for nullable fields or check for default values to manage absence of data.
- **Memory Management**: Use `make` for **slices**, **maps**, and **channels**; use `new` for allocating memory for other types.

