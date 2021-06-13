# Create a Go module

## Make a project directory

```bash
mkdir greetings
cd greetings
```

## Start dependency tracking

```go
go mod init example.com/greetings
```

## Write code

```go
package main

import (
  "fmt"
  "example.com/greetings"
)

func main() {

  message := greetings.Hello("Gladys")
  fmt.Println(message)

}

```
