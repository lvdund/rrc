// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-SessionInfoMulticast-r18 (line 28560).

var mBSSessionInfoMulticastR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionId-r18"},
		{Name: "g-RNTI-r18", Optional: true},
		{Name: "mrb-ListMulticast-r18", Optional: true},
		{Name: "mtch-SchedulingInfo-r18", Optional: true},
		{Name: "mtch-NeighbourCell-r18", Optional: true},
		{Name: "pdsch-ConfigIndex-r18", Optional: true},
		{Name: "mtch-SSB-MappingWindowIndex-r18", Optional: true},
		{Name: "thresholdIndex-r18", Optional: true},
		{Name: "pdcp-SyncIndicator-r18", Optional: true},
		{Name: "stopMonitoringRNTI-r18", Optional: true},
	},
}

var mBSSessionInfoMulticastR18ThresholdIndexR18Constraints = per.Constrained(0, common.MaxNrofThresholdMBS_1_r18)

const (
	MBS_SessionInfoMulticast_r18_Pdcp_SyncIndicator_r18_True = 0
)

var mBSSessionInfoMulticastR18PdcpSyncIndicatorR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MBS_SessionInfoMulticast_r18_Pdcp_SyncIndicator_r18_True},
}

const (
	MBS_SessionInfoMulticast_r18_StopMonitoringRNTI_r18_True = 0
)

var mBSSessionInfoMulticastR18StopMonitoringRNTIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MBS_SessionInfoMulticast_r18_StopMonitoringRNTI_r18_True},
}

type MBS_SessionInfoMulticast_r18 struct {
	Mbs_SessionId_r18               TMGI_r17
	G_RNTI_r18                      *RNTI_Value
	Mrb_ListMulticast_r18           *MRB_ListMulticast_r18
	Mtch_SchedulingInfo_r18         *DRX_ConfigPTM_Index_r17
	Mtch_NeighbourCell_r18          *MTCH_NeighbourCell_r18
	Pdsch_ConfigIndex_r18           *PDSCH_ConfigIndex_r17
	Mtch_SSB_MappingWindowIndex_r18 *MTCH_SSB_MappingWindowIndex_r17
	ThresholdIndex_r18              *int64
	Pdcp_SyncIndicator_r18          *int64
	StopMonitoringRNTI_r18          *int64
}

func (ie *MBS_SessionInfoMulticast_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSSessionInfoMulticastR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.G_RNTI_r18 != nil, ie.Mrb_ListMulticast_r18 != nil, ie.Mtch_SchedulingInfo_r18 != nil, ie.Mtch_NeighbourCell_r18 != nil, ie.Pdsch_ConfigIndex_r18 != nil, ie.Mtch_SSB_MappingWindowIndex_r18 != nil, ie.ThresholdIndex_r18 != nil, ie.Pdcp_SyncIndicator_r18 != nil, ie.StopMonitoringRNTI_r18 != nil}); err != nil {
		return err
	}

	// 3. mbs-SessionId-r18: ref
	{
		if err := ie.Mbs_SessionId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. g-RNTI-r18: ref
	{
		if ie.G_RNTI_r18 != nil {
			if err := ie.G_RNTI_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. mrb-ListMulticast-r18: ref
	{
		if ie.Mrb_ListMulticast_r18 != nil {
			if err := ie.Mrb_ListMulticast_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mtch-SchedulingInfo-r18: ref
	{
		if ie.Mtch_SchedulingInfo_r18 != nil {
			if err := ie.Mtch_SchedulingInfo_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. mtch-NeighbourCell-r18: ref
	{
		if ie.Mtch_NeighbourCell_r18 != nil {
			if err := ie.Mtch_NeighbourCell_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. pdsch-ConfigIndex-r18: ref
	{
		if ie.Pdsch_ConfigIndex_r18 != nil {
			if err := ie.Pdsch_ConfigIndex_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. mtch-SSB-MappingWindowIndex-r18: ref
	{
		if ie.Mtch_SSB_MappingWindowIndex_r18 != nil {
			if err := ie.Mtch_SSB_MappingWindowIndex_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. thresholdIndex-r18: integer
	{
		if ie.ThresholdIndex_r18 != nil {
			if err := e.EncodeInteger(*ie.ThresholdIndex_r18, mBSSessionInfoMulticastR18ThresholdIndexR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. pdcp-SyncIndicator-r18: enumerated
	{
		if ie.Pdcp_SyncIndicator_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_SyncIndicator_r18, mBSSessionInfoMulticastR18PdcpSyncIndicatorR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. stopMonitoringRNTI-r18: enumerated
	{
		if ie.StopMonitoringRNTI_r18 != nil {
			if err := e.EncodeEnumerated(*ie.StopMonitoringRNTI_r18, mBSSessionInfoMulticastR18StopMonitoringRNTIR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBS_SessionInfoMulticast_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSSessionInfoMulticastR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mbs-SessionId-r18: ref
	{
		if err := ie.Mbs_SessionId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. g-RNTI-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.G_RNTI_r18 = new(RNTI_Value)
			if err := ie.G_RNTI_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. mrb-ListMulticast-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mrb_ListMulticast_r18 = new(MRB_ListMulticast_r18)
			if err := ie.Mrb_ListMulticast_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mtch-SchedulingInfo-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Mtch_SchedulingInfo_r18 = new(DRX_ConfigPTM_Index_r17)
			if err := ie.Mtch_SchedulingInfo_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. mtch-NeighbourCell-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mtch_NeighbourCell_r18 = new(MTCH_NeighbourCell_r18)
			if err := ie.Mtch_NeighbourCell_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. pdsch-ConfigIndex-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Pdsch_ConfigIndex_r18 = new(PDSCH_ConfigIndex_r17)
			if err := ie.Pdsch_ConfigIndex_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. mtch-SSB-MappingWindowIndex-r18: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Mtch_SSB_MappingWindowIndex_r18 = new(MTCH_SSB_MappingWindowIndex_r17)
			if err := ie.Mtch_SSB_MappingWindowIndex_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. thresholdIndex-r18: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(mBSSessionInfoMulticastR18ThresholdIndexR18Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdIndex_r18 = &v
		}
	}

	// 11. pdcp-SyncIndicator-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(mBSSessionInfoMulticastR18PdcpSyncIndicatorR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_SyncIndicator_r18 = &idx
		}
	}

	// 12. stopMonitoringRNTI-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(mBSSessionInfoMulticastR18StopMonitoringRNTIR18Constraints)
			if err != nil {
				return err
			}
			ie.StopMonitoringRNTI_r18 = &idx
		}
	}

	return nil
}
