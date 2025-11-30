package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DMRS_UplinkConfig_dmrs_AdditionalPosition_Enum_pos0 aper.Enumerated = 0
	DMRS_UplinkConfig_dmrs_AdditionalPosition_Enum_pos1 aper.Enumerated = 1
	DMRS_UplinkConfig_dmrs_AdditionalPosition_Enum_pos3 aper.Enumerated = 2
)

type DMRS_UplinkConfig_dmrs_AdditionalPosition struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DMRS_UplinkConfig_dmrs_AdditionalPosition) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DMRS_UplinkConfig_dmrs_AdditionalPosition", err)
	}
	return nil
}

func (ie *DMRS_UplinkConfig_dmrs_AdditionalPosition) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DMRS_UplinkConfig_dmrs_AdditionalPosition", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
