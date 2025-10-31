package ies

// MeasObjectWLAN-r13-carrierFreq-r13 ::= CHOICE
const (
	MeasobjectwlanR13CarrierfreqR13ChoiceNothing = iota
	MeasobjectwlanR13CarrierfreqR13ChoiceBandindicatorlistwlanR13
	MeasobjectwlanR13CarrierfreqR13ChoiceCarrierinfolistwlanR13
)

type MeasobjectwlanR13CarrierfreqR13 struct {
	Choice                   uint64
	BandindicatorlistwlanR13 *[]WlanBandindicatorR13 `lb:1,ub:maxWLANBandsR13`
	CarrierinfolistwlanR13   *[]WlanCarrierinfoR13   `lb:1,ub:maxWLANCarrierinfoR13`
}
