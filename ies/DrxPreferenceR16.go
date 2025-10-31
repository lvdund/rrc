package ies

import "rrc/utils"

// DRX-Preference-r16 ::= SEQUENCE
type DrxPreferenceR16 struct {
	PreferreddrxInactivitytimerR16 *DrxPreferenceR16PreferreddrxInactivitytimerR16
	PreferreddrxLongcycleR16       *DrxPreferenceR16PreferreddrxLongcycleR16
	PreferreddrxShortcycleR16      *DrxPreferenceR16PreferreddrxShortcycleR16
	PreferreddrxShortcycletimerR16 *utils.INTEGER `lb:0,ub:16`
}
