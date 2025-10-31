package ies

import "rrc/utils"

// SL-ConfigCommonNR-r16-t400-r16 ::= ENUMERATED
type SlConfigcommonnrR16T400R16 struct {
	Value utils.ENUMERATED
}

const (
	SlConfigcommonnrR16T400R16EnumeratedNothing = iota
	SlConfigcommonnrR16T400R16EnumeratedMs100
	SlConfigcommonnrR16T400R16EnumeratedMs200
	SlConfigcommonnrR16T400R16EnumeratedMs300
	SlConfigcommonnrR16T400R16EnumeratedMs400
	SlConfigcommonnrR16T400R16EnumeratedMs600
	SlConfigcommonnrR16T400R16EnumeratedMs1000
	SlConfigcommonnrR16T400R16EnumeratedMs1500
	SlConfigcommonnrR16T400R16EnumeratedMs2000
)
