package ies

import "rrc/utils"

// SI-OrPSI-GERAN ::= CHOICE
type SiOrpsiGeran interface {
	isSiOrpsiGeran()
}

type SiOrpsiGeranSi struct {
	Value Systeminfolistgeran
}

func (*SiOrpsiGeranSi) isSiOrpsiGeran() {}

type SiOrpsiGeranPsi struct {
	Value Systeminfolistgeran
}

func (*SiOrpsiGeranPsi) isSiOrpsiGeran() {}
