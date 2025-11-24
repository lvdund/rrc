package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RadioLinkMonitoringRS struct {
	RadioLinkMonitoringRS_Id RadioLinkMonitoringRS_Id                `madatory`
	Purpose                  RadioLinkMonitoringRS_purpose           `madatory`
	DetectionResource        RadioLinkMonitoringRS_detectionResource `madatory`
}

func (ie *RadioLinkMonitoringRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.RadioLinkMonitoringRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringRS_Id", err)
	}
	if err = ie.Purpose.Encode(w); err != nil {
		return utils.WrapError("Encode Purpose", err)
	}
	if err = ie.DetectionResource.Encode(w); err != nil {
		return utils.WrapError("Encode DetectionResource", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringRS) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.RadioLinkMonitoringRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringRS_Id", err)
	}
	if err = ie.Purpose.Decode(r); err != nil {
		return utils.WrapError("Decode Purpose", err)
	}
	if err = ie.DetectionResource.Decode(r); err != nil {
		return utils.WrapError("Decode DetectionResource", err)
	}
	return nil
}
