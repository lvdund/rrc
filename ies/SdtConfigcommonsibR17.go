package ies

// SDT-ConfigCommonSIB-r17 ::= SEQUENCE
type SdtConfigcommonsibR17 struct {
	SdtRsrpThresholdR17              *RsrpRange
	SdtLogicalchannelsrDelaytimerR17 *SdtConfigcommonsibR17SdtLogicalchannelsrDelaytimerR17
	SdtDatavolumethresholdR17        SdtConfigcommonsibR17SdtDatavolumethresholdR17
	T319aR17                         SdtConfigcommonsibR17T319aR17
}
