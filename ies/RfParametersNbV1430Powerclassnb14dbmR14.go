package ies

import "rrc/utils"

// RF-Parameters-NB-v1430-powerClassNB-14dBm-r14 ::= ENUMERATED
type RfParametersNbV1430Powerclassnb14dbmR14 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersNbV1430Powerclassnb14dbmR14EnumeratedNothing = iota
	RfParametersNbV1430Powerclassnb14dbmR14EnumeratedSupported
)
