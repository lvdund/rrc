package ies

import "rrc/utils"

// PhyLayerParameters-v1280-alternativeTBS-Indices-r12 ::= ENUMERATED
type PhylayerparametersV1280AlternativetbsIndicesR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1280AlternativetbsIndicesR12EnumeratedNothing = iota
	PhylayerparametersV1280AlternativetbsIndicesR12EnumeratedSupported
)
