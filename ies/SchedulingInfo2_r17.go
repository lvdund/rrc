package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingInfo2_r17 struct {
	Si_BroadcastStatus_r17 SchedulingInfo2_r17_si_BroadcastStatus_r17 `madatory`
	Si_WindowPosition_r17  int64                                      `lb:1,ub:256,madatory`
	Si_Periodicity_r17     SchedulingInfo2_r17_si_Periodicity_r17     `madatory`
	Sib_MappingInfo_r17    SIB_Mapping_v1700                          `madatory`
}

func (ie *SchedulingInfo2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Si_BroadcastStatus_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Si_BroadcastStatus_r17", err)
	}
	if err = w.WriteInteger(ie.Si_WindowPosition_r17, &uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger Si_WindowPosition_r17", err)
	}
	if err = ie.Si_Periodicity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Si_Periodicity_r17", err)
	}
	if err = ie.Sib_MappingInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sib_MappingInfo_r17", err)
	}
	return nil
}

func (ie *SchedulingInfo2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Si_BroadcastStatus_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Si_BroadcastStatus_r17", err)
	}
	var tmp_int_Si_WindowPosition_r17 int64
	if tmp_int_Si_WindowPosition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger Si_WindowPosition_r17", err)
	}
	ie.Si_WindowPosition_r17 = tmp_int_Si_WindowPosition_r17
	if err = ie.Si_Periodicity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Si_Periodicity_r17", err)
	}
	if err = ie.Sib_MappingInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sib_MappingInfo_r17", err)
	}
	return nil
}
