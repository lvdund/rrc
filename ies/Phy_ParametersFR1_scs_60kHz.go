package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFR1_scs_60kHz_Enum_supported aper.Enumerated = 0
)

type Phy_ParametersFR1_scs_60kHz struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersFR1_scs_60kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFR1_scs_60kHz", err)
	}
	return nil
}

func (ie *Phy_ParametersFR1_scs_60kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFR1_scs_60kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
