// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-TimersAndConstantsRemoteUE-r17 (line 16234).

var uETimersAndConstantsRemoteUER17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "t300-RemoteUE-r17", Optional: true},
		{Name: "t301-RemoteUE-r17", Optional: true},
		{Name: "t319-RemoteUE-r17", Optional: true},
	},
}

const (
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms100  = 0
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms200  = 1
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms300  = 2
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms400  = 3
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms600  = 4
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms1000 = 5
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms1500 = 6
	UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms2000 = 7
)

var uETimersAndConstantsRemoteUER17T300RemoteUER17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms100, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms200, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms300, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms400, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms600, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms1000, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms1500, UE_TimersAndConstantsRemoteUE_r17_T300_RemoteUE_r17_Ms2000},
}

const (
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms100  = 0
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms200  = 1
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms300  = 2
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms400  = 3
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms600  = 4
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms1000 = 5
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms1500 = 6
	UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms2000 = 7
)

var uETimersAndConstantsRemoteUER17T301RemoteUER17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms100, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms200, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms300, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms400, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms600, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms1000, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms1500, UE_TimersAndConstantsRemoteUE_r17_T301_RemoteUE_r17_Ms2000},
}

const (
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms100  = 0
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms200  = 1
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms300  = 2
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms400  = 3
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms600  = 4
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms1000 = 5
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms1500 = 6
	UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms2000 = 7
)

var uETimersAndConstantsRemoteUER17T319RemoteUER17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms100, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms200, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms300, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms400, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms600, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms1000, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms1500, UE_TimersAndConstantsRemoteUE_r17_T319_RemoteUE_r17_Ms2000},
}

type UE_TimersAndConstantsRemoteUE_r17 struct {
	T300_RemoteUE_r17 *int64
	T301_RemoteUE_r17 *int64
	T319_RemoteUE_r17 *int64
}

func (ie *UE_TimersAndConstantsRemoteUE_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uETimersAndConstantsRemoteUER17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T300_RemoteUE_r17 != nil, ie.T301_RemoteUE_r17 != nil, ie.T319_RemoteUE_r17 != nil}); err != nil {
		return err
	}

	// 3. t300-RemoteUE-r17: enumerated
	{
		if ie.T300_RemoteUE_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T300_RemoteUE_r17, uETimersAndConstantsRemoteUER17T300RemoteUER17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. t301-RemoteUE-r17: enumerated
	{
		if ie.T301_RemoteUE_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T301_RemoteUE_r17, uETimersAndConstantsRemoteUER17T301RemoteUER17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. t319-RemoteUE-r17: enumerated
	{
		if ie.T319_RemoteUE_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T319_RemoteUE_r17, uETimersAndConstantsRemoteUER17T319RemoteUER17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_TimersAndConstantsRemoteUE_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uETimersAndConstantsRemoteUER17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. t300-RemoteUE-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uETimersAndConstantsRemoteUER17T300RemoteUER17Constraints)
			if err != nil {
				return err
			}
			ie.T300_RemoteUE_r17 = &idx
		}
	}

	// 4. t301-RemoteUE-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uETimersAndConstantsRemoteUER17T301RemoteUER17Constraints)
			if err != nil {
				return err
			}
			ie.T301_RemoteUE_r17 = &idx
		}
	}

	// 5. t319-RemoteUE-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uETimersAndConstantsRemoteUER17T319RemoteUER17Constraints)
			if err != nil {
				return err
			}
			ie.T319_RemoteUE_r17 = &idx
		}
	}

	return nil
}
