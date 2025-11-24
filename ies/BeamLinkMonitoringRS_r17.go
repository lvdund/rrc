package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamLinkMonitoringRS_r17 struct {
	BeamLinkMonitoringRS_Id_r17 BeamLinkMonitoringRS_Id_r17                    `madatory`
	DetectionResource_r17       BeamLinkMonitoringRS_r17_detectionResource_r17 `madatory`
}

func (ie *BeamLinkMonitoringRS_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.BeamLinkMonitoringRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode BeamLinkMonitoringRS_Id_r17", err)
	}
	if err = ie.DetectionResource_r17.Encode(w); err != nil {
		return utils.WrapError("Encode DetectionResource_r17", err)
	}
	return nil
}

func (ie *BeamLinkMonitoringRS_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.BeamLinkMonitoringRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode BeamLinkMonitoringRS_Id_r17", err)
	}
	if err = ie.DetectionResource_r17.Decode(r); err != nil {
		return utils.WrapError("Decode DetectionResource_r17", err)
	}
	return nil
}
