package ies

import "rrc/utils"

// SystemInformationBlockType1-v10l0-IEs ::= SEQUENCE
type Systeminformationblocktype1V10l0Ies struct {
	FreqbandinfoV10l0      *NsPmaxlistV10l0
	MultibandinfolistV10l0 *MultibandinfolistV10l0
	Noncriticalextension   *Systeminformationblocktype1V10x0Ies
}
