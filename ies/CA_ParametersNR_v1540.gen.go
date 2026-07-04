// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1540 (line 17282).

var cAParametersNRV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "simultaneousSRS-AssocCSI-RS-AllCC", Optional: true},
		{Name: "csi-RS-IM-ReceptionForFeedbackPerBandComb", Optional: true},
		{Name: "simultaneousCSI-ReportsAllCC", Optional: true},
		{Name: "dualPA-Architecture", Optional: true},
	},
}

var cAParametersNRV1540SimultaneousSRSAssocCSIRSAllCCConstraints = per.Constrained(5, 32)

var cAParametersNRV1540SimultaneousCSIReportsAllCCConstraints = per.Constrained(5, 32)

const (
	CA_ParametersNR_v1540_DualPA_Architecture_Supported = 0
)

var cAParametersNRV1540DualPAArchitectureConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1540_DualPA_Architecture_Supported},
}

var cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSimultaneousNZP-CSI-RS-ActBWP-AllCC", Optional: true},
		{Name: "totalNumberPortsSimultaneousNZP-CSI-RS-ActBWP-AllCC", Optional: true},
	},
}

type CA_ParametersNR_v1540 struct {
	SimultaneousSRS_AssocCSI_RS_AllCC         *int64
	Csi_RS_IM_ReceptionForFeedbackPerBandComb *struct {
		MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC        *int64
		TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC *int64
	}
	SimultaneousCSI_ReportsAllCC *int64
	DualPA_Architecture          *int64
}

func (ie *CA_ParametersNR_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SimultaneousSRS_AssocCSI_RS_AllCC != nil, ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb != nil, ie.SimultaneousCSI_ReportsAllCC != nil, ie.DualPA_Architecture != nil}); err != nil {
		return err
	}

	// 2. simultaneousSRS-AssocCSI-RS-AllCC: integer
	{
		if ie.SimultaneousSRS_AssocCSI_RS_AllCC != nil {
			if err := e.EncodeInteger(*ie.SimultaneousSRS_AssocCSI_RS_AllCC, cAParametersNRV1540SimultaneousSRSAssocCSIRSAllCCConstraints); err != nil {
				return err
			}
		}
	}

	// 3. csi-RS-IM-ReceptionForFeedbackPerBandComb: sequence
	{
		if ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb != nil {
			c := ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb
			cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombSeq := e.NewSequenceEncoder(cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombConstraints)
			if err := cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombSeq.EncodePreamble([]bool{c.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil, c.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil}); err != nil {
				return err
			}
			if c.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil {
				if err := e.EncodeInteger((*c.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC), per.Constrained(1, 64)); err != nil {
					return err
				}
			}
			if c.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil {
				if err := e.EncodeInteger((*c.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC), per.Constrained(2, 256)); err != nil {
					return err
				}
			}
		}
	}

	// 4. simultaneousCSI-ReportsAllCC: integer
	{
		if ie.SimultaneousCSI_ReportsAllCC != nil {
			if err := e.EncodeInteger(*ie.SimultaneousCSI_ReportsAllCC, cAParametersNRV1540SimultaneousCSIReportsAllCCConstraints); err != nil {
				return err
			}
		}
	}

	// 5. dualPA-Architecture: enumerated
	{
		if ie.DualPA_Architecture != nil {
			if err := e.EncodeEnumerated(*ie.DualPA_Architecture, cAParametersNRV1540DualPAArchitectureConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. simultaneousSRS-AssocCSI-RS-AllCC: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cAParametersNRV1540SimultaneousSRSAssocCSIRSAllCCConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousSRS_AssocCSI_RS_AllCC = &v
		}
	}

	// 3. csi-RS-IM-ReceptionForFeedbackPerBandComb: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb = &struct {
				MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC        *int64
				TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC *int64
			}{}
			c := ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb
			cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombSeq := d.NewSequenceDecoder(cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombConstraints)
			if err := cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombSeq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombSeq.IsComponentPresent(0) {
				c.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 64))
				if err != nil {
					return err
				}
				(*c.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC) = v
			}
			if cAParametersNRV1540CsiRSIMReceptionForFeedbackPerBandCombSeq.IsComponentPresent(1) {
				c.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC = new(int64)
				v, err := d.DecodeInteger(per.Constrained(2, 256))
				if err != nil {
					return err
				}
				(*c.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC) = v
			}
		}
	}

	// 4. simultaneousCSI-ReportsAllCC: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cAParametersNRV1540SimultaneousCSIReportsAllCCConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousCSI_ReportsAllCC = &v
		}
	}

	// 5. dualPA-Architecture: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1540DualPAArchitectureConstraints)
			if err != nil {
				return err
			}
			ie.DualPA_Architecture = &idx
		}
	}

	return nil
}
