# debug

print for debug.

## Import

```go
import "github.com/najeira/debug"
```

## Usage

### Printf

format and print.

```go
debug.Printf("id=%d, name=%s", 123, "Alice")
```

### Print

print any values.

```go
debug.Print(errors.New("some error"))
```

### Logger

set your logger.

```go
debug.Logger = log.New(yourFile, "", log.LstdFlags)
```

disable logging.

```go
debug.Logger = nil
```
