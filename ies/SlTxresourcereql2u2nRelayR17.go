package ies

import "rrc/utils"

// SL-TxResourceReqL2U2N-Relay-r17 ::= SEQUENCE
// Extensible
type SlTxresourcereql2u2nRelayR17 struct {
	SlDestinationidentityl2u2nR17      *SlDestinationidentityR16
	SlTxinterestedfreqlistl2u2nR17     SlTxinterestedfreqlistR16
	SlTypetxsynclistl2u2nR17           []SlTypetxsyncR16 `lb:1,ub:maxNrofFreqSLR16`
	SlLocalidRequestR17                *utils.BOOLEAN
	SlPagingidentityremoteueR17        *SlPagingidentityremoteueR17
	SlCapabilityinformationsidelinkR17 *utils.OCTETSTRING
}
