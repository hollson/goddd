package application

import "mime/multipart"

// 应用层的struct ,主要是为了适配表现层的数据格式。正常情况下，领域模型是无法直接暴露给用户直接使用的。

type FileInfo struct {
    FileName    string `json:"file_name" form:"file_name"`
    FilePath    string `json:"file_path" form:"file_path"`
    ContentType string `json:"content_type" form:"content_type"`
    Size        int    `json:"size" form:"size"`
}

type AddFileForm struct {
    UpFile *multipart.FileHeader `json:"up_file" form:"up_file"`
    Remark string                `json:"remark" form:"remark"`
}
