# Golang String Builder

# 引入依赖

```text
go get -u github.com/golang-infrastructure/go-StringBuilder 
```

# 说明

Golang中的strings.Builder并不是真正的Builder模式，连续拼接要换行就很草， 这个库就是修改了一下内置的strings.Builder的签名并加了几个Append方法，用得更爽一些。

# Example

```go
builder := string_builder.New()
s := builder.AppendString("test").AppendString(" a is: ").AppendInt(100).String()
fmt.Println(s) // test a is: 100
```




