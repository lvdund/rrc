package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TimeUntilReconnection_r16 struct {
	Value uint64 `lb:0,ub:172800,madatory`
}

func (ie *TimeUntilReconnection_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 172800}, false); err != nil {
		return utils.WrapError("Encode TimeUntilReconnection_r16", err)
	}
	return nil
}

func (ie *TimeUntilReconnection_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 172800}, false); err != nil {
		return utils.WrapError("Decode TimeUntilReconnection_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
