package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_BasedPerfMeas_Parameters_r16 struct {
	BarometerMeasReport_r16            *UE_BasedPerfMeas_Parameters_r16_barometerMeasReport_r16            `optional`
	ImmMeasBT_r16                      *UE_BasedPerfMeas_Parameters_r16_immMeasBT_r16                      `optional`
	ImmMeasWLAN_r16                    *UE_BasedPerfMeas_Parameters_r16_immMeasWLAN_r16                    `optional`
	LoggedMeasBT_r16                   *UE_BasedPerfMeas_Parameters_r16_loggedMeasBT_r16                   `optional`
	LoggedMeasurements_r16             *UE_BasedPerfMeas_Parameters_r16_loggedMeasurements_r16             `optional`
	LoggedMeasWLAN_r16                 *UE_BasedPerfMeas_Parameters_r16_loggedMeasWLAN_r16                 `optional`
	OrientationMeasReport_r16          *UE_BasedPerfMeas_Parameters_r16_orientationMeasReport_r16          `optional`
	SpeedMeasReport_r16                *UE_BasedPerfMeas_Parameters_r16_speedMeasReport_r16                `optional`
	Gnss_Location_r16                  *UE_BasedPerfMeas_Parameters_r16_gnss_Location_r16                  `optional`
	UlPDCP_Delay_r16                   *UE_BasedPerfMeas_Parameters_r16_ulPDCP_Delay_r16                   `optional`
	SigBasedLogMDT_OverrideProtect_r17 *UE_BasedPerfMeas_Parameters_r16_sigBasedLogMDT_OverrideProtect_r17 `optional,ext-1`
	MultipleCEF_Report_r17             *UE_BasedPerfMeas_Parameters_r16_multipleCEF_Report_r17             `optional,ext-1`
	ExcessPacketDelay_r17              *UE_BasedPerfMeas_Parameters_r16_excessPacketDelay_r17              `optional,ext-1`
	EarlyMeasLog_r17                   *UE_BasedPerfMeas_Parameters_r16_earlyMeasLog_r17                   `optional,ext-1`
}

func (ie *UE_BasedPerfMeas_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.SigBasedLogMDT_OverrideProtect_r17 != nil || ie.MultipleCEF_Report_r17 != nil || ie.ExcessPacketDelay_r17 != nil || ie.EarlyMeasLog_r17 != nil
	preambleBits := []bool{hasExtensions, ie.BarometerMeasReport_r16 != nil, ie.ImmMeasBT_r16 != nil, ie.ImmMeasWLAN_r16 != nil, ie.LoggedMeasBT_r16 != nil, ie.LoggedMeasurements_r16 != nil, ie.LoggedMeasWLAN_r16 != nil, ie.OrientationMeasReport_r16 != nil, ie.SpeedMeasReport_r16 != nil, ie.Gnss_Location_r16 != nil, ie.UlPDCP_Delay_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BarometerMeasReport_r16 != nil {
		if err = ie.BarometerMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BarometerMeasReport_r16", err)
		}
	}
	if ie.ImmMeasBT_r16 != nil {
		if err = ie.ImmMeasBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ImmMeasBT_r16", err)
		}
	}
	if ie.ImmMeasWLAN_r16 != nil {
		if err = ie.ImmMeasWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ImmMeasWLAN_r16", err)
		}
	}
	if ie.LoggedMeasBT_r16 != nil {
		if err = ie.LoggedMeasBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LoggedMeasBT_r16", err)
		}
	}
	if ie.LoggedMeasurements_r16 != nil {
		if err = ie.LoggedMeasurements_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LoggedMeasurements_r16", err)
		}
	}
	if ie.LoggedMeasWLAN_r16 != nil {
		if err = ie.LoggedMeasWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LoggedMeasWLAN_r16", err)
		}
	}
	if ie.OrientationMeasReport_r16 != nil {
		if err = ie.OrientationMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OrientationMeasReport_r16", err)
		}
	}
	if ie.SpeedMeasReport_r16 != nil {
		if err = ie.SpeedMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpeedMeasReport_r16", err)
		}
	}
	if ie.Gnss_Location_r16 != nil {
		if err = ie.Gnss_Location_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnss_Location_r16", err)
		}
	}
	if ie.UlPDCP_Delay_r16 != nil {
		if err = ie.UlPDCP_Delay_r16.Encode(w); err != nil {
			return utils.WrapError("Encode UlPDCP_Delay_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.SigBasedLogMDT_OverrideProtect_r17 != nil || ie.MultipleCEF_Report_r17 != nil || ie.ExcessPacketDelay_r17 != nil || ie.EarlyMeasLog_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UE_BasedPerfMeas_Parameters_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SigBasedLogMDT_OverrideProtect_r17 != nil, ie.MultipleCEF_Report_r17 != nil, ie.ExcessPacketDelay_r17 != nil, ie.EarlyMeasLog_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SigBasedLogMDT_OverrideProtect_r17 optional
			if ie.SigBasedLogMDT_OverrideProtect_r17 != nil {
				if err = ie.SigBasedLogMDT_OverrideProtect_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SigBasedLogMDT_OverrideProtect_r17", err)
				}
			}
			// encode MultipleCEF_Report_r17 optional
			if ie.MultipleCEF_Report_r17 != nil {
				if err = ie.MultipleCEF_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MultipleCEF_Report_r17", err)
				}
			}
			// encode ExcessPacketDelay_r17 optional
			if ie.ExcessPacketDelay_r17 != nil {
				if err = ie.ExcessPacketDelay_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExcessPacketDelay_r17", err)
				}
			}
			// encode EarlyMeasLog_r17 optional
			if ie.EarlyMeasLog_r17 != nil {
				if err = ie.EarlyMeasLog_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EarlyMeasLog_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *UE_BasedPerfMeas_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var BarometerMeasReport_r16Present bool
	if BarometerMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ImmMeasBT_r16Present bool
	if ImmMeasBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ImmMeasWLAN_r16Present bool
	if ImmMeasWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LoggedMeasBT_r16Present bool
	if LoggedMeasBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LoggedMeasurements_r16Present bool
	if LoggedMeasurements_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LoggedMeasWLAN_r16Present bool
	if LoggedMeasWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OrientationMeasReport_r16Present bool
	if OrientationMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SpeedMeasReport_r16Present bool
	if SpeedMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Gnss_Location_r16Present bool
	if Gnss_Location_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UlPDCP_Delay_r16Present bool
	if UlPDCP_Delay_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BarometerMeasReport_r16Present {
		ie.BarometerMeasReport_r16 = new(UE_BasedPerfMeas_Parameters_r16_barometerMeasReport_r16)
		if err = ie.BarometerMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BarometerMeasReport_r16", err)
		}
	}
	if ImmMeasBT_r16Present {
		ie.ImmMeasBT_r16 = new(UE_BasedPerfMeas_Parameters_r16_immMeasBT_r16)
		if err = ie.ImmMeasBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ImmMeasBT_r16", err)
		}
	}
	if ImmMeasWLAN_r16Present {
		ie.ImmMeasWLAN_r16 = new(UE_BasedPerfMeas_Parameters_r16_immMeasWLAN_r16)
		if err = ie.ImmMeasWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ImmMeasWLAN_r16", err)
		}
	}
	if LoggedMeasBT_r16Present {
		ie.LoggedMeasBT_r16 = new(UE_BasedPerfMeas_Parameters_r16_loggedMeasBT_r16)
		if err = ie.LoggedMeasBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LoggedMeasBT_r16", err)
		}
	}
	if LoggedMeasurements_r16Present {
		ie.LoggedMeasurements_r16 = new(UE_BasedPerfMeas_Parameters_r16_loggedMeasurements_r16)
		if err = ie.LoggedMeasurements_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LoggedMeasurements_r16", err)
		}
	}
	if LoggedMeasWLAN_r16Present {
		ie.LoggedMeasWLAN_r16 = new(UE_BasedPerfMeas_Parameters_r16_loggedMeasWLAN_r16)
		if err = ie.LoggedMeasWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LoggedMeasWLAN_r16", err)
		}
	}
	if OrientationMeasReport_r16Present {
		ie.OrientationMeasReport_r16 = new(UE_BasedPerfMeas_Parameters_r16_orientationMeasReport_r16)
		if err = ie.OrientationMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OrientationMeasReport_r16", err)
		}
	}
	if SpeedMeasReport_r16Present {
		ie.SpeedMeasReport_r16 = new(UE_BasedPerfMeas_Parameters_r16_speedMeasReport_r16)
		if err = ie.SpeedMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpeedMeasReport_r16", err)
		}
	}
	if Gnss_Location_r16Present {
		ie.Gnss_Location_r16 = new(UE_BasedPerfMeas_Parameters_r16_gnss_Location_r16)
		if err = ie.Gnss_Location_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnss_Location_r16", err)
		}
	}
	if UlPDCP_Delay_r16Present {
		ie.UlPDCP_Delay_r16 = new(UE_BasedPerfMeas_Parameters_r16_ulPDCP_Delay_r16)
		if err = ie.UlPDCP_Delay_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UlPDCP_Delay_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SigBasedLogMDT_OverrideProtect_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MultipleCEF_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExcessPacketDelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EarlyMeasLog_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SigBasedLogMDT_OverrideProtect_r17 optional
			if SigBasedLogMDT_OverrideProtect_r17Present {
				ie.SigBasedLogMDT_OverrideProtect_r17 = new(UE_BasedPerfMeas_Parameters_r16_sigBasedLogMDT_OverrideProtect_r17)
				if err = ie.SigBasedLogMDT_OverrideProtect_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SigBasedLogMDT_OverrideProtect_r17", err)
				}
			}
			// decode MultipleCEF_Report_r17 optional
			if MultipleCEF_Report_r17Present {
				ie.MultipleCEF_Report_r17 = new(UE_BasedPerfMeas_Parameters_r16_multipleCEF_Report_r17)
				if err = ie.MultipleCEF_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MultipleCEF_Report_r17", err)
				}
			}
			// decode ExcessPacketDelay_r17 optional
			if ExcessPacketDelay_r17Present {
				ie.ExcessPacketDelay_r17 = new(UE_BasedPerfMeas_Parameters_r16_excessPacketDelay_r17)
				if err = ie.ExcessPacketDelay_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExcessPacketDelay_r17", err)
				}
			}
			// decode EarlyMeasLog_r17 optional
			if EarlyMeasLog_r17Present {
				ie.EarlyMeasLog_r17 = new(UE_BasedPerfMeas_Parameters_r16_earlyMeasLog_r17)
				if err = ie.EarlyMeasLog_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode EarlyMeasLog_r17", err)
				}
			}
		}
	}
	return nil
}
