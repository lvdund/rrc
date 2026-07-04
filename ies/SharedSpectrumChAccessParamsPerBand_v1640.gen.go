// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SharedSpectrumChAccessParamsPerBand-v1640 (line 24951).

var sharedSpectrumChAccessParamsPerBandV1640Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RSRP-AndRSRQ-MeasWithSSB-r16", Optional: true},
		{Name: "csi-RSRP-AndRSRQ-MeasWithoutSSB-r16", Optional: true},
		{Name: "csi-SINR-Meas-r16", Optional: true},
		{Name: "ssb-AndCSI-RS-RLM-r16", Optional: true},
		{Name: "csi-RS-CFRA-ForHO-r16", Optional: true},
	},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1640_Csi_RSRP_AndRSRQ_MeasWithSSB_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1640CsiRSRPAndRSRQMeasWithSSBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1640_Csi_RSRP_AndRSRQ_MeasWithSSB_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1640_Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1640CsiRSRPAndRSRQMeasWithoutSSBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1640_Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1640_Csi_SINR_Meas_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1640CsiSINRMeasR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1640_Csi_SINR_Meas_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1640_Ssb_AndCSI_RS_RLM_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1640SsbAndCSIRSRLMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1640_Ssb_AndCSI_RS_RLM_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1640_Csi_RS_CFRA_ForHO_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1640CsiRSCFRAForHOR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1640_Csi_RS_CFRA_ForHO_r16_Supported},
}

type SharedSpectrumChAccessParamsPerBand_v1640 struct {
	Csi_RSRP_AndRSRQ_MeasWithSSB_r16    *int64
	Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 *int64
	Csi_SINR_Meas_r16                   *int64
	Ssb_AndCSI_RS_RLM_r16               *int64
	Csi_RS_CFRA_ForHO_r16               *int64
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1640) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedSpectrumChAccessParamsPerBandV1640Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16 != nil, ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 != nil, ie.Csi_SINR_Meas_r16 != nil, ie.Ssb_AndCSI_RS_RLM_r16 != nil, ie.Csi_RS_CFRA_ForHO_r16 != nil}); err != nil {
		return err
	}

	// 2. csi-RSRP-AndRSRQ-MeasWithSSB-r16: enumerated
	{
		if ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16, sharedSpectrumChAccessParamsPerBandV1640CsiRSRPAndRSRQMeasWithSSBR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. csi-RSRP-AndRSRQ-MeasWithoutSSB-r16: enumerated
	{
		if ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16, sharedSpectrumChAccessParamsPerBandV1640CsiRSRPAndRSRQMeasWithoutSSBR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. csi-SINR-Meas-r16: enumerated
	{
		if ie.Csi_SINR_Meas_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_SINR_Meas_r16, sharedSpectrumChAccessParamsPerBandV1640CsiSINRMeasR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ssb-AndCSI-RS-RLM-r16: enumerated
	{
		if ie.Ssb_AndCSI_RS_RLM_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_AndCSI_RS_RLM_r16, sharedSpectrumChAccessParamsPerBandV1640SsbAndCSIRSRLMR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. csi-RS-CFRA-ForHO-r16: enumerated
	{
		if ie.Csi_RS_CFRA_ForHO_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RS_CFRA_ForHO_r16, sharedSpectrumChAccessParamsPerBandV1640CsiRSCFRAForHOR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1640) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedSpectrumChAccessParamsPerBandV1640Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. csi-RSRP-AndRSRQ-MeasWithSSB-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1640CsiRSRPAndRSRQMeasWithSSBR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16 = &idx
		}
	}

	// 3. csi-RSRP-AndRSRQ-MeasWithoutSSB-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1640CsiRSRPAndRSRQMeasWithoutSSBR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 = &idx
		}
	}

	// 4. csi-SINR-Meas-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1640CsiSINRMeasR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_SINR_Meas_r16 = &idx
		}
	}

	// 5. ssb-AndCSI-RS-RLM-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1640SsbAndCSIRSRLMR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_AndCSI_RS_RLM_r16 = &idx
		}
	}

	// 6. csi-RS-CFRA-ForHO-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1640CsiRSCFRAForHOR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_CFRA_ForHO_r16 = &idx
		}
	}

	return nil
}
