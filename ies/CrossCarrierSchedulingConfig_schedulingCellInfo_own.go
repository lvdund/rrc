package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig_schedulingCellInfo_own struct {
	Cif_Presence bool `madatory`
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_own) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Cif_Presence); err != nil {
		return utils.WrapError("WriteBoolean Cif_Presence", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_own) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Cif_Presence bool
	if tmp_bool_Cif_Presence, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Cif_Presence", err)
	}
	ie.Cif_Presence = tmp_bool_Cif_Presence
	return nil
}
