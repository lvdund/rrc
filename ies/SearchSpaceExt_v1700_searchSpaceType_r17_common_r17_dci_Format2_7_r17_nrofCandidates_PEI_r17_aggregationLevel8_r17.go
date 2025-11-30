package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17_Enum_n0 aper.Enumerated = 0
	SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17_Enum_n1 aper.Enumerated = 1
	SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17_Enum_n2 aper.Enumerated = 2
)

type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17", err)
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
