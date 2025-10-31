package ies

import "rrc/utils"

// SL-ScheduledConfig-r16 ::= SEQUENCE
// Extensible
type SlScheduledconfigR16 struct {
	SlRntiR16                      RntiValue
	MacMainconfigslR16             *MacMainconfigslR16
	SlCsRntiR16                    *RntiValue
	SlPsfchTopucchR16              *[]utils.INTEGER `lb:1,ub:8`
	SlConfiguredgrantconfiglistR16 *SlConfiguredgrantconfiglistR16
	SlDciToslTransR16              *[]utils.INTEGER `lb:1,ub:8,ext`
}
