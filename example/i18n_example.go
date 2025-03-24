package main

import (
	"context"
	"log"

	"github.com/protoc-gen/protoc-gen-go-errors/errors"
)

// SimpleI18n 是一个简单的i18n实现示例
type SimpleI18n struct {
	templates map[string]map[string]string
}

// Localize 根据上下文和数据本地化错误消息
func (i *SimpleI18n) Localize(ctx context.Context, reason string, data any) string {
	lang := GetLanguageFromContext(ctx)
	if templates, ok := i.templates[reason]; ok {
		if template, ok := templates[lang]; ok {
			// 这里应该使用一个真正的模板引擎来填充数据
			// 简单起见，我们直接返回模板
			return template
		}
		// 如果找不到特定语言的模板，尝试使用英语
		if template, ok := templates["en"]; ok {
			return template
		}
	}
	return ""
}

// GetLanguageFromContext 从上下文中获取语言代码
func GetLanguageFromContext(ctx context.Context) string {
	if ctx == nil {
		return "en"
	}

	// 这里假设语言代码存储在上下文的值中
	if lang, ok := ctx.Value("language").(string); ok {
		return lang
	}
	return "en"
}

func ExampleI18nErrors() {
	// 创建一个简单的i18n实现
	i18n := &SimpleI18n{
		templates: map[string]map[string]string{
			"UnauthorizedError_UNAUTHORIZED": {
				"en": "Authentication required",
				"zh": "需要身份验证",
				"ja": "認証が必要です",
			},
			"ClientError_USER_NOT_FOUND": {
				"en": "User not found",
				"zh": "找不到用户",
				"ja": "ユーザーが見つかりません",
			},
		},
	}

	// 注册i18n管理器
	errors.RegisterI18nManager(i18n)

	// 创建带有英语语言的上下文
	enCtx := context.WithValue(context.Background(), "language", "en")
	// 创建带有中文语言的上下文
	zhCtx := context.WithValue(context.Background(), "language", "zh")
	// 创建带有日语语言的上下文
	jaCtx := context.WithValue(context.Background(), "language", "ja")

	// 使用不同的上下文创建错误
	// 注意：实际代码中使用生成的 ErrorUserNotFoundWithContext 方法
	enErr := errors.NewWithContext(400, "ClientError_USER_NOT_FOUND", enCtx, nil)
	zhErr := errors.NewWithContext(400, "ClientError_USER_NOT_FOUND", zhCtx, nil)
	jaErr := errors.NewWithContext(400, "ClientError_USER_NOT_FOUND", jaCtx, nil)

	// 输出不同语言的错误消息
	log.Printf("English error: %v", enErr)
	log.Printf("Chinese error: %v", zhErr)
	log.Printf("Japanese error: %v", jaErr)

	// 输出效果类似于：
	// English error: error: code = 400 reason = ClientError_USER_NOT_FOUND message = User not found metadata = map[] cause = <nil>
	// Chinese error: error: code = 400 reason = ClientError_USER_NOT_FOUND message = 找不到用户 metadata = map[] cause = <nil>
}
