package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdToAddMod struct {
	MeasId         MeasId         `madatory`
	MeasObjectId   MeasObjectId   `madatory`
	ReportConfigId ReportConfigId `madatory`
}

func (ie *MeasIdToAddMod) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MeasId.Encode(w); err != nil {
		return utils.WrapError("Encode MeasId", err)
	}
	if err = ie.MeasObjectId.Encode(w); err != nil {
		return utils.WrapError("Encode MeasObjectId", err)
	}
	if err = ie.ReportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigId", err)
	}
	return nil
}

func (ie *MeasIdToAddMod) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MeasId.Decode(r); err != nil {
		return utils.WrapError("Decode MeasId", err)
	}
	if err = ie.MeasObjectId.Decode(r); err != nil {
		return utils.WrapError("Decode MeasObjectId", err)
	}
	if err = ie.ReportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode ReportConfigId", err)
	}
	return nil
}
