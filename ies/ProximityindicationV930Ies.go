package ies

import "rrc/utils"

// ProximityIndication-v930-IEs ::= SEQUENCE
type ProximityindicationV930Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
