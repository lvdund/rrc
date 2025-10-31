package ies

// ProximityIndication-r9-IEs ::= SEQUENCE
// Extensible
type ProximityindicationR9 struct {
	TypeR9               ProximityindicationR9IesTypeR9
	CarrierfreqR9        ProximityindicationR9IesCarrierfreqR9
	Noncriticalextension *ProximityindicationV930
}
