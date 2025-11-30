package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MAC_Parameters_v1700 struct {
	Mac_ParametersFR2_2_r17 *MAC_ParametersFR2_2_r17 `optional`
}

func (ie *MAC_Parameters_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mac_ParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mac_ParametersFR2_2_r17 != nil {
		if err = ie.Mac_ParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MAC_Parameters_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Mac_ParametersFR2_2_r17Present bool
	if Mac_ParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Mac_ParametersFR2_2_r17Present {
		ie.Mac_ParametersFR2_2_r17 = new(MAC_ParametersFR2_2_r17)
		if err = ie.Mac_ParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersFR2_2_r17", err)
		}
	}
	return nil
}
