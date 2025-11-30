package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n4  aper.Enumerated = 0
	BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n8  aper.Enumerated = 1
	BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n16 aper.Enumerated = 2
)

type BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17", err)
	}
	return nil
}

func (ie *BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
