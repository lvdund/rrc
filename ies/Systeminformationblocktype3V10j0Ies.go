package ies

import "rrc/utils"

// SystemInformationBlockType3-v10j0-IEs ::= SEQUENCE
type Systeminformationblocktype3V10j0Ies struct {
	FreqbandinfoR10        *NsPmaxlistR10
	MultibandinfolistV10j0 *MultibandinfolistV10j0
	Noncriticalextension   *Systeminformationblocktype3V10l0Ies
}
