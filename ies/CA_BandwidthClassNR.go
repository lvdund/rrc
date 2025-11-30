package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_BandwidthClassNR_Enum_a         aper.Enumerated = 0
	CA_BandwidthClassNR_Enum_b         aper.Enumerated = 1
	CA_BandwidthClassNR_Enum_c         aper.Enumerated = 2
	CA_BandwidthClassNR_Enum_d         aper.Enumerated = 3
	CA_BandwidthClassNR_Enum_e         aper.Enumerated = 4
	CA_BandwidthClassNR_Enum_f         aper.Enumerated = 5
	CA_BandwidthClassNR_Enum_g         aper.Enumerated = 6
	CA_BandwidthClassNR_Enum_h         aper.Enumerated = 7
	CA_BandwidthClassNR_Enum_i         aper.Enumerated = 8
	CA_BandwidthClassNR_Enum_j         aper.Enumerated = 9
	CA_BandwidthClassNR_Enum_k         aper.Enumerated = 10
	CA_BandwidthClassNR_Enum_l         aper.Enumerated = 11
	CA_BandwidthClassNR_Enum_m         aper.Enumerated = 12
	CA_BandwidthClassNR_Enum_n         aper.Enumerated = 13
	CA_BandwidthClassNR_Enum_o         aper.Enumerated = 14
	CA_BandwidthClassNR_Enum_p         aper.Enumerated = 15
	CA_BandwidthClassNR_Enum_q         aper.Enumerated = 16
	CA_BandwidthClassNR_Enum_r2_v1730  aper.Enumerated = 17
	CA_BandwidthClassNR_Enum_r3_v1730  aper.Enumerated = 18
	CA_BandwidthClassNR_Enum_r4_v1730  aper.Enumerated = 19
	CA_BandwidthClassNR_Enum_r5_v1730  aper.Enumerated = 20
	CA_BandwidthClassNR_Enum_r6_v1730  aper.Enumerated = 21
	CA_BandwidthClassNR_Enum_r7_v1730  aper.Enumerated = 22
	CA_BandwidthClassNR_Enum_r8_v1730  aper.Enumerated = 23
	CA_BandwidthClassNR_Enum_r9_v1730  aper.Enumerated = 24
	CA_BandwidthClassNR_Enum_r10_v1730 aper.Enumerated = 25
	CA_BandwidthClassNR_Enum_r11_v1730 aper.Enumerated = 26
	CA_BandwidthClassNR_Enum_r12_v1730 aper.Enumerated = 27
)

type CA_BandwidthClassNR struct {
	Value aper.Enumerated `lb:0,ub:16,madatory`
}

func (ie *CA_BandwidthClassNR) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode CA_BandwidthClassNR", err)
	}
	return nil
}

func (ie *CA_BandwidthClassNR) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode CA_BandwidthClassNR", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
