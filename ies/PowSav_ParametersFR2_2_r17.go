package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_ParametersFR2_2_r17 struct {
	MaxBW_Preference_r17        *PowSav_ParametersFR2_2_r17_maxBW_Preference_r17        `optional`
	MaxMIMO_LayerPreference_r17 *PowSav_ParametersFR2_2_r17_maxMIMO_LayerPreference_r17 `optional`
}

func (ie *PowSav_ParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxBW_Preference_r17 != nil, ie.MaxMIMO_LayerPreference_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxBW_Preference_r17 != nil {
		if err = ie.MaxBW_Preference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBW_Preference_r17", err)
		}
	}
	if ie.MaxMIMO_LayerPreference_r17 != nil {
		if err = ie.MaxMIMO_LayerPreference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxMIMO_LayerPreference_r17", err)
		}
	}
	return nil
}

func (ie *PowSav_ParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var MaxBW_Preference_r17Present bool
	if MaxBW_Preference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMIMO_LayerPreference_r17Present bool
	if MaxMIMO_LayerPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxBW_Preference_r17Present {
		ie.MaxBW_Preference_r17 = new(PowSav_ParametersFR2_2_r17_maxBW_Preference_r17)
		if err = ie.MaxBW_Preference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBW_Preference_r17", err)
		}
	}
	if MaxMIMO_LayerPreference_r17Present {
		ie.MaxMIMO_LayerPreference_r17 = new(PowSav_ParametersFR2_2_r17_maxMIMO_LayerPreference_r17)
		if err = ie.MaxMIMO_LayerPreference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxMIMO_LayerPreference_r17", err)
		}
	}
	return nil
}
