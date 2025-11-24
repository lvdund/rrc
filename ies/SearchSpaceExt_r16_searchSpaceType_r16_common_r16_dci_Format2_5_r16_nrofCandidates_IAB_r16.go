package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16 struct {
	AggregationLevel1_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel1_r16  `optional`
	AggregationLevel2_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16  `optional`
	AggregationLevel4_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel4_r16  `optional`
	AggregationLevel8_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel8_r16  `optional`
	AggregationLevel16_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel16_r16 `optional`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AggregationLevel1_r16 != nil, ie.AggregationLevel2_r16 != nil, ie.AggregationLevel4_r16 != nil, ie.AggregationLevel8_r16 != nil, ie.AggregationLevel16_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AggregationLevel1_r16 != nil {
		if err = ie.AggregationLevel1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel1_r16", err)
		}
	}
	if ie.AggregationLevel2_r16 != nil {
		if err = ie.AggregationLevel2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel2_r16", err)
		}
	}
	if ie.AggregationLevel4_r16 != nil {
		if err = ie.AggregationLevel4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel4_r16", err)
		}
	}
	if ie.AggregationLevel8_r16 != nil {
		if err = ie.AggregationLevel8_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel8_r16", err)
		}
	}
	if ie.AggregationLevel16_r16 != nil {
		if err = ie.AggregationLevel16_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel16_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16) Decode(r *uper.UperReader) error {
	var err error
	var AggregationLevel1_r16Present bool
	if AggregationLevel1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel2_r16Present bool
	if AggregationLevel2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel4_r16Present bool
	if AggregationLevel4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel8_r16Present bool
	if AggregationLevel8_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel16_r16Present bool
	if AggregationLevel16_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AggregationLevel1_r16Present {
		ie.AggregationLevel1_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel1_r16)
		if err = ie.AggregationLevel1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel1_r16", err)
		}
	}
	if AggregationLevel2_r16Present {
		ie.AggregationLevel2_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel2_r16)
		if err = ie.AggregationLevel2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel2_r16", err)
		}
	}
	if AggregationLevel4_r16Present {
		ie.AggregationLevel4_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel4_r16)
		if err = ie.AggregationLevel4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel4_r16", err)
		}
	}
	if AggregationLevel8_r16Present {
		ie.AggregationLevel8_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel8_r16)
		if err = ie.AggregationLevel8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel8_r16", err)
		}
	}
	if AggregationLevel16_r16Present {
		ie.AggregationLevel16_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16_aggregationLevel16_r16)
		if err = ie.AggregationLevel16_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel16_r16", err)
		}
	}
	return nil
}
