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