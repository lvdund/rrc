package ies

import "rrc/utils"

// T316-r16 ::= ENUMERATED
type T316R16 struct {
	Value utils.ENUMERATED
}

const (
	T316R16EnumeratedNothing = iota
	T316R16EnumeratedMs50
	T316R16EnumeratedMs100
	T316R16EnumeratedMs200
	T316R16EnumeratedMs300
	T316R16EnumeratedMs400
	T316R16EnumeratedMs500
	T316R16EnumeratedMs600
	T316R16EnumeratedMs1000
	T316R16EnumeratedMs1500
	T316R16EnumeratedMs2000
)
