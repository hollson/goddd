package application

import (
    "context"
    "fmt"
    "io/ioutil"

    "github.com/hollson/goddd/domain/files"
    "github.com/hollson/goddd/infrastructure"
    "github.com/hollson/goddd/infrastructure/bus"
    "github.com/hollson/goddd/infrastructure/ddd"
    "github.com/hollson/goddd/infrastructure/logs"
    eh "github.com/looplab/eventhorizon"
)

func GetFileById(id string) (vm FileInfo, has bool, err error) {
    obj, has, err := infrastructure.RepoFac.FilesRepo.GetById(id)
    if err != nil {
        logs.Error("FilesRepo GetById ERR:%v", err)
        return
    }
    if !has {
        return
    }

    vm.ContentType = obj.ContentType
    vm.FileName = obj.FileName
    vm.FilePath = obj.FilePath
    vm.Size = obj.Size
    return
}

func AddFile(vm AddFileForm) (fileId string, err error) {
    fileInfo := files.FileValue{}
    f, err := vm.UpFile.Open()
    if err != nil {
        return
    }
    defer f.Close()
    fileInfo.FileBody, err = ioutil.ReadAll(f)
    if err != nil {
        return
    }
    fileInfo.ContentType = vm.UpFile.Header.Get("Content-Type")
    fileInfo.Size = int(vm.UpFile.Size)
    fileInfo.FileName = vm.UpFile.Filename

    fileId, err = files.SingleFilesAgg.AddNewFile(fileInfo)
    if err != nil {
        logs.Error("SingleFilesAgg AddFile ERR:%v", err.Error())
        return
    }
    return
}

func AddFileCommand(vm AddFileForm) (err error) {
    // looplab框架根据命令类型，创建命令实例
    cmd, err := eh.CreateCommand(files.AddFileCmdType)
    if err != nil {
        err = fmt.Errorf("could not create command: %w", err)
        return err
    }

    // 命令进行校验
    if vldt, ok := cmd.(ddd.Validator); ok {
        err = vldt.Verify()
        if err != nil {
            return
        }
    }
    // 通过命令总线，将命令发布出去，至于谁订阅了该命令则不关心
    if err = bus.HandleCommand(context.Background(), cmd); err != nil {
        return err
    }
    return
}
