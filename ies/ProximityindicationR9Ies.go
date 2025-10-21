package ies

import "rrc/utils"

// ProximityIndication-r9-IEs ::= SEQUENCE
// Extensible
type ProximityindicationR9Ies struct {
	TypeR9               utils.ENUMERATED
	CarrierfreqR9        interface{}
	Noncriticalextension *ProximityindicationV930Ies
}
