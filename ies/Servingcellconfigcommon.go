package ies

import "rrc/utils"

// ServingCellConfigCommon ::= SEQUENCE
// Extensible
type Servingcellconfigcommon struct {
	Physcellid                    *Physcellid
	Downlinkconfigcommon          *Downlinkconfigcommon
	Uplinkconfigcommon            *Uplinkconfigcommon
	Supplementaryuplinkconfig     *Uplinkconfigcommon
	NTimingadvanceoffset          *ServingcellconfigcommonNTimingadvanceoffset
	SsbPositionsinburst           *ServingcellconfigcommonSsbPositionsinburst
	SsbPeriodicityservingcell     *ServingcellconfigcommonSsbPeriodicityservingcell
	DmrsTypeaPosition             ServingcellconfigcommonDmrsTypeaPosition
	LteCrsTomatcharound           *utils.Setuprelease[RatematchpatternlteCrs]
	Ratematchpatterntoaddmodlist  *[]Ratematchpattern   `lb:1,ub:maxNrofRateMatchPatterns`
	Ratematchpatterntoreleaselist *[]Ratematchpatternid `lb:1,ub:maxNrofRateMatchPatterns`
	Ssbsubcarrierspacing          *Subcarrierspacing
	TddUlDlConfigurationcommon    *TddUlDlConfigcommon
	SsPbchBlockpower              utils.INTEGER                                         `lb:0,ub:50`
	ChannelaccessmodeR16          *ServingcellconfigcommonChannelaccessmodeR16          `ext`
	DiscoveryburstwindowlengthR16 *ServingcellconfigcommonDiscoveryburstwindowlengthR16 `ext`
	SsbPositionqclR16             *SsbPositionqclRelationR16                            `ext`
	HighspeedconfigR16            *HighspeedconfigR16                                   `ext`
	HighspeedconfigV1700          *HighspeedconfigV1700                                 `ext`
	Channelaccessmode2R17         *ServingcellconfigcommonChannelaccessmode2R17         `ext`
	DiscoveryburstwindowlengthR17 *ServingcellconfigcommonDiscoveryburstwindowlengthR17 `ext`
	SsbPositionqclR17             *SsbPositionqclRelationR17                            `ext`
	Highspeedconfigfr2R17         *Highspeedconfigfr2R17                                `ext`
	UplinkconfigcommonV1700       *UplinkconfigcommonV1700                              `ext`
	NtnConfigR17                  *NtnConfigR17                                         `ext`
	FeatureprioritiesR17          *ServingcellconfigcommonFeatureprioritiesR17          `ext`
}
