package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosRRC_InactiveConfig_r17 struct {
	Srs_PosConfigNUL_r17                    *SRS_PosConfig_r17        `optional`
	Srs_PosConfigSUL_r17                    *SRS_PosConfig_r17        `optional`
	Bwp_NUL_r17                             *BWP                      `optional`
	Bwp_SUL_r17                             *BWP                      `optional`
	InactivePosSRS_TimeAlignmentTimer_r17   *TimeAlignmentTimer       `optional`
	InactivePosSRS_RSRP_ChangeThreshold_r17 *RSRP_ChangeThreshold_r17 `optional`
}

func (ie *SRS_PosRRC_InactiveConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_PosConfigNUL_r17 != nil, ie.Srs_PosConfigSUL_r17 != nil, ie.Bwp_NUL_r17 != nil, ie.Bwp_SUL_r17 != nil, ie.InactivePosSRS_TimeAlignmentTimer_r17 != nil, ie.InactivePosSRS_RSRP_ChangeThreshold_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_PosConfigNUL_r17 != nil {
		if err = ie.Srs_PosConfigNUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosConfigNUL_r17", err)
		}
	}
	if ie.Srs_PosConfigSUL_r17 != nil {
		if err = ie.Srs_PosConfigSUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosConfigSUL_r17", err)
		}
	}
	if ie.Bwp_NUL_r17 != nil {
		if err = ie.Bwp_NUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_NUL_r17", err)
		}
	}
	if ie.Bwp_SUL_r17 != nil {
		if err = ie.Bwp_SUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_SUL_r17", err)
		}
	}
	if ie.InactivePosSRS_TimeAlignmentTimer_r17 != nil {
		if err = ie.InactivePosSRS_TimeAlignmentTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InactivePosSRS_TimeAlignmentTimer_r17", err)
		}
	}
	if ie.InactivePosSRS_RSRP_ChangeThreshold_r17 != nil {
		if err = ie.InactivePosSRS_RSRP_ChangeThreshold_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InactivePosSRS_RSRP_ChangeThreshold_r17", err)
		}
	}
	return nil
}

func (ie *SRS_PosRRC_InactiveConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Srs_PosConfigNUL_r17Present bool
	if Srs_PosConfigNUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_PosConfigSUL_r17Present bool
	if Srs_PosConfigSUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_NUL_r17Present bool
	if Bwp_NUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_SUL_r17Present bool
	if Bwp_SUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InactivePosSRS_TimeAlignmentTimer_r17Present bool
	if InactivePosSRS_TimeAlignmentTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InactivePosSRS_RSRP_ChangeThreshold_r17Present bool
	if InactivePosSRS_RSRP_ChangeThreshold_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_PosConfigNUL_r17Present {
		ie.Srs_PosConfigNUL_r17 = new(SRS_PosConfig_r17)
		if err = ie.Srs_PosConfigNUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_PosConfigNUL_r17", err)
		}
	}
	if Srs_PosConfigSUL_r17Present {
		ie.Srs_PosConfigSUL_r17 = new(SRS_PosConfig_r17)
		if err = ie.Srs_PosConfigSUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_PosConfigSUL_r17", err)
		}
	}
	if Bwp_NUL_r17Present {
		ie.Bwp_NUL_r17 = new(BWP)
		if err = ie.Bwp_NUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_NUL_r17", err)
		}
	}
	if Bwp_SUL_r17Present {
		ie.Bwp_SUL_r17 = new(BWP)
		if err = ie.Bwp_SUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_SUL_r17", err)
		}
	}
	if InactivePosSRS_TimeAlignmentTimer_r17Present {
		ie.InactivePosSRS_TimeAlignmentTimer_r17 = new(TimeAlignmentTimer)
		if err = ie.InactivePosSRS_TimeAlignmentTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InactivePosSRS_TimeAlignmentTimer_r17", err)
		}
	}
	if InactivePosSRS_RSRP_ChangeThreshold_r17Present {
		ie.InactivePosSRS_RSRP_ChangeThreshold_r17 = new(RSRP_ChangeThreshold_r17)
		if err = ie.InactivePosSRS_RSRP_ChangeThreshold_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InactivePosSRS_RSRP_ChangeThreshold_r17", err)
		}
	}
	return nil
}
