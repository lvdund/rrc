package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QCL_Info struct {
	Cell            *ServCellIndex           `optional`
	Bwp_Id          *BWP_Id                  `optional`
	ReferenceSignal QCL_Info_referenceSignal `madatory`
	Qcl_Type        QCL_Info_qcl_Type        `madatory`
}

func (ie *QCL_Info) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Cell != nil, ie.Bwp_Id != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cell != nil {
		if err = ie.Cell.Encode(w); err != nil {
			return utils.WrapError("Encode Cell", err)
		}
	}
	if ie.Bwp_Id != nil {
		if err = ie.Bwp_Id.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_Id", err)
		}
	}
	if err = ie.ReferenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal", err)
	}
	if err = ie.Qcl_Type.Encode(w); err != nil {
		return utils.WrapError("Encode Qcl_Type", err)
	}
	return nil
}

func (ie *QCL_Info) Decode(r *uper.UperReader) error {
	var err error
	var CellPresent bool
	if CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_IdPresent bool
	if Bwp_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CellPresent {
		ie.Cell = new(ServCellIndex)
		if err = ie.Cell.Decode(r); err != nil {
			return utils.WrapError("Decode Cell", err)
		}
	}
	if Bwp_IdPresent {
		ie.Bwp_Id = new(BWP_Id)
		if err = ie.Bwp_Id.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_Id", err)
		}
	}
	if err = ie.ReferenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal", err)
	}
	if err = ie.Qcl_Type.Decode(r); err != nil {
		return utils.WrapError("Decode Qcl_Type", err)
	}
	return nil
}
