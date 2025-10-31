package ies

// AC-BarringPerPLMN-List-r12 ::= SEQUENCE OF AC-BarringPerPLMN-r12
// SIZE (1.. maxPLMN-r11)
type AcBarringperplmnListR12 struct {
	Value []AcBarringperplmnR12 `lb:1,ub:maxPLMNR11`
}
