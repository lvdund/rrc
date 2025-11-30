package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IMS_Parameters_v1700 struct {
	Ims_ParametersFR2_2_r17 *IMS_ParametersFR2_2_r17 `optional`
}

func (ie *IMS_Parameters_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ims_ParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ims_ParametersFR2_2_r17 != nil {
		if err = ie.Ims_ParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_ParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *IMS_Parameters_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Ims_ParametersFR2_2_r17Present bool
	if Ims_ParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ims_ParametersFR2_2_r17Present {
		ie.Ims_ParametersFR2_2_r17 = new(IMS_ParametersFR2_2_r17)
		if err = ie.Ims_ParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_ParametersFR2_2_r17", err)
		}
	}
	return nil
}
