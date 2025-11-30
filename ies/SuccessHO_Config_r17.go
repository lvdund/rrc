package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Config_r17 struct {
	ThresholdPercentageT304_r17     *SuccessHO_Config_r17_thresholdPercentageT304_r17     `optional`
	ThresholdPercentageT310_r17     *SuccessHO_Config_r17_thresholdPercentageT310_r17     `optional`
	ThresholdPercentageT312_r17     *SuccessHO_Config_r17_thresholdPercentageT312_r17     `optional`
	SourceDAPS_FailureReporting_r17 *SuccessHO_Config_r17_sourceDAPS_FailureReporting_r17 `optional`
}

func (ie *SuccessHO_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ThresholdPercentageT304_r17 != nil, ie.ThresholdPercentageT310_r17 != nil, ie.ThresholdPercentageT312_r17 != nil, ie.SourceDAPS_FailureReporting_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ThresholdPercentageT304_r17 != nil {
		if err = ie.ThresholdPercentageT304_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThresholdPercentageT304_r17", err)
		}
	}
	if ie.ThresholdPercentageT310_r17 != nil {
		if err = ie.ThresholdPercentageT310_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThresholdPercentageT310_r17", err)
		}
	}
	if ie.ThresholdPercentageT312_r17 != nil {
		if err = ie.ThresholdPercentageT312_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThresholdPercentageT312_r17", err)
		}
	}
	if ie.SourceDAPS_FailureReporting_r17 != nil {
		if err = ie.SourceDAPS_FailureReporting_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SourceDAPS_FailureReporting_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var ThresholdPercentageT304_r17Present bool
	if ThresholdPercentageT304_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ThresholdPercentageT310_r17Present bool
	if ThresholdPercentageT310_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ThresholdPercentageT312_r17Present bool
	if ThresholdPercentageT312_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SourceDAPS_FailureReporting_r17Present bool
	if SourceDAPS_FailureReporting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ThresholdPercentageT304_r17Present {
		ie.ThresholdPercentageT304_r17 = new(SuccessHO_Config_r17_thresholdPercentageT304_r17)
		if err = ie.ThresholdPercentageT304_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThresholdPercentageT304_r17", err)
		}
	}
	if ThresholdPercentageT310_r17Present {
		ie.ThresholdPercentageT310_r17 = new(SuccessHO_Config_r17_thresholdPercentageT310_r17)
		if err = ie.ThresholdPercentageT310_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThresholdPercentageT310_r17", err)
		}
	}
	if ThresholdPercentageT312_r17Present {
		ie.ThresholdPercentageT312_r17 = new(SuccessHO_Config_r17_thresholdPercentageT312_r17)
		if err = ie.ThresholdPercentageT312_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThresholdPercentageT312_r17", err)
		}
	}
	if SourceDAPS_FailureReporting_r17Present {
		ie.SourceDAPS_FailureReporting_r17 = new(SuccessHO_Config_r17_sourceDAPS_FailureReporting_r17)
		if err = ie.SourceDAPS_FailureReporting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SourceDAPS_FailureReporting_r17", err)
		}
	}
	return nil
}
