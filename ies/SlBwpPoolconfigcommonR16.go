package ies

// SL-BWP-PoolConfigCommon-r16 ::= SEQUENCE
type SlBwpPoolconfigcommonR16 struct {
	SlRxpoolR16               *[]SlResourcepoolR16       `lb:1,ub:maxNrofRXPoolR16`
	SlTxpoolselectednormalR16 *[]SlResourcepoolconfigR16 `lb:1,ub:maxNrofTXPoolR16`
	SlTxpoolexceptionalR16    *SlResourcepoolconfigR16
}
