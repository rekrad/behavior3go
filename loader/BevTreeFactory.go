package loader

import (
	_ "fmt"
	_ "reflect"

	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/actions"
	. "github.com/rekrad/behavior3go/composites"
	. "github.com/rekrad/behavior3go/config"
	. "github.com/rekrad/behavior3go/core"
	. "github.com/rekrad/behavior3go/decorators"
)

var globalBaseMap *b3.RegisterStructMaps

func init() {
	globalBaseMap = createBaseStructMaps()
}

func createBaseStructMaps() *b3.RegisterStructMaps {
	st := b3.NewRegisterStructMaps()
	//actions
	st.Register("Error", &Error{})
	st.Register("Failer", &Failer{})
	st.Register("Runner", &Runner{})
	st.Register("Succeeder", &Succeeder{})
	st.Register("Wait", &Wait{})
	st.Register("Log", &Log{})
	//composites
	st.Register("MemPriority", &MemPriority{})
	st.Register("MemSequence", &MemSequence{})
	st.Register("Priority", &Priority{})
	st.Register("Sequence", &Sequence{})

	//decorators
	st.Register("Inverter", &Inverter{})
	st.Register("Limiter", &Limiter{})
	st.Register("MaxTime", &MaxTime{})
	st.Register("Repeater", &Repeater{})
	st.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})
	return st
}

func CreateBevTreeFromConfig(config *BTTreeCfg, extMap *b3.RegisterStructMaps) *BehaviorTree {
	baseMaps := createBaseStructMaps()
	tree := NewBeTree()
	tree.Load(config, baseMaps, extMap)
	return tree
}

// 这个方法不再每次都进行createBaseStructMaps的操作
func CreateBevTreeFromConfigEx(config *BTTreeCfg, extMap *b3.RegisterStructMaps) *BehaviorTree {
	tree := NewBeTree()
	tree.Load(config, globalBaseMap, extMap)
	return tree
}
