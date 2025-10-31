package ies

import "rrc/utils"

// UERadioPagingInformation-v15e0-IEs-dl-SchedulingOffset-PDSCH-TypeA-TDD-FR2 ::= ENUMERATED
type UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypeaTddFr2 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypeaTddFr2EnumeratedNothing = iota
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypeaTddFr2EnumeratedSupported
)
