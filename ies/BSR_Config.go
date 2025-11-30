package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BSR_Config struct {
	PeriodicBSR_Timer           BSR_Config_periodicBSR_Timer            `madatory`
	RetxBSR_Timer               BSR_Config_retxBSR_Timer                `madatory`
	LogicalChannelSR_DelayTimer *BSR_Config_logicalChannelSR_DelayTimer `optional`
}

func (ie *BSR_Config) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LogicalChannelSR_DelayTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PeriodicBSR_Timer.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodicBSR_Timer", err)
	}
	if err = ie.RetxBSR_Timer.Encode(w); err != nil {
		return utils.WrapError("Encode RetxBSR_Timer", err)
	}
	if ie.LogicalChannelSR_DelayTimer != nil {
		if err = ie.LogicalChannelSR_DelayTimer.Encode(w); err != nil {
			return utils.WrapError("Encode LogicalChannelSR_DelayTimer", err)
		}
	}
	return nil
}

func (ie *BSR_Config) Decode(r *aper.AperReader) error {
	var err error
	var LogicalChannelSR_DelayTimerPresent bool
	if LogicalChannelSR_DelayTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PeriodicBSR_Timer.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodicBSR_Timer", err)
	}
	if err = ie.RetxBSR_Timer.Decode(r); err != nil {
		return utils.WrapError("Decode RetxBSR_Timer", err)
	}
	if LogicalChannelSR_DelayTimerPresent {
		ie.LogicalChannelSR_DelayTimer = new(BSR_Config_logicalChannelSR_DelayTimer)
		if err = ie.LogicalChannelSR_DelayTimer.Decode(r); err != nil {
			return utils.WrapError("Decode LogicalChannelSR_DelayTimer", err)
		}
	}
	return nil
}
