package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_groupBasedBeamReporting_Choice_nothing uint64 = iota
	CSI_ReportConfig_groupBasedBeamReporting_Choice_Enabled
	CSI_ReportConfig_groupBasedBeamReporting_Choice_Disabled
)

type CSI_ReportConfig_groupBasedBeamReporting struct {
	Choice   uint64
	Enabled  aper.NULL `madatory`
	Disabled *CSI_ReportConfig_groupBasedBeamReporting_disabled
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_Enabled:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Enabled", err)
		}
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_Disabled:
		if err = ie.Disabled.Encode(w); err != nil {
			err = utils.WrapError("Encode Disabled", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_Enabled:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Enabled", err)
		}
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_Disabled:
		ie.Disabled = new(CSI_ReportConfig_groupBasedBeamReporting_disabled)
		if err = ie.Disabled.Decode(r); err != nil {
			return utils.WrapError("Decode Disabled", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
