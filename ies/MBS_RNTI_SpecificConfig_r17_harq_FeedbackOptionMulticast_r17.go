package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17_Enum_ack_nack  aper.Enumerated = 0
	MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17_Enum_nack_only aper.Enumerated = 1
)

type MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17", err)
	}
	return nil
}

func (ie *MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
