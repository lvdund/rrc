package ies

// SL-UE-SelectedConfigRP-r16 ::= SEQUENCE
// Extensible
type SlUeSelectedconfigrpR16 struct {
	SlCbrPrioritytxconfiglistR16   *SlCbrPrioritytxconfiglistR16
	SlThresRsrpListR16             *SlThresRsrpListR16
	SlMultireserveresourceR16      *SlUeSelectedconfigrpR16SlMultireserveresourceR16
	SlMaxnumperreserveR16          *SlUeSelectedconfigrpR16SlMaxnumperreserveR16
	SlSensingwindowR16             *SlUeSelectedconfigrpR16SlSensingwindowR16
	SlSelectionwindowlistR16       *SlSelectionwindowlistR16
	SlResourcereserveperiodlistR16 *[]SlResourcereserveperiodR16 `lb:1,ub:16`
	SlRsForsensingR16              SlUeSelectedconfigrpR16SlRsForsensingR16
	SlCbrPrioritytxconfiglistV1650 *SlCbrPrioritytxconfiglistV1650 `ext`
}
