package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasGapConfig struct {
	GapFR2                               *GapConfig                            `optional,setuprelease`
	GapFR1                               *GapConfig                            `optional,ext-1,setuprelease`
	GapUE                                *GapConfig                            `optional,ext-1,setuprelease`
	GapToAddModList_r17                  []GapConfig_r17                       `lb:1,ub:maxNrofGapId_r17,optional,ext-2`
	GapToReleaseList_r17                 []MeasGapId_r17                       `lb:1,ub:maxNrofGapId_r17,optional,ext-2`
	PosMeasGapPreConfigToAddModList_r17  *PosMeasGapPreConfigToAddModList_r17  `optional,ext-2`
	PosMeasGapPreConfigToReleaseList_r17 *PosMeasGapPreConfigToReleaseList_r17 `optional,ext-2`
}

func (ie *MeasGapConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.GapFR1 != nil || ie.GapUE != nil || len(ie.GapToAddModList_r17) > 0 || len(ie.GapToReleaseList_r17) > 0 || ie.PosMeasGapPreConfigToAddModList_r17 != nil || ie.PosMeasGapPreConfigToReleaseList_r17 != nil
	preambleBits := []bool{hasExtensions, ie.GapFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.GapFR2 != nil {
		tmp_GapFR2 := utils.SetupRelease[*GapConfig]{
			Setup: ie.GapFR2,
		}
		if err = tmp_GapFR2.Encode(w); err != nil {
			return utils.WrapError("Encode GapFR2", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.GapFR1 != nil || ie.GapUE != nil, len(ie.GapToAddModList_r17) > 0 || len(ie.GapToReleaseList_r17) > 0 || ie.PosMeasGapPreConfigToAddModList_r17 != nil || ie.PosMeasGapPreConfigToReleaseList_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasGapConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.GapFR1 != nil, ie.GapUE != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode GapFR1 optional
			if ie.GapFR1 != nil {
				tmp_GapFR1 := utils.SetupRelease[*GapConfig]{
					Setup: ie.GapFR1,
				}
				if err = tmp_GapFR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GapFR1", err)
				}
			}
			// encode GapUE optional
			if ie.GapUE != nil {
				tmp_GapUE := utils.SetupRelease[*GapConfig]{
					Setup: ie.GapUE,
				}
				if err = tmp_GapUE.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GapUE", err)
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
			optionals_ext_2 := []bool{len(ie.GapToAddModList_r17) > 0, len(ie.GapToReleaseList_r17) > 0, ie.PosMeasGapPreConfigToAddModList_r17 != nil, ie.PosMeasGapPreConfigToReleaseList_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode GapToAddModList_r17 optional
			if len(ie.GapToAddModList_r17) > 0 {
				tmp_GapToAddModList_r17 := utils.NewSequence[*GapConfig_r17]([]*GapConfig_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				for _, i := range ie.GapToAddModList_r17 {
					tmp_GapToAddModList_r17.Value = append(tmp_GapToAddModList_r17.Value, &i)
				}
				if err = tmp_GapToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GapToAddModList_r17", err)
				}
			}
			// encode GapToReleaseList_r17 optional
			if len(ie.GapToReleaseList_r17) > 0 {
				tmp_GapToReleaseList_r17 := utils.NewSequence[*MeasGapId_r17]([]*MeasGapId_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				for _, i := range ie.GapToReleaseList_r17 {
					tmp_GapToReleaseList_r17.Value = append(tmp_GapToReleaseList_r17.Value, &i)
				}
				if err = tmp_GapToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GapToReleaseList_r17", err)
				}
			}
			// encode PosMeasGapPreConfigToAddModList_r17 optional
			if ie.PosMeasGapPreConfigToAddModList_r17 != nil {
				if err = ie.PosMeasGapPreConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PosMeasGapPreConfigToAddModList_r17", err)
				}
			}
			// encode PosMeasGapPreConfigToReleaseList_r17 optional
			if ie.PosMeasGapPreConfigToReleaseList_r17 != nil {
				if err = ie.PosMeasGapPreConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PosMeasGapPreConfigToReleaseList_r17", err)
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

func (ie *MeasGapConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var GapFR2Present bool
	if GapFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if GapFR2Present {
		tmp_GapFR2 := utils.SetupRelease[*GapConfig]{}
		if err = tmp_GapFR2.Decode(r); err != nil {
			return utils.WrapError("Decode GapFR2", err)
		}
		ie.GapFR2 = tmp_GapFR2.Setup
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

			GapFR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GapUEPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode GapFR1 optional
			if GapFR1Present {
				tmp_GapFR1 := utils.SetupRelease[*GapConfig]{}
				if err = tmp_GapFR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode GapFR1", err)
				}
				ie.GapFR1 = tmp_GapFR1.Setup
			}
			// decode GapUE optional
			if GapUEPresent {
				tmp_GapUE := utils.SetupRelease[*GapConfig]{}
				if err = tmp_GapUE.Decode(extReader); err != nil {
					return utils.WrapError("Decode GapUE", err)
				}
				ie.GapUE = tmp_GapUE.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			GapToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GapToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PosMeasGapPreConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PosMeasGapPreConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode GapToAddModList_r17 optional
			if GapToAddModList_r17Present {
				tmp_GapToAddModList_r17 := utils.NewSequence[*GapConfig_r17]([]*GapConfig_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				fn_GapToAddModList_r17 := func() *GapConfig_r17 {
					return new(GapConfig_r17)
				}
				if err = tmp_GapToAddModList_r17.Decode(extReader, fn_GapToAddModList_r17); err != nil {
					return utils.WrapError("Decode GapToAddModList_r17", err)
				}
				ie.GapToAddModList_r17 = []GapConfig_r17{}
				for _, i := range tmp_GapToAddModList_r17.Value {
					ie.GapToAddModList_r17 = append(ie.GapToAddModList_r17, *i)
				}
			}
			// decode GapToReleaseList_r17 optional
			if GapToReleaseList_r17Present {
				tmp_GapToReleaseList_r17 := utils.NewSequence[*MeasGapId_r17]([]*MeasGapId_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				fn_GapToReleaseList_r17 := func() *MeasGapId_r17 {
					return new(MeasGapId_r17)
				}
				if err = tmp_GapToReleaseList_r17.Decode(extReader, fn_GapToReleaseList_r17); err != nil {
					return utils.WrapError("Decode GapToReleaseList_r17", err)
				}
				ie.GapToReleaseList_r17 = []MeasGapId_r17{}
				for _, i := range tmp_GapToReleaseList_r17.Value {
					ie.GapToReleaseList_r17 = append(ie.GapToReleaseList_r17, *i)
				}
			}
			// decode PosMeasGapPreConfigToAddModList_r17 optional
			if PosMeasGapPreConfigToAddModList_r17Present {
				ie.PosMeasGapPreConfigToAddModList_r17 = new(PosMeasGapPreConfigToAddModList_r17)
				if err = ie.PosMeasGapPreConfigToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PosMeasGapPreConfigToAddModList_r17", err)
				}
			}
			// decode PosMeasGapPreConfigToReleaseList_r17 optional
			if PosMeasGapPreConfigToReleaseList_r17Present {
				ie.PosMeasGapPreConfigToReleaseList_r17 = new(PosMeasGapPreConfigToReleaseList_r17)
				if err = ie.PosMeasGapPreConfigToReleaseList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PosMeasGapPreConfigToReleaseList_r17", err)
				}
			}
		}
	}
	return nil
}
