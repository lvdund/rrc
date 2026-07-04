// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBSMulticastConfiguration-r18-IEs (line 672).

var mBSMulticastConfigurationR18IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionInfoListMulticast-r18", Optional: true},
		{Name: "mbs-NeighbourCellList-r18", Optional: true},
		{Name: "drx-ConfigPTM-List-r18", Optional: true},
		{Name: "pdsch-ConfigMTCH-r18", Optional: true},
		{Name: "mtch-SSB-MappingWindowList-r18", Optional: true},
		{Name: "thresholdMBS-List-r18", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var mBSMulticastConfigurationR18IEsDrxConfigPTMListR18Constraints = per.SizeRange(1, common.MaxNrofDRX_ConfigPTM_r17)

var mBSMulticastConfigurationR18IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type MBSMulticastConfiguration_r18_IEs struct {
	Mbs_SessionInfoListMulticast_r18 *MBS_SessionInfoListMulticast_r18
	Mbs_NeighbourCellList_r18        *MBS_NeighbourCellList_r17
	Drx_ConfigPTM_List_r18           []DRX_ConfigPTM_r17
	Pdsch_ConfigMTCH_r18             *PDSCH_ConfigBroadcast_r17
	Mtch_SSB_MappingWindowList_r18   *MTCH_SSB_MappingWindowList_r17
	ThresholdMBS_List_r18            *ThresholdMBS_List_r18
	LateNonCriticalExtension         []byte
	NonCriticalExtension             *struct{}
}

func (ie *MBSMulticastConfiguration_r18_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSMulticastConfigurationR18IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_SessionInfoListMulticast_r18 != nil, ie.Mbs_NeighbourCellList_r18 != nil, ie.Drx_ConfigPTM_List_r18 != nil, ie.Pdsch_ConfigMTCH_r18 != nil, ie.Mtch_SSB_MappingWindowList_r18 != nil, ie.ThresholdMBS_List_r18 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mbs-SessionInfoListMulticast-r18: ref
	{
		if ie.Mbs_SessionInfoListMulticast_r18 != nil {
			if err := ie.Mbs_SessionInfoListMulticast_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mbs-NeighbourCellList-r18: ref
	{
		if ie.Mbs_NeighbourCellList_r18 != nil {
			if err := ie.Mbs_NeighbourCellList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. drx-ConfigPTM-List-r18: sequence-of
	{
		if ie.Drx_ConfigPTM_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mBSMulticastConfigurationR18IEsDrxConfigPTMListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Drx_ConfigPTM_List_r18))); err != nil {
				return err
			}
			for i := range ie.Drx_ConfigPTM_List_r18 {
				if err := ie.Drx_ConfigPTM_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. pdsch-ConfigMTCH-r18: ref
	{
		if ie.Pdsch_ConfigMTCH_r18 != nil {
			if err := ie.Pdsch_ConfigMTCH_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mtch-SSB-MappingWindowList-r18: ref
	{
		if ie.Mtch_SSB_MappingWindowList_r18 != nil {
			if err := ie.Mtch_SSB_MappingWindowList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. thresholdMBS-List-r18: ref
	{
		if ie.ThresholdMBS_List_r18 != nil {
			if err := ie.ThresholdMBS_List_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, mBSMulticastConfigurationR18IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 9. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *MBSMulticastConfiguration_r18_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSMulticastConfigurationR18IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-SessionInfoListMulticast-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_SessionInfoListMulticast_r18 = new(MBS_SessionInfoListMulticast_r18)
			if err := ie.Mbs_SessionInfoListMulticast_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mbs-NeighbourCellList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mbs_NeighbourCellList_r18 = new(MBS_NeighbourCellList_r17)
			if err := ie.Mbs_NeighbourCellList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. drx-ConfigPTM-List-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(mBSMulticastConfigurationR18IEsDrxConfigPTMListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Drx_ConfigPTM_List_r18 = make([]DRX_ConfigPTM_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Drx_ConfigPTM_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. pdsch-ConfigMTCH-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Pdsch_ConfigMTCH_r18 = new(PDSCH_ConfigBroadcast_r17)
			if err := ie.Pdsch_ConfigMTCH_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mtch-SSB-MappingWindowList-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mtch_SSB_MappingWindowList_r18 = new(MTCH_SSB_MappingWindowList_r17)
			if err := ie.Mtch_SSB_MappingWindowList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. thresholdMBS-List-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.ThresholdMBS_List_r18 = new(ThresholdMBS_List_r18)
			if err := ie.ThresholdMBS_List_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(mBSMulticastConfigurationR18IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 9. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
