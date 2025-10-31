package ies

// PDCP-Parameters ::= SEQUENCE
// Extensible
type PdcpParameters struct {
	SupportedrohcProfiles            PdcpParametersSupportedrohcProfiles
	MaxnumberrohcContextsessions     PdcpParametersMaxnumberrohcContextsessions
	UplinkonlyrohcProfiles           *PdcpParametersUplinkonlyrohcProfiles
	ContinuerohcContext              *PdcpParametersContinuerohcContext
	Outoforderdelivery               *PdcpParametersOutoforderdelivery
	Shortsn                          *PdcpParametersShortsn
	PdcpDuplicationsrb               *PdcpParametersPdcpDuplicationsrb
	PdcpDuplicationmcgOrscgDrb       *PdcpParametersPdcpDuplicationmcgOrscgDrb
	DrbIabR16                        *PdcpParametersDrbIabR16                        `ext`
	NonDrbIabR16                     *PdcpParametersNonDrbIabR16                     `ext`
	ExtendeddiscardtimerR16          *PdcpParametersExtendeddiscardtimerR16          `ext`
	ContinueehcContextR16            *PdcpParametersContinueehcContextR16            `ext`
	EhcR16                           *PdcpParametersEhcR16                           `ext`
	MaxnumberehcContextsR16          *PdcpParametersMaxnumberehcContextsR16          `ext`
	JointehcRohcConfigR16            *PdcpParametersJointehcRohcConfigR16            `ext`
	PdcpDuplicationmorethantworlcR16 *PdcpParametersPdcpDuplicationmorethantworlcR16 `ext`
	LongsnRedcapR17                  *PdcpParametersLongsnRedcapR17                  `ext`
	UdcR17                           *PdcpParametersUdcR17                           `ext`
}
