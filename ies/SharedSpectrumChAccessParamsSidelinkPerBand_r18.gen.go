// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SharedSpectrumChAccessParamsSidelinkPerBand-r18 (line 24979).

var sharedSpectrumChAccessParamsSidelinkPerBandR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DynamicChannelAccess-r18", Optional: true},
		{Name: "sl-DynamicMultiChannelAccess-r18", Optional: true},
		{Name: "sl-LBT-Option1-r18", Optional: true},
		{Name: "sl-LBT-Option2-r18", Optional: true},
		{Name: "sl-ResourceAllocMode1-r18", Optional: true},
		{Name: "sl-Interlace-RB-TxRx-r18", Optional: true},
		{Name: "sl-PSFCH-MultiOccasion-r18", Optional: true},
		{Name: "sl-ContiguousRB-TxRx-r18", Optional: true},
		{Name: "sl-PSFCH-MultiContiguousRB-r18", Optional: true},
		{Name: "sl-PSFCH-MultiNonContiguousRB-r18", Optional: true},
		{Name: "sl-MultiplePRB-CommonInterlacePSFCH-r18", Optional: true},
		{Name: "sl-MultiplePRB-DedicatedInterlacePSFCH-r18", Optional: true},
	},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_DynamicChannelAccess_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlDynamicChannelAccessR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_DynamicChannelAccess_r18_Supported},
}

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlDynamicMultiChannelAccessR18Constraints = per.Constrained(2, 5)

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_LBT_Option1_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlLBTOption1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_LBT_Option1_r18_Supported},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_LBT_Option2_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlLBTOption2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_LBT_Option2_r18_Supported},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_ResourceAllocMode1_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlResourceAllocMode1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_ResourceAllocMode1_r18_Supported},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_Interlace_RB_TxRx_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlInterlaceRBTxRxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_Interlace_RB_TxRx_r18_Supported},
}

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiOccasionR18Constraints = per.Constrained(1, 4)

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_ContiguousRB_TxRx_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlContiguousRBTxRxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_ContiguousRB_TxRx_r18_Supported},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_PSFCH_MultiContiguousRB_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiContiguousRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_PSFCH_MultiContiguousRB_r18_Supported},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_PSFCH_MultiNonContiguousRB_r18_Supported = 0
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiNonContiguousRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_PSFCH_MultiNonContiguousRB_r18_Supported},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N4  = 0
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N5  = 1
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N8  = 2
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N15 = 3
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N16 = 4
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N20 = 5
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlMultiplePRBCommonInterlacePSFCHR18TxTotalPRBPSFCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N4, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N5, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N8, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N15, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N16, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Tx_TotalPRB_PSFCH_r18_N20},
}

const (
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N5  = 0
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N6  = 1
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N15 = 2
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N16 = 3
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N25 = 4
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N26 = 5
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N32 = 6
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N35 = 7
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N45 = 8
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N46 = 9
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N50 = 10
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N64 = 11
	SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N65 = 12
)

var sharedSpectrumChAccessParamsSidelinkPerBandR18SlMultiplePRBCommonInterlacePSFCHR18RxTotalPRBPSFCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N5, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N6, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N15, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N16, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N25, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N26, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N32, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N35, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N45, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N46, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N50, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N64, SharedSpectrumChAccessParamsSidelinkPerBand_r18_Sl_MultiplePRB_CommonInterlacePSFCH_r18_Rx_TotalPRB_PSFCH_r18_N65},
}

type SharedSpectrumChAccessParamsSidelinkPerBand_r18 struct {
	Sl_DynamicChannelAccess_r18             *int64
	Sl_DynamicMultiChannelAccess_r18        *int64
	Sl_LBT_Option1_r18                      *int64
	Sl_LBT_Option2_r18                      *int64
	Sl_ResourceAllocMode1_r18               *int64
	Sl_Interlace_RB_TxRx_r18                *int64
	Sl_PSFCH_MultiOccasion_r18              *int64
	Sl_ContiguousRB_TxRx_r18                *int64
	Sl_PSFCH_MultiContiguousRB_r18          *int64
	Sl_PSFCH_MultiNonContiguousRB_r18       *int64
	Sl_MultiplePRB_CommonInterlacePSFCH_r18 *struct {
		Tx_TotalPRB_PSFCH_r18 int64
		Rx_TotalPRB_PSFCH_r18 int64
	}
	Sl_MultiplePRB_DedicatedInterlacePSFCH_r18 *struct {
		Tx_TotalPRB_PSFCH_r18 int64
		Rx_TotalPRB_PSFCH_r18 int64
	}
}

func (ie *SharedSpectrumChAccessParamsSidelinkPerBand_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedSpectrumChAccessParamsSidelinkPerBandR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DynamicChannelAccess_r18 != nil, ie.Sl_DynamicMultiChannelAccess_r18 != nil, ie.Sl_LBT_Option1_r18 != nil, ie.Sl_LBT_Option2_r18 != nil, ie.Sl_ResourceAllocMode1_r18 != nil, ie.Sl_Interlace_RB_TxRx_r18 != nil, ie.Sl_PSFCH_MultiOccasion_r18 != nil, ie.Sl_ContiguousRB_TxRx_r18 != nil, ie.Sl_PSFCH_MultiContiguousRB_r18 != nil, ie.Sl_PSFCH_MultiNonContiguousRB_r18 != nil, ie.Sl_MultiplePRB_CommonInterlacePSFCH_r18 != nil, ie.Sl_MultiplePRB_DedicatedInterlacePSFCH_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-DynamicChannelAccess-r18: enumerated
	{
		if ie.Sl_DynamicChannelAccess_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_DynamicChannelAccess_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlDynamicChannelAccessR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-DynamicMultiChannelAccess-r18: integer
	{
		if ie.Sl_DynamicMultiChannelAccess_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_DynamicMultiChannelAccess_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlDynamicMultiChannelAccessR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-LBT-Option1-r18: enumerated
	{
		if ie.Sl_LBT_Option1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LBT_Option1_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlLBTOption1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-LBT-Option2-r18: enumerated
	{
		if ie.Sl_LBT_Option2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LBT_Option2_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlLBTOption2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-ResourceAllocMode1-r18: enumerated
	{
		if ie.Sl_ResourceAllocMode1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ResourceAllocMode1_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlResourceAllocMode1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-Interlace-RB-TxRx-r18: enumerated
	{
		if ie.Sl_Interlace_RB_TxRx_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Interlace_RB_TxRx_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlInterlaceRBTxRxR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-PSFCH-MultiOccasion-r18: integer
	{
		if ie.Sl_PSFCH_MultiOccasion_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PSFCH_MultiOccasion_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-ContiguousRB-TxRx-r18: enumerated
	{
		if ie.Sl_ContiguousRB_TxRx_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ContiguousRB_TxRx_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlContiguousRBTxRxR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-PSFCH-MultiContiguousRB-r18: enumerated
	{
		if ie.Sl_PSFCH_MultiContiguousRB_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PSFCH_MultiContiguousRB_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiContiguousRBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-PSFCH-MultiNonContiguousRB-r18: enumerated
	{
		if ie.Sl_PSFCH_MultiNonContiguousRB_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PSFCH_MultiNonContiguousRB_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiNonContiguousRBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-MultiplePRB-CommonInterlacePSFCH-r18: sequence
	{
		if ie.Sl_MultiplePRB_CommonInterlacePSFCH_r18 != nil {
			c := ie.Sl_MultiplePRB_CommonInterlacePSFCH_r18
			if err := e.EncodeEnumerated(c.Tx_TotalPRB_PSFCH_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlMultiplePRBCommonInterlacePSFCHR18TxTotalPRBPSFCHR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Rx_TotalPRB_PSFCH_r18, sharedSpectrumChAccessParamsSidelinkPerBandR18SlMultiplePRBCommonInterlacePSFCHR18RxTotalPRBPSFCHR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sl-MultiplePRB-DedicatedInterlacePSFCH-r18: sequence
	{
		if ie.Sl_MultiplePRB_DedicatedInterlacePSFCH_r18 != nil {
			c := ie.Sl_MultiplePRB_DedicatedInterlacePSFCH_r18
			if err := e.EncodeInteger(c.Tx_TotalPRB_PSFCH_r18, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Rx_TotalPRB_PSFCH_r18, per.Constrained(1, 5)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SharedSpectrumChAccessParamsSidelinkPerBand_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedSpectrumChAccessParamsSidelinkPerBandR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-DynamicChannelAccess-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlDynamicChannelAccessR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DynamicChannelAccess_r18 = &idx
		}
	}

	// 3. sl-DynamicMultiChannelAccess-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sharedSpectrumChAccessParamsSidelinkPerBandR18SlDynamicMultiChannelAccessR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DynamicMultiChannelAccess_r18 = &v
		}
	}

	// 4. sl-LBT-Option1-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlLBTOption1R18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LBT_Option1_r18 = &idx
		}
	}

	// 5. sl-LBT-Option2-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlLBTOption2R18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LBT_Option2_r18 = &idx
		}
	}

	// 6. sl-ResourceAllocMode1-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlResourceAllocMode1R18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ResourceAllocMode1_r18 = &idx
		}
	}

	// 7. sl-Interlace-RB-TxRx-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlInterlaceRBTxRxR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Interlace_RB_TxRx_r18 = &idx
		}
	}

	// 8. sl-PSFCH-MultiOccasion-r18: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_MultiOccasion_r18 = &v
		}
	}

	// 9. sl-ContiguousRB-TxRx-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlContiguousRBTxRxR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ContiguousRB_TxRx_r18 = &idx
		}
	}

	// 10. sl-PSFCH-MultiContiguousRB-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiContiguousRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_MultiContiguousRB_r18 = &idx
		}
	}

	// 11. sl-PSFCH-MultiNonContiguousRB-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlPSFCHMultiNonContiguousRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_MultiNonContiguousRB_r18 = &idx
		}
	}

	// 12. sl-MultiplePRB-CommonInterlacePSFCH-r18: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.Sl_MultiplePRB_CommonInterlacePSFCH_r18 = &struct {
				Tx_TotalPRB_PSFCH_r18 int64
				Rx_TotalPRB_PSFCH_r18 int64
			}{}
			c := ie.Sl_MultiplePRB_CommonInterlacePSFCH_r18
			{
				v, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlMultiplePRBCommonInterlacePSFCHR18TxTotalPRBPSFCHR18Constraints)
				if err != nil {
					return err
				}
				c.Tx_TotalPRB_PSFCH_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsSidelinkPerBandR18SlMultiplePRBCommonInterlacePSFCHR18RxTotalPRBPSFCHR18Constraints)
				if err != nil {
					return err
				}
				c.Rx_TotalPRB_PSFCH_r18 = v
			}
		}
	}

	// 13. sl-MultiplePRB-DedicatedInterlacePSFCH-r18: sequence
	{
		if seq.IsComponentPresent(11) {
			ie.Sl_MultiplePRB_DedicatedInterlacePSFCH_r18 = &struct {
				Tx_TotalPRB_PSFCH_r18 int64
				Rx_TotalPRB_PSFCH_r18 int64
			}{}
			c := ie.Sl_MultiplePRB_DedicatedInterlacePSFCH_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.Tx_TotalPRB_PSFCH_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 5))
				if err != nil {
					return err
				}
				c.Rx_TotalPRB_PSFCH_r18 = v
			}
		}
	}

	return nil
}
