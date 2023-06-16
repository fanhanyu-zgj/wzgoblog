package requests

import (
	"wz/app/models/article"

	"github.com/thedevsaddam/govalidator"
)

func VlidateArticleForm(data article.Article) map[string][]string {
	// 1. 表单规则
	rules := govalidator.MapData{
		"title": []string{"required", "min_cn:3", "max_cn:40"},
		"body":  []string{"required", "min_cn:10"},
	}
	// 2. 定制错误信息
	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min:标题长度需大于 3",
			"max:标题长度需小于 40",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min:标题长度需大于 10",
		},
	}
	// 3. 配置选项
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // Struct 便签标识符
		Messages:      messages,
	}
	// 4. 开始认证
	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
