package models

type Image struct {
	PandoraObj
	CategoryID uint64
	SubjectID  uint64
	Base64     string
}
