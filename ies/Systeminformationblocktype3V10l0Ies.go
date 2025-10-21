package ies

import "rrc/utils"

// SystemInformationBlockType3-v10l0-IEs ::= SEQUENCE
type Systeminformationblocktype3V10l0Ies struct {
	FreqbandinfoV10l0      *NsPmaxlistV10l0
	MultibandinfolistV10l0 *MultibandinfolistV10l0
	Noncriticalextension   *interface{}
}
