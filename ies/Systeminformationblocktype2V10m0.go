package ies

// SystemInformationBlockType2-v10m0-IEs ::= SEQUENCE
type Systeminformationblocktype2V10m0 struct {
	FreqinfoV10l0          *Systeminformationblocktype2V10m0IesFreqinfoV10l0
	MultibandinfolistV10l0 *[]AdditionalspectrumemissionV10l0 `lb:1,ub:maxMultiBands`
	Noncriticalextension   *Systeminformationblocktype2V10n0
}
