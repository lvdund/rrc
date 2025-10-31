package ies

import "rrc/utils"

// SupportedUDC-r15-supportedStandardDic-r15 ::= ENUMERATED
type SupportedudcR15SupportedstandarddicR15 struct {
	Value utils.ENUMERATED
}

const (
	SupportedudcR15SupportedstandarddicR15EnumeratedNothing = iota
	SupportedudcR15SupportedstandarddicR15EnumeratedSupported
)
