package ies

import "rrc/utils"

// UERadioPagingInformation-v15e0-IEs-dl-SchedulingOffset-PDSCH-TypeB-FDD-FR1 ::= ENUMERATED
type UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebFddFr1 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebFddFr1EnumeratedNothing = iota
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypebFddFr1EnumeratedSupported
)
