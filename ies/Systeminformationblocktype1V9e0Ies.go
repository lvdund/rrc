package ies

import "rrc/utils"

// SystemInformationBlockType1-v9e0-IEs ::= SEQUENCE
type Systeminformationblocktype1V9e0Ies struct {
	FreqbandindicatorV9e0 *FreqbandindicatorV9e0
	MultibandinfolistV9e0 *MultibandinfolistV9e0
	Noncriticalextension  *Systeminformationblocktype1V10j0Ies
}
