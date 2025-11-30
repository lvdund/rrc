package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16_Enum_formats0_2_And_1_2                aper.Enumerated = 0
	SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16_Enum_formats0_1_And_1_1And_0_2_And_1_2 aper.Enumerated = 1
)

type SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
