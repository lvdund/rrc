package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17_Enum_n50  aper.Enumerated = 0
	CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17_Enum_n60  aper.Enumerated = 1
	CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17_Enum_n70  aper.Enumerated = 2
	CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17_Enum_n80  aper.Enumerated = 3
	CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17_Enum_n90  aper.Enumerated = 4
	CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17_Enum_n100 aper.Enumerated = 5
)

type CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
