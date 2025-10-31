package ies

// SystemInformationBlockType3-v10l0-IEs ::= SEQUENCE
type Systeminformationblocktype3V10l0 struct {
	FreqbandinfoV10l0      *NsPmaxlistV10l0
	MultibandinfolistV10l0 *MultibandinfolistV10l0
	Noncriticalextension   *Systeminformationblocktype3V10l0IesNoncriticalextension
}
