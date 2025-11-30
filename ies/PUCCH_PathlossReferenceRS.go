package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS struct {
	Pucch_PathlossReferenceRS_Id PUCCH_PathlossReferenceRS_Id              `madatory`
	ReferenceSignal              PUCCH_PathlossReferenceRS_referenceSignal `madatory`
}

func (ie *PUCCH_PathlossReferenceRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pucch_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.ReferenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pucch_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.ReferenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal", err)
	}
	return nil
}
