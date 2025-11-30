package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSRS_RSRP_r16 struct {
	Srs_ResourceId_r16  SRS_ResourceId     `madatory`
	Srs_RSRP_Result_r16 SRS_RSRP_Range_r16 `madatory`
}

func (ie *MeasResultSRS_RSRP_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Srs_ResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_ResourceId_r16", err)
	}
	if err = ie.Srs_RSRP_Result_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_RSRP_Result_r16", err)
	}
	return nil
}

func (ie *MeasResultSRS_RSRP_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Srs_ResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_ResourceId_r16", err)
	}
	if err = ie.Srs_RSRP_Result_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_RSRP_Result_r16", err)
	}
	return nil
}
