package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SimulSRS_ForAntennaSwitching_r16 struct {
	SupportSRS_xTyR_xLessThanY_r16  *SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xLessThanY_r16  `optional`
	SupportSRS_xTyR_xEqualToY_r16   *SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xEqualToY_r16   `optional`
	SupportSRS_AntennaSwitching_r16 *SimulSRS_ForAntennaSwitching_r16_supportSRS_AntennaSwitching_r16 `optional`
}

func (ie *SimulSRS_ForAntennaSwitching_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportSRS_xTyR_xLessThanY_r16 != nil, ie.SupportSRS_xTyR_xEqualToY_r16 != nil, ie.SupportSRS_AntennaSwitching_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportSRS_xTyR_xLessThanY_r16 != nil {
		if err = ie.SupportSRS_xTyR_xLessThanY_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportSRS_xTyR_xLessThanY_r16", err)
		}
	}
	if ie.SupportSRS_xTyR_xEqualToY_r16 != nil {
		if err = ie.SupportSRS_xTyR_xEqualToY_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportSRS_xTyR_xEqualToY_r16", err)
		}
	}
	if ie.SupportSRS_AntennaSwitching_r16 != nil {
		if err = ie.SupportSRS_AntennaSwitching_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportSRS_AntennaSwitching_r16", err)
		}
	}
	return nil
}

func (ie *SimulSRS_ForAntennaSwitching_r16) Decode(r *uper.UperReader) error {
	var err error
	var SupportSRS_xTyR_xLessThanY_r16Present bool
	if SupportSRS_xTyR_xLessThanY_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportSRS_xTyR_xEqualToY_r16Present bool
	if SupportSRS_xTyR_xEqualToY_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportSRS_AntennaSwitching_r16Present bool
	if SupportSRS_AntennaSwitching_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportSRS_xTyR_xLessThanY_r16Present {
		ie.SupportSRS_xTyR_xLessThanY_r16 = new(SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xLessThanY_r16)
		if err = ie.SupportSRS_xTyR_xLessThanY_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportSRS_xTyR_xLessThanY_r16", err)
		}
	}
	if SupportSRS_xTyR_xEqualToY_r16Present {
		ie.SupportSRS_xTyR_xEqualToY_r16 = new(SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xEqualToY_r16)
		if err = ie.SupportSRS_xTyR_xEqualToY_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportSRS_xTyR_xEqualToY_r16", err)
		}
	}
	if SupportSRS_AntennaSwitching_r16Present {
		ie.SupportSRS_AntennaSwitching_r16 = new(SimulSRS_ForAntennaSwitching_r16_supportSRS_AntennaSwitching_r16)
		if err = ie.SupportSRS_AntennaSwitching_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportSRS_AntennaSwitching_r16", err)
		}
	}
	return nil
}
