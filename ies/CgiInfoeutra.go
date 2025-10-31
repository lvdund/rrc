package ies

import "rrc/utils"

// CGI-InfoEUTRA ::= SEQUENCE
type CgiInfoeutra struct {
	CgiInfoEpc                *CgiInfoeutraCgiInfoEpc
	CgiInfo5gc                *[]CellaccessrelatedinfoEutra5gc `lb:1,ub:maxPLMN`
	Freqbandindicator         Freqbandindicatoreutra
	Multibandinfolist         *Multibandinfolisteutra
	Freqbandindicatorpriority *utils.BOOLEAN
}
