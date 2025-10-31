package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1610-eutra-SI-AcquisitionForHO-ENDC-r16 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1610EutraSiAcquisitionforhoEndcR16 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1610EutraSiAcquisitionforhoEndcR16EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1610EutraSiAcquisitionforhoEndcR16EnumeratedSupported
)
