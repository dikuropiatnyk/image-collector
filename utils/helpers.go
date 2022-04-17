package utils

import "encoding/base64"

func ConvertToBase64(bytes []byte, result *string) {
	*result = base64.StdEncoding.EncodeToString(bytes)
}
