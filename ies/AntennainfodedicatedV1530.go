package ies

import "rrc/utils"

// AntennaInfoDedicated-v1530 ::= CHOICE
type AntennainfodedicatedV1530 interface {
	isAntennainfodedicatedV1530()
}

type AntennainfodedicatedV1530Release struct {
	Value struct{}
}

func (*AntennainfodedicatedV1530Release) isAntennainfodedicatedV1530() {}

type AntennainfodedicatedV1530Setup struct {
	Value interface{}
}

func (*AntennainfodedicatedV1530Setup) isAntennainfodedicatedV1530() {}
