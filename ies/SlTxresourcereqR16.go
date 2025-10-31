package ies

import "rrc/utils"

// SL-TxResourceReq-r16 ::= SEQUENCE
type SlTxresourcereqR16 struct {
	SlDestinationidentityR16           SlDestinationidentityR16
	SlCasttypeR16                      SlTxresourcereqR16SlCasttypeR16
	SlRlcModeindicationlistR16         *[]SlRlcModeindicationR16 `lb:1,ub:maxNrofSLRBR16`
	SlQosInfolistR16                   *[]SlQosInfoR16           `lb:1,ub:maxNrofSLQfisperdestR16`
	SlTypetxsynclistR16                *[]SlTypetxsyncR16        `lb:1,ub:maxNrofFreqSLR16`
	SlTxinterestedfreqlistR16          *SlTxinterestedfreqlistR16
	SlCapabilityinformationsidelinkR16 *utils.OCTETSTRING
}
