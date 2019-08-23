/*
从原生工程文件加载
*/
package main

import (
	"fmt"
	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/config"
	. "github.com/rekrad/behavior3go/core"
	. "github.com/rekrad/behavior3go/examples/share"
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


func main() {
	projectConfig, ok := LoadRawProjectCfg("example.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))
	maps.Register("TestRunner", new(TestRunner))

	var firstTree *BehaviorTree
	//载入
	for _, v := range projectConfig.Data.Trees {
		tree := CreateBevTreeFromConfig(&v, maps)
		tree.Print()
		if firstTree == nil {
			firstTree = tree
		}
	}

	//输入板
	board := NewBlackboard()
	//循环每一帧
	for i := 0; i < 5; i++ {
		firstTree.Tick(i, board)
	}
}
