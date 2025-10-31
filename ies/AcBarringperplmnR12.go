package ies

import "rrc/utils"

// AC-BarringPerPLMN-r12 ::= SEQUENCE
type AcBarringperplmnR12 struct {
	PlmnIdentityindexR12          utils.INTEGER `lb:0,ub:maxPLMNR11`
	AcBarringinfoR12              *AcBarringperplmnR12AcBarringinfoR12
	AcBarringskipformmtelvoiceR12 *bool
	AcBarringskipformmtelvideoR12 *bool
	AcBarringskipforsmsR12        *bool
	AcBarringforcsfbR12           *AcBarringconfig
	SsacBarringformmtelVoiceR12   *AcBarringconfig
	SsacBarringformmtelVideoR12   *AcBarringconfig
}
