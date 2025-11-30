package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceLocation_r17 struct {
	Value []byte `madatory`
}

func (ie *ReferenceLocation_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode ReferenceLocation_r17", err)
	}
	return nil
}

func (ie *ReferenceLocation_r17) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode ReferenceLocation_r17", err)
	}
	ie.Value = v
	return nil
}
