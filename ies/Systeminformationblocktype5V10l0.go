package ies

// SystemInformationBlockType5-v10l0-IEs ::= SEQUENCE
type Systeminformationblocktype5V10l0 struct {
	InterfreqcarrierfreqlistV10l0 *[]InterfreqcarrierfreqinfoV10l0 `lb:1,ub:maxFreq`
	Noncriticalextension          *Systeminformationblocktype5V13a0
}
