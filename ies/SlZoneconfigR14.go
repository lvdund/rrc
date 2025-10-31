package ies

import "rrc/utils"

// SL-ZoneConfig-r14 ::= SEQUENCE
type SlZoneconfigR14 struct {
	ZonelengthR14     SlZoneconfigR14ZonelengthR14
	ZonewidthR14      SlZoneconfigR14ZonewidthR14
	ZoneidlongimodR14 utils.INTEGER `lb:0,ub:4`
	ZoneidlatimodR14  utils.INTEGER `lb:0,ub:4`
}
