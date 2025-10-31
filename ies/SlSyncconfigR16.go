package ies

import "rrc/utils"

// SL-SyncConfig-r16 ::= SEQUENCE
// Extensible
type SlSyncconfigR16 struct {
	SlSyncrefminhystR16     *SlSyncconfigR16SlSyncrefminhystR16
	SlSyncrefdiffhystR16    *SlSyncconfigR16SlSyncrefdiffhystR16
	SlFiltercoefficientR16  *Filtercoefficient
	SlSsbTimeallocation1R16 *SlSsbTimeallocationR16
	SlSsbTimeallocation2R16 *SlSsbTimeallocationR16
	SlSsbTimeallocation3R16 *SlSsbTimeallocationR16
	SlSsidR16               *utils.INTEGER `lb:0,ub:671`
	TxparametersR16         *SlSyncconfigR16TxparametersR16
	GnssSyncR16             *utils.BOOLEAN
}
