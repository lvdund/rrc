package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SRS struct {
	Resource  SRS_ResourceId `madatory`
	UplinkBWP BWP_Id         `madatory`
}

func (ie *PUCCH_SRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Resource.Encode(w); err != nil {
		return utils.WrapError("Encode Resource", err)
	}
	if err = ie.UplinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkBWP", err)
	}
	return nil
}

func (ie *PUCCH_SRS) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Resource.Decode(r); err != nil {
		return utils.WrapError("Decode Resource", err)
	}
	if err = ie.UplinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkBWP", err)
	}
	return nil
}
