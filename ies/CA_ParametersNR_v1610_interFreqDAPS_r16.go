package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610_interFreqDAPS_r16 struct {
	InterFreqAsyncDAPS_r16                        *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqAsyncDAPS_r16                        `optional`
	InterFreqDiffSCS_DAPS_r16                     *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDiffSCS_DAPS_r16                     `optional`
	InterFreqMultiUL_TransmissionDAPS_r16         *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqMultiUL_TransmissionDAPS_r16         `optional`
	InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode1_r16 `optional`
	InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode2_r16 `optional`
	InterFreqDynamicPowerSharingDAPS_r16          *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDynamicPowerSharingDAPS_r16          `optional`
	InterFreqUL_TransCancellationDAPS_r16         *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqUL_TransCancellationDAPS_r16         `optional`
}

func (ie *CA_ParametersNR_v1610_interFreqDAPS_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.InterFreqAsyncDAPS_r16 != nil, ie.InterFreqDiffSCS_DAPS_r16 != nil, ie.InterFreqMultiUL_TransmissionDAPS_r16 != nil, ie.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 != nil, ie.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 != nil, ie.InterFreqDynamicPowerSharingDAPS_r16 != nil, ie.InterFreqUL_TransCancellationDAPS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InterFreqAsyncDAPS_r16 != nil {
		if err = ie.InterFreqAsyncDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqAsyncDAPS_r16", err)
		}
	}
	if ie.InterFreqDiffSCS_DAPS_r16 != nil {
		if err = ie.InterFreqDiffSCS_DAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqDiffSCS_DAPS_r16", err)
		}
	}
	if ie.InterFreqMultiUL_TransmissionDAPS_r16 != nil {
		if err = ie.InterFreqMultiUL_TransmissionDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqMultiUL_TransmissionDAPS_r16", err)
		}
	}
	if ie.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 != nil {
		if err = ie.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqSemiStaticPowerSharingDAPS_Mode1_r16", err)
		}
	}
	if ie.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 != nil {
		if err = ie.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqSemiStaticPowerSharingDAPS_Mode2_r16", err)
		}
	}
	if ie.InterFreqDynamicPowerSharingDAPS_r16 != nil {
		if err = ie.InterFreqDynamicPowerSharingDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqDynamicPowerSharingDAPS_r16", err)
		}
	}
	if ie.InterFreqUL_TransCancellationDAPS_r16 != nil {
		if err = ie.InterFreqUL_TransCancellationDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqUL_TransCancellationDAPS_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_interFreqDAPS_r16) Decode(r *aper.AperReader) error {
	var err error
	var InterFreqAsyncDAPS_r16Present bool
	if InterFreqAsyncDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqDiffSCS_DAPS_r16Present bool
	if InterFreqDiffSCS_DAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqMultiUL_TransmissionDAPS_r16Present bool
	if InterFreqMultiUL_TransmissionDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqSemiStaticPowerSharingDAPS_Mode1_r16Present bool
	if InterFreqSemiStaticPowerSharingDAPS_Mode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqSemiStaticPowerSharingDAPS_Mode2_r16Present bool
	if InterFreqSemiStaticPowerSharingDAPS_Mode2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqDynamicPowerSharingDAPS_r16Present bool
	if InterFreqDynamicPowerSharingDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqUL_TransCancellationDAPS_r16Present bool
	if InterFreqUL_TransCancellationDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InterFreqAsyncDAPS_r16Present {
		ie.InterFreqAsyncDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqAsyncDAPS_r16)
		if err = ie.InterFreqAsyncDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqAsyncDAPS_r16", err)
		}
	}
	if InterFreqDiffSCS_DAPS_r16Present {
		ie.InterFreqDiffSCS_DAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDiffSCS_DAPS_r16)
		if err = ie.InterFreqDiffSCS_DAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqDiffSCS_DAPS_r16", err)
		}
	}
	if InterFreqMultiUL_TransmissionDAPS_r16Present {
		ie.InterFreqMultiUL_TransmissionDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqMultiUL_TransmissionDAPS_r16)
		if err = ie.InterFreqMultiUL_TransmissionDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqMultiUL_TransmissionDAPS_r16", err)
		}
	}
	if InterFreqSemiStaticPowerSharingDAPS_Mode1_r16Present {
		ie.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode1_r16)
		if err = ie.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqSemiStaticPowerSharingDAPS_Mode1_r16", err)
		}
	}
	if InterFreqSemiStaticPowerSharingDAPS_Mode2_r16Present {
		ie.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode2_r16)
		if err = ie.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqSemiStaticPowerSharingDAPS_Mode2_r16", err)
		}
	}
	if InterFreqDynamicPowerSharingDAPS_r16Present {
		ie.InterFreqDynamicPowerSharingDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDynamicPowerSharingDAPS_r16)
		if err = ie.InterFreqDynamicPowerSharingDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqDynamicPowerSharingDAPS_r16", err)
		}
	}
	if InterFreqUL_TransCancellationDAPS_r16Present {
		ie.InterFreqUL_TransCancellationDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqUL_TransCancellationDAPS_r16)
		if err = ie.InterFreqUL_TransCancellationDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqUL_TransCancellationDAPS_r16", err)
		}
	}
	return nil
}
