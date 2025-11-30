package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfigMN struct {
	MeasuredFrequenciesMN []NR_FreqInfo                    `lb:1,ub:maxMeasFreqsMN,optional`
	MeasGapConfig         *GapConfig                       `optional,setuprelease`
	GapPurpose            *MeasConfigMN_gapPurpose         `optional`
	MeasGapConfigFR2      *GapConfig                       `optional,ext-1,setuprelease`
	InterFreqNoGap_r16    *MeasConfigMN_interFreqNoGap_r16 `optional,ext-2`
}

func (ie *MeasConfigMN) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.MeasGapConfigFR2 != nil || ie.InterFreqNoGap_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.MeasuredFrequenciesMN) > 0, ie.MeasGapConfig != nil, ie.GapPurpose != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.MeasuredFrequenciesMN) > 0 {
		tmp_MeasuredFrequenciesMN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, aper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
		for _, i := range ie.MeasuredFrequenciesMN {
			tmp_MeasuredFrequenciesMN.Value = append(tmp_MeasuredFrequenciesMN.Value, &i)
		}
		if err = tmp_MeasuredFrequenciesMN.Encode(w); err != nil {
			return utils.WrapError("Encode MeasuredFrequenciesMN", err)
		}
	}
	if ie.MeasGapConfig != nil {
		tmp_MeasGapConfig := utils.SetupRelease[*GapConfig]{
			Setup: ie.MeasGapConfig,
		}
		if err = tmp_MeasGapConfig.Encode(w); err != nil {
			return utils.WrapError("Encode MeasGapConfig", err)
		}
	}
	if ie.GapPurpose != nil {
		if err = ie.GapPurpose.Encode(w); err != nil {
			return utils.WrapError("Encode GapPurpose", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.MeasGapConfigFR2 != nil, ie.InterFreqNoGap_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasConfigMN", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MeasGapConfigFR2 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MeasGapConfigFR2 optional
			if ie.MeasGapConfigFR2 != nil {
				tmp_MeasGapConfigFR2 := utils.SetupRelease[*GapConfig]{
					Setup: ie.MeasGapConfigFR2,
				}
				if err = tmp_MeasGapConfigFR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasGapConfigFR2", err)
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
			optionals_ext_2 := []bool{ie.InterFreqNoGap_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InterFreqNoGap_r16 optional
			if ie.InterFreqNoGap_r16 != nil {
				if err = ie.InterFreqNoGap_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFreqNoGap_r16", err)
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

func (ie *MeasConfigMN) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasuredFrequenciesMNPresent bool
	if MeasuredFrequenciesMNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasGapConfigPresent bool
	if MeasGapConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GapPurposePresent bool
	if GapPurposePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasuredFrequenciesMNPresent {
		tmp_MeasuredFrequenciesMN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, aper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
		fn_MeasuredFrequenciesMN := func() *NR_FreqInfo {
			return new(NR_FreqInfo)
		}
		if err = tmp_MeasuredFrequenciesMN.Decode(r, fn_MeasuredFrequenciesMN); err != nil {
			return utils.WrapError("Decode MeasuredFrequenciesMN", err)
		}
		ie.MeasuredFrequenciesMN = []NR_FreqInfo{}
		for _, i := range tmp_MeasuredFrequenciesMN.Value {
			ie.MeasuredFrequenciesMN = append(ie.MeasuredFrequenciesMN, *i)
		}
	}
	if MeasGapConfigPresent {
		tmp_MeasGapConfig := utils.SetupRelease[*GapConfig]{}
		if err = tmp_MeasGapConfig.Decode(r); err != nil {
			return utils.WrapError("Decode MeasGapConfig", err)
		}
		ie.MeasGapConfig = tmp_MeasGapConfig.Setup
	}
	if GapPurposePresent {
		ie.GapPurpose = new(MeasConfigMN_gapPurpose)
		if err = ie.GapPurpose.Decode(r); err != nil {
			return utils.WrapError("Decode GapPurpose", err)
		}
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

			MeasGapConfigFR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MeasGapConfigFR2 optional
			if MeasGapConfigFR2Present {
				tmp_MeasGapConfigFR2 := utils.SetupRelease[*GapConfig]{}
				if err = tmp_MeasGapConfigFR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasGapConfigFR2", err)
				}
				ie.MeasGapConfigFR2 = tmp_MeasGapConfigFR2.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InterFreqNoGap_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InterFreqNoGap_r16 optional
			if InterFreqNoGap_r16Present {
				ie.InterFreqNoGap_r16 = new(MeasConfigMN_interFreqNoGap_r16)
				if err = ie.InterFreqNoGap_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFreqNoGap_r16", err)
				}
			}
		}
	}
	return nil
}
