package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CI_ConfigurationPerServingCell_r16 struct {
	ServingCellId                    ServCellIndex                                                        `madatory`
	PositionInDCI_r16                int64                                                                `lb:0,ub:maxCI_DCI_PayloadSize_1_r16,madatory`
	PositionInDCI_ForSUL_r16         *int64                                                               `lb:0,ub:maxCI_DCI_PayloadSize_1_r16,optional`
	Ci_PayloadSize_r16               CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16                `madatory`
	TimeFrequencyRegion_r16          *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16          `lb:0,ub:37949,optional`
	UplinkCancellationPriority_v1610 *CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610 `optional,ext`
}

func (ie *CI_ConfigurationPerServingCell_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PositionInDCI_ForSUL_r16 != nil, ie.TimeFrequencyRegion_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode ServingCellId", err)
	}
	if err = w.WriteInteger(ie.PositionInDCI_r16, &aper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
		return utils.WrapError("WriteInteger PositionInDCI_r16", err)
	}
	if ie.PositionInDCI_ForSUL_r16 != nil {
		if err = w.WriteInteger(*ie.PositionInDCI_ForSUL_r16, &aper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Encode PositionInDCI_ForSUL_r16", err)
		}
	}
	if err = ie.Ci_PayloadSize_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ci_PayloadSize_r16", err)
	}
	if ie.TimeFrequencyRegion_r16 != nil {
		if err = ie.TimeFrequencyRegion_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TimeFrequencyRegion_r16", err)
		}
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16) Decode(r *aper.AperReader) error {
	var err error
	var PositionInDCI_ForSUL_r16Present bool
	if PositionInDCI_ForSUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeFrequencyRegion_r16Present bool
	if TimeFrequencyRegion_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode ServingCellId", err)
	}
	var tmp_int_PositionInDCI_r16 int64
	if tmp_int_PositionInDCI_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
		return utils.WrapError("ReadInteger PositionInDCI_r16", err)
	}
	ie.PositionInDCI_r16 = tmp_int_PositionInDCI_r16
	if PositionInDCI_ForSUL_r16Present {
		var tmp_int_PositionInDCI_ForSUL_r16 int64
		if tmp_int_PositionInDCI_ForSUL_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Decode PositionInDCI_ForSUL_r16", err)
		}
		ie.PositionInDCI_ForSUL_r16 = &tmp_int_PositionInDCI_ForSUL_r16
	}
	if err = ie.Ci_PayloadSize_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ci_PayloadSize_r16", err)
	}
	if TimeFrequencyRegion_r16Present {
		ie.TimeFrequencyRegion_r16 = new(CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16)
		if err = ie.TimeFrequencyRegion_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TimeFrequencyRegion_r16", err)
		}
	}
	return nil
}
