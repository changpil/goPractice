# Basic Type
```
bool, string, int, float32, float64
```
# Aggregate Type ( Types from other types)
```
array, struct
```
# Reference Type ( store memory address)
```
pointer, slice, channel, function
```
# Interface Type
```
func DoSomething(t interface{}) {
    switch reflect.TypeOf(t).name() {
        // chang actin based on type expected
    }
}
Dosomething(interface1)
Dosomething(interface2)
```
# Go CLI commands
```
go build
go install
go get
go doc
go test
```
