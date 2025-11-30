package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PathlossReferenceRS_r16 struct {
	Srs_PathlossReferenceRS_Id_r16 SRS_PathlossReferenceRS_Id_r16 `madatory`
	PathlossReferenceRS_r16        PathlossReferenceRS_Config     `madatory`
}

func (ie *PathlossReferenceRS_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Srs_PathlossReferenceRS_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.PathlossReferenceRS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PathlossReferenceRS_r16", err)
	}
	return nil
}

func (ie *PathlossReferenceRS_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Srs_PathlossReferenceRS_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.PathlossReferenceRS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PathlossReferenceRS_r16", err)
	}
	return nil
}
