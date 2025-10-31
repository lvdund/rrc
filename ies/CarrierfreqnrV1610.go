package ies

// CarrierFreqNR-v1610 ::= SEQUENCE
type CarrierfreqnrV1610 struct {
	Smtc2LpR16                *MtcSsb2LpNrR16
	SsbPositionqclCommonnrR16 *SsbPositionqclRelationnrR16
	WhitecelllistnrR16        *WhitecelllistnrR16
	HighspeedcarriernrR16     *bool
}
