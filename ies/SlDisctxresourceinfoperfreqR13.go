package ies

// SL-DiscTxResourceInfoPerFreq-r13 ::= SEQUENCE
// Extensible
type SlDisctxresourceinfoperfreqR13 struct {
	DisctxcarrierfreqR13         ArfcnValueeutraR9
	DisctxresourcesR13           *SlDisctxresourceR13
	DisctxresourcespsR13         *SlDisctxresourceR13
	DisctxrefcarrierdedicatedR13 *SlDisctxrefcarrierdedicatedR13
	DisccellselectioninfoR13     *CellselectioninfonfreqR13
}
