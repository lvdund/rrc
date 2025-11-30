package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17_Enum_n8  aper.Enumerated = 0
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17_Enum_n16 aper.Enumerated = 1
)

type BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17", err)
	}
	return nil
}

func (ie *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
