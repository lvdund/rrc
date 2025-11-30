package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MCCH_Config_r17 struct {
	Mcch_RepetitionPeriodAndOffset_r17 MCCH_RepetitionPeriodAndOffset_r17          `madatory`
	Mcch_WindowStartSlot_r17           int64                                       `lb:0,ub:79,madatory`
	Mcch_WindowDuration_r17            *MCCH_Config_r17_mcch_WindowDuration_r17    `optional`
	Mcch_ModificationPeriod_r17        MCCH_Config_r17_mcch_ModificationPeriod_r17 `madatory`
}

func (ie *MCCH_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mcch_WindowDuration_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Mcch_RepetitionPeriodAndOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mcch_RepetitionPeriodAndOffset_r17", err)
	}
	if err = w.WriteInteger(ie.Mcch_WindowStartSlot_r17, &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
		return utils.WrapError("WriteInteger Mcch_WindowStartSlot_r17", err)
	}
	if ie.Mcch_WindowDuration_r17 != nil {
		if err = ie.Mcch_WindowDuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mcch_WindowDuration_r17", err)
		}
	}
	if err = ie.Mcch_ModificationPeriod_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mcch_ModificationPeriod_r17", err)
	}
	return nil
}

func (ie *MCCH_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var Mcch_WindowDuration_r17Present bool
	if Mcch_WindowDuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Mcch_RepetitionPeriodAndOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mcch_RepetitionPeriodAndOffset_r17", err)
	}
	var tmp_int_Mcch_WindowStartSlot_r17 int64
	if tmp_int_Mcch_WindowStartSlot_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
		return utils.WrapError("ReadInteger Mcch_WindowStartSlot_r17", err)
	}
	ie.Mcch_WindowStartSlot_r17 = tmp_int_Mcch_WindowStartSlot_r17
	if Mcch_WindowDuration_r17Present {
		ie.Mcch_WindowDuration_r17 = new(MCCH_Config_r17_mcch_WindowDuration_r17)
		if err = ie.Mcch_WindowDuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mcch_WindowDuration_r17", err)
		}
	}
	if err = ie.Mcch_ModificationPeriod_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mcch_ModificationPeriod_r17", err)
	}
	return nil
}
