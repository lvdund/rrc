package ies

// SL-RadioBearerConfig-r16 ::= SEQUENCE
// Extensible
type SlRadiobearerconfigR16 struct {
	SlrbUuConfigindexR16 SlrbUuConfigindexR16
	SlSdapConfigR16      *SlSdapConfigR16
	SlPdcpConfigR16      *SlPdcpConfigR16
	SlTransrangeR16      *SlRadiobearerconfigR16SlTransrangeR16
}
