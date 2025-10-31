package ies

import "rrc/utils"

// SL-UE-SelectedConfig-r16 ::= SEQUENCE
// Extensible
type SlUeSelectedconfigR16 struct {
	SlPsschTxconfiglistR16     *SlPsschTxconfiglistR16
	SlProbresourcekeepR16      *SlUeSelectedconfigR16SlProbresourcekeepR16
	SlReselectafterR16         *SlUeSelectedconfigR16SlReselectafterR16
	SlCbrCommontxconfiglistR16 *SlCbrCommontxconfiglistR16
	UlPrioritizationthresR16   *utils.INTEGER `lb:0,ub:16`
	SlPrioritizationthresR16   *utils.INTEGER `lb:0,ub:8`
}
