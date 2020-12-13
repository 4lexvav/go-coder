#### Usage:
`go run main.go <input string> <mode>`

```
Modes:
    - encode
    - decode
```

#### Example:
`go run main.go "hello world" decode`

#### Output:
```
Original: hello world
Result: h2ll4 w4rld
```

#### Example 2:
`go run main.go "h3 th2r2" decode`

#### Output:
```
Original: h3 th2r2
Result: hi there
```