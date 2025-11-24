package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1700_IEs struct {
	NeedForGapNCSG_InfoNR_r17    *NeedForGapNCSG_InfoNR_r17            `optional`
	NeedForGapNCSG_InfoEUTRA_r17 *NeedForGapNCSG_InfoEUTRA_r17         `optional`
	SelectedCondRRCReconfig_r17  *CondReconfigId_r16                   `optional`
	NonCriticalExtension         *RRCReconfigurationComplete_v1720_IEs `optional`
}

func (ie *RRCReconfigurationComplete_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NeedForGapNCSG_InfoNR_r17 != nil, ie.NeedForGapNCSG_InfoEUTRA_r17 != nil, ie.SelectedCondRRCReconfig_r17 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NeedForGapNCSG_InfoNR_r17 != nil {
		if err = ie.NeedForGapNCSG_InfoNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapNCSG_InfoNR_r17", err)
		}
	}
	if ie.NeedForGapNCSG_InfoEUTRA_r17 != nil {
		if err = ie.NeedForGapNCSG_InfoEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapNCSG_InfoEUTRA_r17", err)
		}
	}
	if ie.SelectedCondRRCReconfig_r17 != nil {
		if err = ie.SelectedCondRRCReconfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SelectedCondRRCReconfig_r17", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var NeedForGapNCSG_InfoNR_r17Present bool
	if NeedForGapNCSG_InfoNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapNCSG_InfoEUTRA_r17Present bool
	if NeedForGapNCSG_InfoEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SelectedCondRRCReconfig_r17Present bool
	if SelectedCondRRCReconfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if NeedForGapNCSG_InfoNR_r17Present {
		ie.NeedForGapNCSG_InfoNR_r17 = new(NeedForGapNCSG_InfoNR_r17)
		if err = ie.NeedForGapNCSG_InfoNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapNCSG_InfoNR_r17", err)
		}
	}
	if NeedForGapNCSG_InfoEUTRA_r17Present {
		ie.NeedForGapNCSG_InfoEUTRA_r17 = new(NeedForGapNCSG_InfoEUTRA_r17)
		if err = ie.NeedForGapNCSG_InfoEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapNCSG_InfoEUTRA_r17", err)
		}
	}
	if SelectedCondRRCReconfig_r17Present {
		ie.SelectedCondRRCReconfig_r17 = new(CondReconfigId_r16)
		if err = ie.SelectedCondRRCReconfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SelectedCondRRCReconfig_r17", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1720_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
