package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17 struct {
	MaxNumListDL_TCI_r17 *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListDL_TCI_r17 `optional`
	MaxNumListUL_TCI_r17 *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListUL_TCI_r17 `optional`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumListDL_TCI_r17 != nil, ie.MaxNumListUL_TCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxNumListDL_TCI_r17 != nil {
		if err = ie.MaxNumListDL_TCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumListDL_TCI_r17", err)
		}
	}
	if ie.MaxNumListUL_TCI_r17 != nil {
		if err = ie.MaxNumListUL_TCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumListUL_TCI_r17", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17) Decode(r *aper.AperReader) error {
	var err error
	var MaxNumListDL_TCI_r17Present bool
	if MaxNumListDL_TCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumListUL_TCI_r17Present bool
	if MaxNumListUL_TCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxNumListDL_TCI_r17Present {
		ie.MaxNumListDL_TCI_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListDL_TCI_r17)
		if err = ie.MaxNumListDL_TCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumListDL_TCI_r17", err)
		}
	}
	if MaxNumListUL_TCI_r17Present {
		ie.MaxNumListUL_TCI_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListUL_TCI_r17)
		if err = ie.MaxNumListUL_TCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumListUL_TCI_r17", err)
		}
	}
	return nil
}
