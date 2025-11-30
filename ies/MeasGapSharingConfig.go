package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasGapSharingConfig struct {
	GapSharingFR2 *MeasGapSharingScheme `optional,setuprelease`
	GapSharingFR1 *MeasGapSharingScheme `optional,ext-1,setuprelease`
	GapSharingUE  *MeasGapSharingScheme `optional,ext-1,setuprelease`
}

func (ie *MeasGapSharingConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.GapSharingFR1 != nil || ie.GapSharingUE != nil
	preambleBits := []bool{hasExtensions, ie.GapSharingFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.GapSharingFR2 != nil {
		tmp_GapSharingFR2 := utils.SetupRelease[*MeasGapSharingScheme]{
			Setup: ie.GapSharingFR2,
		}
		if err = tmp_GapSharingFR2.Encode(w); err != nil {
			return utils.WrapError("Encode GapSharingFR2", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.GapSharingFR1 != nil || ie.GapSharingUE != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasGapSharingConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.GapSharingFR1 != nil, ie.GapSharingUE != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode GapSharingFR1 optional
			if ie.GapSharingFR1 != nil {
				tmp_GapSharingFR1 := utils.SetupRelease[*MeasGapSharingScheme]{
					Setup: ie.GapSharingFR1,
				}
				if err = tmp_GapSharingFR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GapSharingFR1", err)
				}
			}
			// encode GapSharingUE optional
			if ie.GapSharingUE != nil {
				tmp_GapSharingUE := utils.SetupRelease[*MeasGapSharingScheme]{
					Setup: ie.GapSharingUE,
				}
				if err = tmp_GapSharingUE.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GapSharingUE", err)
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

func (ie *MeasGapSharingConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var GapSharingFR2Present bool
	if GapSharingFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if GapSharingFR2Present {
		tmp_GapSharingFR2 := utils.SetupRelease[*MeasGapSharingScheme]{}
		if err = tmp_GapSharingFR2.Decode(r); err != nil {
			return utils.WrapError("Decode GapSharingFR2", err)
		}
		ie.GapSharingFR2 = tmp_GapSharingFR2.Setup
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

			GapSharingFR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GapSharingUEPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode GapSharingFR1 optional
			if GapSharingFR1Present {
				tmp_GapSharingFR1 := utils.SetupRelease[*MeasGapSharingScheme]{}
				if err = tmp_GapSharingFR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode GapSharingFR1", err)
				}
				ie.GapSharingFR1 = tmp_GapSharingFR1.Setup
			}
			// decode GapSharingUE optional
			if GapSharingUEPresent {
				tmp_GapSharingUE := utils.SetupRelease[*MeasGapSharingScheme]{}
				if err = tmp_GapSharingUE.Decode(extReader); err != nil {
					return utils.WrapError("Decode GapSharingUE", err)
				}
				ie.GapSharingUE = tmp_GapSharingUE.Setup
			}
		}
	}
	return nil
}
