package ies

import "rrc/utils"

// SL-FreqConfig-r16 ::= SEQUENCE
type SlFreqconfigR16 struct {
	SlFreqIdR16                  SlFreqIdR16
	SlScsSpecificcarrierlistR16  []ScsSpecificcarrier `lb:1,ub:maxSCSs`
	SlAbsolutefrequencypointaR16 *ArfcnValuenr
	SlAbsolutefrequencyssbR16    *ArfcnValuenr
	Frequencyshift7p5khzslR16    *utils.BOOLEAN
	ValuenR16                    utils.INTEGER     `lb:0,ub:1`
	SlBwpToreleaselistR16        *[]BwpId          `lb:1,ub:maxNrofSLBwpsR16`
	SlBwpToaddmodlistR16         *[]SlBwpConfigR16 `lb:1,ub:maxNrofSLBwpsR16`
	SlSyncconfiglistR16          *SlSyncconfiglistR16
	SlSyncpriorityR16            *SlFreqconfigR16SlSyncpriorityR16
}
