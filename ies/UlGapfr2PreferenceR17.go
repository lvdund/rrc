package ies

import "rrc/utils"

// UL-GapFR2-Preference-r17 ::= SEQUENCE
type UlGapfr2PreferenceR17 struct {
	UlGapfr2PatternpreferenceR17 *utils.INTEGER `lb:0,ub:3`
}
