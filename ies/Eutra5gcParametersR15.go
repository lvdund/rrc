package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-r15 ::= SEQUENCE
type Eutra5gcParametersR15 struct {
	Eutra5gcR15                      *utils.ENUMERATED
	EutraEpcHoEutra5gcR15            *utils.ENUMERATED
	HoEutra5gcFddTddR15              *utils.ENUMERATED
	HoInterfreqeutra5gcR15           *utils.ENUMERATED
	ImsVoiceovermcgBearereutra5gcR15 *utils.ENUMERATED
	InactivestateR15                 *utils.ENUMERATED
	ReflectiveqosR15                 *utils.ENUMERATED
}
