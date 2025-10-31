package ies

import "rrc/utils"

// UERadioPagingInformation-v15e0-IEs-dl-SchedulingOffset-PDSCH-TypeA-FDD-FR1 ::= ENUMERATED
type UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypeaFddFr1 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypeaFddFr1EnumeratedNothing = iota
	UeradiopaginginformationV15e0IesDlSchedulingoffsetPdschTypeaFddFr1EnumeratedSupported
)
