# 日志

- import

```go
import "github.com/fideism/golang-tools/logger"
```

- use

不设置 `channel`
```go
logger.NewEntry().WithFields(logrus.Fields{
		"kwy": "value",
	}).Info("message")
```

设置 `channel`
```go
logger.NewEntry(logger.Channel("sms")).Info("设置channel为 sms")
```

统一设置
```go
logger.SetOption(logger.Channel("sms"))

logger.Entry().Info("channel sms")
```

- channle
使用 `logger.Channel("sms")` 指定特定的`channel`，方便日志收集