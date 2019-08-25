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
}

type agent struct {
	bTree *BehaviorTree
	blackboard *Blackboard
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