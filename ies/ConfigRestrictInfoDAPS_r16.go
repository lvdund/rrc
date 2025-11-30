package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoDAPS_r16 struct {
	PowerCoordination_r16 *ConfigRestrictInfoDAPS_r16_powerCoordination_r16 `optional`
}

func (ie *ConfigRestrictInfoDAPS_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PowerCoordination_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PowerCoordination_r16 != nil {
		if err = ie.PowerCoordination_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PowerCoordination_r16", err)
		}
	}
	return nil
}

func (ie *ConfigRestrictInfoDAPS_r16) Decode(r *aper.AperReader) error {
	var err error
	var PowerCoordination_r16Present bool
	if PowerCoordination_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PowerCoordination_r16Present {
		ie.PowerCoordination_r16 = new(ConfigRestrictInfoDAPS_r16_powerCoordination_r16)
		if err = ie.PowerCoordination_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PowerCoordination_r16", err)
		}
	}
	return nil
}
