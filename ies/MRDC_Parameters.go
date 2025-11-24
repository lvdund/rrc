package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters struct {
	SingleUL_Transmission         *MRDC_Parameters_singleUL_Transmission         `optional`
	DynamicPowerSharingENDC       *MRDC_Parameters_dynamicPowerSharingENDC       `optional`
	Tdm_Pattern                   *MRDC_Parameters_tdm_Pattern                   `optional`
	Ul_SharingEUTRA_NR            *MRDC_Parameters_ul_SharingEUTRA_NR            `optional`
	Ul_SwitchingTimeEUTRA_NR      *MRDC_Parameters_ul_SwitchingTimeEUTRA_NR      `optional`
	SimultaneousRxTxInterBandENDC *MRDC_Parameters_simultaneousRxTxInterBandENDC `optional`
	AsyncIntraBandENDC            *MRDC_Parameters_asyncIntraBandENDC            `optional`
	DualPA_Architecture           *MRDC_Parameters_dualPA_Architecture           `optional,ext-1`
	IntraBandENDC_Support         *MRDC_Parameters_intraBandENDC_Support         `optional,ext-1`
	Ul_TimingAlignmentEUTRA_NR    *MRDC_Parameters_ul_TimingAlignmentEUTRA_NR    `optional,ext-1`
}

func (ie *MRDC_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.DualPA_Architecture != nil || ie.IntraBandENDC_Support != nil || ie.Ul_TimingAlignmentEUTRA_NR != nil
	preambleBits := []bool{hasExtensions, ie.SingleUL_Transmission != nil, ie.DynamicPowerSharingENDC != nil, ie.Tdm_Pattern != nil, ie.Ul_SharingEUTRA_NR != nil, ie.Ul_SwitchingTimeEUTRA_NR != nil, ie.SimultaneousRxTxInterBandENDC != nil, ie.AsyncIntraBandENDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SingleUL_Transmission != nil {
		if err = ie.SingleUL_Transmission.Encode(w); err != nil {
			return utils.WrapError("Encode SingleUL_Transmission", err)
		}
	}
	if ie.DynamicPowerSharingENDC != nil {
		if err = ie.DynamicPowerSharingENDC.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPowerSharingENDC", err)
		}
	}
	if ie.Tdm_Pattern != nil {
		if err = ie.Tdm_Pattern.Encode(w); err != nil {
			return utils.WrapError("Encode Tdm_Pattern", err)
		}
	}
	if ie.Ul_SharingEUTRA_NR != nil {
		if err = ie.Ul_SharingEUTRA_NR.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_SharingEUTRA_NR", err)
		}
	}
	if ie.Ul_SwitchingTimeEUTRA_NR != nil {
		if err = ie.Ul_SwitchingTimeEUTRA_NR.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_SwitchingTimeEUTRA_NR", err)
		}
	}
	if ie.SimultaneousRxTxInterBandENDC != nil {
		if err = ie.SimultaneousRxTxInterBandENDC.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTxInterBandENDC", err)
		}
	}
	if ie.AsyncIntraBandENDC != nil {
		if err = ie.AsyncIntraBandENDC.Encode(w); err != nil {
			return utils.WrapError("Encode AsyncIntraBandENDC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.DualPA_Architecture != nil || ie.IntraBandENDC_Support != nil || ie.Ul_TimingAlignmentEUTRA_NR != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MRDC_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.DualPA_Architecture != nil, ie.IntraBandENDC_Support != nil, ie.Ul_TimingAlignmentEUTRA_NR != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DualPA_Architecture optional
			if ie.DualPA_Architecture != nil {
				if err = ie.DualPA_Architecture.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DualPA_Architecture", err)
				}
			}
			// encode IntraBandENDC_Support optional
			if ie.IntraBandENDC_Support != nil {
				if err = ie.IntraBandENDC_Support.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraBandENDC_Support", err)
				}
			}
			// encode Ul_TimingAlignmentEUTRA_NR optional
			if ie.Ul_TimingAlignmentEUTRA_NR != nil {
				if err = ie.Ul_TimingAlignmentEUTRA_NR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_TimingAlignmentEUTRA_NR", err)
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

func (ie *MRDC_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SingleUL_TransmissionPresent bool
	if SingleUL_TransmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicPowerSharingENDCPresent bool
	if DynamicPowerSharingENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdm_PatternPresent bool
	if Tdm_PatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_SharingEUTRA_NRPresent bool
	if Ul_SharingEUTRA_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_SwitchingTimeEUTRA_NRPresent bool
	if Ul_SwitchingTimeEUTRA_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousRxTxInterBandENDCPresent bool
	if SimultaneousRxTxInterBandENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AsyncIntraBandENDCPresent bool
	if AsyncIntraBandENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SingleUL_TransmissionPresent {
		ie.SingleUL_Transmission = new(MRDC_Parameters_singleUL_Transmission)
		if err = ie.SingleUL_Transmission.Decode(r); err != nil {
			return utils.WrapError("Decode SingleUL_Transmission", err)
		}
	}
	if DynamicPowerSharingENDCPresent {
		ie.DynamicPowerSharingENDC = new(MRDC_Parameters_dynamicPowerSharingENDC)
		if err = ie.DynamicPowerSharingENDC.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicPowerSharingENDC", err)
		}
	}
	if Tdm_PatternPresent {
		ie.Tdm_Pattern = new(MRDC_Parameters_tdm_Pattern)
		if err = ie.Tdm_Pattern.Decode(r); err != nil {
			return utils.WrapError("Decode Tdm_Pattern", err)
		}
	}
	if Ul_SharingEUTRA_NRPresent {
		ie.Ul_SharingEUTRA_NR = new(MRDC_Parameters_ul_SharingEUTRA_NR)
		if err = ie.Ul_SharingEUTRA_NR.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_SharingEUTRA_NR", err)
		}
	}
	if Ul_SwitchingTimeEUTRA_NRPresent {
		ie.Ul_SwitchingTimeEUTRA_NR = new(MRDC_Parameters_ul_SwitchingTimeEUTRA_NR)
		if err = ie.Ul_SwitchingTimeEUTRA_NR.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_SwitchingTimeEUTRA_NR", err)
		}
	}
	if SimultaneousRxTxInterBandENDCPresent {
		ie.SimultaneousRxTxInterBandENDC = new(MRDC_Parameters_simultaneousRxTxInterBandENDC)
		if err = ie.SimultaneousRxTxInterBandENDC.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTxInterBandENDC", err)
		}
	}
	if AsyncIntraBandENDCPresent {
		ie.AsyncIntraBandENDC = new(MRDC_Parameters_asyncIntraBandENDC)
		if err = ie.AsyncIntraBandENDC.Decode(r); err != nil {
			return utils.WrapError("Decode AsyncIntraBandENDC", err)
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

			DualPA_ArchitecturePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraBandENDC_SupportPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_TimingAlignmentEUTRA_NRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DualPA_Architecture optional
			if DualPA_ArchitecturePresent {
				ie.DualPA_Architecture = new(MRDC_Parameters_dualPA_Architecture)
				if err = ie.DualPA_Architecture.Decode(extReader); err != nil {
					return utils.WrapError("Decode DualPA_Architecture", err)
				}
			}
			// decode IntraBandENDC_Support optional
			if IntraBandENDC_SupportPresent {
				ie.IntraBandENDC_Support = new(MRDC_Parameters_intraBandENDC_Support)
				if err = ie.IntraBandENDC_Support.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraBandENDC_Support", err)
				}
			}
			// decode Ul_TimingAlignmentEUTRA_NR optional
			if Ul_TimingAlignmentEUTRA_NRPresent {
				ie.Ul_TimingAlignmentEUTRA_NR = new(MRDC_Parameters_ul_TimingAlignmentEUTRA_NR)
				if err = ie.Ul_TimingAlignmentEUTRA_NR.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_TimingAlignmentEUTRA_NR", err)
				}
			}
		}
	}
	return nil
}
