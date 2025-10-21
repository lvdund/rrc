package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo-v9e0 ::= SEQUENCE
type InterfreqcarrierfreqinfoV9e0 struct {
	DlCarrierfreqV9e0     *ArfcnValueeutraV9e0
	MultibandinfolistV9e0 *MultibandinfolistV9e0
}
