package base64_utils

import "encoding/base64"

//EncodeFromString
func EncodeFromString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//EncodeFromBytes
func EncodeFromBytes(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

//DecodeToBytes
func DecodeToBytes(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//DecodeFromString
func DecodeToString(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
