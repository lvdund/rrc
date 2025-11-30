package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TAR_Config_r17 struct {
	OffsetThresholdTA_r17 *TAR_Config_r17_offsetThresholdTA_r17 `optional`
	TimingAdvanceSR_r17   *TAR_Config_r17_timingAdvanceSR_r17   `optional`
}

func (ie *TAR_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OffsetThresholdTA_r17 != nil, ie.TimingAdvanceSR_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OffsetThresholdTA_r17 != nil {
		if err = ie.OffsetThresholdTA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetThresholdTA_r17", err)
		}
	}
	if ie.TimingAdvanceSR_r17 != nil {
		if err = ie.TimingAdvanceSR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TimingAdvanceSR_r17", err)
		}
	}
	return nil
}

func (ie *TAR_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var OffsetThresholdTA_r17Present bool
	if OffsetThresholdTA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimingAdvanceSR_r17Present bool
	if TimingAdvanceSR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OffsetThresholdTA_r17Present {
		ie.OffsetThresholdTA_r17 = new(TAR_Config_r17_offsetThresholdTA_r17)
		if err = ie.OffsetThresholdTA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetThresholdTA_r17", err)
		}
	}
	if TimingAdvanceSR_r17Present {
		ie.TimingAdvanceSR_r17 = new(TAR_Config_r17_timingAdvanceSR_r17)
		if err = ie.TimingAdvanceSR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TimingAdvanceSR_r17", err)
		}
	}
	return nil
}
