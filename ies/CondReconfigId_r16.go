package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CondReconfigId_r16 struct {
	Value uint64 `lb:1,ub:maxNrofCondCells_r16,madatory`
}

func (ie *CondReconfigId_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false); err != nil {
		return utils.WrapError("Encode CondReconfigId_r16", err)
	}
	return nil
}

func (ie *CondReconfigId_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false); err != nil {
		return utils.WrapError("Decode CondReconfigId_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
