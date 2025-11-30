package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1730 struct {
	IntraSlotTDM_UnicastGroupCommonPDSCH_r17 *FeatureSetDownlinkPerCC_v1730_intraSlotTDM_UnicastGroupCommonPDSCH_r17 `optional`
	Sps_MulticastSCell_r17                   *FeatureSetDownlinkPerCC_v1730_sps_MulticastSCell_r17                   `optional`
	Sps_MulticastSCellMultiConfig_r17        *int64                                                                  `lb:1,ub:8,optional`
	Dci_BroadcastWith16Repetitions_r17       *FeatureSetDownlinkPerCC_v1730_dci_BroadcastWith16Repetitions_r17       `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17 != nil, ie.Sps_MulticastSCell_r17 != nil, ie.Sps_MulticastSCellMultiConfig_r17 != nil, ie.Dci_BroadcastWith16Repetitions_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17 != nil {
		if err = ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode IntraSlotTDM_UnicastGroupCommonPDSCH_r17", err)
		}
	}
	if ie.Sps_MulticastSCell_r17 != nil {
		if err = ie.Sps_MulticastSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sps_MulticastSCell_r17", err)
		}
	}
	if ie.Sps_MulticastSCellMultiConfig_r17 != nil {
		if err = w.WriteInteger(*ie.Sps_MulticastSCellMultiConfig_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sps_MulticastSCellMultiConfig_r17", err)
		}
	}
	if ie.Dci_BroadcastWith16Repetitions_r17 != nil {
		if err = ie.Dci_BroadcastWith16Repetitions_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_BroadcastWith16Repetitions_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1730) Decode(r *aper.AperReader) error {
	var err error
	var IntraSlotTDM_UnicastGroupCommonPDSCH_r17Present bool
	if IntraSlotTDM_UnicastGroupCommonPDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sps_MulticastSCell_r17Present bool
	if Sps_MulticastSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sps_MulticastSCellMultiConfig_r17Present bool
	if Sps_MulticastSCellMultiConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_BroadcastWith16Repetitions_r17Present bool
	if Dci_BroadcastWith16Repetitions_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if IntraSlotTDM_UnicastGroupCommonPDSCH_r17Present {
		ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17 = new(FeatureSetDownlinkPerCC_v1730_intraSlotTDM_UnicastGroupCommonPDSCH_r17)
		if err = ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode IntraSlotTDM_UnicastGroupCommonPDSCH_r17", err)
		}
	}
	if Sps_MulticastSCell_r17Present {
		ie.Sps_MulticastSCell_r17 = new(FeatureSetDownlinkPerCC_v1730_sps_MulticastSCell_r17)
		if err = ie.Sps_MulticastSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sps_MulticastSCell_r17", err)
		}
	}
	if Sps_MulticastSCellMultiConfig_r17Present {
		var tmp_int_Sps_MulticastSCellMultiConfig_r17 int64
		if tmp_int_Sps_MulticastSCellMultiConfig_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sps_MulticastSCellMultiConfig_r17", err)
		}
		ie.Sps_MulticastSCellMultiConfig_r17 = &tmp_int_Sps_MulticastSCellMultiConfig_r17
	}
	if Dci_BroadcastWith16Repetitions_r17Present {
		ie.Dci_BroadcastWith16Repetitions_r17 = new(FeatureSetDownlinkPerCC_v1730_dci_BroadcastWith16Repetitions_r17)
		if err = ie.Dci_BroadcastWith16Repetitions_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_BroadcastWith16Repetitions_r17", err)
		}
	}
	return nil
}
