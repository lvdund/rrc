package ies

import "rrc/utils"

// SPS-ConfigUL ::= CHOICE
// Extensible
type SpsConfigul interface {
	isSpsConfigul()
}

type SpsConfigulRelease struct {
	Value struct{}
}

func (*SpsConfigulRelease) isSpsConfigul() {}

type SpsConfigulSetup struct {
	Value interface{}
}

func (*SpsConfigulSetup) isSpsConfigul() {}
