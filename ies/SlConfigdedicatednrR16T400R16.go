package ies

import "rrc/utils"

// SL-ConfigDedicatedNR-r16-t400-r16 ::= ENUMERATED
type SlConfigdedicatednrR16T400R16 struct {
	Value utils.ENUMERATED
}

const (
	SlConfigdedicatednrR16T400R16EnumeratedNothing = iota
	SlConfigdedicatednrR16T400R16EnumeratedMs100
	SlConfigdedicatednrR16T400R16EnumeratedMs200
	SlConfigdedicatednrR16T400R16EnumeratedMs300
	SlConfigdedicatednrR16T400R16EnumeratedMs400
	SlConfigdedicatednrR16T400R16EnumeratedMs600
	SlConfigdedicatednrR16T400R16EnumeratedMs1000
	SlConfigdedicatednrR16T400R16EnumeratedMs1500
	SlConfigdedicatednrR16T400R16EnumeratedMs2000
)
