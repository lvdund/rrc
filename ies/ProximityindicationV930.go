package ies

import "rrc/utils"

// ProximityIndication-v930-IEs ::= SEQUENCE
type ProximityindicationV930 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *ProximityindicationV930IesNoncriticalextension
}
