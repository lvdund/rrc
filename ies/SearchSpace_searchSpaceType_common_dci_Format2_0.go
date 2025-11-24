package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common_dci_Format2_0 struct {
	NrofCandidates_SFI *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI `optional`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofCandidates_SFI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofCandidates_SFI != nil {
		if err = ie.NrofCandidates_SFI.Encode(w); err != nil {
			return utils.WrapError("Encode NrofCandidates_SFI", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0) Decode(r *uper.UperReader) error {
	var err error
	var NrofCandidates_SFIPresent bool
	if NrofCandidates_SFIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofCandidates_SFIPresent {
		ie.NrofCandidates_SFI = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI)
		if err = ie.NrofCandidates_SFI.Decode(r); err != nil {
			return utils.WrapError("Decode NrofCandidates_SFI", err)
		}
	}
	return nil
}
