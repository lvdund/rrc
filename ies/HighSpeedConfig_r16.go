package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedConfig_r16 struct {
	HighSpeedMeasFlag_r16  *HighSpeedConfig_r16_highSpeedMeasFlag_r16  `optional`
	HighSpeedDemodFlag_r16 *HighSpeedConfig_r16_highSpeedDemodFlag_r16 `optional`
}

func (ie *HighSpeedConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.HighSpeedMeasFlag_r16 != nil, ie.HighSpeedDemodFlag_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.HighSpeedMeasFlag_r16 != nil {
		if err = ie.HighSpeedMeasFlag_r16.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedMeasFlag_r16", err)
		}
	}
	if ie.HighSpeedDemodFlag_r16 != nil {
		if err = ie.HighSpeedDemodFlag_r16.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedDemodFlag_r16", err)
		}
	}
	return nil
}

func (ie *HighSpeedConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var HighSpeedMeasFlag_r16Present bool
	if HighSpeedMeasFlag_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedDemodFlag_r16Present bool
	if HighSpeedDemodFlag_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if HighSpeedMeasFlag_r16Present {
		ie.HighSpeedMeasFlag_r16 = new(HighSpeedConfig_r16_highSpeedMeasFlag_r16)
		if err = ie.HighSpeedMeasFlag_r16.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedMeasFlag_r16", err)
		}
	}
	if HighSpeedDemodFlag_r16Present {
		ie.HighSpeedDemodFlag_r16 = new(HighSpeedConfig_r16_highSpeedDemodFlag_r16)
		if err = ie.HighSpeedDemodFlag_r16.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedDemodFlag_r16", err)
		}
	}
	return nil
}
