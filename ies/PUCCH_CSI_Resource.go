package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_CSI_Resource struct {
	UplinkBandwidthPartId BWP_Id           `madatory`
	Pucch_Resource        PUCCH_ResourceId `madatory`
}

func (ie *PUCCH_CSI_Resource) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.UplinkBandwidthPartId.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkBandwidthPartId", err)
	}
	if err = ie.Pucch_Resource.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_Resource", err)
	}
	return nil
}

func (ie *PUCCH_CSI_Resource) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.UplinkBandwidthPartId.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkBandwidthPartId", err)
	}
	if err = ie.Pucch_Resource.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_Resource", err)
	}
	return nil
}
