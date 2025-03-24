# protoc-gen-go-errors

`protoc-gen-go-errors` 是一个 protoc 插件，用于从 protobuf 枚举生成错误处理代码。

## 功能

- 从 protobuf 枚举自动生成错误处理代码
- 支持基于 HTTP 状态码的错误映射
- 支持错误检查和创建
- 支持国际化(i18n)错误消息

## 新增特性: 国际化(i18n)错误支持

最新版本增加了对国际化错误消息的支持，使您可以根据上下文生成不同语言的错误消息。

### 使用方法

1. 首先，实现 `errors.I18nMessage` 接口:

```go
type I18nMessage interface {
    // Localize 根据上下文和数据对错误原因进行本地化
    Localize(ctx context.Context, reason string, data map[string]any) string
}
```

2. 注册你的 i18n 管理器:

```go
// 创建 i18n 实现
i18n := &YourI18nImplementation{}
// 注册到错误系统
errors.RegisterI18nManager(i18n)
```

3. 使用生成的 WithContext 方法创建错误:

```go
// 例如使用带有语言信息的上下文创建错误
ctx := context.WithValue(context.Background(), "language", "zh")
err := ErrorUserNotFoundWithContext(ctx, map[string]any{"id": "12345"})
```

## 示例

请查看 `example/i18n_example.go` 获取完整的国际化错误处理示例。