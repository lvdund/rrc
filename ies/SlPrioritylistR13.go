package ies

// SL-PriorityList-r13 ::= SEQUENCE OF SL-Priority-r13
// SIZE (1..maxSL-Prio-r13)
type SlPrioritylistR13 struct {
	Value []SlPriorityR13 `lb:1,ub:maxSLPrioR13`
}
