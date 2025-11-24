package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610_format_v1610_occ_v1610 struct {
	Occ_Length_v1610 *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610 `optional`
	Occ_Index_v1610  *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610  `optional`
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Occ_Length_v1610 != nil, ie.Occ_Index_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Occ_Length_v1610 != nil {
		if err = ie.Occ_Length_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Occ_Length_v1610", err)
		}
	}
	if ie.Occ_Index_v1610 != nil {
		if err = ie.Occ_Index_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Occ_Index_v1610", err)
		}
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610) Decode(r *uper.UperReader) error {
	var err error
	var Occ_Length_v1610Present bool
	if Occ_Length_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Occ_Index_v1610Present bool
	if Occ_Index_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Occ_Length_v1610Present {
		ie.Occ_Length_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610)
		if err = ie.Occ_Length_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Occ_Length_v1610", err)
		}
	}
	if Occ_Index_v1610Present {
		ie.Occ_Index_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610)
		if err = ie.Occ_Index_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Occ_Index_v1610", err)
		}
	}
	return nil
}
