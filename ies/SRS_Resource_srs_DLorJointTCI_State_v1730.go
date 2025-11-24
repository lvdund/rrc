package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_srs_DLorJointTCI_State_v1730 struct {
	CellAndBWP_r17 ServingCellAndBWP_Id_r17 `madatory`
}

func (ie *SRS_Resource_srs_DLorJointTCI_State_v1730) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CellAndBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CellAndBWP_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_srs_DLorJointTCI_State_v1730) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CellAndBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CellAndBWP_r17", err)
	}
	return nil
}
