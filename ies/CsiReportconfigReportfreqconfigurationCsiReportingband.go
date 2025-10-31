package ies

import "rrc/utils"

// CSI-ReportConfig-reportFreqConfiguration-csi-ReportingBand ::= CHOICE
// Extensible
const (
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceNothing = iota
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands3
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands4
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands5
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands6
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands7
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands8
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands9
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands10
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands11
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands12
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands13
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands14
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands15
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands16
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands17
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands18
	CsiReportconfigReportfreqconfigurationCsiReportingbandChoiceSubbands19V1530
)

type CsiReportconfigReportfreqconfigurationCsiReportingband struct {
	Choice          uint64
	Subbands3       *utils.BITSTRING `lb:3,ub:3`
	Subbands4       *utils.BITSTRING `lb:4,ub:4`
	Subbands5       *utils.BITSTRING `lb:5,ub:5`
	Subbands6       *utils.BITSTRING `lb:6,ub:6`
	Subbands7       *utils.BITSTRING `lb:7,ub:7`
	Subbands8       *utils.BITSTRING `lb:8,ub:8`
	Subbands9       *utils.BITSTRING `lb:9,ub:9`
	Subbands10      *utils.BITSTRING `lb:10,ub:10`
	Subbands11      *utils.BITSTRING `lb:11,ub:11`
	Subbands12      *utils.BITSTRING `lb:12,ub:12`
	Subbands13      *utils.BITSTRING `lb:13,ub:13`
	Subbands14      *utils.BITSTRING `lb:14,ub:14`
	Subbands15      *utils.BITSTRING `lb:15,ub:15`
	Subbands16      *utils.BITSTRING `lb:16,ub:16`
	Subbands17      *utils.BITSTRING `lb:17,ub:17`
	Subbands18      *utils.BITSTRING `lb:18,ub:18`
	Subbands19V1530 *utils.BITSTRING `lb:19,ub:19`
}
