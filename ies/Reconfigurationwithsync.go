package ies

// ReconfigurationWithSync ::= SEQUENCE
// Extensible
type Reconfigurationwithsync struct {
	Spcellconfigcommon       *Servingcellconfigcommon
	NewueIdentity            RntiValue
	T304                     ReconfigurationwithsyncT304
	RachConfigdedicated      *ReconfigurationwithsyncRachConfigdedicated
	Smtc                     *SsbMtc                   `ext`
	DapsUplinkpowerconfigR16 *DapsUplinkpowerconfigR16 `ext`
	SlPathswitchconfigR17    *SlPathswitchconfigR17    `ext`
}
