package ies

import "rrc/utils"

// BW-Preference-r14 ::= SEQUENCE
type BwPreferenceR14 struct {
	DlPreferenceR14 *utils.ENUMERATED
	UlPreferenceR14 *utils.ENUMERATED
}
