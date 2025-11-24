package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16 struct {
	MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResWithinSlotAcrossCC_AcrossFR_r16 `optional`
	MaxNumberResAcrossCC_AcrossFR_r16           *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16           `optional`
}

func (ie *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 != nil, ie.MaxNumberResAcrossCC_AcrossFR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 != nil {
		if err = ie.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberResWithinSlotAcrossCC_AcrossFR_r16", err)
		}
	}
	if ie.MaxNumberResAcrossCC_AcrossFR_r16 != nil {
		if err = ie.MaxNumberResAcrossCC_AcrossFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberResAcrossCC_AcrossFR_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16) Decode(r *uper.UperReader) error {
	var err error
	var MaxNumberResWithinSlotAcrossCC_AcrossFR_r16Present bool
	if MaxNumberResWithinSlotAcrossCC_AcrossFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberResAcrossCC_AcrossFR_r16Present bool
	if MaxNumberResAcrossCC_AcrossFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxNumberResWithinSlotAcrossCC_AcrossFR_r16Present {
		ie.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 = new(Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResWithinSlotAcrossCC_AcrossFR_r16)
		if err = ie.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberResWithinSlotAcrossCC_AcrossFR_r16", err)
		}
	}
	if MaxNumberResAcrossCC_AcrossFR_r16Present {
		ie.MaxNumberResAcrossCC_AcrossFR_r16 = new(Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16)
		if err = ie.MaxNumberResAcrossCC_AcrossFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberResAcrossCC_AcrossFR_r16", err)
		}
	}
	return nil
}
