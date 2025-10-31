package ies

// SystemInformationBlockType5-v9e0-IEs ::= SEQUENCE
type Systeminformationblocktype5V9e0 struct {
	InterfreqcarrierfreqlistV9e0 *[]InterfreqcarrierfreqinfoV9e0 `lb:1,ub:maxFreq`
	Noncriticalextension         *Systeminformationblocktype5V10j0
}
