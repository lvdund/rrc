package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n5  aper.Enumerated = 0
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n15 aper.Enumerated = 1
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n25 aper.Enumerated = 2
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n32 aper.Enumerated = 3
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n35 aper.Enumerated = 4
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n45 aper.Enumerated = 5
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n50 aper.Enumerated = 6
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n64 aper.Enumerated = 7
)

type BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17", err)
	}
	return nil
}

func (ie *BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
