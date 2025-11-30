package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_Parameters_v1700 struct {
	PowSav_ParametersFR2_2_r17 *PowSav_ParametersFR2_2_r17 `optional`
}

func (ie *PowSav_Parameters_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PowSav_ParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PowSav_ParametersFR2_2_r17 != nil {
		if err = ie.PowSav_ParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PowSav_ParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *PowSav_Parameters_v1700) Decode(r *aper.AperReader) error {
	var err error
	var PowSav_ParametersFR2_2_r17Present bool
	if PowSav_ParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PowSav_ParametersFR2_2_r17Present {
		ie.PowSav_ParametersFR2_2_r17 = new(PowSav_ParametersFR2_2_r17)
		if err = ie.PowSav_ParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PowSav_ParametersFR2_2_r17", err)
		}
	}
	return nil
}
