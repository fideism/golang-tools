# Golang Tools

基础包

### 日志

```go
github.com/fideism/golang-tools/logger
```

使用 `logger.Channel("sms")` 指定特定的`channel`，方便日志收集

具体使用 [Readme](./logger)

### 参数

`Params util.Params`

```go
import "github.com/fideism/golang-tools/util"

// Params map[string]interface{}
type Params map[string]interface{}

// Set 设置值
func (p Params) Set(k string, v interface{})

// Get 获取值
func (p Params) Get(k string) (v interface{})

// GetString 强制获取k对应的v string类型
func (p Params) GetString(k string) string

// Exists 判断是否存在
func (p Params) Exists(k string) bool

//具体使用
p := util.Params{
    "openid": "xx",
}

// also can
p.Set("notify_url", "https://github.com/fideism/golang-wechat")

// InterfaceToString 不定类型强转string

InterfaceToString
```
