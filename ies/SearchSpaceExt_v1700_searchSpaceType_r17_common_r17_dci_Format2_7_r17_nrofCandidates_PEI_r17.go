package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17 struct {
	AggregationLevel4_r17  *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel4_r17  `optional`
	AggregationLevel8_r17  *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17  `optional`
	AggregationLevel16_r17 *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17 `optional`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AggregationLevel4_r17 != nil, ie.AggregationLevel8_r17 != nil, ie.AggregationLevel16_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AggregationLevel4_r17 != nil {
		if err = ie.AggregationLevel4_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel4_r17", err)
		}
	}
	if ie.AggregationLevel8_r17 != nil {
		if err = ie.AggregationLevel8_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel8_r17", err)
		}
	}
	if ie.AggregationLevel16_r17 != nil {
		if err = ie.AggregationLevel16_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel16_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17) Decode(r *uper.UperReader) error {
	var err error
	var AggregationLevel4_r17Present bool
	if AggregationLevel4_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel8_r17Present bool
	if AggregationLevel8_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel16_r17Present bool
	if AggregationLevel16_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AggregationLevel4_r17Present {
		ie.AggregationLevel4_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel4_r17)
		if err = ie.AggregationLevel4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel4_r17", err)
		}
	}
	if AggregationLevel8_r17Present {
		ie.AggregationLevel8_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17)
		if err = ie.AggregationLevel8_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel8_r17", err)
		}
	}
	if AggregationLevel16_r17Present {
		ie.AggregationLevel16_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17)
		if err = ie.AggregationLevel16_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel16_r17", err)
		}
	}
	return nil
}
