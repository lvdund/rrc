package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH_Enum_notSupported aper.Enumerated = 0
)

type Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
