package ies

import "rrc/utils"

// FreqHoppingParameters-r13 ::= SEQUENCE
type FreqhoppingparametersR13 struct {
	Dummy                                 *FreqhoppingparametersR13Dummy
	Dummy2                                *FreqhoppingparametersR13Dummy2
	Dummy3                                *FreqhoppingparametersR13Dummy3
	IntervalUlhoppingconfigcommonmodeaR13 *FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodeaR13
	IntervalUlhoppingconfigcommonmodebR13 *FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodebR13
	Dummy4                                *utils.INTEGER `lb:0,ub:maxAvailNarrowBandsR13`
}
