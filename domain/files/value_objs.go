package files

import (
	"github.com/hollson/goddd/interfaces"
)

type FileInfos struct {
	interfaces.FileInfo
	FileBody []byte
}
