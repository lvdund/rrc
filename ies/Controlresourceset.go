package ies

import "rrc/utils"

// ControlResourceSet ::= SEQUENCE
// Extensible
type Controlresourceset struct {
	Controlresourcesetid        Controlresourcesetid
	Frequencydomainresources    utils.BITSTRING `lb:45,ub:45`
	Duration                    utils.INTEGER   `lb:0,ub:maxCoReSetDuration`
	CceRegMappingtype           *ControlresourcesetCceRegMappingtype
	Precodergranularity         ControlresourcesetPrecodergranularity
	TciStatespdcchToaddlist     *[]TciStateid `lb:1,ub:maxNrofTCIStatespdcch`
	TciStatespdcchToreleaselist *[]TciStateid `lb:1,ub:maxNrofTCIStatespdcch`
	TciPresentindci             *ControlresourcesetTciPresentindci
	PdcchDmrsScramblingid       *utils.INTEGER                              `lb:0,ub:65535`
	RbOffsetR16                 *utils.INTEGER                              `lb:0,ub:5,ext`
	TciPresentdci12R16          *utils.INTEGER                              `lb:0,ub:3,ext`
	CoresetpoolindexR16         *utils.INTEGER                              `lb:0,ub:1,ext`
	ControlresourcesetidV1610   *ControlresourcesetidV1610                  `ext`
	FollowunifiedtciStateR17    *ControlresourcesetFollowunifiedtciStateR17 `ext`
}
