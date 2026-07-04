// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBSBroadcastConfiguration-r17-IEs (line 621).

var mBSBroadcastConfigurationR17IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionInfoList-r17", Optional: true},
		{Name: "mbs-NeighbourCellList-r17", Optional: true},
		{Name: "drx-ConfigPTM-List-r17", Optional: true},
		{Name: "pdsch-ConfigMTCH-r17", Optional: true},
		{Name: "mtch-SSB-MappingWindowList-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var mBSBroadcastConfigurationR17IEsDrxConfigPTMListR17Constraints = per.SizeRange(1, common.MaxNrofDRX_ConfigPTM_r17)

var mBSBroadcastConfigurationR17IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type MBSBroadcastConfiguration_r17_IEs struct {
	Mbs_SessionInfoList_r17        *MBS_SessionInfoList_r17
	Mbs_NeighbourCellList_r17      *MBS_NeighbourCellList_r17
	Drx_ConfigPTM_List_r17         []DRX_ConfigPTM_r17
	Pdsch_ConfigMTCH_r17           *PDSCH_ConfigBroadcast_r17
	Mtch_SSB_MappingWindowList_r17 *MTCH_SSB_MappingWindowList_r17
	LateNonCriticalExtension       []byte
	NonCriticalExtension           *MBSBroadcastConfiguration_v1900_IEs
}

func (ie *MBSBroadcastConfiguration_r17_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSBroadcastConfigurationR17IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_SessionInfoList_r17 != nil, ie.Mbs_NeighbourCellList_r17 != nil, ie.Drx_ConfigPTM_List_r17 != nil, ie.Pdsch_ConfigMTCH_r17 != nil, ie.Mtch_SSB_MappingWindowList_r17 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mbs-SessionInfoList-r17: ref
	{
		if ie.Mbs_SessionInfoList_r17 != nil {
			if err := ie.Mbs_SessionInfoList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mbs-NeighbourCellList-r17: ref
	{
		if ie.Mbs_NeighbourCellList_r17 != nil {
			if err := ie.Mbs_NeighbourCellList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. drx-ConfigPTM-List-r17: sequence-of
	{
		if ie.Drx_ConfigPTM_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mBSBroadcastConfigurationR17IEsDrxConfigPTMListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Drx_ConfigPTM_List_r17))); err != nil {
				return err
			}
			for i := range ie.Drx_ConfigPTM_List_r17 {
				if err := ie.Drx_ConfigPTM_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. pdsch-ConfigMTCH-r17: ref
	{
		if ie.Pdsch_ConfigMTCH_r17 != nil {
			if err := ie.Pdsch_ConfigMTCH_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mtch-SSB-MappingWindowList-r17: ref
	{
		if ie.Mtch_SSB_MappingWindowList_r17 != nil {
			if err := ie.Mtch_SSB_MappingWindowList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, mBSBroadcastConfigurationR17IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBSBroadcastConfiguration_r17_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSBroadcastConfigurationR17IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-SessionInfoList-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_SessionInfoList_r17 = new(MBS_SessionInfoList_r17)
			if err := ie.Mbs_SessionInfoList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mbs-NeighbourCellList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mbs_NeighbourCellList_r17 = new(MBS_NeighbourCellList_r17)
			if err := ie.Mbs_NeighbourCellList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. drx-ConfigPTM-List-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(mBSBroadcastConfigurationR17IEsDrxConfigPTMListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Drx_ConfigPTM_List_r17 = make([]DRX_ConfigPTM_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Drx_ConfigPTM_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. pdsch-ConfigMTCH-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Pdsch_ConfigMTCH_r17 = new(PDSCH_ConfigBroadcast_r17)
			if err := ie.Pdsch_ConfigMTCH_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mtch-SSB-MappingWindowList-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mtch_SSB_MappingWindowList_r17 = new(MTCH_SSB_MappingWindowList_r17)
			if err := ie.Mtch_SSB_MappingWindowList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(mBSBroadcastConfigurationR17IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = new(MBSBroadcastConfiguration_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
