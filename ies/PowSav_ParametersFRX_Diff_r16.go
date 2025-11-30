package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_ParametersFRX_Diff_r16 struct {
	MaxBW_Preference_r16        *PowSav_ParametersFRX_Diff_r16_maxBW_Preference_r16        `optional`
	MaxMIMO_LayerPreference_r16 *PowSav_ParametersFRX_Diff_r16_maxMIMO_LayerPreference_r16 `optional`
}

func (ie *PowSav_ParametersFRX_Diff_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxBW_Preference_r16 != nil, ie.MaxMIMO_LayerPreference_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxBW_Preference_r16 != nil {
		if err = ie.MaxBW_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBW_Preference_r16", err)
		}
	}
	if ie.MaxMIMO_LayerPreference_r16 != nil {
		if err = ie.MaxMIMO_LayerPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxMIMO_LayerPreference_r16", err)
		}
	}
	return nil
}

func (ie *PowSav_ParametersFRX_Diff_r16) Decode(r *aper.AperReader) error {
	var err error
	var MaxBW_Preference_r16Present bool
	if MaxBW_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMIMO_LayerPreference_r16Present bool
	if MaxMIMO_LayerPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxBW_Preference_r16Present {
		ie.MaxBW_Preference_r16 = new(PowSav_ParametersFRX_Diff_r16_maxBW_Preference_r16)
		if err = ie.MaxBW_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBW_Preference_r16", err)
		}
	}
	if MaxMIMO_LayerPreference_r16Present {
		ie.MaxMIMO_LayerPreference_r16 = new(PowSav_ParametersFRX_Diff_r16_maxMIMO_LayerPreference_r16)
		if err = ie.MaxMIMO_LayerPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxMIMO_LayerPreference_r16", err)
		}
	}
	return nil
}
