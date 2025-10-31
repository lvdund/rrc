package ies

// SL-TxResourceReq-v1700 ::= SEQUENCE
// Extensible
type SlTxresourcereqV1700 struct {
	SlDrxInfofromrxlistR17 *[]SlDrxConfigucSemistaticR17 `lb:1,ub:maxNrofSLRxinfosetR17`
	SlDrxIndicationR17     *SlTxresourcereqV1700SlDrxIndicationR17
}
