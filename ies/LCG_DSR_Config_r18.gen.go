// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LCG-DSR-Config-r18 (line 9071).

var lCGDSRConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lcg-Id-r18"},
		{Name: "remainingTimeThreshold-r18"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var lCGDSRConfigR18RemainingTimeThresholdR18Constraints = per.Constrained(1, 64)

var lCGDSRConfigR18ExtMultiEntryDSRR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dsr-ReportingThresList-r19"},
		{Name: "dsr-ReportNonDelayCriticalData-r19", Optional: true},
	},
}

var lCGDSRConfigR18ExtMultiEntryDSRR19DsrReportingThresListR19Constraints = per.SizeRange(1, common.MaxNrOfDSR_ReportingThres_r19)

const (
	LCG_DSR_Config_r18_Ext_MultiEntryDSR_r19_Dsr_ReportNonDelayCriticalData_r19_Enabled = 0
)

var lCGDSRConfigR18ExtMultiEntryDSRR19DsrReportNonDelayCriticalDataR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LCG_DSR_Config_r18_Ext_MultiEntryDSR_r19_Dsr_ReportNonDelayCriticalData_r19_Enabled},
}

type LCG_DSR_Config_r18 struct {
	Lcg_Id_r18                 LCG_Id_r18
	RemainingTimeThreshold_r18 int64
	MultiEntryDSR_r19          *struct {
		Dsr_ReportingThresList_r19         []DSR_ReportingThreshold_r19
		Dsr_ReportNonDelayCriticalData_r19 *int64
	}
}

func (ie *LCG_DSR_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lCGDSRConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MultiEntryDSR_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. lcg-Id-r18: ref
	{
		if err := ie.Lcg_Id_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. remainingTimeThreshold-r18: integer
	{
		if err := e.EncodeInteger(ie.RemainingTimeThreshold_r18, lCGDSRConfigR18RemainingTimeThresholdR18Constraints); err != nil {
			return err
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
					{Name: "multiEntryDSR-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultiEntryDSR_r19 != nil}); err != nil {
				return err
			}

			if ie.MultiEntryDSR_r19 != nil {
				c := ie.MultiEntryDSR_r19
				lCGDSRConfigR18ExtMultiEntryDSRR19Seq := ex.NewSequenceEncoder(lCGDSRConfigR18ExtMultiEntryDSRR19Constraints)
				if err := lCGDSRConfigR18ExtMultiEntryDSRR19Seq.EncodePreamble([]bool{c.Dsr_ReportNonDelayCriticalData_r19 != nil}); err != nil {
					return err
				}
				{
					seqOf := ex.NewSequenceOfEncoder(lCGDSRConfigR18ExtMultiEntryDSRR19DsrReportingThresListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Dsr_ReportingThresList_r19))); err != nil {
						return err
					}
					for i := range c.Dsr_ReportingThresList_r19 {
						if err := c.Dsr_ReportingThresList_r19[i].Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.Dsr_ReportNonDelayCriticalData_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Dsr_ReportNonDelayCriticalData_r19), lCGDSRConfigR18ExtMultiEntryDSRR19DsrReportNonDelayCriticalDataR19Constraints); err != nil {
						return err
					}
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

func (ie *LCG_DSR_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lCGDSRConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. lcg-Id-r18: ref
	{
		if err := ie.Lcg_Id_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. remainingTimeThreshold-r18: integer
	{
		v1, err := d.DecodeInteger(lCGDSRConfigR18RemainingTimeThresholdR18Constraints)
		if err != nil {
			return err
		}
		ie.RemainingTimeThreshold_r18 = v1
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
				{Name: "multiEntryDSR-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MultiEntryDSR_r19 = &struct {
				Dsr_ReportingThresList_r19         []DSR_ReportingThreshold_r19
				Dsr_ReportNonDelayCriticalData_r19 *int64
			}{}
			c := ie.MultiEntryDSR_r19
			lCGDSRConfigR18ExtMultiEntryDSRR19Seq := dx.NewSequenceDecoder(lCGDSRConfigR18ExtMultiEntryDSRR19Constraints)
			if err := lCGDSRConfigR18ExtMultiEntryDSRR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				seqOf := dx.NewSequenceOfDecoder(lCGDSRConfigR18ExtMultiEntryDSRR19DsrReportingThresListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Dsr_ReportingThresList_r19 = make([]DSR_ReportingThreshold_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.Dsr_ReportingThresList_r19[i].Decode(dx); err != nil {
						return err
					}
				}
			}
			if lCGDSRConfigR18ExtMultiEntryDSRR19Seq.IsComponentPresent(1) {
				c.Dsr_ReportNonDelayCriticalData_r19 = new(int64)
				v, err := dx.DecodeEnumerated(lCGDSRConfigR18ExtMultiEntryDSRR19DsrReportNonDelayCriticalDataR19Constraints)
				if err != nil {
					return err
				}
				(*c.Dsr_ReportNonDelayCriticalData_r19) = v
			}
		}
	}

	return nil
}
