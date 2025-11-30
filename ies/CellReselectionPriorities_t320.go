package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellReselectionPriorities_t320_Enum_min5   aper.Enumerated = 0
	CellReselectionPriorities_t320_Enum_min10  aper.Enumerated = 1
	CellReselectionPriorities_t320_Enum_min20  aper.Enumerated = 2
	CellReselectionPriorities_t320_Enum_min30  aper.Enumerated = 3
	CellReselectionPriorities_t320_Enum_min60  aper.Enumerated = 4
	CellReselectionPriorities_t320_Enum_min120 aper.Enumerated = 5
	CellReselectionPriorities_t320_Enum_min180 aper.Enumerated = 6
	CellReselectionPriorities_t320_Enum_spare1 aper.Enumerated = 7
)

type CellReselectionPriorities_t320 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CellReselectionPriorities_t320) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CellReselectionPriorities_t320", err)
	}
	return nil
}

func (ie *CellReselectionPriorities_t320) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CellReselectionPriorities_t320", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
