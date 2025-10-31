package ies

// SystemInformationBlockType5-v10j0-IEs ::= SEQUENCE
type Systeminformationblocktype5V10j0 struct {
	InterfreqcarrierfreqlistV10j0 *[]InterfreqcarrierfreqinfoV10j0 `lb:1,ub:maxFreq`
	Noncriticalextension          *Systeminformationblocktype5V10l0
}
