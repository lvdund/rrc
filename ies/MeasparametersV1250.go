package ies

import "rrc/utils"

// MeasParameters-v1250 ::= SEQUENCE
type MeasparametersV1250 struct {
	Timert312R12                 *utils.ENUMERATED
	AlternativetimetotriggerR12  *utils.ENUMERATED
	IncmoneutraR12               *utils.ENUMERATED
	IncmonutraR12                *utils.ENUMERATED
	ExtendedmaxmeasidR12         *utils.ENUMERATED
	ExtendedrsrqLowerrangeR12    *utils.ENUMERATED
	RsrqOnallsymbolsR12          *utils.ENUMERATED
	CrsDiscoverysignalsmeasR12   *utils.ENUMERATED
	CsiRsDiscoverysignalsmeasR12 *utils.ENUMERATED
}
