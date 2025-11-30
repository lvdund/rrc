package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellReselectionSubPriority_Enum_oDot2 aper.Enumerated = 0
	CellReselectionSubPriority_Enum_oDot4 aper.Enumerated = 1
	CellReselectionSubPriority_Enum_oDot6 aper.Enumerated = 2
	CellReselectionSubPriority_Enum_oDot8 aper.Enumerated = 3
)

type CellReselectionSubPriority struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CellReselectionSubPriority) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CellReselectionSubPriority", err)
	}
	return nil
}

func (ie *CellReselectionSubPriority) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CellReselectionSubPriority", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
