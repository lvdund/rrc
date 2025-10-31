package ies

import "rrc/utils"

// PCCH-Config-firstPDCCH-MonitoringOccasionOfPO ::= CHOICE
const (
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceNothing = iota
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs15khzonet
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs30khzonetScs15khzhalft
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs60khzonetScs30khzhalftScs15khzquartert
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs120khzonetScs60khzhalftScs30khzquartertScs15khzoneeightht
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs120khzhalftScs60khzquartertScs30khzoneeighthtScs15khzonesixteentht
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs480khzonetScs120khzquartertScs60khzoneeighthtScs30khzonesixteentht
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs480khzhalftScs120khzoneeighthtScs60khzonesixteentht
	PcchConfigFirstpdcchMonitoringoccasionofpoChoiceScs480khzquartertScs120khzonesixteentht
)

type PcchConfigFirstpdcchMonitoringoccasionofpo struct {
	Choice                                                                uint64
	Scs15khzonet                                                          *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs30khzonetScs15khzhalft                                             *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs60khzonetScs30khzhalftScs15khzquartert                             *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzonetScs60khzhalftScs30khzquartertScs15khzoneeightht          *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzhalftScs60khzquartertScs30khzoneeighthtScs15khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs480khzonetScs120khzquartertScs60khzoneeighthtScs30khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs480khzhalftScs120khzoneeighthtScs60khzonesixteentht                *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs480khzquartertScs120khzonesixteentht                               *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
}
