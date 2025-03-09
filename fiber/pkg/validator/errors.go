package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatError(errors *validate.Errors) string {
	var res string
	for k, v := range errors.Errors {
		res = res + k + ": " + strings.Join(v, ", ") + "\n"
	}
	return res
}
