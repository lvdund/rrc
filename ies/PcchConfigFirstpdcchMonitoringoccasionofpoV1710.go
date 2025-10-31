package ies

import "rrc/utils"

// PCCH-Config-firstPDCCH-MonitoringOccasionOfPO-v1710 ::= CHOICE
const (
	PcchConfigFirstpdcchMonitoringoccasionofpoV1710ChoiceNothing = iota
	PcchConfigFirstpdcchMonitoringoccasionofpoV1710ChoiceScs480khzoneeightht
	PcchConfigFirstpdcchMonitoringoccasionofpoV1710ChoiceScs480khzonesixteentht
)

type PcchConfigFirstpdcchMonitoringoccasionofpoV1710 struct {
	Choice                 uint64
	Scs480khzoneeightht    *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
	Scs480khzonesixteentht *[]utils.INTEGER `lb:1,ub:maxPOPerpf`
}
