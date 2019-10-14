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

// 带有IAgent接口的action
type AgentAction struct {
	Action
	Agent IAgent
}

func (this *AgentAction) OnTick(tick *Tick) {
	if this.Agent == nil {
		this.Agent = tick.GetTarget().(IAgent)
	}
}

// 巡逻
type patrol struct {
	AgentAction
}

func (this *patrol) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	this.Agent.Patrol()
	return b3.SUCCESS
}

// ai察觉
type alert struct {
	AgentAction
}

func (this *alert) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	this.Agent.Alert()
	return b3.FAILURE
}

// 追赶
type chase struct {
	AgentAction
}

func (this *chase) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	this.Agent.Chase()
	return b3.SUCCESS
}

// 远程攻击
type remoteAttack struct {
	AgentAction
}

func (this *remoteAttack) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	this.Agent.RemoteAttack()
	return b3.SUCCESS
}

// 近战攻击
type melee struct {
	AgentAction
}

func (this *melee) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	this.Agent.Melee()
	return b3.SUCCESS
}

// 逃跑
type runaway struct {
	AgentAction
}

func (this *runaway) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
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

// 敌人不在附近
type HostileNotNearby struct {
	AgentAction
}

func (this *HostileNotNearby) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)

	// 获取最近的敌人的距离
	distance := this.Agent.GetNearestHostileDistance()
	// 小于0则表示附近没有敌人
	if distance < 0 {
		return b3.SUCCESS
	}

	// 获取agent的alert范围
	alertRange := this.Agent.GetAlertRange()
	if  distance > alertRange {
		return b3.SUCCESS
	}

	return b3.FAILURE
}

// 警觉
type HostileAlert struct {
	AgentAction
}

func (this *HostileAlert) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)

	// 获取最近的敌人的距离
	distance := this.Agent.GetNearestHostileDistance()
	alertRange := this.Agent.GetAlertRange()
	canBeSeenDistance := this.Agent.GetSeeFieldRange()
	if canBeSeenDistance < distance && distance <= alertRange {
		return b3.SUCCESS
	}

	return b3.FAILURE
}

// 可见
type HostileCanBeSeen struct {
	AgentAction
}

func (this *HostileCanBeSeen) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	distance := this.Agent.GetNearestHostileDistance()
	canBeSeenDistance := this.Agent.GetSeeFieldRange()
	remoteAttackRange := this.Agent.GetRemoteAttackRange()
	if remoteAttackRange < distance && distance <= canBeSeenDistance {
		return b3.SUCCESS
	}

	return b3.FAILURE
}

// 在远程攻击范围
type HostileInRemoteAttackRange struct {
	AgentAction
}

func (this *HostileInRemoteAttackRange) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	distance := this.Agent.GetNearestHostileDistance()
	remoteAttackRange := this.Agent.GetRemoteAttackRange()
	meleeAttackRange := this.Agent.GetMeleeAttackRange()
	if meleeAttackRange < distance && distance <= remoteAttackRange {
		return b3.SUCCESS
	}

	return b3.FAILURE
}

// 在近战攻击范围
type HostileInMeleeRange struct {
	AgentAction
}

func (this *HostileInMeleeRange) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	distance := this.Agent.GetNearestHostileDistance()
	meleeAttackRange := this.Agent.GetMeleeAttackRange()
	if 0 <= distance && distance <= meleeAttackRange {
		return b3.SUCCESS
	}

	return b3.FAILURE
}

// 滴血量
type LowHp struct {
	AgentAction
}

func  (this *LowHp) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	if this.Agent.IsLowHp() {
		return b3.SUCCESS
	}
	return b3.FAILURE
}

type MeleeAttack1 struct {
	AgentAction
}

func (this *MeleeAttack1) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}

type MeleeAttack2 struct {
	AgentAction
}

func (this *MeleeAttack2) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}

type MeleeAttack3 struct {
	AgentAction
}

func (this *MeleeAttack3) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}

type TestRunner1 struct {
	AgentAction
	i int
}

func (this *TestRunner1) OnTick(tick *Tick) b3.Status {
	this.i++
	if this.i == 1 {
		log.Println("i == 1 return b3.SUCCESS")
		return b3.SUCCESS
	} else if this.i == 2 {
		log.Println("i == 2 return b3.FAILURE")
		return b3.FAILURE
	} else if this.i == 3 {
		log.Println("i == 3 return b3.FAILURE")
		return b3.FAILURE
	} else if this.i == 7 {
		log.Println("i == 7 return b3.FAILURE")
		return b3.FAILURE
	} 
	log.Printf("i == %v return b3.SUCCESS\n", this.i)
	return b3.SUCCESS
}

type TestRunner2 struct {
	AgentAction
}

func (this *TestRunner2) OnTick(tick *Tick) b3.Status {
	log.Println("TestRunner2.OnTick")
	return b3.RUNNING
}

type HitA struct {
	AgentAction
}

func (this *HitA) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	log.Println("HitA");
	return b3.SUCCESS
}


type HitB struct {
	AgentAction
}

func (this *HitB) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	log.Println("HitB");
	return b3.SUCCESS
}

type HitC struct {
	AgentAction
}

func (this *HitC) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	log.Println("HitC");
	return b3.SUCCESS
}

type HitD struct {
	AgentAction
}

func (this *HitD) OnTick(tick *Tick) b3.Status {
	this.AgentAction.OnTick(tick)
	log.Println("HitD");
	return b3.SUCCESS
}