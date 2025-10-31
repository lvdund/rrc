package ies

// RLC-Parameters ::= SEQUENCE
// Extensible
type RlcParameters struct {
	AmWithshortsn              *RlcParametersAmWithshortsn
	UmWithshortsn              *RlcParametersUmWithshortsn
	UmWithlongsn               *RlcParametersUmWithlongsn
	ExtendedtPollretransmitR16 *RlcParametersExtendedtPollretransmitR16 `ext`
	ExtendedtStatusprohibitR16 *RlcParametersExtendedtStatusprohibitR16 `ext`
	AmWithlongsnRedcapR17      *RlcParametersAmWithlongsnRedcapR17      `ext`
}
