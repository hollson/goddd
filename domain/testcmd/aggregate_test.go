package testcmd

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hollson/goddd/infrastructure/bus"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"

)

///这种实现，才更契合聚合根的思路！
//聚合根具有全局的唯一标识，而实体只有在聚合内部有唯一的本地标识，值对象没有唯一标识，

func init() {
	SingletonAggTestOnly = &testOnlyAggregate{
		AggregateBase: events.NewAggregateBase("ProcessorAggregate", uuid.Nil),
	}
	bus.SetDealer(SingletonAggTestOnly, TestOnlyCmdType)

}

const ProcessorAggregateType = eh.AggregateType("AggregateType_Processor")

///只有聚合根才可以处理CMD、处理Enevt、激活领域事件。直接持有实体，本身也是实体！

//想更多的表达"继承"建议使用匿名成员。具名成员表示组合。
type testOnlyAggregate struct {
	*events.AggregateBase //DDD框架约束
}

//单例聚合根的特殊实例，ID为NIL
var SingletonAggTestOnly *testOnlyAggregate

//Command异步执行，不需要返回值的
func (a *testOnlyAggregate) HandleCommand(ctx context.Context, cmd eh.Command) (err error) {
	return fmt.Errorf("couldn't handle command")
}

//Command同步执行，需要返回值的
func (a *testOnlyAggregate) DealCommand(ctx context.Context, cmd eh.Command) (interface{}, error) {
	switch cmd := cmd.(type) {
	case *TestOnlyCmd:
		return a.DoCmd(cmd)
	}
	return nil, fmt.Errorf("couldn't Dealer command")
}

func (a *testOnlyAggregate) ApplyEvent(ctx context.Context, event eh.Event) (err error) {

	return
}

func (a *testOnlyAggregate) DoCmd(cmd *TestOnlyCmd) (rst interface{}, err error) {
	rst = "OK"
	switch cmd.CmdType {
	case "fake":
	}

	return
}
