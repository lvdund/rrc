package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17_Enum_enabled aper.Enumerated = 0
)

type DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17", err)
	}
	return nil
}

func (ie *DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
