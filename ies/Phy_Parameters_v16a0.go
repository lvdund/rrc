package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_Parameters_v16a0 struct {
	Phy_ParametersCommon_v16a0 *Phy_ParametersCommon_v16a0 `optional`
}

func (ie *Phy_Parameters_v16a0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Phy_ParametersCommon_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Phy_ParametersCommon_v16a0 != nil {
		if err = ie.Phy_ParametersCommon_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersCommon_v16a0", err)
		}
	}
	return nil
}

func (ie *Phy_Parameters_v16a0) Decode(r *aper.AperReader) error {
	var err error
	var Phy_ParametersCommon_v16a0Present bool
	if Phy_ParametersCommon_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Phy_ParametersCommon_v16a0Present {
		ie.Phy_ParametersCommon_v16a0 = new(Phy_ParametersCommon_v16a0)
		if err = ie.Phy_ParametersCommon_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersCommon_v16a0", err)
		}
	}
	return nil
}
