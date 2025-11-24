package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16 struct {
	Dci_Format2_4_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16 `optional`
	Dci_Format2_5_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16 `optional`
	Dci_Format2_6_r16 interface{}                                                          `optional`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dci_Format2_4_r16 != nil, ie.Dci_Format2_5_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dci_Format2_4_r16 != nil {
		if err = ie.Dci_Format2_4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_Format2_4_r16", err)
		}
	}
	if ie.Dci_Format2_5_r16 != nil {
		if err = ie.Dci_Format2_5_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_Format2_5_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16) Decode(r *uper.UperReader) error {
	var err error
	var Dci_Format2_4_r16Present bool
	if Dci_Format2_4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_Format2_5_r16Present bool
	if Dci_Format2_5_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dci_Format2_4_r16Present {
		ie.Dci_Format2_4_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16)
		if err = ie.Dci_Format2_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_Format2_4_r16", err)
		}
	}
	if Dci_Format2_5_r16Present {
		ie.Dci_Format2_5_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16)
		if err = ie.Dci_Format2_5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_Format2_5_r16", err)
		}
	}
	return nil
}
