// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ERedCap-ConfigCommonSIB-r18 (line 2100).

var eRedCapConfigCommonSIBR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellBarred-eRedCap-r18"},
	},
}

const (
	ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap1Rx_r18_Barred    = 0
	ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap1Rx_r18_NotBarred = 1
)

var eRedCapConfigCommonSIBR18CellBarredERedCapR18CellBarredERedCap1RxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap1Rx_r18_Barred, ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap1Rx_r18_NotBarred},
}

const (
	ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap2Rx_r18_Barred    = 0
	ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap2Rx_r18_NotBarred = 1
)

var eRedCapConfigCommonSIBR18CellBarredERedCapR18CellBarredERedCap2RxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap2Rx_r18_Barred, ERedCap_ConfigCommonSIB_r18_CellBarred_ERedCap_r18_CellBarred_ERedCap2Rx_r18_NotBarred},
}

type ERedCap_ConfigCommonSIB_r18 struct {
	CellBarred_ERedCap_r18 struct {
		CellBarred_ERedCap1Rx_r18 int64
		CellBarred_ERedCap2Rx_r18 int64
	}
}

func (ie *ERedCap_ConfigCommonSIB_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eRedCapConfigCommonSIBR18Constraints)
	_ = seq

	// 1. cellBarred-eRedCap-r18: sequence
	{
		{
			c := &ie.CellBarred_ERedCap_r18
			if err := e.EncodeEnumerated(c.CellBarred_ERedCap1Rx_r18, eRedCapConfigCommonSIBR18CellBarredERedCapR18CellBarredERedCap1RxR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.CellBarred_ERedCap2Rx_r18, eRedCapConfigCommonSIBR18CellBarredERedCapR18CellBarredERedCap2RxR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ERedCap_ConfigCommonSIB_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eRedCapConfigCommonSIBR18Constraints)
	_ = seq

	// 1. cellBarred-eRedCap-r18: sequence
	{
		{
			c := &ie.CellBarred_ERedCap_r18
			{
				v, err := d.DecodeEnumerated(eRedCapConfigCommonSIBR18CellBarredERedCapR18CellBarredERedCap1RxR18Constraints)
				if err != nil {
					return err
				}
				c.CellBarred_ERedCap1Rx_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(eRedCapConfigCommonSIBR18CellBarredERedCapR18CellBarredERedCap2RxR18Constraints)
				if err != nil {
					return err
				}
				c.CellBarred_ERedCap2Rx_r18 = v
			}
		}
	}

	return nil
}
