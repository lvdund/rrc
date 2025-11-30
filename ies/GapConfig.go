package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type GapConfig struct {
	GapOffset                 int64                           `lb:0,ub:159,madatory`
	Mgl                       GapConfig_mgl                   `madatory`
	Mgrp                      GapConfig_mgrp                  `madatory`
	Mgta                      GapConfig_mgta                  `madatory`
	RefServCellIndicator      *GapConfig_refServCellIndicator `optional,ext-1`
	RefFR2ServCellAsyncCA_r16 *ServCellIndex                  `optional,ext-2`
	Mgl_r16                   *GapConfig_mgl_r16              `optional,ext-2`
}

func (ie *GapConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.RefServCellIndicator != nil || ie.RefFR2ServCellAsyncCA_r16 != nil || ie.Mgl_r16 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.GapOffset, &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger GapOffset", err)
	}
	if err = ie.Mgl.Encode(w); err != nil {
		return utils.WrapError("Encode Mgl", err)
	}
	if err = ie.Mgrp.Encode(w); err != nil {
		return utils.WrapError("Encode Mgrp", err)
	}
	if err = ie.Mgta.Encode(w); err != nil {
		return utils.WrapError("Encode Mgta", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.RefServCellIndicator != nil, ie.RefFR2ServCellAsyncCA_r16 != nil || ie.Mgl_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap GapConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.RefServCellIndicator != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RefServCellIndicator optional
			if ie.RefServCellIndicator != nil {
				if err = ie.RefServCellIndicator.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RefServCellIndicator", err)
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
			optionals_ext_2 := []bool{ie.RefFR2ServCellAsyncCA_r16 != nil, ie.Mgl_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RefFR2ServCellAsyncCA_r16 optional
			if ie.RefFR2ServCellAsyncCA_r16 != nil {
				if err = ie.RefFR2ServCellAsyncCA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RefFR2ServCellAsyncCA_r16", err)
				}
			}
			// encode Mgl_r16 optional
			if ie.Mgl_r16 != nil {
				if err = ie.Mgl_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mgl_r16", err)
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

func (ie *GapConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_GapOffset int64
	if tmp_int_GapOffset, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger GapOffset", err)
	}
	ie.GapOffset = tmp_int_GapOffset
	if err = ie.Mgl.Decode(r); err != nil {
		return utils.WrapError("Decode Mgl", err)
	}
	if err = ie.Mgrp.Decode(r); err != nil {
		return utils.WrapError("Decode Mgrp", err)
	}
	if err = ie.Mgta.Decode(r); err != nil {
		return utils.WrapError("Decode Mgta", err)
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

			RefServCellIndicatorPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RefServCellIndicator optional
			if RefServCellIndicatorPresent {
				ie.RefServCellIndicator = new(GapConfig_refServCellIndicator)
				if err = ie.RefServCellIndicator.Decode(extReader); err != nil {
					return utils.WrapError("Decode RefServCellIndicator", err)
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

			RefFR2ServCellAsyncCA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mgl_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RefFR2ServCellAsyncCA_r16 optional
			if RefFR2ServCellAsyncCA_r16Present {
				ie.RefFR2ServCellAsyncCA_r16 = new(ServCellIndex)
				if err = ie.RefFR2ServCellAsyncCA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RefFR2ServCellAsyncCA_r16", err)
				}
			}
			// decode Mgl_r16 optional
			if Mgl_r16Present {
				ie.Mgl_r16 = new(GapConfig_mgl_r16)
				if err = ie.Mgl_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mgl_r16", err)
				}
			}
		}
	}
	return nil
}
