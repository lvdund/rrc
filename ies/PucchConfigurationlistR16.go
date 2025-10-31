package ies

// PUCCH-ConfigurationList-r16 ::= SEQUENCE OF PUCCH-Config
// SIZE (1..2)
type PucchConfigurationlistR16 struct {
	Value []PucchConfig `lb:1,ub:2`
}
