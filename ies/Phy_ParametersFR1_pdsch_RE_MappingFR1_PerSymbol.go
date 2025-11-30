package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol_Enum_n10 aper.Enumerated = 0
	Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol_Enum_n20 aper.Enumerated = 1
)

type Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol", err)
	}
	return nil
}

func (ie *Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
