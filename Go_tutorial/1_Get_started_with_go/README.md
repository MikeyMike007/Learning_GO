# Get Started With Go

## Create a new project called hello

### Create a new directory

```bash
mkdir hello
cd hello
```

### Enable Tracking for your code

When your code imports packages contained in other modules, you manage those
dependencies through your code's own module. That module is defined by a go.mod
file that tracks the modules that provide those packages. That go.mod file stays
with your code, including in your source code repository.

To enable dependency tracking run:

```go
go mod init example.com/hello
```

If you plan to publish your module for others to use, the module path must be a
location from which Go tools can download your modlue, such as Github.

### Write code

```go
package main // Declare a main package

import "fmt" // Import fmt package which contains functions for I/O

func main() {
  // Call Println function from fmt to print
  fmt.Println("Hello World")

}
```

### Run your code

```go
go run .
```

## Call code in an external package

### Modify previous code

```go
package main // Declare a main package

import "fmt" // Import fmt package which contains functions for I/O

import "rsc.io/quote" // Imports the quote package from rsc.io

func main() {
  // Call Println function from fmt to print
  fmt.Println(quote.Go())

}
```

### Add new module requirments and sums

```go
go mod tidy
```

### Run the code

```go
go run .
Dont communicate by sharing memory, share memory by communicating
```
