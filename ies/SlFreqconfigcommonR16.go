package ies

import "rrc/utils"

// SL-FreqConfigCommon-r16 ::= SEQUENCE
// Extensible
type SlFreqconfigcommonR16 struct {
	SlScsSpecificcarrierlistR16  []ScsSpecificcarrier `lb:1,ub:maxSCSs`
	SlAbsolutefrequencypointaR16 ArfcnValuenr
	SlAbsolutefrequencyssbR16    *ArfcnValuenr
	Frequencyshift7p5khzslR16    *utils.BOOLEAN
	ValuenR16                    utils.INTEGER           `lb:0,ub:1`
	SlBwpListR16                 *[]SlBwpConfigcommonR16 `lb:1,ub:maxNrofSLBwpsR16`
	SlSyncpriorityR16            *SlFreqconfigcommonR16SlSyncpriorityR16
	SlNbassyncR16                *utils.BOOLEAN
	SlSyncconfiglistR16          *SlSyncconfiglistR16
}
