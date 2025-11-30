package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n0 aper.Enumerated = 0
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n1 aper.Enumerated = 1
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n2 aper.Enumerated = 2
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n3 aper.Enumerated = 3
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n4 aper.Enumerated = 4
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n5 aper.Enumerated = 5
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n6 aper.Enumerated = 6
	SearchSpace_nrofCandidates_aggregationLevel2_Enum_n8 aper.Enumerated = 7
)

type SearchSpace_nrofCandidates_aggregationLevel2 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SearchSpace_nrofCandidates_aggregationLevel2) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_nrofCandidates_aggregationLevel2", err)
	}
	return nil
}

func (ie *SearchSpace_nrofCandidates_aggregationLevel2) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_nrofCandidates_aggregationLevel2", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
