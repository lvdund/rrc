package ies

import "rrc/utils"

// MeasAndMobParametersCommon-rrm-RelaxationRRC-ConnectedRedCap-r17 ::= ENUMERATED
type MeasandmobparameterscommonRrmRelaxationrrcConnectedredcapR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonRrmRelaxationrrcConnectedredcapR17EnumeratedNothing = iota
	MeasandmobparameterscommonRrmRelaxationrrcConnectedredcapR17EnumeratedSupported
)
