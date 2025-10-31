package ies

// SystemInformationBlockType6-v8h0-IEs ::= SEQUENCE
type Systeminformationblocktype6V8h0 struct {
	CarrierfreqlistutraFddV8h0 *[]CarrierfreqinfoutraFddV8h0 `lb:1,ub:maxUTRAFddCarrier`
	Noncriticalextension       *Systeminformationblocktype6V8h0IesNoncriticalextension
}
