package ies

import "rrc/utils"

// SystemInformationBlockType1-v10j0-IEs ::= SEQUENCE
type Systeminformationblocktype1V10j0Ies struct {
	FreqbandinfoR10        *NsPmaxlistR10
	MultibandinfolistV10j0 *MultibandinfolistV10j0
	Noncriticalextension   *Systeminformationblocktype1V10l0Ies
}
