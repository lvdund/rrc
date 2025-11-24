package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResource_r16 struct {
	Srs_PosResourceId_r16      SRS_PosResourceId_r16                          `madatory`
	TransmissionComb_r16       SRS_PosResource_r16_transmissionComb_r16       `lb:0,ub:1,madatory`
	ResourceMapping_r16        SRS_PosResource_r16_resourceMapping_r16        `lb:0,ub:13,madatory,ext`
	FreqDomainShift_r16        int64                                          `lb:0,ub:268,madatory,ext`
	FreqHopping_r16            SRS_PosResource_r16_freqHopping_r16            `lb:0,ub:63,madatory,ext`
	GroupOrSequenceHopping_r16 SRS_PosResource_r16_groupOrSequenceHopping_r16 `madatory,ext`
	ResourceType_r16           Resourcetype_r16_SRS_PosResource_r16           `madatory,ext`
	SequenceId_r16             int64                                          `lb:0,ub:65535,madatory,ext`
	SpatialRelationInfoPos_r16 *SRS_SpatialRelationInfoPos_r16                `optional,ext`
}

func (ie *SRS_PosResource_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Srs_PosResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_PosResourceId_r16", err)
	}
	if err = ie.TransmissionComb_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TransmissionComb_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Srs_PosResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_PosResourceId_r16", err)
	}
	if err = ie.TransmissionComb_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TransmissionComb_r16", err)
	}
	return nil
}
