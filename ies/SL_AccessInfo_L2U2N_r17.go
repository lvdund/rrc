package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_AccessInfo_L2U2N_r17 struct {
	CellAccessRelatedInfo_r17 CellAccessRelatedInfo  `madatory`
	Sl_ServingCellInfo_r17    SL_ServingCellInfo_r17 `madatory`
}

func (ie *SL_AccessInfo_L2U2N_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CellAccessRelatedInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CellAccessRelatedInfo_r17", err)
	}
	if err = ie.Sl_ServingCellInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ServingCellInfo_r17", err)
	}
	return nil
}

func (ie *SL_AccessInfo_L2U2N_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CellAccessRelatedInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CellAccessRelatedInfo_r17", err)
	}
	if err = ie.Sl_ServingCellInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ServingCellInfo_r17", err)
	}
	return nil
}
