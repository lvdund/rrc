// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SyncConfig-r16 (line 28310).

var sLSyncConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SyncRefMinHyst-r16", Optional: true},
		{Name: "sl-SyncRefDiffHyst-r16", Optional: true},
		{Name: "sl-FilterCoefficient-r16", Optional: true},
		{Name: "sl-SSB-TimeAllocation1-r16", Optional: true},
		{Name: "sl-SSB-TimeAllocation2-r16", Optional: true},
		{Name: "sl-SSB-TimeAllocation3-r16", Optional: true},
		{Name: "sl-SSID-r16", Optional: true},
		{Name: "txParameters-r16"},
		{Name: "gnss-Sync-r16", Optional: true},
	},
}

const (
	SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB0  = 0
	SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB3  = 1
	SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB6  = 2
	SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB9  = 3
	SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB12 = 4
)

var sLSyncConfigR16SlSyncRefMinHystR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB0, SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB3, SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB6, SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB9, SL_SyncConfig_r16_Sl_SyncRefMinHyst_r16_DB12},
}

const (
	SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB0   = 0
	SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB3   = 1
	SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB6   = 2
	SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB9   = 3
	SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB12  = 4
	SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DBinf = 5
)

var sLSyncConfigR16SlSyncRefDiffHystR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB0, SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB3, SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB6, SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB9, SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DB12, SL_SyncConfig_r16_Sl_SyncRefDiffHyst_r16_DBinf},
}

var sLSyncConfigR16SlSSIDR16Constraints = per.Constrained(0, 671)

const (
	SL_SyncConfig_r16_Gnss_Sync_r16_True = 0
)

var sLSyncConfigR16GnssSyncR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SyncConfig_r16_Gnss_Sync_r16_True},
}

var sLSyncConfigR16TxParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "syncTxThreshIC-r16", Optional: true},
		{Name: "syncTxThreshOoC-r16", Optional: true},
		{Name: "syncInfoReserved-r16", Optional: true},
	},
}

type SL_SyncConfig_r16 struct {
	Sl_SyncRefMinHyst_r16      *int64
	Sl_SyncRefDiffHyst_r16     *int64
	Sl_FilterCoefficient_r16   *FilterCoefficient
	Sl_SSB_TimeAllocation1_r16 *SL_SSB_TimeAllocation_r16
	Sl_SSB_TimeAllocation2_r16 *SL_SSB_TimeAllocation_r16
	Sl_SSB_TimeAllocation3_r16 *SL_SSB_TimeAllocation_r16
	Sl_SSID_r16                *int64
	TxParameters_r16           struct {
		SyncTxThreshIC_r16   *SL_RSRP_Range_r16
		SyncTxThreshOoC_r16  *SL_RSRP_Range_r16
		SyncInfoReserved_r16 *per.BitString
	}
	Gnss_Sync_r16 *int64
}

func (ie *SL_SyncConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSyncConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_SyncRefMinHyst_r16 != nil, ie.Sl_SyncRefDiffHyst_r16 != nil, ie.Sl_FilterCoefficient_r16 != nil, ie.Sl_SSB_TimeAllocation1_r16 != nil, ie.Sl_SSB_TimeAllocation2_r16 != nil, ie.Sl_SSB_TimeAllocation3_r16 != nil, ie.Sl_SSID_r16 != nil, ie.Gnss_Sync_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-SyncRefMinHyst-r16: enumerated
	{
		if ie.Sl_SyncRefMinHyst_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SyncRefMinHyst_r16, sLSyncConfigR16SlSyncRefMinHystR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-SyncRefDiffHyst-r16: enumerated
	{
		if ie.Sl_SyncRefDiffHyst_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SyncRefDiffHyst_r16, sLSyncConfigR16SlSyncRefDiffHystR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-FilterCoefficient-r16: ref
	{
		if ie.Sl_FilterCoefficient_r16 != nil {
			if err := ie.Sl_FilterCoefficient_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-SSB-TimeAllocation1-r16: ref
	{
		if ie.Sl_SSB_TimeAllocation1_r16 != nil {
			if err := ie.Sl_SSB_TimeAllocation1_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. sl-SSB-TimeAllocation2-r16: ref
	{
		if ie.Sl_SSB_TimeAllocation2_r16 != nil {
			if err := ie.Sl_SSB_TimeAllocation2_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. sl-SSB-TimeAllocation3-r16: ref
	{
		if ie.Sl_SSB_TimeAllocation3_r16 != nil {
			if err := ie.Sl_SSB_TimeAllocation3_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. sl-SSID-r16: integer
	{
		if ie.Sl_SSID_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_SSID_r16, sLSyncConfigR16SlSSIDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. txParameters-r16: sequence
	{
		{
			c := &ie.TxParameters_r16
			sLSyncConfigR16TxParametersR16Seq := e.NewSequenceEncoder(sLSyncConfigR16TxParametersR16Constraints)
			if err := sLSyncConfigR16TxParametersR16Seq.EncodePreamble([]bool{c.SyncTxThreshIC_r16 != nil, c.SyncTxThreshOoC_r16 != nil, c.SyncInfoReserved_r16 != nil}); err != nil {
				return err
			}
			if c.SyncTxThreshIC_r16 != nil {
				if err := c.SyncTxThreshIC_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.SyncTxThreshOoC_r16 != nil {
				if err := c.SyncTxThreshOoC_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.SyncInfoReserved_r16 != nil {
				if err := e.EncodeBitString((*c.SyncInfoReserved_r16), per.FixedSize(2)); err != nil {
					return err
				}
			}
		}
	}

	// 11. gnss-Sync-r16: enumerated
	{
		if ie.Gnss_Sync_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Gnss_Sync_r16, sLSyncConfigR16GnssSyncR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_SyncConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSyncConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-SyncRefMinHyst-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLSyncConfigR16SlSyncRefMinHystR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncRefMinHyst_r16 = &idx
		}
	}

	// 4. sl-SyncRefDiffHyst-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLSyncConfigR16SlSyncRefDiffHystR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncRefDiffHyst_r16 = &idx
		}
	}

	// 5. sl-FilterCoefficient-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_FilterCoefficient_r16 = new(FilterCoefficient)
			if err := ie.Sl_FilterCoefficient_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-SSB-TimeAllocation1-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_SSB_TimeAllocation1_r16 = new(SL_SSB_TimeAllocation_r16)
			if err := ie.Sl_SSB_TimeAllocation1_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. sl-SSB-TimeAllocation2-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sl_SSB_TimeAllocation2_r16 = new(SL_SSB_TimeAllocation_r16)
			if err := ie.Sl_SSB_TimeAllocation2_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. sl-SSB-TimeAllocation3-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Sl_SSB_TimeAllocation3_r16 = new(SL_SSB_TimeAllocation_r16)
			if err := ie.Sl_SSB_TimeAllocation3_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. sl-SSID-r16: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sLSyncConfigR16SlSSIDR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SSID_r16 = &v
		}
	}

	// 10. txParameters-r16: sequence
	{
		{
			c := &ie.TxParameters_r16
			sLSyncConfigR16TxParametersR16Seq := d.NewSequenceDecoder(sLSyncConfigR16TxParametersR16Constraints)
			if err := sLSyncConfigR16TxParametersR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if sLSyncConfigR16TxParametersR16Seq.IsComponentPresent(0) {
				c.SyncTxThreshIC_r16 = new(SL_RSRP_Range_r16)
				if err := (*c.SyncTxThreshIC_r16).Decode(d); err != nil {
					return err
				}
			}
			if sLSyncConfigR16TxParametersR16Seq.IsComponentPresent(1) {
				c.SyncTxThreshOoC_r16 = new(SL_RSRP_Range_r16)
				if err := (*c.SyncTxThreshOoC_r16).Decode(d); err != nil {
					return err
				}
			}
			if sLSyncConfigR16TxParametersR16Seq.IsComponentPresent(2) {
				c.SyncInfoReserved_r16 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				(*c.SyncInfoReserved_r16) = v
			}
		}
	}

	// 11. gnss-Sync-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(sLSyncConfigR16GnssSyncR16Constraints)
			if err != nil {
				return err
			}
			ie.Gnss_Sync_r16 = &idx
		}
	}

	return nil
}
