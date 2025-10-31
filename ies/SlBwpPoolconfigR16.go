package ies

// SL-BWP-PoolConfig-r16 ::= SEQUENCE
type SlBwpPoolconfigR16 struct {
	SlRxpoolR16               *[]SlResourcepoolR16 `lb:1,ub:maxNrofRXPoolR16`
	SlTxpoolselectednormalR16 *SlTxpooldedicatedR16
	SlTxpoolschedulingR16     *SlTxpooldedicatedR16
	SlTxpoolexceptionalR16    *SlResourcepoolconfigR16
}
