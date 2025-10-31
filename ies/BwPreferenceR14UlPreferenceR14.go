package ies

import "rrc/utils"

// BW-Preference-r14-ul-Preference-r14 ::= ENUMERATED
type BwPreferenceR14UlPreferenceR14 struct {
	Value utils.ENUMERATED
}

const (
	BwPreferenceR14UlPreferenceR14EnumeratedNothing = iota
	BwPreferenceR14UlPreferenceR14EnumeratedMhz1dot4
	BwPreferenceR14UlPreferenceR14EnumeratedMhz5
)
