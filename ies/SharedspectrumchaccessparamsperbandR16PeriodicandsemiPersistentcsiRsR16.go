package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-periodicAndSemi-PersistentCSI-RS-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16PeriodicandsemiPersistentcsiRsR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16PeriodicandsemiPersistentcsiRsR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16PeriodicandsemiPersistentcsiRsR16EnumeratedSupported
)
