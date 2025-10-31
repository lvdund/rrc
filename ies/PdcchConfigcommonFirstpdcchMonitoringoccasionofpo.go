package ies

import "rrc/utils"

// PDCCH-ConfigCommon-firstPDCCH-MonitoringOccasionOfPO ::= CHOICE
const (
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceNothing = iota
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs15khzonet
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs30khzonetScs15khzhalft
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs60khzonetScs30khzhalftScs15khzquartert
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs120khzonetScs60khzhalftScs30khzquartertScs15khzoneeightht
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs120khzhalftScs60khzquartertScs30khzoneeighthtScs15khzonesixteentht
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs120khzquartertScs60khzoneeighthtScs30khzonesixteentht
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs120khzoneeighthtScs60khzonesixteentht
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoChoiceScs120khzonesixteentht
)

type PdcchConfigcommonFirstpdcchMonitoringoccasionofpo struct {
	Choice                                                                uint64
	Scs15khzonet                                                          *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs30khzonetScs15khzhalft                                             *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs60khzonetScs30khzhalftScs15khzquartert                             *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzonetScs60khzhalftScs30khzquartertScs15khzoneeightht          *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzhalftScs60khzquartertScs30khzoneeighthtScs15khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzquartertScs60khzoneeighthtScs30khzonesixteentht              *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzoneeighthtScs60khzonesixteentht                              *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs120khzonesixteentht                                                *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
}
