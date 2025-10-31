package ies

import "rrc/utils"

// PCCH-Config ::= SEQUENCE
// Extensible
type PcchConfig struct {
	Defaultpagingcycle                       Pagingcycle
	Nandpagingframeoffset                    PcchConfigNandpagingframeoffset
	Ns                                       PcchConfigNs
	FirstpdcchMonitoringoccasionofpo         *PcchConfigFirstpdcchMonitoringoccasionofpo
	NrofpdcchMonitoringoccasionperssbInpoR16 *utils.INTEGER                                   `lb:0,ub:4,ext`
	RanpaginginidlepoR17                     *utils.BOOLEAN                                   `ext`
	FirstpdcchMonitoringoccasionofpoV1710    *PcchConfigFirstpdcchMonitoringoccasionofpoV1710 `ext`
}
