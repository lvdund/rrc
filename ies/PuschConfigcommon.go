package ies

import "rrc/utils"

// PUSCH-ConfigCommon ::= SEQUENCE
// Extensible
type PuschConfigcommon struct {
	Grouphoppingenabledtransformprecoding *PuschConfigcommonGrouphoppingenabledtransformprecoding
	PuschTimedomainallocationlist         *PuschTimedomainresourceallocationlist
	Msg3Deltapreamble                     *utils.INTEGER `lb:0,ub:6`
	P0Nominalwithgrant                    *utils.INTEGER `lb:0,ub:24`
}
