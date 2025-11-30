package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16 struct {
	NrofCandidates_IAB_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16 `optional`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofCandidates_IAB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofCandidates_IAB_r16 != nil {
		if err = ie.NrofCandidates_IAB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NrofCandidates_IAB_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16) Decode(r *aper.AperReader) error {
	var err error
	var NrofCandidates_IAB_r16Present bool
	if NrofCandidates_IAB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofCandidates_IAB_r16Present {
		ie.NrofCandidates_IAB_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16_nrofCandidates_IAB_r16)
		if err = ie.NrofCandidates_IAB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NrofCandidates_IAB_r16", err)
		}
	}
	return nil
}
