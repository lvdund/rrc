package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PathlossReferenceRS struct {
	Pusch_PathlossReferenceRS_Id PUSCH_PathlossReferenceRS_Id              `madatory`
	ReferenceSignal              PUSCH_PathlossReferenceRS_referenceSignal `madatory`
}

func (ie *PUSCH_PathlossReferenceRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pusch_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Pusch_PathlossReferenceRS_Id", err)
	}
	if err = ie.ReferenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal", err)
	}
	return nil
}

func (ie *PUSCH_PathlossReferenceRS) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pusch_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Pusch_PathlossReferenceRS_Id", err)
	}
	if err = ie.ReferenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal", err)
	}
	return nil
}
