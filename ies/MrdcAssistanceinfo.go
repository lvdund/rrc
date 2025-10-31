package ies

import "rrc/utils"

// MRDC-AssistanceInfo ::= SEQUENCE
// Extensible
type MrdcAssistanceinfo struct {
	Affectedcarrierfreqcombinfolistmrdc []Affectedcarrierfreqcombinfomrdc `lb:1,ub:maxNrofCombIDC`
	OverheatingassistancescgR16         *utils.OCTETSTRING                `ext`
	OverheatingassistancescgFr22R17     *utils.OCTETSTRING                `ext`
}
