package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_nrofCandidates struct {
	AggregationLevel1  SearchSpace_nrofCandidates_aggregationLevel1  `madatory`
	AggregationLevel2  SearchSpace_nrofCandidates_aggregationLevel2  `madatory`
	AggregationLevel4  SearchSpace_nrofCandidates_aggregationLevel4  `madatory`
	AggregationLevel8  SearchSpace_nrofCandidates_aggregationLevel8  `madatory`
	AggregationLevel16 SearchSpace_nrofCandidates_aggregationLevel16 `madatory`
}

func (ie *SearchSpace_nrofCandidates) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.AggregationLevel1.Encode(w); err != nil {
		return utils.WrapError("Encode AggregationLevel1", err)
	}
	if err = ie.AggregationLevel2.Encode(w); err != nil {
		return utils.WrapError("Encode AggregationLevel2", err)
	}
	if err = ie.AggregationLevel4.Encode(w); err != nil {
		return utils.WrapError("Encode AggregationLevel4", err)
	}
	if err = ie.AggregationLevel8.Encode(w); err != nil {
		return utils.WrapError("Encode AggregationLevel8", err)
	}
	if err = ie.AggregationLevel16.Encode(w); err != nil {
		return utils.WrapError("Encode AggregationLevel16", err)
	}
	return nil
}

func (ie *SearchSpace_nrofCandidates) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.AggregationLevel1.Decode(r); err != nil {
		return utils.WrapError("Decode AggregationLevel1", err)
	}
	if err = ie.AggregationLevel2.Decode(r); err != nil {
		return utils.WrapError("Decode AggregationLevel2", err)
	}
	if err = ie.AggregationLevel4.Decode(r); err != nil {
		return utils.WrapError("Decode AggregationLevel4", err)
	}
	if err = ie.AggregationLevel8.Decode(r); err != nil {
		return utils.WrapError("Decode AggregationLevel8", err)
	}
	if err = ie.AggregationLevel16.Decode(r); err != nil {
		return utils.WrapError("Decode AggregationLevel16", err)
	}
	return nil
}
