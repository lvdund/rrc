package ies

import "rrc/utils"

// PhyLayerParameters-v1430-alternativeTBS-Index-r14 ::= ENUMERATED
type PhylayerparametersV1430AlternativetbsIndexR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430AlternativetbsIndexR14EnumeratedNothing = iota
	PhylayerparametersV1430AlternativetbsIndexR14EnumeratedSupported
)
