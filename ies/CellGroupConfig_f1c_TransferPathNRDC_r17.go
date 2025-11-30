package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellGroupConfig_f1c_TransferPathNRDC_r17_Enum_mcg  aper.Enumerated = 0
	CellGroupConfig_f1c_TransferPathNRDC_r17_Enum_scg  aper.Enumerated = 1
	CellGroupConfig_f1c_TransferPathNRDC_r17_Enum_both aper.Enumerated = 2
)

type CellGroupConfig_f1c_TransferPathNRDC_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CellGroupConfig_f1c_TransferPathNRDC_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CellGroupConfig_f1c_TransferPathNRDC_r17", err)
	}
	return nil
}

func (ie *CellGroupConfig_f1c_TransferPathNRDC_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CellGroupConfig_f1c_TransferPathNRDC_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
