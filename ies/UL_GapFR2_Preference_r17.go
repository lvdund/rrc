package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_GapFR2_Preference_r17 struct {
	Ul_GapFR2_PatternPreference_r17 *int64 `lb:0,ub:3,optional`
}

func (ie *UL_GapFR2_Preference_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_GapFR2_PatternPreference_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_GapFR2_PatternPreference_r17 != nil {
		if err = w.WriteInteger(*ie.Ul_GapFR2_PatternPreference_r17, &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode Ul_GapFR2_PatternPreference_r17", err)
		}
	}
	return nil
}

func (ie *UL_GapFR2_Preference_r17) Decode(r *aper.AperReader) error {
	var err error
	var Ul_GapFR2_PatternPreference_r17Present bool
	if Ul_GapFR2_PatternPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_GapFR2_PatternPreference_r17Present {
		var tmp_int_Ul_GapFR2_PatternPreference_r17 int64
		if tmp_int_Ul_GapFR2_PatternPreference_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Ul_GapFR2_PatternPreference_r17", err)
		}
		ie.Ul_GapFR2_PatternPreference_r17 = &tmp_int_Ul_GapFR2_PatternPreference_r17
	}
	return nil
}
