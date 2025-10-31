package ies

import "rrc/utils"

// CGI-InfoEUTRALogging ::= SEQUENCE
type CgiInfoeutralogging struct {
	PlmnIdentityEutra5gc     *PlmnIdentity
	TrackingareacodeEutra5gc *Trackingareacode
	CellidentityEutra5gc     *utils.BITSTRING `lb:28,ub:28`
	PlmnIdentityEutraEpc     *PlmnIdentity
	TrackingareacodeEutraEpc *utils.BITSTRING `lb:16,ub:16`
	CellidentityEutraEpc     *utils.BITSTRING `lb:28,ub:28`
}
