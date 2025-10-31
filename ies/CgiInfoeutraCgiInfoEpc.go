package ies

// CGI-InfoEUTRA-cgi-info-EPC ::= SEQUENCE
type CgiInfoeutraCgiInfoEpc struct {
	CgiInfoEpcLegacy CellaccessrelatedinfoEutraEpc
	CgiInfoEpcList   *[]CellaccessrelatedinfoEutraEpc `lb:1,ub:maxPLMN`
}
