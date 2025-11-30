package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_eutra          aper.Enumerated = 0
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_utra_fdd_v1610 aper.Enumerated = 1
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_spare2         aper.Enumerated = 2
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_spare1         aper.Enumerated = 3
)

type MobilityFromNRCommand_IEs_targetRAT_Type struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MobilityFromNRCommand_IEs_targetRAT_Type) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MobilityFromNRCommand_IEs_targetRAT_Type", err)
	}
	return nil
}

func (ie *MobilityFromNRCommand_IEs_targetRAT_Type) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MobilityFromNRCommand_IEs_targetRAT_Type", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
