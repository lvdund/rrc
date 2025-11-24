package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSet_nr struct {
	DownlinkSetNR FeatureSetDownlinkId `madatory`
	UplinkSetNR   FeatureSetUplinkId   `madatory`
}

func (ie *FeatureSet_nr) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.DownlinkSetNR.Encode(w); err != nil {
		return utils.WrapError("Encode DownlinkSetNR", err)
	}
	if err = ie.UplinkSetNR.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkSetNR", err)
	}
	return nil
}

func (ie *FeatureSet_nr) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.DownlinkSetNR.Decode(r); err != nil {
		return utils.WrapError("Decode DownlinkSetNR", err)
	}
	if err = ie.UplinkSetNR.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkSetNR", err)
	}
	return nil
}
