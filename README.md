# debug

print for debug.

## Import

```go
import "github.com/najeira/debug"
```

## Usage

### Print

print any values.

```go
debug.Print(errors.New("some error"))
```

format and print.

```go
debug.Print("id=%d, name=%s", 123, "Alice")
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
