// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-SessionInfo-r17 (line 28502).

var mBSSessionInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionId-r17"},
		{Name: "g-RNTI-r17"},
		{Name: "mrb-ListBroadcast-r17"},
		{Name: "mtch-SchedulingInfo-r17", Optional: true},
		{Name: "mtch-NeighbourCell-r17", Optional: true},
		{Name: "pdsch-ConfigIndex-r17", Optional: true},
		{Name: "mtch-SSB-MappingWindowIndex-r17", Optional: true},
	},
}

var mBSSessionInfoR17MtchNeighbourCellR17Constraints = per.FixedSize(common.MaxNeighCellMBS_r17)

type MBS_SessionInfo_r17 struct {
	Mbs_SessionId_r17               TMGI_r17
	G_RNTI_r17                      RNTI_Value
	Mrb_ListBroadcast_r17           MRB_ListBroadcast_r17
	Mtch_SchedulingInfo_r17         *DRX_ConfigPTM_Index_r17
	Mtch_NeighbourCell_r17          *per.BitString
	Pdsch_ConfigIndex_r17           *PDSCH_ConfigIndex_r17
	Mtch_SSB_MappingWindowIndex_r17 *MTCH_SSB_MappingWindowIndex_r17
}

func (ie *MBS_SessionInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSSessionInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mtch_SchedulingInfo_r17 != nil, ie.Mtch_NeighbourCell_r17 != nil, ie.Pdsch_ConfigIndex_r17 != nil, ie.Mtch_SSB_MappingWindowIndex_r17 != nil}); err != nil {
		return err
	}

	// 2. mbs-SessionId-r17: ref
	{
		if err := ie.Mbs_SessionId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. g-RNTI-r17: ref
	{
		if err := ie.G_RNTI_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. mrb-ListBroadcast-r17: ref
	{
		if err := ie.Mrb_ListBroadcast_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. mtch-SchedulingInfo-r17: ref
	{
		if ie.Mtch_SchedulingInfo_r17 != nil {
			if err := ie.Mtch_SchedulingInfo_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mtch-NeighbourCell-r17: bit-string
	{
		if ie.Mtch_NeighbourCell_r17 != nil {
			if err := e.EncodeBitString(*ie.Mtch_NeighbourCell_r17, mBSSessionInfoR17MtchNeighbourCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. pdsch-ConfigIndex-r17: ref
	{
		if ie.Pdsch_ConfigIndex_r17 != nil {
			if err := ie.Pdsch_ConfigIndex_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. mtch-SSB-MappingWindowIndex-r17: ref
	{
		if ie.Mtch_SSB_MappingWindowIndex_r17 != nil {
			if err := ie.Mtch_SSB_MappingWindowIndex_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBS_SessionInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSSessionInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-SessionId-r17: ref
	{
		if err := ie.Mbs_SessionId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. g-RNTI-r17: ref
	{
		if err := ie.G_RNTI_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. mrb-ListBroadcast-r17: ref
	{
		if err := ie.Mrb_ListBroadcast_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. mtch-SchedulingInfo-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Mtch_SchedulingInfo_r17 = new(DRX_ConfigPTM_Index_r17)
			if err := ie.Mtch_SchedulingInfo_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mtch-NeighbourCell-r17: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(mBSSessionInfoR17MtchNeighbourCellR17Constraints)
			if err != nil {
				return err
			}
			ie.Mtch_NeighbourCell_r17 = &v
		}
	}

	// 7. pdsch-ConfigIndex-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Pdsch_ConfigIndex_r17 = new(PDSCH_ConfigIndex_r17)
			if err := ie.Pdsch_ConfigIndex_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. mtch-SSB-MappingWindowIndex-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Mtch_SSB_MappingWindowIndex_r17 = new(MTCH_SSB_MappingWindowIndex_r17)
			if err := ie.Mtch_SSB_MappingWindowIndex_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
