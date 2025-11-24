package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersSidelinkXDD_Diff_r16 struct {
	MultipleSR_ConfigurationsSidelink_r16   *MAC_ParametersSidelinkXDD_Diff_r16_multipleSR_ConfigurationsSidelink_r16   `optional`
	LogicalChannelSR_DelayTimerSidelink_r16 *MAC_ParametersSidelinkXDD_Diff_r16_logicalChannelSR_DelayTimerSidelink_r16 `optional`
}

func (ie *MAC_ParametersSidelinkXDD_Diff_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MultipleSR_ConfigurationsSidelink_r16 != nil, ie.LogicalChannelSR_DelayTimerSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MultipleSR_ConfigurationsSidelink_r16 != nil {
		if err = ie.MultipleSR_ConfigurationsSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleSR_ConfigurationsSidelink_r16", err)
		}
	}
	if ie.LogicalChannelSR_DelayTimerSidelink_r16 != nil {
		if err = ie.LogicalChannelSR_DelayTimerSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogicalChannelSR_DelayTimerSidelink_r16", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersSidelinkXDD_Diff_r16) Decode(r *uper.UperReader) error {
	var err error
	var MultipleSR_ConfigurationsSidelink_r16Present bool
	if MultipleSR_ConfigurationsSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogicalChannelSR_DelayTimerSidelink_r16Present bool
	if LogicalChannelSR_DelayTimerSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MultipleSR_ConfigurationsSidelink_r16Present {
		ie.MultipleSR_ConfigurationsSidelink_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16_multipleSR_ConfigurationsSidelink_r16)
		if err = ie.MultipleSR_ConfigurationsSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleSR_ConfigurationsSidelink_r16", err)
		}
	}
	if LogicalChannelSR_DelayTimerSidelink_r16Present {
		ie.LogicalChannelSR_DelayTimerSidelink_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16_logicalChannelSR_DelayTimerSidelink_r16)
		if err = ie.LogicalChannelSR_DelayTimerSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogicalChannelSR_DelayTimerSidelink_r16", err)
		}
	}
	return nil
}
