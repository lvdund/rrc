// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-UplinkTxSwitch-r16 (line 16828).

var bandCombinationUplinkTxSwitchR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-r16"},
		{Name: "bandCombination-v1540", Optional: true},
		{Name: "bandCombination-v1560", Optional: true},
		{Name: "bandCombination-v1570", Optional: true},
		{Name: "bandCombination-v1580", Optional: true},
		{Name: "bandCombination-v1590", Optional: true},
		{Name: "bandCombination-v1610", Optional: true},
		{Name: "supportedBandPairListNR-r16"},
		{Name: "uplinkTxSwitching-OptionSupport-r16", Optional: true},
		{Name: "uplinkTxSwitching-PowerBoosting-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var bandCombinationUplinkTxSwitchR16SupportedBandPairListNRR16Constraints = per.SizeRange(1, common.MaxULTxSwitchingBandPairs)

const (
	BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_OptionSupport_r16_SwitchedUL = 0
	BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_OptionSupport_r16_DualUL     = 1
	BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_OptionSupport_r16_Both       = 2
)

var bandCombinationUplinkTxSwitchR16UplinkTxSwitchingOptionSupportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_OptionSupport_r16_SwitchedUL, BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_OptionSupport_r16_DualUL, BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_OptionSupport_r16_Both},
}

const (
	BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_PowerBoosting_r16_Supported = 0
)

var bandCombinationUplinkTxSwitchR16UplinkTxSwitchingPowerBoostingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_r16_UplinkTxSwitching_PowerBoosting_r16_Supported},
}

const (
	BandCombination_UplinkTxSwitch_r16_Ext_UplinkTxSwitching_PUSCH_TransCoherence_r16_NonCoherent  = 0
	BandCombination_UplinkTxSwitch_r16_Ext_UplinkTxSwitching_PUSCH_TransCoherence_r16_FullCoherent = 1
)

var bandCombinationUplinkTxSwitchR16ExtUplinkTxSwitchingPUSCHTransCoherenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_r16_Ext_UplinkTxSwitching_PUSCH_TransCoherence_r16_NonCoherent, BandCombination_UplinkTxSwitch_r16_Ext_UplinkTxSwitching_PUSCH_TransCoherence_r16_FullCoherent},
}

type BandCombination_UplinkTxSwitch_r16 struct {
	BandCombination_r16                        BandCombination
	BandCombination_v1540                      *BandCombination_v1540
	BandCombination_v1560                      *BandCombination_v1560
	BandCombination_v1570                      *BandCombination_v1570
	BandCombination_v1580                      *BandCombination_v1580
	BandCombination_v1590                      *BandCombination_v1590
	BandCombination_v1610                      *BandCombination_v1610
	SupportedBandPairListNR_r16                []ULTxSwitchingBandPair_r16
	UplinkTxSwitching_OptionSupport_r16        *int64
	UplinkTxSwitching_PowerBoosting_r16        *int64
	UplinkTxSwitching_PUSCH_TransCoherence_r16 *int64
}

func (ie *BandCombination_UplinkTxSwitch_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1540 != nil, ie.BandCombination_v1560 != nil, ie.BandCombination_v1570 != nil, ie.BandCombination_v1580 != nil, ie.BandCombination_v1590 != nil, ie.BandCombination_v1610 != nil, ie.UplinkTxSwitching_OptionSupport_r16 != nil, ie.UplinkTxSwitching_PowerBoosting_r16 != nil}); err != nil {
		return err
	}

	// 3. bandCombination-r16: ref
	{
		if err := ie.BandCombination_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. bandCombination-v1540: ref
	{
		if ie.BandCombination_v1540 != nil {
			if err := ie.BandCombination_v1540.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bandCombination-v1560: ref
	{
		if ie.BandCombination_v1560 != nil {
			if err := ie.BandCombination_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. bandCombination-v1570: ref
	{
		if ie.BandCombination_v1570 != nil {
			if err := ie.BandCombination_v1570.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. bandCombination-v1580: ref
	{
		if ie.BandCombination_v1580 != nil {
			if err := ie.BandCombination_v1580.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. bandCombination-v1590: ref
	{
		if ie.BandCombination_v1590 != nil {
			if err := ie.BandCombination_v1590.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. bandCombination-v1610: ref
	{
		if ie.BandCombination_v1610 != nil {
			if err := ie.BandCombination_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. supportedBandPairListNR-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchR16SupportedBandPairListNRR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.SupportedBandPairListNR_r16))); err != nil {
			return err
		}
		for i := range ie.SupportedBandPairListNR_r16 {
			if err := ie.SupportedBandPairListNR_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. uplinkTxSwitching-OptionSupport-r16: enumerated
	{
		if ie.UplinkTxSwitching_OptionSupport_r16 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitching_OptionSupport_r16, bandCombinationUplinkTxSwitchR16UplinkTxSwitchingOptionSupportR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. uplinkTxSwitching-PowerBoosting-r16: enumerated
	{
		if ie.UplinkTxSwitching_PowerBoosting_r16 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitching_PowerBoosting_r16, bandCombinationUplinkTxSwitchR16UplinkTxSwitchingPowerBoostingR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "uplinkTxSwitching-PUSCH-TransCoherence-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil}); err != nil {
				return err
			}

			if ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitching_PUSCH_TransCoherence_r16, bandCombinationUplinkTxSwitchR16ExtUplinkTxSwitchingPUSCHTransCoherenceR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bandCombination-r16: ref
	{
		if err := ie.BandCombination_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. bandCombination-v1540: ref
	{
		if seq.IsComponentPresent(1) {
			ie.BandCombination_v1540 = new(BandCombination_v1540)
			if err := ie.BandCombination_v1540.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bandCombination-v1560: ref
	{
		if seq.IsComponentPresent(2) {
			ie.BandCombination_v1560 = new(BandCombination_v1560)
			if err := ie.BandCombination_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. bandCombination-v1570: ref
	{
		if seq.IsComponentPresent(3) {
			ie.BandCombination_v1570 = new(BandCombination_v1570)
			if err := ie.BandCombination_v1570.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. bandCombination-v1580: ref
	{
		if seq.IsComponentPresent(4) {
			ie.BandCombination_v1580 = new(BandCombination_v1580)
			if err := ie.BandCombination_v1580.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. bandCombination-v1590: ref
	{
		if seq.IsComponentPresent(5) {
			ie.BandCombination_v1590 = new(BandCombination_v1590)
			if err := ie.BandCombination_v1590.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. bandCombination-v1610: ref
	{
		if seq.IsComponentPresent(6) {
			ie.BandCombination_v1610 = new(BandCombination_v1610)
			if err := ie.BandCombination_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. supportedBandPairListNR-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchR16SupportedBandPairListNRR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SupportedBandPairListNR_r16 = make([]ULTxSwitchingBandPair_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SupportedBandPairListNR_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. uplinkTxSwitching-OptionSupport-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(bandCombinationUplinkTxSwitchR16UplinkTxSwitchingOptionSupportR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_OptionSupport_r16 = &idx
		}
	}

	// 12. uplinkTxSwitching-PowerBoosting-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(bandCombinationUplinkTxSwitchR16UplinkTxSwitchingPowerBoostingR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_PowerBoosting_r16 = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "uplinkTxSwitching-PUSCH-TransCoherence-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandCombinationUplinkTxSwitchR16ExtUplinkTxSwitchingPUSCHTransCoherenceR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 = &v
		}
	}

	return nil
}
