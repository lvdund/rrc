package ies

import "rrc/utils"

// UERadioPagingInformation-v15e0-IEs-dl-SchedulingOffset-PDSCH-TypeB-TDD-FR1 ::= ENUMERATED
type UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebTddFr1 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebTddFr1EnumeratedNothing = iota
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebTddFr1EnumeratedSupported
)
