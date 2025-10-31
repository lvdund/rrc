package ies

import "rrc/utils"

// NTN-Config-r17 ::= SEQUENCE
// Extensible
type NtnConfigR17 struct {
	EpochtimeR17                 *EpochtimeR17
	NtnUlsyncvaliditydurationR17 *NtnConfigR17NtnUlsyncvaliditydurationR17
	CellspecifickoffsetR17       *utils.INTEGER `lb:0,ub:1023`
	KmacR17                      *utils.INTEGER `lb:0,ub:512`
	TaInfoR17                    *TaInfoR17
	NtnPolarizationdlR17         *NtnConfigR17NtnPolarizationdlR17
	NtnPolarizationulR17         *NtnConfigR17NtnPolarizationulR17
	EphemerisinfoR17             *EphemerisinfoR17
	TaReportR17                  *NtnConfigR17TaReportR17
}
