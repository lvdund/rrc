package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PathlossReferenceRS_r16 struct {
	Pusch_PathlossReferenceRS_Id_r16 PUSCH_PathlossReferenceRS_Id_v1610                `madatory`
	ReferenceSignal_r16              PUSCH_PathlossReferenceRS_r16_referenceSignal_r16 `madatory`
}

func (ie *PUSCH_PathlossReferenceRS_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pusch_PathlossReferenceRS_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pusch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.ReferenceSignal_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal_r16", err)
	}
	return nil
}

func (ie *PUSCH_PathlossReferenceRS_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pusch_PathlossReferenceRS_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pusch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.ReferenceSignal_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal_r16", err)
	}
	return nil
}
