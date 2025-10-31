package ies

// SL-UE-AssistanceInformationNR-r16 ::= SEQUENCE OF SL-TrafficPatternInfo-r16
// SIZE (1..maxNrofTrafficPattern-r16)
type SlUeAssistanceinformationnrR16 struct {
	Value []SlTrafficpatterninfoR16 `lb:1,ub:maxNrofTrafficPatternR16`
}
