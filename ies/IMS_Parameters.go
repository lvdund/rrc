package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IMS_Parameters struct {
	Ims_ParametersCommon   *IMS_ParametersCommon   `optional`
	Ims_ParametersFRX_Diff *IMS_ParametersFRX_Diff `optional`
}

func (ie *IMS_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ims_ParametersCommon != nil, ie.Ims_ParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ims_ParametersCommon != nil {
		if err = ie.Ims_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_ParametersCommon", err)
		}
	}
	if ie.Ims_ParametersFRX_Diff != nil {
		if err = ie.Ims_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *IMS_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var Ims_ParametersCommonPresent bool
	if Ims_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ims_ParametersFRX_DiffPresent bool
	if Ims_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ims_ParametersCommonPresent {
		ie.Ims_ParametersCommon = new(IMS_ParametersCommon)
		if err = ie.Ims_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_ParametersCommon", err)
		}
	}
	if Ims_ParametersFRX_DiffPresent {
		ie.Ims_ParametersFRX_Diff = new(IMS_ParametersFRX_Diff)
		if err = ie.Ims_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}
