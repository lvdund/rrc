package ies

import "rrc/utils"

// UEInformationRequest-v1530-IEs ::= SEQUENCE
type UeinformationrequestV1530Ies struct {
	IdlemodemeasurementreqR15 *utils.ENUMERATED
	FlightpathinforeqR15      *FlightpathinforeportconfigR15
	Noncriticalextension      *interface{}
}
