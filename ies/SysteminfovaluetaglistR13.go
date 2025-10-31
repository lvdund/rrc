package ies

// SystemInfoValueTagList-r13 ::= SEQUENCE OF SystemInfoValueTagSI-r13
// SIZE (1..maxSI-Message)
type SysteminfovaluetaglistR13 struct {
	Value []SysteminfovaluetagsiR13 `lb:1,ub:maxSIMessage`
}
