package localization

type ILocalizer interface {
	Localize(id string, t ...string) string
	SetLanguage(langu string)
}
