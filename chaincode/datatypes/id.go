package datatypes

import (
	"strconv"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var cur = 0

var id = assets.DataType{
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		cur += 1
		return strconv.Itoa(cur), cur, nil
	},
}
