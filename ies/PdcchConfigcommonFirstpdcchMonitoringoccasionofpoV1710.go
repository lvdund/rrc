package ies

import "rrc/utils"

// PDCCH-ConfigCommon-firstPDCCH-MonitoringOccasionOfPO-v1710 ::= CHOICE
const (
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoV1710ChoiceNothing = iota
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoV1710ChoiceScs480khzoneeightht
	PdcchConfigcommonFirstpdcchMonitoringoccasionofpoV1710ChoiceScs480khzonesixteentht
)

type PdcchConfigcommonFirstpdcchMonitoringoccasionofpoV1710 struct {
	Choice                 uint64
	Scs480khzoneeightht    *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs480khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
}
