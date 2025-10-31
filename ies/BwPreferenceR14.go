package ies

// BW-Preference-r14 ::= SEQUENCE
type BwPreferenceR14 struct {
	DlPreferenceR14 *BwPreferenceR14DlPreferenceR14
	UlPreferenceR14 *BwPreferenceR14UlPreferenceR14
}
