package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_semiOpenLoopCSI_Enum_supported aper.Enumerated = 0
)

type Phy_ParametersFRX_Diff_semiOpenLoopCSI struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersFRX_Diff_semiOpenLoopCSI) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_semiOpenLoopCSI", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_semiOpenLoopCSI) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_semiOpenLoopCSI", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
