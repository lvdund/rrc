package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_ue_Specific struct {
	Dci_Formats        SearchSpace_searchSpaceType_ue_Specific_dci_Formats         `madatory`
	Dci_Formats_MT_r16 *SearchSpace_searchSpaceType_ue_Specific_dci_Formats_MT_r16 `optional`
	Dci_FormatsSL_r16  *SearchSpace_searchSpaceType_ue_Specific_dci_FormatsSL_r16  `optional`
	Dci_FormatsExt_r16 *SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16 `optional`
}

func (ie *SearchSpace_searchSpaceType_ue_Specific) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dci_Formats_MT_r16 != nil, ie.Dci_FormatsSL_r16 != nil, ie.Dci_FormatsExt_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Dci_Formats.Encode(w); err != nil {
		return utils.WrapError("Encode Dci_Formats", err)
	}
	if ie.Dci_Formats_MT_r16 != nil {
		if err = ie.Dci_Formats_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_Formats_MT_r16", err)
		}
	}
	if ie.Dci_FormatsSL_r16 != nil {
		if err = ie.Dci_FormatsSL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_FormatsSL_r16", err)
		}
	}
	if ie.Dci_FormatsExt_r16 != nil {
		if err = ie.Dci_FormatsExt_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_FormatsExt_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_ue_Specific) Decode(r *uper.UperReader) error {
	var err error
	var Dci_Formats_MT_r16Present bool
	if Dci_Formats_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_FormatsSL_r16Present bool
	if Dci_FormatsSL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_FormatsExt_r16Present bool
	if Dci_FormatsExt_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Dci_Formats.Decode(r); err != nil {
		return utils.WrapError("Decode Dci_Formats", err)
	}
	if Dci_Formats_MT_r16Present {
		ie.Dci_Formats_MT_r16 = new(SearchSpace_searchSpaceType_ue_Specific_dci_Formats_MT_r16)
		if err = ie.Dci_Formats_MT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_Formats_MT_r16", err)
		}
	}
	if Dci_FormatsSL_r16Present {
		ie.Dci_FormatsSL_r16 = new(SearchSpace_searchSpaceType_ue_Specific_dci_FormatsSL_r16)
		if err = ie.Dci_FormatsSL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_FormatsSL_r16", err)
		}
	}
	if Dci_FormatsExt_r16Present {
		ie.Dci_FormatsExt_r16 = new(SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16)
		if err = ie.Dci_FormatsExt_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_FormatsExt_r16", err)
		}
	}
	return nil
}
