package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSchedulingInfo_r16 struct {
	OffsetToSI_Used_r16       *PosSchedulingInfo_r16_offsetToSI_Used_r16      `optional`
	PosSI_Periodicity_r16     PosSchedulingInfo_r16_posSI_Periodicity_r16     `madatory`
	PosSI_BroadcastStatus_r16 PosSchedulingInfo_r16_posSI_BroadcastStatus_r16 `madatory`
	PosSIB_MappingInfo_r16    PosSIB_MappingInfo_r16                          `madatory`
}

func (ie *PosSchedulingInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OffsetToSI_Used_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OffsetToSI_Used_r16 != nil {
		if err = ie.OffsetToSI_Used_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetToSI_Used_r16", err)
		}
	}
	if err = ie.PosSI_Periodicity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSI_Periodicity_r16", err)
	}
	if err = ie.PosSI_BroadcastStatus_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSI_BroadcastStatus_r16", err)
	}
	if err = ie.PosSIB_MappingInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSIB_MappingInfo_r16", err)
	}
	return nil
}

func (ie *PosSchedulingInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var OffsetToSI_Used_r16Present bool
	if OffsetToSI_Used_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OffsetToSI_Used_r16Present {
		ie.OffsetToSI_Used_r16 = new(PosSchedulingInfo_r16_offsetToSI_Used_r16)
		if err = ie.OffsetToSI_Used_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetToSI_Used_r16", err)
		}
	}
	if err = ie.PosSI_Periodicity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PosSI_Periodicity_r16", err)
	}
	if err = ie.PosSI_BroadcastStatus_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PosSI_BroadcastStatus_r16", err)
	}
	if err = ie.PosSIB_MappingInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PosSIB_MappingInfo_r16", err)
	}
	return nil
}
