package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hollson/goddd/application"
	"github.com/hollson/goddd/infrastructure/errorext"
)

func GetFileHandler(c *gin.Context) {
	rst, has, err := application.GetFileById(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	if !has {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", rst.FileName))
	c.Header("Content-Type", rst.ContentType)
	c.File(rst.FilePath)
}

func AddFileHandler(c *gin.Context) {
	var parm application.AddFileForm
	if err := c.ShouldBind(&parm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return
	}
	if parm.UpFile == nil {
		c.Error(errorext.NewCodeError(101, "文件无效", nil))
		return
	}
	fileId, err := application.AddFile(parm)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, fileId)
}
