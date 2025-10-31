package ies

import "rrc/utils"

// MeasGapConfig ::= SEQUENCE
// Extensible
type Measgapconfig struct {
	Gapfr2                              *utils.Setuprelease[Gapconfig]
	Gapfr1                              *utils.Setuprelease[Gapconfig]       `ext`
	Gapue                               *utils.Setuprelease[Gapconfig]       `ext`
	GaptoaddmodlistR17                  *[]GapconfigR17                      `lb:1,ub:maxNrofGapIdR17,ext`
	GaptoreleaselistR17                 *[]MeasgapidR17                      `lb:1,ub:maxNrofGapIdR17,ext`
	PosmeasgappreconfigtoaddmodlistR17  *PosmeasgappreconfigtoaddmodlistR17  `ext`
	PosmeasgappreconfigtoreleaselistR17 *PosmeasgappreconfigtoreleaselistR17 `ext`
}
