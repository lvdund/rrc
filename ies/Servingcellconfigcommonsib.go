package ies

import "rrc/utils"

// ServingCellConfigCommonSIB ::= SEQUENCE
// Extensible
type Servingcellconfigcommonsib struct {
	Downlinkconfigcommon            Downlinkconfigcommonsib
	Uplinkconfigcommon              *Uplinkconfigcommonsib
	Supplementaryuplink             *Uplinkconfigcommonsib
	NTimingadvanceoffset            *ServingcellconfigcommonsibNTimingadvanceoffset
	SsbPositionsinburst             *ServingcellconfigcommonsibSsbPositionsinburst
	SsbPeriodicityservingcell       ServingcellconfigcommonsibSsbPeriodicityservingcell
	TddUlDlConfigurationcommon      *TddUlDlConfigcommon
	SsPbchBlockpower                utils.INTEGER                                              `lb:0,ub:50`
	ChannelaccessmodeR16            *ServingcellconfigcommonsibChannelaccessmodeR16            `ext`
	DiscoveryburstwindowlengthR16   *ServingcellconfigcommonsibDiscoveryburstwindowlengthR16   `ext`
	HighspeedconfigR16              *HighspeedconfigR16                                        `ext`
	Channelaccessmode2R17           *ServingcellconfigcommonsibChannelaccessmode2R17           `ext`
	DiscoveryburstwindowlengthV1700 *ServingcellconfigcommonsibDiscoveryburstwindowlengthV1700 `ext`
	Highspeedconfigfr2R17           *Highspeedconfigfr2R17                                     `ext`
	UplinkconfigcommonV1700         *UplinkconfigcommonsibV1700                                `ext`
	EnhancedmeasurementleoR17       *utils.BOOLEAN                                             `ext`
}
