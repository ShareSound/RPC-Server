package shared

import(
	"github.com/dchest/uniuri"
)

func GetDefaultSecureHash() string { return uniuri.New() }
func GetSecureHash(length int) string { return uniuri.NewLen(length) }
