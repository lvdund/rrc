package ies

// MobilityControlInfo ::= SEQUENCE
// Extensible
type Mobilitycontrolinfo struct {
	Targetphyscellid           Physcellid
	Carrierfreq                *Carrierfreqeutra
	Carrierbandwidth           *Carrierbandwidtheutra
	Additionalspectrumemission *Additionalspectrumemission
	T304                       MobilitycontrolinfoT304
	NewueIdentity              CRnti
	Radioresourceconfigcommon  Radioresourceconfigcommon
	RachConfigdedicated        *RachConfigdedicated
}
