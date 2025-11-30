package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PerRAAttemptInfo_r16 struct {
	ContentionDetected_r16   *bool                                          `optional`
	DlRSRPAboveThreshold_r16 *bool                                          `optional`
	FallbackToFourStepRA_r17 *PerRAAttemptInfo_r16_fallbackToFourStepRA_r17 `optional,ext-1`
}

func (ie *PerRAAttemptInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.FallbackToFourStepRA_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ContentionDetected_r16 != nil, ie.DlRSRPAboveThreshold_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ContentionDetected_r16 != nil {
		if err = w.WriteBoolean(*ie.ContentionDetected_r16); err != nil {
			return utils.WrapError("Encode ContentionDetected_r16", err)
		}
	}
	if ie.DlRSRPAboveThreshold_r16 != nil {
		if err = w.WriteBoolean(*ie.DlRSRPAboveThreshold_r16); err != nil {
			return utils.WrapError("Encode DlRSRPAboveThreshold_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.FallbackToFourStepRA_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PerRAAttemptInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.FallbackToFourStepRA_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FallbackToFourStepRA_r17 optional
			if ie.FallbackToFourStepRA_r17 != nil {
				if err = ie.FallbackToFourStepRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FallbackToFourStepRA_r17", err)
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

func (ie *PerRAAttemptInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ContentionDetected_r16Present bool
	if ContentionDetected_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DlRSRPAboveThreshold_r16Present bool
	if DlRSRPAboveThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ContentionDetected_r16Present {
		var tmp_bool_ContentionDetected_r16 bool
		if tmp_bool_ContentionDetected_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode ContentionDetected_r16", err)
		}
		ie.ContentionDetected_r16 = &tmp_bool_ContentionDetected_r16
	}
	if DlRSRPAboveThreshold_r16Present {
		var tmp_bool_DlRSRPAboveThreshold_r16 bool
		if tmp_bool_DlRSRPAboveThreshold_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode DlRSRPAboveThreshold_r16", err)
		}
		ie.DlRSRPAboveThreshold_r16 = &tmp_bool_DlRSRPAboveThreshold_r16
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FallbackToFourStepRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FallbackToFourStepRA_r17 optional
			if FallbackToFourStepRA_r17Present {
				ie.FallbackToFourStepRA_r17 = new(PerRAAttemptInfo_r16_fallbackToFourStepRA_r17)
				if err = ie.FallbackToFourStepRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FallbackToFourStepRA_r17", err)
				}
			}
		}
	}
	return nil
}
