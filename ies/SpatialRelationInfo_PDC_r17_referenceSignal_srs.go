package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SpatialRelationInfo_PDC_r17_referenceSignal_srs struct {
	ResourceId SRS_ResourceId `madatory`
	UplinkBWP  BWP_Id         `madatory`
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal_srs) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceId", err)
	}
	if err = ie.UplinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkBWP", err)
	}
	return nil
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal_srs) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceId", err)
	}
	if err = ie.UplinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkBWP", err)
	}
	return nil
}
