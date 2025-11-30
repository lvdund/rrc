package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MBS_FSAI_r17 struct {
	Value []byte `lb:3,ub:3,madatory`
}

func (ie *MBS_FSAI_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MBS_FSAI_r17", err)
	}
	return nil
}

func (ie *MBS_FSAI_r17) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MBS_FSAI_r17", err)
	}
	ie.Value = v
	return nil
}
