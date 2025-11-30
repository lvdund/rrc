package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSet_eutra struct {
	DownlinkSetEUTRA FeatureSetEUTRA_DownlinkId `madatory`
	UplinkSetEUTRA   FeatureSetEUTRA_UplinkId   `madatory`
}

func (ie *FeatureSet_eutra) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.DownlinkSetEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode DownlinkSetEUTRA", err)
	}
	if err = ie.UplinkSetEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkSetEUTRA", err)
	}
	return nil
}

func (ie *FeatureSet_eutra) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.DownlinkSetEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode DownlinkSetEUTRA", err)
	}
	if err = ie.UplinkSetEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkSetEUTRA", err)
	}
	return nil
}
