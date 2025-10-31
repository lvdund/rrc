package ies

// SL-BWP-DiscPoolConfigCommon-r17 ::= SEQUENCE
// Extensible
type SlBwpDiscpoolconfigcommonR17 struct {
	SlDiscrxpoolR17         *[]SlResourcepoolR16       `lb:1,ub:maxNrofRXPoolR16`
	SlDisctxpoolselectedR17 *[]SlResourcepoolconfigR16 `lb:1,ub:maxNrofTXPoolR16`
}
