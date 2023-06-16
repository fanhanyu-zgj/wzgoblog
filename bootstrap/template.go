package bootstrap

import (
	"embed"
	"wz/pkg/view"
)

func SetupTemplate(tmpFS embed.FS) {
	view.TplFS = tmpFS
}
