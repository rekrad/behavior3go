package main

import (
	"fmt"
	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/config"
	. "github.com/rekrad/behavior3go/core"
	// . "github.com/rekrad/behavior3go/examples/share"
	. "github.com/rekrad/behavior3go/loader"
	_ "log"
)

// 根据行为树名称创建一个agent
func CreateNpcAgent(btName string, projectConfig *RawProjectCfg) *agent {
	agent := &agent{ blackboard: NewBlackboard()}
	agent.nearestHostileDistance = 115
	agent.hp = 100
	
	// 注册自定义节点类型 
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

	// agent := CreateNpcAgent("HostileNPC", projectConfig)
	// for i := 0; i < 30; i++ {
	// 	log.Println("i: ", i)
	// 	agent.Tick()
	// 	agent.hp -= 4
	// 	if agent.nearestHostileDistance > 0 {
	// 		agent.nearestHostileDistance -= 5
	// 		if agent.nearestHostileDistance < 0 {
	// 			agent.nearestHostileDistance = 0
	// 		}
	// 	}
	// }

	agent2 := CreateNpcAgent("TestTree", projectConfig)
	for i := 0; i < 15; i++ {
		agent2.Tick()
	}
}
