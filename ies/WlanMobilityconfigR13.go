package ies

// WLAN-MobilityConfig-r13 ::= SEQUENCE
// Extensible
type WlanMobilityconfigR13 struct {
	WlanToreleaselistR13      *WlanIdListR13
	WlanToaddlistR13          *WlanIdListR13
	AssociationtimerR13       *WlanMobilityconfigR13AssociationtimerR13
	SuccessreportrequestedR13 *bool
}
