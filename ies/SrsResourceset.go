package ies

import "rrc/utils"

// SRS-ResourceSet ::= SEQUENCE
// Extensible
type SrsResourceset struct {
	SrsResourcesetid                SrsResourcesetid
	SrsResourceidlist               *[]SrsResourceid            `lb:1,ub:maxNrofSRSResourcesperset`
	Resourcetype                    *SrsResourcesetResourcetype `ext`
	Usage                           SrsResourcesetUsage
	Alpha                           *Alpha
	P0                              *utils.INTEGER `lb:0,ub:24`
	Pathlossreferencers             *PathlossreferencersConfig
	SrsPowercontroladjustmentstates *SrsResourcesetSrsPowercontroladjustmentstates
	PathlossreferencerslistR16      *utils.Setuprelease[PathlossreferencerslistR16] `ext`
	UsagepdcR17                     *utils.BOOLEAN                                  `ext`
	AvailableslotoffsetlistR17      *[]AvailableslotoffsetR17                       `lb:1,ub:4,ext`
	FollowunifiedtciStatesrsR17     *SrsResourcesetFollowunifiedtciStatesrsR17      `ext`
}
