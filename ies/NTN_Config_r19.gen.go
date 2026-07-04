// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-Config-r19 (line 10794).

var nTNConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ephemerisInfo-r19", Optional: true},
		{Name: "epochTime-r19", Optional: true},
		{Name: "ntn-UlSyncValidityDuration-r19", Optional: true},
		{Name: "kmac-r19", Optional: true},
		{Name: "ta-InfoIoT-r19", Optional: true},
	},
}

const (
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S5   = 0
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S10  = 1
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S15  = 2
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S20  = 3
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S25  = 4
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S30  = 5
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S35  = 6
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S40  = 7
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S45  = 8
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S50  = 9
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S55  = 10
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S60  = 11
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S120 = 12
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S180 = 13
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S240 = 14
	NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S900 = 15
)

var nTNConfigR19NtnUlSyncValidityDurationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S5, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S10, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S15, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S20, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S25, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S30, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S35, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S40, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S45, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S50, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S55, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S60, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S120, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S180, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S240, NTN_Config_r19_Ntn_UlSyncValidityDuration_r19_S900},
}

var nTNConfigR19KmacR19Constraints = per.Constrained(1, 512)

type NTN_Config_r19 struct {
	EphemerisInfo_r19              *EphemerisInfo_r17
	EpochTime_r19                  *EpochTime_r17
	Ntn_UlSyncValidityDuration_r19 *int64
	Kmac_r19                       *int64
	Ta_InfoIoT_r19                 *TA_InfoIoT_r19
}

func (ie *NTN_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EphemerisInfo_r19 != nil, ie.EpochTime_r19 != nil, ie.Ntn_UlSyncValidityDuration_r19 != nil, ie.Kmac_r19 != nil, ie.Ta_InfoIoT_r19 != nil}); err != nil {
		return err
	}

	// 3. ephemerisInfo-r19: ref
	{
		if ie.EphemerisInfo_r19 != nil {
			if err := ie.EphemerisInfo_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. epochTime-r19: ref
	{
		if ie.EpochTime_r19 != nil {
			if err := ie.EpochTime_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ntn-UlSyncValidityDuration-r19: enumerated
	{
		if ie.Ntn_UlSyncValidityDuration_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_UlSyncValidityDuration_r19, nTNConfigR19NtnUlSyncValidityDurationR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. kmac-r19: integer
	{
		if ie.Kmac_r19 != nil {
			if err := e.EncodeInteger(*ie.Kmac_r19, nTNConfigR19KmacR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ta-InfoIoT-r19: ref
	{
		if ie.Ta_InfoIoT_r19 != nil {
			if err := ie.Ta_InfoIoT_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ephemerisInfo-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.EphemerisInfo_r19 = new(EphemerisInfo_r17)
			if err := ie.EphemerisInfo_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. epochTime-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.EpochTime_r19 = new(EpochTime_r17)
			if err := ie.EpochTime_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ntn-UlSyncValidityDuration-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(nTNConfigR19NtnUlSyncValidityDurationR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_UlSyncValidityDuration_r19 = &idx
		}
	}

	// 6. kmac-r19: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(nTNConfigR19KmacR19Constraints)
			if err != nil {
				return err
			}
			ie.Kmac_r19 = &v
		}
	}

	// 7. ta-InfoIoT-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ta_InfoIoT_r19 = new(TA_InfoIoT_r19)
			if err := ie.Ta_InfoIoT_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
