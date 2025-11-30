package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n16  aper.Enumerated = 0
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n32  aper.Enumerated = 1
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n48  aper.Enumerated = 2
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n64  aper.Enumerated = 3
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n80  aper.Enumerated = 4
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n96  aper.Enumerated = 5
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n112 aper.Enumerated = 6
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n128 aper.Enumerated = 7
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n144 aper.Enumerated = 8
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n160 aper.Enumerated = 9
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n176 aper.Enumerated = 10
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n192 aper.Enumerated = 11
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n208 aper.Enumerated = 12
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n224 aper.Enumerated = 13
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n240 aper.Enumerated = 14
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n256 aper.Enumerated = 15
)

type Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot", err)
	}
	return nil
}

func (ie *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
