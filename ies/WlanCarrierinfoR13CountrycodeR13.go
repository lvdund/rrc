package ies

import "rrc/utils"

// WLAN-CarrierInfo-r13-countryCode-r13 ::= utils.ENUMERATED // Extensible
type WlanCarrierinfoR13CountrycodeR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanCarrierinfoR13CountrycodeR13EnumeratedNothing = iota
	WlanCarrierinfoR13CountrycodeR13EnumeratedUnitedstates
	WlanCarrierinfoR13CountrycodeR13EnumeratedEurope
	WlanCarrierinfoR13CountrycodeR13EnumeratedJapan
	WlanCarrierinfoR13CountrycodeR13EnumeratedGlobal
)
