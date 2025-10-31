package ies

// SL-TxResourceReqDisc-r17 ::= SEQUENCE
// Extensible
type SlTxresourcereqdiscR17 struct {
	SlDestinationidentitydiscR17  SlDestinationidentityR16
	SlSourceidentityrelayueR17    *SlSourceidentityR17
	SlCasttypediscR17             SlTxresourcereqdiscR17SlCasttypediscR17
	SlTxinterestedfreqlistdiscR17 SlTxinterestedfreqlistR16
	SlTypetxsynclistdiscR17       []SlTypetxsyncR16 `lb:1,ub:maxNrofFreqSLR16`
	SlDiscoverytypeR17            SlTxresourcereqdiscR17SlDiscoverytypeR17
}
