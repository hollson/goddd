package files

import (
    "github.com/hollson/goddd/interfaces"
)

// 值对象(value object)
type FileValue struct {
    interfaces.FileInfo
    FileBody []byte
}
