package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_pdsch_MappingTypeB_Enum_supported aper.Enumerated = 0
)

type Phy_ParametersCommon_pdsch_MappingTypeB struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersCommon_pdsch_MappingTypeB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_pdsch_MappingTypeB", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_pdsch_MappingTypeB) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_pdsch_MappingTypeB", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
