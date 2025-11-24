package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common struct {
	Dci_Format0_0_AndFormat1_0 interface{}                                       `optional`
	Dci_Format2_0              *SearchSpace_searchSpaceType_common_dci_Format2_0 `optional`
	Dci_Format2_1              interface{}                                       `optional`
	Dci_Format2_2              interface{}                                       `optional`
	Dci_Format2_3              *SearchSpace_searchSpaceType_common_dci_Format2_3 `optional`
}

func (ie *SearchSpace_searchSpaceType_common) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dci_Format2_0 != nil, ie.Dci_Format2_3 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dci_Format2_0 != nil {
		if err = ie.Dci_Format2_0.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_Format2_0", err)
		}
	}
	if ie.Dci_Format2_3 != nil {
		if err = ie.Dci_Format2_3.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_Format2_3", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common) Decode(r *uper.UperReader) error {
	var err error
	var Dci_Format2_0Present bool
	if Dci_Format2_0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_Format2_3Present bool
	if Dci_Format2_3Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dci_Format2_0Present {
		ie.Dci_Format2_0 = new(SearchSpace_searchSpaceType_common_dci_Format2_0)
		if err = ie.Dci_Format2_0.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_Format2_0", err)
		}
	}
	if Dci_Format2_3Present {
		ie.Dci_Format2_3 = new(SearchSpace_searchSpaceType_common_dci_Format2_3)
		if err = ie.Dci_Format2_3.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_Format2_3", err)
		}
	}
	return nil
}
