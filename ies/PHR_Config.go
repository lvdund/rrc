package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PHR_Config struct {
	Phr_PeriodicTimer        PHR_Config_phr_PeriodicTimer        `madatory`
	Phr_ProhibitTimer        PHR_Config_phr_ProhibitTimer        `madatory`
	Phr_Tx_PowerFactorChange PHR_Config_phr_Tx_PowerFactorChange `madatory`
	MultiplePHR              bool                                `madatory`
	Dummy                    bool                                `madatory`
	Phr_Type2OtherCell       bool                                `madatory`
	Phr_ModeOtherCG          PHR_Config_phr_ModeOtherCG          `madatory`
	Mpe_Reporting_FR2_r16    *MPE_Config_FR2_r16                 `optional,ext-1,setuprelease`
	Mpe_Reporting_FR2_r17    *MPE_Config_FR2_r17                 `optional,ext-2,setuprelease`
	TwoPHRMode_r17           *PHR_Config_twoPHRMode_r17          `optional,ext-2`
}

func (ie *PHR_Config) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Mpe_Reporting_FR2_r16 != nil || ie.Mpe_Reporting_FR2_r17 != nil || ie.TwoPHRMode_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Phr_PeriodicTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Phr_PeriodicTimer", err)
	}
	if err = ie.Phr_ProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Phr_ProhibitTimer", err)
	}
	if err = ie.Phr_Tx_PowerFactorChange.Encode(w); err != nil {
		return utils.WrapError("Encode Phr_Tx_PowerFactorChange", err)
	}
	if err = w.WriteBoolean(ie.MultiplePHR); err != nil {
		return utils.WrapError("WriteBoolean MultiplePHR", err)
	}
	if err = w.WriteBoolean(ie.Dummy); err != nil {
		return utils.WrapError("WriteBoolean Dummy", err)
	}
	if err = w.WriteBoolean(ie.Phr_Type2OtherCell); err != nil {
		return utils.WrapError("WriteBoolean Phr_Type2OtherCell", err)
	}
	if err = ie.Phr_ModeOtherCG.Encode(w); err != nil {
		return utils.WrapError("Encode Phr_ModeOtherCG", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Mpe_Reporting_FR2_r16 != nil, ie.Mpe_Reporting_FR2_r17 != nil || ie.TwoPHRMode_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PHR_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Mpe_Reporting_FR2_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Mpe_Reporting_FR2_r16 optional
			if ie.Mpe_Reporting_FR2_r16 != nil {
				tmp_Mpe_Reporting_FR2_r16 := utils.SetupRelease[*MPE_Config_FR2_r16]{
					Setup: ie.Mpe_Reporting_FR2_r16,
				}
				if err = tmp_Mpe_Reporting_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpe_Reporting_FR2_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Mpe_Reporting_FR2_r17 != nil, ie.TwoPHRMode_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Mpe_Reporting_FR2_r17 optional
			if ie.Mpe_Reporting_FR2_r17 != nil {
				tmp_Mpe_Reporting_FR2_r17 := utils.SetupRelease[*MPE_Config_FR2_r17]{
					Setup: ie.Mpe_Reporting_FR2_r17,
				}
				if err = tmp_Mpe_Reporting_FR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpe_Reporting_FR2_r17", err)
				}
			}
			// encode TwoPHRMode_r17 optional
			if ie.TwoPHRMode_r17 != nil {
				if err = ie.TwoPHRMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoPHRMode_r17", err)
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

func (ie *PHR_Config) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Phr_PeriodicTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Phr_PeriodicTimer", err)
	}
	if err = ie.Phr_ProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Phr_ProhibitTimer", err)
	}
	if err = ie.Phr_Tx_PowerFactorChange.Decode(r); err != nil {
		return utils.WrapError("Decode Phr_Tx_PowerFactorChange", err)
	}
	var tmp_bool_MultiplePHR bool
	if tmp_bool_MultiplePHR, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean MultiplePHR", err)
	}
	ie.MultiplePHR = tmp_bool_MultiplePHR
	var tmp_bool_Dummy bool
	if tmp_bool_Dummy, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Dummy", err)
	}
	ie.Dummy = tmp_bool_Dummy
	var tmp_bool_Phr_Type2OtherCell bool
	if tmp_bool_Phr_Type2OtherCell, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Phr_Type2OtherCell", err)
	}
	ie.Phr_Type2OtherCell = tmp_bool_Phr_Type2OtherCell
	if err = ie.Phr_ModeOtherCG.Decode(r); err != nil {
		return utils.WrapError("Decode Phr_ModeOtherCG", err)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Mpe_Reporting_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Mpe_Reporting_FR2_r16 optional
			if Mpe_Reporting_FR2_r16Present {
				tmp_Mpe_Reporting_FR2_r16 := utils.SetupRelease[*MPE_Config_FR2_r16]{}
				if err = tmp_Mpe_Reporting_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mpe_Reporting_FR2_r16", err)
				}
				ie.Mpe_Reporting_FR2_r16 = tmp_Mpe_Reporting_FR2_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Mpe_Reporting_FR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TwoPHRMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Mpe_Reporting_FR2_r17 optional
			if Mpe_Reporting_FR2_r17Present {
				tmp_Mpe_Reporting_FR2_r17 := utils.SetupRelease[*MPE_Config_FR2_r17]{}
				if err = tmp_Mpe_Reporting_FR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mpe_Reporting_FR2_r17", err)
				}
				ie.Mpe_Reporting_FR2_r17 = tmp_Mpe_Reporting_FR2_r17.Setup
			}
			// decode TwoPHRMode_r17 optional
			if TwoPHRMode_r17Present {
				ie.TwoPHRMode_r17 = new(PHR_Config_twoPHRMode_r17)
				if err = ie.TwoPHRMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoPHRMode_r17", err)
				}
			}
		}
	}
	return nil
}
