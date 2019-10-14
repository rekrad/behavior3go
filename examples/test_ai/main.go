package main

import (
	"fmt"
	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/config"
	. "github.com/rekrad/behavior3go/core"
	// . "github.com/rekrad/behavior3go/examples/share"
	. "github.com/rekrad/behavior3go/loader"
	"log"
	"time"
)

func registerTestNode() *b3.RegisterStructMaps {
	maps := b3.NewRegisterStructMaps()
	maps.Register("Patrol", &patrol{})
	maps.Register("Alert", &alert{})
	maps.Register("Chase", &chase{})
	maps.Register("RemoteAttack", &remoteAttack{})
	maps.Register("Melee", &melee{})
	maps.Register("Runaway", &runaway{})
	maps.Register("GreaterThan", &greaterThan{})
	maps.Register("HostileNotNearby", &HostileNotNearby{})
	maps.Register("HostileAlert", &HostileAlert{})
	maps.Register("HostileCanBeSeen", &HostileCanBeSeen{})
	maps.Register("HostileInRemoteAttackRange", &HostileInRemoteAttackRange{})
	maps.Register("HostileInMeleeRange", &HostileInMeleeRange{})
	maps.Register("LowHp", &LowHp{})
	maps.Register("DebugLog", &DebugLog{})
	maps.Register("MeleeAttack1", &MeleeAttack1{})
	maps.Register("MeleeAttack2", &MeleeAttack2{})
	maps.Register("MeleeAttack3", &MeleeAttack3{})
	maps.Register("TestRunner1", &TestRunner1{})
	maps.Register("TestRunner2", &TestRunner2{})
	maps.Register("TestCondition", &TestCondition{})
	maps.Register("HitA", &HitA{})
	maps.Register("HitB", &HitB{})
	maps.Register("HitC", &HitC{})
	maps.Register("HitD", &HitD{})

	return maps
}

func createAgent(btree *BehaviorTree) *agent {
	agent := &agent{ 
		blackboard: NewBlackboard(),
		bTree: btree,
	}
	return agent
}

// 根据行为树名称创建一个agent
func CreateAgentFromProjectCfg(btName string, projectConfig *RawProjectCfg) *agent {
	// 注册自定义节点类型 
	maps := registerTestNode()
	// 根据树名加载行为树
	var bTree *BehaviorTree
	for _, v := range projectConfig.Data.Trees {
		if v.Title == btName {
			bTree = CreateBevTreeFromConfig(&v, maps)
			bTree.Print()
			break
		}
	}
	// 创建agent
	agent := createAgent(bTree)
	agent.nearestHostileDistance = 115
	agent.hp = 100

	return agent
}

func TestProjectCfgTree(projectCfg string, treename string) {
	projectConfig, ok := LoadRawProjectCfg(projectCfg)
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}
	agent := CreateAgentFromProjectCfg(treename, projectConfig)
	for i := 0; i < 30; i++ {
		log.Println("i: ", i)
		agent.Tick()
		agent.hp -= 4
		if agent.nearestHostileDistance > 0 {
			agent.nearestHostileDistance -= 5
			if agent.nearestHostileDistance < 0 {
				agent.nearestHostileDistance = 0
			}
		}
	}
}

func TestTreeCfg(treecfg string) {
	treeConfig, ok := LoadTreeCfg(treecfg)
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}
	//自定义节点注册
	maps := registerTestNode()
	//载入
	btree := CreateBevTreeFromConfig(treeConfig, maps)
	btree.Print()
	agent := createAgent(btree)
	
	//每隔100ms循环一帧
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 100)
		agent.Tick()
	}
}

func main() {
	TestProjectCfgTree("test_ai.b3", "HostileNPC")
	TestTreeCfg("conhit.b3")
}
