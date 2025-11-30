package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RadioBearerConfig struct {
	Srb_ToAddModList      *SRB_ToAddModList                     `optional`
	Srb3_ToRelease        *RadioBearerConfig_srb3_ToRelease     `optional`
	Drb_ToAddModList      *DRB_ToAddModList                     `optional`
	Drb_ToReleaseList     *DRB_ToReleaseList                    `optional`
	SecurityConfig        *SecurityConfig                       `optional`
	Mrb_ToAddModList_r17  *MRB_ToAddModList_r17                 `optional,ext-1`
	Mrb_ToReleaseList_r17 *MRB_ToReleaseList_r17                `optional,ext-1`
	Srb4_ToAddMod_r17     *SRB_ToAddMod                         `optional,ext-1`
	Srb4_ToRelease_r17    *RadioBearerConfig_srb4_ToRelease_r17 `optional,ext-1`
}

func (ie *RadioBearerConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Mrb_ToAddModList_r17 != nil || ie.Mrb_ToReleaseList_r17 != nil || ie.Srb4_ToAddMod_r17 != nil || ie.Srb4_ToRelease_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Srb_ToAddModList != nil, ie.Srb3_ToRelease != nil, ie.Drb_ToAddModList != nil, ie.Drb_ToReleaseList != nil, ie.SecurityConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srb_ToAddModList != nil {
		if err = ie.Srb_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Srb_ToAddModList", err)
		}
	}
	if ie.Srb3_ToRelease != nil {
		if err = ie.Srb3_ToRelease.Encode(w); err != nil {
			return utils.WrapError("Encode Srb3_ToRelease", err)
		}
	}
	if ie.Drb_ToAddModList != nil {
		if err = ie.Drb_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Drb_ToAddModList", err)
		}
	}
	if ie.Drb_ToReleaseList != nil {
		if err = ie.Drb_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Drb_ToReleaseList", err)
		}
	}
	if ie.SecurityConfig != nil {
		if err = ie.SecurityConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SecurityConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Mrb_ToAddModList_r17 != nil || ie.Mrb_ToReleaseList_r17 != nil || ie.Srb4_ToAddMod_r17 != nil || ie.Srb4_ToRelease_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RadioBearerConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Mrb_ToAddModList_r17 != nil, ie.Mrb_ToReleaseList_r17 != nil, ie.Srb4_ToAddMod_r17 != nil, ie.Srb4_ToRelease_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Mrb_ToAddModList_r17 optional
			if ie.Mrb_ToAddModList_r17 != nil {
				if err = ie.Mrb_ToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mrb_ToAddModList_r17", err)
				}
			}
			// encode Mrb_ToReleaseList_r17 optional
			if ie.Mrb_ToReleaseList_r17 != nil {
				if err = ie.Mrb_ToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mrb_ToReleaseList_r17", err)
				}
			}
			// encode Srb4_ToAddMod_r17 optional
			if ie.Srb4_ToAddMod_r17 != nil {
				if err = ie.Srb4_ToAddMod_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srb4_ToAddMod_r17", err)
				}
			}
			// encode Srb4_ToRelease_r17 optional
			if ie.Srb4_ToRelease_r17 != nil {
				if err = ie.Srb4_ToRelease_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srb4_ToRelease_r17", err)
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

func (ie *RadioBearerConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Srb_ToAddModListPresent bool
	if Srb_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srb3_ToReleasePresent bool
	if Srb3_ToReleasePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Drb_ToAddModListPresent bool
	if Drb_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Drb_ToReleaseListPresent bool
	if Drb_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SecurityConfigPresent bool
	if SecurityConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Srb_ToAddModListPresent {
		ie.Srb_ToAddModList = new(SRB_ToAddModList)
		if err = ie.Srb_ToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode Srb_ToAddModList", err)
		}
	}
	if Srb3_ToReleasePresent {
		ie.Srb3_ToRelease = new(RadioBearerConfig_srb3_ToRelease)
		if err = ie.Srb3_ToRelease.Decode(r); err != nil {
			return utils.WrapError("Decode Srb3_ToRelease", err)
		}
	}
	if Drb_ToAddModListPresent {
		ie.Drb_ToAddModList = new(DRB_ToAddModList)
		if err = ie.Drb_ToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode Drb_ToAddModList", err)
		}
	}
	if Drb_ToReleaseListPresent {
		ie.Drb_ToReleaseList = new(DRB_ToReleaseList)
		if err = ie.Drb_ToReleaseList.Decode(r); err != nil {
			return utils.WrapError("Decode Drb_ToReleaseList", err)
		}
	}
	if SecurityConfigPresent {
		ie.SecurityConfig = new(SecurityConfig)
		if err = ie.SecurityConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityConfig", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Mrb_ToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mrb_ToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srb4_ToAddMod_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srb4_ToRelease_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Mrb_ToAddModList_r17 optional
			if Mrb_ToAddModList_r17Present {
				ie.Mrb_ToAddModList_r17 = new(MRB_ToAddModList_r17)
				if err = ie.Mrb_ToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mrb_ToAddModList_r17", err)
				}
			}
			// decode Mrb_ToReleaseList_r17 optional
			if Mrb_ToReleaseList_r17Present {
				ie.Mrb_ToReleaseList_r17 = new(MRB_ToReleaseList_r17)
				if err = ie.Mrb_ToReleaseList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mrb_ToReleaseList_r17", err)
				}
			}
			// decode Srb4_ToAddMod_r17 optional
			if Srb4_ToAddMod_r17Present {
				ie.Srb4_ToAddMod_r17 = new(SRB_ToAddMod)
				if err = ie.Srb4_ToAddMod_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srb4_ToAddMod_r17", err)
				}
			}
			// decode Srb4_ToRelease_r17 optional
			if Srb4_ToRelease_r17Present {
				ie.Srb4_ToRelease_r17 = new(RadioBearerConfig_srb4_ToRelease_r17)
				if err = ie.Srb4_ToRelease_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srb4_ToRelease_r17", err)
				}
			}
		}
	}
	return nil
}
