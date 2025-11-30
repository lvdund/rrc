package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db_3 aper.Enumerated = 0
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db0  aper.Enumerated = 1
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db3  aper.Enumerated = 2
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db6  aper.Enumerated = 3
)

type NZP_CSI_RS_Resource_powerControlOffsetSS struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *NZP_CSI_RS_Resource_powerControlOffsetSS) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode NZP_CSI_RS_Resource_powerControlOffsetSS", err)
	}
	return nil
}

func (ie *NZP_CSI_RS_Resource_powerControlOffsetSS) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode NZP_CSI_RS_Resource_powerControlOffsetSS", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
