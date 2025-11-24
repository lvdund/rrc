package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_Parameters_r16 struct {
	PowSav_ParametersCommon_r16   *PowSav_ParametersCommon_r16   `optional`
	PowSav_ParametersFRX_Diff_r16 *PowSav_ParametersFRX_Diff_r16 `optional`
}

func (ie *PowSav_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PowSav_ParametersCommon_r16 != nil, ie.PowSav_ParametersFRX_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PowSav_ParametersCommon_r16 != nil {
		if err = ie.PowSav_ParametersCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PowSav_ParametersCommon_r16", err)
		}
	}
	if ie.PowSav_ParametersFRX_Diff_r16 != nil {
		if err = ie.PowSav_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PowSav_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}

func (ie *PowSav_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var PowSav_ParametersCommon_r16Present bool
	if PowSav_ParametersCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PowSav_ParametersFRX_Diff_r16Present bool
	if PowSav_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PowSav_ParametersCommon_r16Present {
		ie.PowSav_ParametersCommon_r16 = new(PowSav_ParametersCommon_r16)
		if err = ie.PowSav_ParametersCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PowSav_ParametersCommon_r16", err)
		}
	}
	if PowSav_ParametersFRX_Diff_r16Present {
		ie.PowSav_ParametersFRX_Diff_r16 = new(PowSav_ParametersFRX_Diff_r16)
		if err = ie.PowSav_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PowSav_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}
