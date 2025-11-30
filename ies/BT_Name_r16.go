package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BT_Name_r16 struct {
	Value []byte `lb:1,ub:248,madatory`
}

func (ie *BT_Name_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 1, Ub: 248}, false); err != nil {
		return utils.WrapError("Encode BT_Name_r16", err)
	}
	return nil
}

func (ie *BT_Name_r16) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 248}, false); err != nil {
		return utils.WrapError("Decode BT_Name_r16", err)
	}
	ie.Value = v
	return nil
}
