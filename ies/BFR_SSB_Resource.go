package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BFR_SSB_Resource struct {
	Ssb              SSB_Index `madatory`
	Ra_PreambleIndex int64     `lb:0,ub:63,madatory`
}

func (ie *BFR_SSB_Resource) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ssb.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb", err)
	}
	if err = w.WriteInteger(ie.Ra_PreambleIndex, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger Ra_PreambleIndex", err)
	}
	return nil
}

func (ie *BFR_SSB_Resource) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ssb.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb", err)
	}
	var tmp_int_Ra_PreambleIndex int64
	if tmp_int_Ra_PreambleIndex, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger Ra_PreambleIndex", err)
	}
	ie.Ra_PreambleIndex = tmp_int_Ra_PreambleIndex
	return nil
}
