package ies

// SSB-MTC ::= SEQUENCE
type SsbMtc struct {
	Periodicityandoffset SsbMtcPeriodicityandoffset
	Duration             SsbMtcDuration
}
