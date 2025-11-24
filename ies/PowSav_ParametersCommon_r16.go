package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_ParametersCommon_r16 struct {
	Drx_Preference_r16                *PowSav_ParametersCommon_r16_drx_Preference_r16                `optional`
	MaxCC_Preference_r16              *PowSav_ParametersCommon_r16_maxCC_Preference_r16              `optional`
	ReleasePreference_r16             *PowSav_ParametersCommon_r16_releasePreference_r16             `optional`
	MinSchedulingOffsetPreference_r16 *PowSav_ParametersCommon_r16_minSchedulingOffsetPreference_r16 `optional`
}

func (ie *PowSav_ParametersCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Drx_Preference_r16 != nil, ie.MaxCC_Preference_r16 != nil, ie.ReleasePreference_r16 != nil, ie.MinSchedulingOffsetPreference_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drx_Preference_r16 != nil {
		if err = ie.Drx_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_Preference_r16", err)
		}
	}
	if ie.MaxCC_Preference_r16 != nil {
		if err = ie.MaxCC_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCC_Preference_r16", err)
		}
	}
	if ie.ReleasePreference_r16 != nil {
		if err = ie.ReleasePreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReleasePreference_r16", err)
		}
	}
	if ie.MinSchedulingOffsetPreference_r16 != nil {
		if err = ie.MinSchedulingOffsetPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MinSchedulingOffsetPreference_r16", err)
		}
	}
	return nil
}

func (ie *PowSav_ParametersCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var Drx_Preference_r16Present bool
	if Drx_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCC_Preference_r16Present bool
	if MaxCC_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReleasePreference_r16Present bool
	if ReleasePreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MinSchedulingOffsetPreference_r16Present bool
	if MinSchedulingOffsetPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Drx_Preference_r16Present {
		ie.Drx_Preference_r16 = new(PowSav_ParametersCommon_r16_drx_Preference_r16)
		if err = ie.Drx_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_Preference_r16", err)
		}
	}
	if MaxCC_Preference_r16Present {
		ie.MaxCC_Preference_r16 = new(PowSav_ParametersCommon_r16_maxCC_Preference_r16)
		if err = ie.MaxCC_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCC_Preference_r16", err)
		}
	}
	if ReleasePreference_r16Present {
		ie.ReleasePreference_r16 = new(PowSav_ParametersCommon_r16_releasePreference_r16)
		if err = ie.ReleasePreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReleasePreference_r16", err)
		}
	}
	if MinSchedulingOffsetPreference_r16Present {
		ie.MinSchedulingOffsetPreference_r16 = new(PowSav_ParametersCommon_r16_minSchedulingOffsetPreference_r16)
		if err = ie.MinSchedulingOffsetPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MinSchedulingOffsetPreference_r16", err)
		}
	}
	return nil
}
