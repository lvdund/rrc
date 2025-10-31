package ies

// SL-DiscTxPowerInfoList-r12 ::= SEQUENCE OF SL-DiscTxPowerInfo-r12
// SIZE (maxSL-DiscPowerClass-r12)
type SlDisctxpowerinfolistR12 struct {
	Value []SlDisctxpowerinfoR12 `lb:maxSLDiscpowerclassR12,ub:maxSLDiscpowerclassR12`
}
