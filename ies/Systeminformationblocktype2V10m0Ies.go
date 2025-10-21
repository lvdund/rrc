package ies

import "rrc/utils"

// SystemInformationBlockType2-v10m0-IEs ::= SEQUENCE
type Systeminformationblocktype2V10m0Ies struct {
	FreqinfoV10l0          *interface{}
	MultibandinfolistV10l0 *interface{}
	Noncriticalextension   *Systeminformationblocktype2V10n0Ies
}
