package ies

// SystemInformationBlockType5-v8h0-IEs ::= SEQUENCE
type Systeminformationblocktype5V8h0 struct {
	InterfreqcarrierfreqlistV8h0 *[]InterfreqcarrierfreqinfoV8h0 `lb:1,ub:maxFreq`
	Noncriticalextension         *Systeminformationblocktype5V9e0
}
