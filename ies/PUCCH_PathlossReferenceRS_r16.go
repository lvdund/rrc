package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS_r16 struct {
	Pucch_PathlossReferenceRS_Id_r16 PUCCH_PathlossReferenceRS_Id_v1610                `madatory`
	ReferenceSignal_r16              PUCCH_PathlossReferenceRS_r16_referenceSignal_r16 `madatory`
}

func (ie *PUCCH_PathlossReferenceRS_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pucch_PathlossReferenceRS_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.ReferenceSignal_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal_r16", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pucch_PathlossReferenceRS_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.ReferenceSignal_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal_r16", err)
	}
	return nil
}
