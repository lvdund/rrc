package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SubSlot_Config_r16 struct {
	Sub_SlotConfig_NCP_r16 *SubSlot_Config_r16_sub_SlotConfig_NCP_r16 `optional`
	Sub_SlotConfig_ECP_r16 *SubSlot_Config_r16_sub_SlotConfig_ECP_r16 `optional`
}

func (ie *SubSlot_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sub_SlotConfig_NCP_r16 != nil, ie.Sub_SlotConfig_ECP_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sub_SlotConfig_NCP_r16 != nil {
		if err = ie.Sub_SlotConfig_NCP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sub_SlotConfig_NCP_r16", err)
		}
	}
	if ie.Sub_SlotConfig_ECP_r16 != nil {
		if err = ie.Sub_SlotConfig_ECP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sub_SlotConfig_ECP_r16", err)
		}
	}
	return nil
}

func (ie *SubSlot_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sub_SlotConfig_NCP_r16Present bool
	if Sub_SlotConfig_NCP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sub_SlotConfig_ECP_r16Present bool
	if Sub_SlotConfig_ECP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sub_SlotConfig_NCP_r16Present {
		ie.Sub_SlotConfig_NCP_r16 = new(SubSlot_Config_r16_sub_SlotConfig_NCP_r16)
		if err = ie.Sub_SlotConfig_NCP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sub_SlotConfig_NCP_r16", err)
		}
	}
	if Sub_SlotConfig_ECP_r16Present {
		ie.Sub_SlotConfig_ECP_r16 = new(SubSlot_Config_r16_sub_SlotConfig_ECP_r16)
		if err = ie.Sub_SlotConfig_ECP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sub_SlotConfig_ECP_r16", err)
		}
	}
	return nil
}
