package files

import (
    "io/ioutil"
    "time"

    "github.com/google/uuid"
    "github.com/hollson/goddd/infrastructure/helper"
    "github.com/hollson/goddd/interfaces"
)

// 实体
type fileEntity struct {
    *interfaces.FileInfo
}

// 根据值对象创建实体
// func newEntity(fileInfo interfaces.FileInfo) *fileEntity {
//     return &fileEntity{
//         FileInfo: &fileInfo,
//     }
// }

func newEntity(fileInfo FileValue) *fileEntity {
    return &fileEntity{
        FileInfo: &fileInfo.FileInfo,
    }
}


func (en *fileEntity) EntityID() uuid.UUID {
    uid, err := uuid.Parse(en.Id)
    if err != nil {
        uid = uuid.Nil
    }
    return uid
}

func (en *fileEntity) AddFile(body []byte) (err error) {
    filePath := FileRootPath + time.Now().Format("2006/01/02/") + en.Id + en.FileName
    helper.MakesureFileExist(filePath)
    err = ioutil.WriteFile(filePath, body, 0755)
    if err != nil {
        return
    }
    en.FilePath = filePath
    en.Status = 1
    _, err = fileRepos.AddObj(en.FileInfo)
    return
}
