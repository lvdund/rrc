package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n50   aper.Enumerated = 0
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n75   aper.Enumerated = 1
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n100  aper.Enumerated = 2
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n125  aper.Enumerated = 3
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n150  aper.Enumerated = 4
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n175  aper.Enumerated = 5
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n200  aper.Enumerated = 6
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n225  aper.Enumerated = 7
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n250  aper.Enumerated = 8
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n275  aper.Enumerated = 9
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n300  aper.Enumerated = 10
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n350  aper.Enumerated = 11
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n400  aper.Enumerated = 12
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n450  aper.Enumerated = 13
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n500  aper.Enumerated = 14
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_spare aper.Enumerated = 15
)

type NAICS_Capability_Entry_numberOfAggregatedPRB struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *NAICS_Capability_Entry_numberOfAggregatedPRB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode NAICS_Capability_Entry_numberOfAggregatedPRB", err)
	}
	return nil
}

func (ie *NAICS_Capability_Entry_numberOfAggregatedPRB) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode NAICS_Capability_Entry_numberOfAggregatedPRB", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
