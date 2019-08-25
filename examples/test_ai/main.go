package main

import (
	"fmt"
	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/config"
	. "github.com/rekrad/behavior3go/core"
	// . "github.com/rekrad/behavior3go/examples/share"
	. "github.com/rekrad/behavior3go/loader"
)

// 自定义action节点
type TestRunner struct {
	Action
	info string
	info2 int
	info3 string
	info4 int
}

func (this *TestRunner) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("prob1")
	this.info2 = setting.GetPropertyAsInt("prob2")
}

func (this *TestRunner) OnTick(tick *Tick) b3.Status {
	fmt.Printf("TestRunner.OnTick: prob1: %v prob2 %v p1: %v p2: %v\n", this.info, this.info2, this.info3, this.info4)
	return b3.SUCCESS
}


// 根据行为树名称创建一个agent
func CreateNpcAgent(btName string, projectConfig *RawProjectCfg) *agent {
	agent := &agent{ blackboard: NewBlackboard()}
	maps := b3.NewRegisterStructMaps()

	
	// CreateAgentAction("Patrol", agent, maps)
	// CreateAgentAction("Alert", agent, maps)
	// CreateAgentAction("Chase", agent, maps)
	// CreateAgentAction("RemoteAttack", agent, maps)
	// CreateAgentAction("Melee", agent, maps)
	// CreateAgentAction("Runaway", agent, maps)
	// CreateAgentAction("GreaterThan", agent, maps)

	// 初始化并注册自定义节点
	patrol := &patrol{}
	patrol.Init(agent)
	maps.Register("Patrol", patrol)

	alert := &alert{}
	alert.Init(agent)
	maps.Register("Alert", alert)

	chase := &chase{}
	chase.Init(agent)
	maps.Register("Chase", chase)

	remoteAttack := &remoteAttack{}
	remoteAttack.Init(agent)
	maps.Register("RemoteAttack", remoteAttack)

	melee := &melee{}
	melee.Init(agent)
	maps.Register("Melee", melee)

	runaway := &runaway{}
	runaway.Init(agent)
	maps.Register("Runaway", runaway)

	greaterThan := &greaterThan{}
	greaterThan.Init(agent)
	maps.Register("GreaterThan", greaterThan)

	// 根据树名加载行为树
	var bTree *BehaviorTree
	for _, v := range projectConfig.Data.Trees {
		if v.Title == btName {
			bTree = CreateBevTreeFromConfig(&v, maps)
			bTree.Print()
			break
		}
	}

	// 设置agen的行为树
	agent.bTree = bTree

	return agent
}


func main() {
	projectConfig, ok := LoadRawProjectCfg("test_ai.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	agent := CreateNpcAgent("HostileNPC", projectConfig)
	for i := 0; i < 10; i++ {
		agent.Tick()
	}

	// //自定义节点注册
	// maps := b3.NewRegisterStructMaps()
	// maps.Register("Log", new(LogTest))
	// maps.Register("TestRunner", new(TestRunner))

	// var firstTree *BehaviorTree
	// //载入
	// for _, v := range projectConfig.Data.Trees {
	// 	tree := CreateBevTreeFromConfig(&v, maps)
	// 	tree.Print()
	// 	if firstTree == nil {
	// 		firstTree = tree
	// 	}
	// }

	// //输入板
	// board := NewBlackboard()
	// //循环每一帧
	// for i := 0; i < 5; i++ {
	// 	firstTree.Tick(i, board)
	// }
}
