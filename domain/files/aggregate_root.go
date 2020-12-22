package files

import (
    "context"
    "fmt"

    "github.com/google/uuid"
    "github.com/hollson/goddd/infrastructure/bus"
    "github.com/hollson/goddd/infrastructure/logs"
    "github.com/hollson/goddd/interfaces"
    eh "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/aggregatestore/events"
)

// 单例聚合根的特殊实例，ID为NIL
var SingleFilesAgg *Agg

func init() {
    sthAgg := &Agg{
        AggregateBase: events.NewAggregateBase(AgentAggregateType, uuid.Nil),
    }
    bus.RegisterHandler(AddFileCmdType, sthAgg)

    SingleFilesAgg = sthAgg
}

// 想更多的表达"继承"建议使用匿名成员。具名成员表示组合。
type Agg struct {
    *events.AggregateBase // DDD框架约束
}

// Command异步执行，不需要返回值的
func (a *Agg) HandleCommand(ctx context.Context, cmd eh.Command) (err error) {
    switch cmd := cmd.(type) {
    case *AddFileCmd:
        vo := FileValue{
            FileInfo: interfaces.FileInfo{
                Id:          uuid.New().String(),
                FileName:    cmd.FileName,
                Size:        cmd.Size,
                ContentType: cmd.ContentType,
            },
        }
        en := newEntity(vo)
        err = en.AddFile(cmd.FileBody)
        if err != nil {
            logs.Error("新增文件出错：%s ", err.Error())
        }
    default:
        err = fmt.Errorf("couldn't handle command")
    }
    return
}

func (a *Agg) ApplyEvent(ctx context.Context, event eh.Event) (err error) {

    return
}

// Command同步执行，需要返回值的
func (a *Agg) DealCommand(ctx context.Context, cmd eh.Command) (interface{}, error) {
    return nil, fmt.Errorf("couldn't Dealer command")
}

// ///聚合根对外开放的能力

func (a *Agg) AddNewFile(vo FileValue) (fileId string, err error) {
    fileId = uuid.New().String()
    vo.Id = fileId
    en := newEntity(vo)
    err = en.AddFile(vo.FileBody)
    return
}
