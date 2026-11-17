# GO LANGUAGE – TIME PACKAGE

## 1. Introduction

- Go provides a built-in package called `time` to work with date and time.
- It is used to get the current time, format dates, calculate time differences, create timers and add delays.
- To use the time package, we must import it.

**Syntax:**
```go
import "time"
```

## 2. Important Functions of Time Package

| Function | Description |
|---|---|
| `time.Now()` | Returns the current date and time. |
| `time.Sleep()` | Pauses program execution for a specific duration. |
| `time.Date()` | Creates a date and time. |
| `time.Format()` | Formats date and time. |
| `time.Parse()` | Converts a string into time. |
| `time.Since()` | Calculates elapsed time. |
| `time.After()` | Returns a channel that receives after a duration. |
| `time.NewTimer()` | Creates a timer. |

## 3. Program to Display Current Date and Time

**Program:**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()

    fmt.Println("Current Date and Time:", t)
}
```

**Explanation:**
- `time.Now()` returns the current date and time.
- The result is stored in variable `t`.
- `fmt.Println()` displays the current date and time.

## 4. Program to Display Formatted Date and Time

**Program:**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()

    fmt.Println("Date:", t.Format("02-01-2006"))
    fmt.Println("Time:", t.Format("15:04:05"))
    fmt.Println("Day:", t.Format("Monday"))
}
```

**Explanation:**
- `time.Now()` gets the current date and time.
- `Format()` converts the time into the required format.
- `02-01-2006` represents day-month-year.
- `15:04:05` represents hour-minute-second.
- `Monday` displays the day of the week.

**Important:** Go uses a special reference date for formatting:

`Mon Jan 2 15:04:05 MST 2006`

Common formatting values:

| Value | Meaning |
|---|---|
| 02 | Day |
| 01 | Month |
| 2006 | Year |
| 15 | Hour (24-hour format) |
| 04 | Minutes |
| 05 | Seconds |
| Monday | Day name |
| January