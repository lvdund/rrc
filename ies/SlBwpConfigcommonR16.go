package ies

// SL-BWP-ConfigCommon-r16 ::= SEQUENCE
// Extensible
type SlBwpConfigcommonR16 struct {
	SlBwpGenericR16              *SlBwpGenericR16
	SlBwpPoolconfigcommonR16     *SlBwpPoolconfigcommonR16
	SlBwpPoolconfigcommonpsR17   *SlBwpPoolconfigcommonR16     `ext`
	SlBwpDiscpoolconfigcommonR17 *SlBwpDiscpoolconfigcommonR17 `ext`
}
