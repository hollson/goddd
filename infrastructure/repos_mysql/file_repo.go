package repos_mysql

import (
	"time"

	"github.com/hollson/goddd/interfaces"
)

type fileRepo struct {
}

func NewfileRepo() *fileRepo {
	return new(fileRepo)
}

//新增，
func (r *fileRepo) AddObj(obj *interfaces.FileInfo) (num int64, err error) {
	now := time.Now().Unix()
	obj.CreateAt = now
	obj.UpdateAt = now
	_, err = writeEngine.
		Insert(obj)
	return
}

//单条查询
func (r *fileRepo) GetById(id string) (obj interfaces.FileInfo, has bool, err error) {
	has, err = readEngine.
		Where("id=?", id).
		Get(&obj)
	return
}

func (r *fileRepo) Find(parm interfaces.FindParmFiles) (objs []interfaces.FileInfo, total int64, err error) {
	objs = make([]interfaces.FileInfo, 0)
	err = readEngine.
		Desc("created").
		Limit(parm.PageSize, parm.Page*parm.PageSize).
		Find(&objs)
	return
}
