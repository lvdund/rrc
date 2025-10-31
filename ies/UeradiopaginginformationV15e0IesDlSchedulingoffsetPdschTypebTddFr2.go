package ies

import "rrc/utils"

// UERadioPagingInformation-v15e0-IEs-dl-SchedulingOffset-PDSCH-TypeB-TDD-FR2 ::= ENUMERATED
type UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebTddFr2 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebTddFr2EnumeratedNothing = iota
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebTddFr2EnumeratedSupported
)
