package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4_Enum_n1 aper.Enumerated = 0
	SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4_Enum_n2 aper.Enumerated = 1
)

type SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
