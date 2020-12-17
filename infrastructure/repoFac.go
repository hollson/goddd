package infrastructure

import (
	"github.com/hollson/goddd/infrastructure/repos_mysql"
	"github.com/hollson/goddd/infrastructure/repos_redis"
	"github.com/hollson/goddd/interfaces"
)

func init() {
	RepoFac.CaptchaRepo = repos_redis.NewcaptchaRepo()

	RepoFac.FilesRepo = repos_mysql.NewfileRepo()
}

var (
	RepoFac *RepoFactory = &RepoFactory{}
	Empty   interface{}  = struct{}{}
)

type RepoFactory struct {
	CaptchaRepo interfaces.ICaptchaRepo
	FilesRepo   interfaces.IFileInfoRepo
}
