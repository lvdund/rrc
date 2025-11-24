package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1630 struct {
	OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16                    *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16                    `optional`
	OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16        *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16        `optional`
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 `optional`
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16    *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16    `optional`
	Dummy                                                    *FeatureSetUplink_v1630_dummy                                                    `optional`
	PartialCancellationPUCCH_PUSCH_PRACH_TX_r16              *FeatureSetUplink_v1630_partialCancellationPUCCH_PUSCH_PRACH_TX_r16              `optional`
}

func (ie *FeatureSetUplink_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16 != nil, ie.Dummy != nil, ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16 != nil {
		if err = ie.OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16", err)
		}
	}
	if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16 != nil {
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16", err)
		}
	}
	if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 != nil {
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16", err)
		}
	}
	if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16 != nil {
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16", err)
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16 != nil {
		if err = ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PartialCancellationPUCCH_PUSCH_PRACH_TX_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1630) Decode(r *uper.UperReader) error {
	var err error
	var OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16Present bool
	if OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16Present bool
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16Present bool
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16Present bool
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PartialCancellationPUCCH_PUSCH_PRACH_TX_r16Present bool
	if PartialCancellationPUCCH_PUSCH_PRACH_TX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16Present {
		ie.OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16)
		if err = ie.OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetSRS_CB_PUSCH_Ant_Switch_fr1_r16", err)
		}
	}
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16Present {
		ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16)
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16", err)
		}
	}
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16Present {
		ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16)
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16", err)
		}
	}
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16Present {
		ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16)
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16", err)
		}
	}
	if DummyPresent {
		ie.Dummy = new(FeatureSetUplink_v1630_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if PartialCancellationPUCCH_PUSCH_PRACH_TX_r16Present {
		ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16 = new(FeatureSetUplink_v1630_partialCancellationPUCCH_PUSCH_PRACH_TX_r16)
		if err = ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PartialCancellationPUCCH_PUSCH_PRACH_TX_r16", err)
		}
	}
	return nil
}
