package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasTiming_frequencyAndTiming struct {
	CarrierFreq                        ARFCN_ValueNR        `madatory`
	SsbSubcarrierSpacing               SubcarrierSpacing    `madatory`
	Ssb_MeasurementTimingConfiguration SSB_MTC              `madatory`
	Ss_RSSI_Measurement                *SS_RSSI_Measurement `optional`
}

func (ie *MeasTiming_frequencyAndTiming) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ss_RSSI_Measurement != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if err = ie.SsbSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SsbSubcarrierSpacing", err)
	}
	if err = ie.Ssb_MeasurementTimingConfiguration.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_MeasurementTimingConfiguration", err)
	}
	if ie.Ss_RSSI_Measurement != nil {
		if err = ie.Ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_RSSI_Measurement", err)
		}
	}
	return nil
}

func (ie *MeasTiming_frequencyAndTiming) Decode(r *aper.AperReader) error {
	var err error
	var Ss_RSSI_MeasurementPresent bool
	if Ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	if err = ie.SsbSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SsbSubcarrierSpacing", err)
	}
	if err = ie.Ssb_MeasurementTimingConfiguration.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_MeasurementTimingConfiguration", err)
	}
	if Ss_RSSI_MeasurementPresent {
		ie.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.Ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_RSSI_Measurement", err)
		}
	}
	return nil
}
