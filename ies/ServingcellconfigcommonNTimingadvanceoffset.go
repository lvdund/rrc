package ies

import "rrc/utils"

// ServingCellConfigCommon-n-TimingAdvanceOffset ::= ENUMERATED
type ServingcellconfigcommonNTimingadvanceoffset struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigcommonNTimingadvanceoffsetEnumeratedNothing = iota
	ServingcellconfigcommonNTimingadvanceoffsetEnumeratedN0
	ServingcellconfigcommonNTimingadvanceoffsetEnumeratedN25600
	ServingcellconfigcommonNTimingadvanceoffsetEnumeratedN39936
)
