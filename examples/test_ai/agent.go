package main

import (
	. "github.com/rekrad/behavior3go/core"
	"log"
)

type IAgent interface {
	Patrol()
	Alert()
	Chase()
	RemoteAttack()
	Melee()
	Runaway()
	// 获取最近的敌人的距离
	GetNearestHostileDistance() int
	// 获取警觉距离
	GetAlertRange() int
	// 获取视野距离
	GetSeeFieldRange() int
	// 获取远程攻击距离
	GetRemoteAttackRange() int
	// 获取近战攻击距离
	GetMeleeAttackRange() int
	// 是否低血量
	IsLowHp() bool
}

type agent struct {
	bTree *BehaviorTree
	blackboard *Blackboard

	// 最近敌人的距离
	nearestHostileDistance int
	// 血量
	hp int
}

func (this *agent) Tick() {
	this.bTree.Tick(this, this.blackboard)
}

func (this *agent) Patrol() {
	log.Println("Patrol")
}

func (this *agent) Alert() {
	log.Println("Alert")
}

func (this *agent) Chase() {
	log.Println("Chase")
}

func (this *agent) RemoteAttack() {
	log.Println("RemoteAttack")
}

func (this *agent) Melee() {
	log.Println("Melee")
}

func (this *agent) Runaway() {
	log.Println("Runaway")
}

// 获取最近的敌人的距离
func (this *agent) GetNearestHostileDistance() int {
	return this.nearestHostileDistance
}

// 获取警觉距离
func (this *agent) GetAlertRange() int {
	return 100
}

// 获取视野距离
func (this *agent) GetSeeFieldRange() int {
	return 70
}

// 获取远程攻击距离
func (this *agent) GetRemoteAttackRange() int {
	return 30
}

// 获取近战攻击距离
func (this *agent) GetMeleeAttackRange() int {
	return 8
}

// 是否低血量
func (this *agent) IsLowHp() bool {
	return this.hp < 10
}