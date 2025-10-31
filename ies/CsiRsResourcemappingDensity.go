package ies

import "rrc/utils"

// CSI-RS-ResourceMapping-density ::= CHOICE
const (
	CsiRsResourcemappingDensityChoiceNothing = iota
	CsiRsResourcemappingDensityChoiceDot5
	CsiRsResourcemappingDensityChoiceOne
	CsiRsResourcemappingDensityChoiceThree
	CsiRsResourcemappingDensityChoiceSpare
)

type CsiRsResourcemappingDensity struct {
	Choice uint64
	Dot5   *utils.ENUMERATED
	One    *struct{}
	Three  *struct{}
	Spare  *struct{}
}
