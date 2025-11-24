package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingInfo struct {
	Si_BroadcastStatus SchedulingInfo_si_BroadcastStatus `madatory`
	Si_Periodicity     SchedulingInfo_si_Periodicity     `madatory`
	Sib_MappingInfo    SIB_Mapping                       `madatory`
}

func (ie *SchedulingInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Si_BroadcastStatus.Encode(w); err != nil {
		return utils.WrapError("Encode Si_BroadcastStatus", err)
	}
	if err = ie.Si_Periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode Si_Periodicity", err)
	}
	if err = ie.Sib_MappingInfo.Encode(w); err != nil {
		return utils.WrapError("Encode Sib_MappingInfo", err)
	}
	return nil
}

func (ie *SchedulingInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Si_BroadcastStatus.Decode(r); err != nil {
		return utils.WrapError("Decode Si_BroadcastStatus", err)
	}
	if err = ie.Si_Periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode Si_Periodicity", err)
	}
	if err = ie.Sib_MappingInfo.Decode(r); err != nil {
		return utils.WrapError("Decode Sib_MappingInfo", err)
	}
	return nil
}
