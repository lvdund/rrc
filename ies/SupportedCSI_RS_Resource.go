package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SupportedCSI_RS_Resource struct {
	MaxNumberTxPortsPerResource SupportedCSI_RS_Resource_maxNumberTxPortsPerResource `madatory`
	MaxNumberResourcesPerBand   int64                                                `lb:1,ub:64,madatory`
	TotalNumberTxPortsPerBand   int64                                                `lb:2,ub:256,madatory`
}

func (ie *SupportedCSI_RS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberTxPortsPerResource.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberTxPortsPerResource", err)
	}
	if err = w.WriteInteger(ie.MaxNumberResourcesPerBand, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberResourcesPerBand", err)
	}
	if err = w.WriteInteger(ie.TotalNumberTxPortsPerBand, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger TotalNumberTxPortsPerBand", err)
	}
	return nil
}

func (ie *SupportedCSI_RS_Resource) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberTxPortsPerResource.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberTxPortsPerResource", err)
	}
	var tmp_int_MaxNumberResourcesPerBand int64
	if tmp_int_MaxNumberResourcesPerBand, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberResourcesPerBand", err)
	}
	ie.MaxNumberResourcesPerBand = tmp_int_MaxNumberResourcesPerBand
	var tmp_int_TotalNumberTxPortsPerBand int64
	if tmp_int_TotalNumberTxPortsPerBand, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger TotalNumberTxPortsPerBand", err)
	}
	ie.TotalNumberTxPortsPerBand = tmp_int_TotalNumberTxPortsPerBand
	return nil
}
