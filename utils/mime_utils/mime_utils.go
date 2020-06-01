package mime_utils

import (
	"path/filepath"
	"strings"
)

func GetMimeType(p string) string {
	ext := filepath.Ext(p)
	if ext != "" {
		ext = ext[1:]
	}
	ext = strings.ToLower(ext)
	mType := MimeTypes[ext]
	if mType == "" {
		return "application/octet-stream"
	}
	return mType
}
