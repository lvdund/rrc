// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SyncAllowed-r16 (line 28021).

var sLSyncAllowedR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gnss-Sync-r16", Optional: true},
		{Name: "gnbEnb-Sync-r16", Optional: true},
		{Name: "ue-Sync-r16", Optional: true},
	},
}

const (
	SL_SyncAllowed_r16_Gnss_Sync_r16_True = 0
)

var sLSyncAllowedR16GnssSyncR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SyncAllowed_r16_Gnss_Sync_r16_True},
}

const (
	SL_SyncAllowed_r16_GnbEnb_Sync_r16_True = 0
)

var sLSyncAllowedR16GnbEnbSyncR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SyncAllowed_r16_GnbEnb_Sync_r16_True},
}

const (
	SL_SyncAllowed_r16_Ue_Sync_r16_True = 0
)

var sLSyncAllowedR16UeSyncR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SyncAllowed_r16_Ue_Sync_r16_True},
}

type SL_SyncAllowed_r16 struct {
	Gnss_Sync_r16   *int64
	GnbEnb_Sync_r16 *int64
	Ue_Sync_r16     *int64
}

func (ie *SL_SyncAllowed_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSyncAllowedR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gnss_Sync_r16 != nil, ie.GnbEnb_Sync_r16 != nil, ie.Ue_Sync_r16 != nil}); err != nil {
		return err
	}

	// 2. gnss-Sync-r16: enumerated
	{
		if ie.Gnss_Sync_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Gnss_Sync_r16, sLSyncAllowedR16GnssSyncR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. gnbEnb-Sync-r16: enumerated
	{
		if ie.GnbEnb_Sync_r16 != nil {
			if err := e.EncodeEnumerated(*ie.GnbEnb_Sync_r16, sLSyncAllowedR16GnbEnbSyncR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ue-Sync-r16: enumerated
	{
		if ie.Ue_Sync_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ue_Sync_r16, sLSyncAllowedR16UeSyncR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_SyncAllowed_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSyncAllowedR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gnss-Sync-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLSyncAllowedR16GnssSyncR16Constraints)
			if err != nil {
				return err
			}
			ie.Gnss_Sync_r16 = &idx
		}
	}

	// 3. gnbEnb-Sync-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLSyncAllowedR16GnbEnbSyncR16Constraints)
			if err != nil {
				return err
			}
			ie.GnbEnb_Sync_r16 = &idx
		}
	}

	// 4. ue-Sync-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLSyncAllowedR16UeSyncR16Constraints)
			if err != nil {
				return err
			}
			ie.Ue_Sync_r16 = &idx
		}
	}

	return nil
}
