package ies

import "rrc/utils"

// SidelinkUEInformation-v1310-IEs-discTxResourceReq-v1310 ::= SEQUENCE
type SidelinkueinformationV1310IesDisctxresourcereqV1310 struct {
	CarrierfreqdisctxR13        *utils.INTEGER `lb:0,ub:maxFreq`
	DisctxresourcereqaddfreqR13 *SlDisctxresourcereqperfreqlistR13
}
