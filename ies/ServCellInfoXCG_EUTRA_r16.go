package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoXCG_EUTRA_r16 struct {
	Dl_CarrierFreq_EUTRA_r16        *ARFCN_ValueEUTRA                `optional`
	Ul_CarrierFreq_EUTRA_r16        *ARFCN_ValueEUTRA                `optional`
	TransmissionBandwidth_EUTRA_r16 *TransmissionBandwidth_EUTRA_r16 `optional`
}

func (ie *ServCellInfoXCG_EUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_CarrierFreq_EUTRA_r16 != nil, ie.Ul_CarrierFreq_EUTRA_r16 != nil, ie.TransmissionBandwidth_EUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_CarrierFreq_EUTRA_r16 != nil {
		if err = ie.Dl_CarrierFreq_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_CarrierFreq_EUTRA_r16", err)
		}
	}
	if ie.Ul_CarrierFreq_EUTRA_r16 != nil {
		if err = ie.Ul_CarrierFreq_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_CarrierFreq_EUTRA_r16", err)
		}
	}
	if ie.TransmissionBandwidth_EUTRA_r16 != nil {
		if err = ie.TransmissionBandwidth_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TransmissionBandwidth_EUTRA_r16", err)
		}
	}
	return nil
}

func (ie *ServCellInfoXCG_EUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var Dl_CarrierFreq_EUTRA_r16Present bool
	if Dl_CarrierFreq_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_CarrierFreq_EUTRA_r16Present bool
	if Ul_CarrierFreq_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TransmissionBandwidth_EUTRA_r16Present bool
	if TransmissionBandwidth_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_CarrierFreq_EUTRA_r16Present {
		ie.Dl_CarrierFreq_EUTRA_r16 = new(ARFCN_ValueEUTRA)
		if err = ie.Dl_CarrierFreq_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_CarrierFreq_EUTRA_r16", err)
		}
	}
	if Ul_CarrierFreq_EUTRA_r16Present {
		ie.Ul_CarrierFreq_EUTRA_r16 = new(ARFCN_ValueEUTRA)
		if err = ie.Ul_CarrierFreq_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_CarrierFreq_EUTRA_r16", err)
		}
	}
	if TransmissionBandwidth_EUTRA_r16Present {
		ie.TransmissionBandwidth_EUTRA_r16 = new(TransmissionBandwidth_EUTRA_r16)
		if err = ie.TransmissionBandwidth_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TransmissionBandwidth_EUTRA_r16", err)
		}
	}
	return nil
}
