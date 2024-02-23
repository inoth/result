# result
尝试参考rust Result错误处理方式

```sh
  go get -u github.com/inoth/result
```

```go
func add(a, b int) Result[int] {
	return result.New(a + b, /*errors.New("some err")]*/)
}

var sum int

sum = add(1, 2).Unwrap()

sum = add(1, 2).Expect(errors.New("input a err"))

add(1, 2).Match(
	func(val int) {
		fmt.printf("success output: %d\n", val)
    sum = val
	},
	func(err error) {
		fmt.printf(err.Error())
	})
```
```go
var f *os.File

f = result.Must(os.Open("test.txt")).Unwrap()

f = result.Must(os.Open("test.txt")).Expect(errors.New("input a err"))

result.Must(os.Open("test.txt")).Match(
  func(val *os.File) {
    fmt.printf("file name %s\n",val.Name())
    f = val
  },func(err error) {
    fmt.printf(err.Error())
  }
)
```