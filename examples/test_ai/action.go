package main

import (
	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/core"
	. "github.com/rekrad/behavior3go/config"
	// . "github.com/rekrad/behavior3go/loader"
	// "fmt"
	"log"
)

type IAgentAction interface {
	Init(agent IAgent)
}

type AgentAction struct {
	Action
	Agent IAgent
}

// 初始化
func (this *AgentAction) Init(agent IAgent) {
	this.Agent = agent
}

// 巡逻
type patrol struct {
	AgentAction
}

func (this *patrol) OnTick(tick *Tick) b3.Status {
	//agent := IAgent(tick.GetTarget())
	agent := tick.GetTarget().(IAgent)
	agent.Patrol()
	return b3.SUCCESS
}

// ai察觉
type alert struct {
	AgentAction
}

func (this *alert) OnTick(tick *Tick) b3.Status {
	agent := tick.GetTarget().(IAgent)
	agent.Alert()
	return b3.FAILURE
}

// 追赶
type chase struct {
	AgentAction
}

func (this *chase) OnTick(tick *Tick) b3.Status {
	this.Agent.Chase()
	return b3.SUCCESS
}

// 远程攻击
type remoteAttack struct {
	AgentAction
}

func (this *remoteAttack) OnTick(tick *Tick) b3.Status {
	this.Agent.RemoteAttack()
	return b3.SUCCESS
}

// 近战攻击
type melee struct {
	AgentAction
}

func (this *melee) OnTick(tick *Tick) b3.Status {
	this.Agent.Melee()
	return b3.SUCCESS
}

// 逃跑
type runaway struct {
	AgentAction
}

func (this *runaway) OnTick(tick *Tick) b3.Status {
	this.Agent.Runaway()
	return b3.SUCCESS
}

// 大于
type greaterThan struct {
	AgentAction
	leftOpVal string
	rightOpVal string
}

func (this *greaterThan) Initialize(setting *BTNodeCfg) {
	this.AgentAction.Initialize(setting)
	this.leftOpVal = setting.GetPropertyAsString("LeftOp")
	this.rightOpVal = setting.GetPropertyAsString("RightOp")
}

func (this *greaterThan) OnTick(tick *Tick) b3.Status {
	// if (this.leftOpVal > this.rightOpVal) {
	// 	return b3.SUCCESS
	// } else {
	// 	return b3.FAILURE
	// }
	return b3.SUCCESS
}

// 输出日志
type DebugLog struct {
	AgentAction
	info string
}

func (this *DebugLog) Initialize(setting *BTNodeCfg) {
	this.AgentAction.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *DebugLog) OnTick(tick *Tick) b3.Status {
	log.Println("DebugLog:", this.info)
	return b3.SUCCESS
}
