package ies

import "rrc/utils"

// UERadioPagingInformation-v1700-IEs-numberOfRxRedCap-r17 ::= ENUMERATED
type UeradiopaginginformationV1700IesNumberofrxredcapR17 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV1700IesNumberofrxredcapR17EnumeratedNothing = iota
	UeradiopaginginformationV1700IesNumberofrxredcapR17EnumeratedOne
	UeradiopaginginformationV1700IesNumberofrxredcapR17EnumeratedTwo
)
