package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DummyB struct {
	MaxNumberTxPortsPerResource    DummyB_maxNumberTxPortsPerResource `madatory`
	MaxNumberResources             int64                              `lb:1,ub:64,madatory`
	TotalNumberTxPorts             int64                              `lb:2,ub:256,madatory`
	SupportedCodebookMode          DummyB_supportedCodebookMode       `madatory`
	MaxNumberCSI_RS_PerResourceSet int64                              `lb:1,ub:8,madatory`
}

func (ie *DummyB) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumberTxPortsPerResource.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberTxPortsPerResource", err)
	}
	if err = w.WriteInteger(ie.MaxNumberResources, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberResources", err)
	}
	if err = w.WriteInteger(ie.TotalNumberTxPorts, &aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger TotalNumberTxPorts", err)
	}
	if err = ie.SupportedCodebookMode.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedCodebookMode", err)
	}
	if err = w.WriteInteger(ie.MaxNumberCSI_RS_PerResourceSet, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberCSI_RS_PerResourceSet", err)
	}
	return nil
}

func (ie *DummyB) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumberTxPortsPerResource.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberTxPortsPerResource", err)
	}
	var tmp_int_MaxNumberResources int64
	if tmp_int_MaxNumberResources, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberResources", err)
	}
	ie.MaxNumberResources = tmp_int_MaxNumberResources
	var tmp_int_TotalNumberTxPorts int64
	if tmp_int_TotalNumberTxPorts, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger TotalNumberTxPorts", err)
	}
	ie.TotalNumberTxPorts = tmp_int_TotalNumberTxPorts
	if err = ie.SupportedCodebookMode.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedCodebookMode", err)
	}
	var tmp_int_MaxNumberCSI_RS_PerResourceSet int64
	if tmp_int_MaxNumberCSI_RS_PerResourceSet, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberCSI_RS_PerResourceSet", err)
	}
	ie.MaxNumberCSI_RS_PerResourceSet = tmp_int_MaxNumberCSI_RS_PerResourceSet
	return nil
}
