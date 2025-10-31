package ies

// SL-BWP-DiscPoolConfig-r17 ::= SEQUENCE
type SlBwpDiscpoolconfigR17 struct {
	SlDiscrxpoolR17           *[]SlResourcepoolR16 `lb:1,ub:maxNrofRXPoolR16`
	SlDisctxpoolselectedR17   *SlTxpooldedicatedR16
	SlDisctxpoolschedulingR17 *SlTxpooldedicatedR16
}
