package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Q_QualMin struct {
	Value uint64 `lb:-43,ub:-12,madatory`
}

func (ie *Q_QualMin) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: -43, Ub: -12}, false); err != nil {
		return utils.WrapError("Encode Q_QualMin", err)
	}
	return nil
}

func (ie *Q_QualMin) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: -43, Ub: -12}, false); err != nil {
		return utils.WrapError("Decode Q_QualMin", err)
	}
	ie.Value = uint64(v)
	return nil
}
