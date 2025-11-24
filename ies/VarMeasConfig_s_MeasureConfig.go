package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	VarMeasConfig_s_MeasureConfig_Choice_nothing uint64 = iota
	VarMeasConfig_s_MeasureConfig_Choice_Ssb_RSRP
	VarMeasConfig_s_MeasureConfig_Choice_Csi_RSRP
)

type VarMeasConfig_s_MeasureConfig struct {
	Choice   uint64
	Ssb_RSRP *RSRP_Range
	Csi_RSRP *RSRP_Range
}

func (ie *VarMeasConfig_s_MeasureConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VarMeasConfig_s_MeasureConfig_Choice_Ssb_RSRP:
		if err = ie.Ssb_RSRP.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_RSRP", err)
		}
	case VarMeasConfig_s_MeasureConfig_Choice_Csi_RSRP:
		if err = ie.Csi_RSRP.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RSRP", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *VarMeasConfig_s_MeasureConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VarMeasConfig_s_MeasureConfig_Choice_Ssb_RSRP:
		ie.Ssb_RSRP = new(RSRP_Range)
		if err = ie.Ssb_RSRP.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RSRP", err)
		}
	case VarMeasConfig_s_MeasureConfig_Choice_Csi_RSRP:
		ie.Csi_RSRP = new(RSRP_Range)
		if err = ie.Csi_RSRP.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RSRP", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
