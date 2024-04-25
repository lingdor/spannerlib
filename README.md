# spannerlib
a common base lib for golang
# Features
## 1. Exception processing
### 1.1 Must get value and panic when wrong
Normal
```go
num,numErr:=strconv.Itoa("123")
if numErr!=nil {
    panic(numErr)
}
age,ageErr:=strconv.Itoa("18")
if ageErr!=nil {
    panic(ageErr)
}

```
Must get
```go
num := E.Must1(strconv.Itoa("123"))
//or
age := E.Catch1(strconv.Itoa("123")).IfErrorDataFunc(func()any{
    return "number was input wrong"
}).Must()
//or
num:=E.Catch1(strconv.Itoa("123")).Do(func(err error){
    panic(err)
})

```
### 1.2. Wrap Error
```go
err:=fmt.Errorf("123")
err:=errors.Wrap(err,0,"msg")

fmt.printf("%v",err)
```
output
```go
Exception MSG
testing.tRunner(/usr/local/go/src/testing/testing.go:1689)
```

## 2. String processing 
### 2.1. StartWith

```go
if str.StartWith("hello world","hello") {
//true
}
```
### 2.2. StringPick

```go
fmt.Println(E.Must1(StringPick("<html><body>123</body></html>", "<body>", "</body>")))
```
output:
```text

123
```

