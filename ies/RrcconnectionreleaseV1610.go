package ies

// RRCConnectionRelease-v1610-IEs ::= SEQUENCE
type RrcconnectionreleaseV1610 struct {
	FulliRntiR16             *IRntiR15
	ShortiRntiR16            *ShortiRntiR15
	PurConfigR16             *SetupreleasePurConfigR16
	RrcInactiveconfigV1610   *RrcInactiveconfigV1610
	ReleaseidlemeasconfigR16 *bool
	AltfreqprioritiesR16     *bool
	T323R16                  *RrcconnectionreleaseV1610IesT323R16
	Noncriticalextension     *RrcconnectionreleaseV1650
}
