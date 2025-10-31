package ies

import "rrc/utils"

// CSI-RS-ResourceMapping-frequencyDomainAllocation ::= CHOICE
const (
	CsiRsResourcemappingFrequencydomainallocationChoiceNothing = iota
	CsiRsResourcemappingFrequencydomainallocationChoiceRow1
	CsiRsResourcemappingFrequencydomainallocationChoiceRow2
	CsiRsResourcemappingFrequencydomainallocationChoiceRow4
	CsiRsResourcemappingFrequencydomainallocationChoiceOther
)

type CsiRsResourcemappingFrequencydomainallocation struct {
	Choice uint64
	Row1   *utils.BITSTRING `lb:4,ub:4`
	Row2   *utils.BITSTRING `lb:12,ub:12`
	Row4   *utils.BITSTRING `lb:3,ub:3`
	Other  *utils.BITSTRING `lb:6,ub:6`
}
