package ies

// AS-Config ::= SEQUENCE
// Extensible
type AsConfig struct {
	Sourcemeasconfig                  Measconfig
	Sourceradioresourceconfig         Radioresourceconfigdedicated
	Sourcesecurityalgorithmconfig     Securityalgorithmconfig
	SourceueIdentity                  CRnti
	Sourcemasterinformationblock      Masterinformationblock
	Sourcesysteminformationblocktype1 Systeminformationblocktype1
	Sourcesysteminformationblocktype2 Systeminformationblocktype2
	Antennainfocommon                 Antennainfocommon
	SourcedlCarrierfreq               ArfcnValueeutra
}
