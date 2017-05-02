package conversion

import (
	"encoding/base64"
	"github.com/lttviet/cryptopals/stringutil"
)

func HexToBase64(hexstr string) string {
	byteArr := stringutil.DecodeHexStr(hexstr)
	return base64.StdEncoding.EncodeToString(byteArr)
}
