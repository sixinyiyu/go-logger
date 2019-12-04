## go-logger

### 介绍
封装zap的简单日志操作工具

go 1.13

### 使用

```go
logger := log.NewLogger("./log/biz.out", "debug")

logger.Info(" Hello World Logger !")
```
