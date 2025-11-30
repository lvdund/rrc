package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber_Enum_n4  aper.Enumerated = 0
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber_Enum_n8  aper.Enumerated = 1
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber_Enum_n16 aper.Enumerated = 2
)

type BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber", err)
	}
	return nil
}

func (ie *BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
