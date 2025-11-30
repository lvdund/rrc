package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common_dci_Format2_3 struct {
	Dummy1 *SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1 `optional`
	Dummy2 SearchSpace_searchSpaceType_common_dci_Format2_3_dummy2  `madatory`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dummy1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if err = ie.Dummy2.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy2", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3) Decode(r *aper.AperReader) error {
	var err error
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dummy1Present {
		ie.Dummy1 = new(SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if err = ie.Dummy2.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy2", err)
	}
	return nil
}
