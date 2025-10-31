package ies

// SL-DiscTxInfoInterFreqListAdd-r13 ::= SEQUENCE
// Extensible
type SlDisctxinfointerfreqlistaddR13 struct {
	DisctxfreqtoaddmodlistR13  *[]SlDisctxresourceinfoperfreqR13 `lb:1,ub:maxFreq`
	DisctxfreqtoreleaselistR13 *[]ArfcnValueeutraR9              `lb:1,ub:maxFreq`
}
