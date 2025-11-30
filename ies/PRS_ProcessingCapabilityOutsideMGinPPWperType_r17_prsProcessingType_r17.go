package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17_Enum_type1A aper.Enumerated = 0
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17_Enum_type1B aper.Enumerated = 1
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17_Enum_type2  aper.Enumerated = 2
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17", err)
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
