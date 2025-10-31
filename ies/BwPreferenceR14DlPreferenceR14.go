package ies

import "rrc/utils"

// BW-Preference-r14-dl-Preference-r14 ::= ENUMERATED
type BwPreferenceR14DlPreferenceR14 struct {
	Value utils.ENUMERATED
}

const (
	BwPreferenceR14DlPreferenceR14EnumeratedNothing = iota
	BwPreferenceR14DlPreferenceR14EnumeratedMhz1dot4
	BwPreferenceR14DlPreferenceR14EnumeratedMhz5
	BwPreferenceR14DlPreferenceR14EnumeratedMhz20
)
