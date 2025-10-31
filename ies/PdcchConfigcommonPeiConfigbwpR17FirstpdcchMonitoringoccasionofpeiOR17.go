package ies

import "rrc/utils"

// PDCCH-ConfigCommon-pei-ConfigBWP-r17-firstPDCCH-MonitoringOccasionOfPEI-O-r17 ::= CHOICE
const (
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceNothing = iota
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs15khzonet
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs30khzonetScs15khzhalft
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs60khzonetScs30khzhalftScs15khzquartert
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs120khzonetScs60khzhalftScs30khzquartertScs15khzoneeightht
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs120khzhalftScs60khzquartertScs30khzoneeighthtScs15khzonesixteentht
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs480khzonetScs120khzquartertScs60khzoneeighthtScs30khzonesixteentht
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs480khzhalftScs120khzoneeighthtScs60khzonesixteentht
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs480khzquartertScs120khzonesixteentht
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs480khzoneeightht
	PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17ChoiceScs480khzonesixteentht
)

type PdcchConfigcommonPeiConfigbwpR17FirstpdcchMonitoringoccasionofpeiOR17 struct {
	Choice                                                                uint64
	Scs15khzonet                                                          *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs30khzonetScs15khzhalft                                             *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs60khzonetScs30khzhalftScs15khzquartert                             *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs120khzonetScs60khzhalftScs30khzquartertScs15khzoneeightht          *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs120khzhalftScs60khzquartertScs30khzoneeighthtScs15khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs480khzonetScs120khzquartertScs60khzoneeighthtScs30khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs480khzhalftScs120khzoneeighthtScs60khzonesixteentht                *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs480khzquartertScs120khzonesixteentht                               *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs480khzoneeightht                                                   *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
	Scs480khzonesixteentht                                                *[]utils.INTEGER `lb:1,ub:maxPEIPerpfR17`
}
