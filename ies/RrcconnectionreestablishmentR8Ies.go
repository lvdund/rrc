package ies

import "rrc/utils"

// RRCConnectionReestablishment-r8-IEs ::= SEQUENCE
type RrcconnectionreestablishmentR8Ies struct {
	Radioresourceconfigdedicated Radioresourceconfigdedicated
	Nexthopchainingcount         Nexthopchainingcount
	Noncriticalextension         *RrcconnectionreestablishmentV8a0Ies
}
