package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DeactivatedSCG_Config_r17 struct {
	Bfd_and_RLM_r17 bool `madatory`
}

func (ie *DeactivatedSCG_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Bfd_and_RLM_r17); err != nil {
		return utils.WrapError("WriteBoolean Bfd_and_RLM_r17", err)
	}
	return nil
}

func (ie *DeactivatedSCG_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Bfd_and_RLM_r17 bool
	if tmp_bool_Bfd_and_RLM_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Bfd_and_RLM_r17", err)
	}
	ie.Bfd_and_RLM_r17 = tmp_bool_Bfd_and_RLM_r17
	return nil
}
