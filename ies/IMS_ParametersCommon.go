package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IMS_ParametersCommon struct {
	VoiceOverEUTRA_5GC             *IMS_ParametersCommon_voiceOverEUTRA_5GC             `optional`
	VoiceOverSCG_BearerEUTRA_5GC   *IMS_ParametersCommon_voiceOverSCG_BearerEUTRA_5GC   `optional,ext-1`
	VoiceFallbackIndicationEPS_r16 *IMS_ParametersCommon_voiceFallbackIndicationEPS_r16 `optional,ext-2`
}

func (ie *IMS_ParametersCommon) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.VoiceOverSCG_BearerEUTRA_5GC != nil || ie.VoiceFallbackIndicationEPS_r16 != nil
	preambleBits := []bool{hasExtensions, ie.VoiceOverEUTRA_5GC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.VoiceOverEUTRA_5GC != nil {
		if err = ie.VoiceOverEUTRA_5GC.Encode(w); err != nil {
			return utils.WrapError("Encode VoiceOverEUTRA_5GC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.VoiceOverSCG_BearerEUTRA_5GC != nil, ie.VoiceFallbackIndicationEPS_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap IMS_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.VoiceOverSCG_BearerEUTRA_5GC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode VoiceOverSCG_BearerEUTRA_5GC optional
			if ie.VoiceOverSCG_BearerEUTRA_5GC != nil {
				if err = ie.VoiceOverSCG_BearerEUTRA_5GC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode VoiceOverSCG_BearerEUTRA_5GC", err)
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
			optionals_ext_2 := []bool{ie.VoiceFallbackIndicationEPS_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode VoiceFallbackIndicationEPS_r16 optional
			if ie.VoiceFallbackIndicationEPS_r16 != nil {
				if err = ie.VoiceFallbackIndicationEPS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode VoiceFallbackIndicationEPS_r16", err)
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

func (ie *IMS_ParametersCommon) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var VoiceOverEUTRA_5GCPresent bool
	if VoiceOverEUTRA_5GCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if VoiceOverEUTRA_5GCPresent {
		ie.VoiceOverEUTRA_5GC = new(IMS_ParametersCommon_voiceOverEUTRA_5GC)
		if err = ie.VoiceOverEUTRA_5GC.Decode(r); err != nil {
			return utils.WrapError("Decode VoiceOverEUTRA_5GC", err)
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

			VoiceOverSCG_BearerEUTRA_5GCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode VoiceOverSCG_BearerEUTRA_5GC optional
			if VoiceOverSCG_BearerEUTRA_5GCPresent {
				ie.VoiceOverSCG_BearerEUTRA_5GC = new(IMS_ParametersCommon_voiceOverSCG_BearerEUTRA_5GC)
				if err = ie.VoiceOverSCG_BearerEUTRA_5GC.Decode(extReader); err != nil {
					return utils.WrapError("Decode VoiceOverSCG_BearerEUTRA_5GC", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			VoiceFallbackIndicationEPS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode VoiceFallbackIndicationEPS_r16 optional
			if VoiceFallbackIndicationEPS_r16Present {
				ie.VoiceFallbackIndicationEPS_r16 = new(IMS_ParametersCommon_voiceFallbackIndicationEPS_r16)
				if err = ie.VoiceFallbackIndicationEPS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode VoiceFallbackIndicationEPS_r16", err)
				}
			}
		}
	}
	return nil
}
