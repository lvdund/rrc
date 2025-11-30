package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH_Enum_supported aper.Enumerated = 0
)

type Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH", err)
	}
	return nil
}

func (ie *Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
