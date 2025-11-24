package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoDAPS_v1640 struct {
	SourceFeatureSetPerDownlinkCC_r16 FeatureSetDownlinkPerCC_Id `madatory`
	SourceFeatureSetPerUplinkCC_r16   FeatureSetUplinkPerCC_Id   `madatory`
}

func (ie *ConfigRestrictInfoDAPS_v1640) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.SourceFeatureSetPerDownlinkCC_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SourceFeatureSetPerDownlinkCC_r16", err)
	}
	if err = ie.SourceFeatureSetPerUplinkCC_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SourceFeatureSetPerUplinkCC_r16", err)
	}
	return nil
}

func (ie *ConfigRestrictInfoDAPS_v1640) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.SourceFeatureSetPerDownlinkCC_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SourceFeatureSetPerDownlinkCC_r16", err)
	}
	if err = ie.SourceFeatureSetPerUplinkCC_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SourceFeatureSetPerUplinkCC_r16", err)
	}
	return nil
}
