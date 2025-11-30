package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceId struct {
	Value uint64 `lb:0,ub:maxNrofSearchSpaces_1,madatory`
}

func (ie *SearchSpaceId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofSearchSpaces_1}, false); err != nil {
		return utils.WrapError("Encode SearchSpaceId", err)
	}
	return nil
}

func (ie *SearchSpaceId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofSearchSpaces_1}, false); err != nil {
		return utils.WrapError("Decode SearchSpaceId", err)
	}
	ie.Value = uint64(v)
	return nil
}
