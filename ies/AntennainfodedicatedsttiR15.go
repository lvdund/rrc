package ies

import "rrc/utils"

// AntennaInfoDedicatedSTTI-r15 ::= CHOICE
type AntennainfodedicatedsttiR15 interface {
	isAntennainfodedicatedsttiR15()
}

type AntennainfodedicatedsttiR15Release struct {
	Value struct{}
}

func (*AntennainfodedicatedsttiR15Release) isAntennainfodedicatedsttiR15() {}

type AntennainfodedicatedsttiR15Setup struct {
	Value interface{}
}

func (*AntennainfodedicatedsttiR15Setup) isAntennainfodedicatedsttiR15() {}
