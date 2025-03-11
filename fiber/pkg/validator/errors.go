package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatError(errors *validate.Errors) string {
	var res string
	for _, v := range errors.Errors {
		res = res + strings.Join(v, ", ") + ", "
	}
	return res
}
