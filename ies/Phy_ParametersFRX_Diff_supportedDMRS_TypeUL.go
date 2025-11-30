package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_supportedDMRS_TypeUL_Enum_type1     aper.Enumerated = 0
	Phy_ParametersFRX_Diff_supportedDMRS_TypeUL_Enum_type1And2 aper.Enumerated = 1
)

type Phy_ParametersFRX_Diff_supportedDMRS_TypeUL struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *Phy_ParametersFRX_Diff_supportedDMRS_TypeUL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_supportedDMRS_TypeUL", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_supportedDMRS_TypeUL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_supportedDMRS_TypeUL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
