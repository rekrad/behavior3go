package main

import (
	b3 "github.com/rekrad/behavior3go"
	"github.com/rekrad/behavior3go/core"
	// . "github.com/rekrad/behavior3go/config"
	// . "github.com/rekrad/behavior3go/loader"
	// "fmt"
	"log"
)

type TestCondition struct {
	core.Condition
}


func (this *TestCondition) OnTick(tick *core.Tick) b3.Status {
	log.Println("TestCondition.OnTick")
	return b3.SUCCESS
}

