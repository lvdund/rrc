package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig_schedulingCellInfo_other struct {
	SchedulingCellId     ServCellIndex `madatory`
	Cif_InSchedulingCell int64         `lb:1,ub:7,madatory`
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_other) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.SchedulingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode SchedulingCellId", err)
	}
	if err = w.WriteInteger(ie.Cif_InSchedulingCell, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger Cif_InSchedulingCell", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_other) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.SchedulingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode SchedulingCellId", err)
	}
	var tmp_int_Cif_InSchedulingCell int64
	if tmp_int_Cif_InSchedulingCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger Cif_InSchedulingCell", err)
	}
	ie.Cif_InSchedulingCell = tmp_int_Cif_InSchedulingCell
	return nil
}
