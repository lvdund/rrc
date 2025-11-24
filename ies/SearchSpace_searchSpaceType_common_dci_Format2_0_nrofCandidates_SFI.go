package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI struct {
	AggregationLevel1  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel1  `optional`
	AggregationLevel2  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel2  `optional`
	AggregationLevel4  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4  `optional`
	AggregationLevel8  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8  `optional`
	AggregationLevel16 *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel16 `optional`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AggregationLevel1 != nil, ie.AggregationLevel2 != nil, ie.AggregationLevel4 != nil, ie.AggregationLevel8 != nil, ie.AggregationLevel16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AggregationLevel1 != nil {
		if err = ie.AggregationLevel1.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel1", err)
		}
	}
	if ie.AggregationLevel2 != nil {
		if err = ie.AggregationLevel2.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel2", err)
		}
	}
	if ie.AggregationLevel4 != nil {
		if err = ie.AggregationLevel4.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel4", err)
		}
	}
	if ie.AggregationLevel8 != nil {
		if err = ie.AggregationLevel8.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel8", err)
		}
	}
	if ie.AggregationLevel16 != nil {
		if err = ie.AggregationLevel16.Encode(w); err != nil {
			return utils.WrapError("Encode AggregationLevel16", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI) Decode(r *uper.UperReader) error {
	var err error
	var AggregationLevel1Present bool
	if AggregationLevel1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel2Present bool
	if AggregationLevel2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel4Present bool
	if AggregationLevel4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel8Present bool
	if AggregationLevel8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AggregationLevel16Present bool
	if AggregationLevel16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AggregationLevel1Present {
		ie.AggregationLevel1 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel1)
		if err = ie.AggregationLevel1.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel1", err)
		}
	}
	if AggregationLevel2Present {
		ie.AggregationLevel2 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel2)
		if err = ie.AggregationLevel2.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel2", err)
		}
	}
	if AggregationLevel4Present {
		ie.AggregationLevel4 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4)
		if err = ie.AggregationLevel4.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel4", err)
		}
	}
	if AggregationLevel8Present {
		ie.AggregationLevel8 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8)
		if err = ie.AggregationLevel8.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel8", err)
		}
	}
	if AggregationLevel16Present {
		ie.AggregationLevel16 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel16)
		if err = ie.AggregationLevel16.Decode(r); err != nil {
			return utils.WrapError("Decode AggregationLevel16", err)
		}
	}
	return nil
}
