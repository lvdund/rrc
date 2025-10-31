package ies

import "rrc/utils"

// T312-r16 ::= ENUMERATED
type T312R16 struct {
	Value utils.ENUMERATED
}

const (
	T312R16EnumeratedNothing = iota
	T312R16EnumeratedMs0
	T312R16EnumeratedMs50
	T312R16EnumeratedMs100
	T312R16EnumeratedMs200
	T312R16EnumeratedMs300
	T312R16EnumeratedMs400
	T312R16EnumeratedMs500
	T312R16EnumeratedMs1000
)
