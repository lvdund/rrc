package ies

import "rrc/utils"

// CSI-RS-Resource-Mobility-frequencyDomainAllocation ::= CHOICE
const (
	CsiRsResourceMobilityFrequencydomainallocationChoiceNothing = iota
	CsiRsResourceMobilityFrequencydomainallocationChoiceRow1
	CsiRsResourceMobilityFrequencydomainallocationChoiceRow2
)

type CsiRsResourceMobilityFrequencydomainallocation struct {
	Choice uint64
	Row1   *utils.BITSTRING `lb:4,ub:4`
	Row2   *utils.BITSTRING `lb:12,ub:12`
}
