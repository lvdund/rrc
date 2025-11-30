package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16_Enum_n1 aper.Enumerated = 0
	SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16_Enum_n2 aper.Enumerated = 1
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16", err)
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
