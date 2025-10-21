package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1430-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1430Ies struct {
	SlV2xConfigdedicatedR14                 *SlV2xConfigdedicatedR14
	ScelltoaddmodlistextV1430               *ScelltoaddmodlistextV1430
	PerccGapindicationrequestR14            *utils.ENUMERATED
	Systeminformationblocktype2dedicatedR14 *utils.OCTETSTRING
	Noncriticalextension                    *RrcconnectionreconfigurationV1510Ies
}
