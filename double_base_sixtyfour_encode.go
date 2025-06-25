package sdk

import "encoding/base64"

func doubleBase64Encode(s string) string {
	first := base64.StdEncoding.EncodeToString([]byte(s))
	return base64.StdEncoding.EncodeToString([]byte(first))
}
