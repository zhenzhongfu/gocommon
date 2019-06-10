package wrap

import (
	"github.com/astaxie/beego/validation"

	"github.com/zhenzhongfu/gocommon/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Infoln(err.Key, err.Message)
	}

	return
}
