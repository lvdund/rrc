package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms3  aper.Enumerated = 0
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms4  aper.Enumerated = 1
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms6  aper.Enumerated = 2
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms10 aper.Enumerated = 3
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms20 aper.Enumerated = 4
)

type MUSIM_GapInfo_r17_musim_GapLength_r17 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MUSIM_GapInfo_r17_musim_GapLength_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MUSIM_GapInfo_r17_musim_GapLength_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapInfo_r17_musim_GapLength_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MUSIM_GapInfo_r17_musim_GapLength_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
