package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 struct {
	MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17               *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17               `optional`
	MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17               *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17               `optional`
	MaxNumOfSRSposResourceSets_r17                              *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSRSposResourceSets_r17                              `optional`
	MaxNumOfPeriodicSRSposResources_r17                         *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResources_r17                         `optional`
	MaxNumOfPeriodicSRSposResourcesPerSlot_r17                  *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17                  `optional`
	DifferentNumerologyBetweenSRSposAndInitialBWP_r17           *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentNumerologyBetweenSRSposAndInitialBWP_r17           `optional`
	SrsPosWithoutRestrictionOnBWP_r17                           *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_srsPosWithoutRestrictionOnBWP_r17                           `optional`
	MaxNumOfPeriodicAndSemipersistentSRSposResources_r17        *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17        `optional`
	MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 `optional`
	DifferentCenterFreqBetweenSRSposAndInitialBWP_r17           *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentCenterFreqBetweenSRSposAndInitialBWP_r17           `optional`
	SwitchingTimeSRS_TX_OtherTX_r17                             *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17                             `optional`
	MaxNumOfSemiPersistentSRSposResources_r17                   *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResources_r17                   `optional`
	MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17            *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17            `optional`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17 != nil, ie.MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17 != nil, ie.MaxNumOfSRSposResourceSets_r17 != nil, ie.MaxNumOfPeriodicSRSposResources_r17 != nil, ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17 != nil, ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17 != nil, ie.SrsPosWithoutRestrictionOnBWP_r17 != nil, ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17 != nil, ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 != nil, ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17 != nil, ie.SwitchingTimeSRS_TX_OtherTX_r17 != nil, ie.MaxNumOfSemiPersistentSRSposResources_r17 != nil, ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17 != nil {
		if err = ie.MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
		}
	}
	if ie.MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17 != nil {
		if err = ie.MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17", err)
		}
	}
	if ie.MaxNumOfSRSposResourceSets_r17 != nil {
		if err = ie.MaxNumOfSRSposResourceSets_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfSRSposResourceSets_r17", err)
		}
	}
	if ie.MaxNumOfPeriodicSRSposResources_r17 != nil {
		if err = ie.MaxNumOfPeriodicSRSposResources_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfPeriodicSRSposResources_r17", err)
		}
	}
	if ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17 != nil {
		if err = ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfPeriodicSRSposResourcesPerSlot_r17", err)
		}
	}
	if ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17 != nil {
		if err = ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DifferentNumerologyBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if ie.SrsPosWithoutRestrictionOnBWP_r17 != nil {
		if err = ie.SrsPosWithoutRestrictionOnBWP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SrsPosWithoutRestrictionOnBWP_r17", err)
		}
	}
	if ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17 != nil {
		if err = ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfPeriodicAndSemipersistentSRSposResources_r17", err)
		}
	}
	if ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 != nil {
		if err = ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	if ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17 != nil {
		if err = ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DifferentCenterFreqBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if ie.SwitchingTimeSRS_TX_OtherTX_r17 != nil {
		if err = ie.SwitchingTimeSRS_TX_OtherTX_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SwitchingTimeSRS_TX_OtherTX_r17", err)
		}
	}
	if ie.MaxNumOfSemiPersistentSRSposResources_r17 != nil {
		if err = ie.MaxNumOfSemiPersistentSRSposResources_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfSemiPersistentSRSposResources_r17", err)
		}
	}
	if ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 != nil {
		if err = ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17) Decode(r *uper.UperReader) error {
	var err error
	var MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17Present bool
	if MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17Present bool
	if MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfSRSposResourceSets_r17Present bool
	if MaxNumOfSRSposResourceSets_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfPeriodicSRSposResources_r17Present bool
	if MaxNumOfPeriodicSRSposResources_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfPeriodicSRSposResourcesPerSlot_r17Present bool
	if MaxNumOfPeriodicSRSposResourcesPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DifferentNumerologyBetweenSRSposAndInitialBWP_r17Present bool
	if DifferentNumerologyBetweenSRSposAndInitialBWP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SrsPosWithoutRestrictionOnBWP_r17Present bool
	if SrsPosWithoutRestrictionOnBWP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfPeriodicAndSemipersistentSRSposResources_r17Present bool
	if MaxNumOfPeriodicAndSemipersistentSRSposResources_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17Present bool
	if MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DifferentCenterFreqBetweenSRSposAndInitialBWP_r17Present bool
	if DifferentCenterFreqBetweenSRSposAndInitialBWP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SwitchingTimeSRS_TX_OtherTX_r17Present bool
	if SwitchingTimeSRS_TX_OtherTX_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfSemiPersistentSRSposResources_r17Present bool
	if MaxNumOfSemiPersistentSRSposResources_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17Present bool
	if MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17Present {
		ie.MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17)
		if err = ie.MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
		}
	}
	if MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17Present {
		ie.MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17)
		if err = ie.MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxSRSposBandwidthForEachSCS_withinCC_FR2_r17", err)
		}
	}
	if MaxNumOfSRSposResourceSets_r17Present {
		ie.MaxNumOfSRSposResourceSets_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSRSposResourceSets_r17)
		if err = ie.MaxNumOfSRSposResourceSets_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfSRSposResourceSets_r17", err)
		}
	}
	if MaxNumOfPeriodicSRSposResources_r17Present {
		ie.MaxNumOfPeriodicSRSposResources_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResources_r17)
		if err = ie.MaxNumOfPeriodicSRSposResources_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfPeriodicSRSposResources_r17", err)
		}
	}
	if MaxNumOfPeriodicSRSposResourcesPerSlot_r17Present {
		ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17)
		if err = ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfPeriodicSRSposResourcesPerSlot_r17", err)
		}
	}
	if DifferentNumerologyBetweenSRSposAndInitialBWP_r17Present {
		ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentNumerologyBetweenSRSposAndInitialBWP_r17)
		if err = ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DifferentNumerologyBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if SrsPosWithoutRestrictionOnBWP_r17Present {
		ie.SrsPosWithoutRestrictionOnBWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_srsPosWithoutRestrictionOnBWP_r17)
		if err = ie.SrsPosWithoutRestrictionOnBWP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SrsPosWithoutRestrictionOnBWP_r17", err)
		}
	}
	if MaxNumOfPeriodicAndSemipersistentSRSposResources_r17Present {
		ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17)
		if err = ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfPeriodicAndSemipersistentSRSposResources_r17", err)
		}
	}
	if MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17Present {
		ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17)
		if err = ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	if DifferentCenterFreqBetweenSRSposAndInitialBWP_r17Present {
		ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentCenterFreqBetweenSRSposAndInitialBWP_r17)
		if err = ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DifferentCenterFreqBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if SwitchingTimeSRS_TX_OtherTX_r17Present {
		ie.SwitchingTimeSRS_TX_OtherTX_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17)
		if err = ie.SwitchingTimeSRS_TX_OtherTX_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SwitchingTimeSRS_TX_OtherTX_r17", err)
		}
	}
	if MaxNumOfSemiPersistentSRSposResources_r17Present {
		ie.MaxNumOfSemiPersistentSRSposResources_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResources_r17)
		if err = ie.MaxNumOfSemiPersistentSRSposResources_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfSemiPersistentSRSposResources_r17", err)
		}
	}
	if MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17Present {
		ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17)
		if err = ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	return nil
}
