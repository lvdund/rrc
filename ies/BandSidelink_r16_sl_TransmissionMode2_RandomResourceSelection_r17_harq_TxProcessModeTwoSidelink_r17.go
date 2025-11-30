package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17_Enum_n8  aper.Enumerated = 0
	BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17_Enum_n16 aper.Enumerated = 1
)

type BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17", err)
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
