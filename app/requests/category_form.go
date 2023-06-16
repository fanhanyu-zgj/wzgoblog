package requests

import (
	"wz/app/models/category"

	"github.com/thedevsaddam/govalidator"
)

func VlidateCategoryForm(data category.Category) map[string][]string {
	// 1. 表单规则
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}
	// 2. 定制错误信息
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min:分类名称长度需大于 2",
			"max:分类名称长度需小于 8",
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
