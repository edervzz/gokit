package localization

import "fmt"

func LocalizerInstance(langu string, messageMap map[string]map[string]string) ILocalizer {
	return &localizer{
		messages: messageMap,
		langu:    langu,
	}
}

func (l *localizer) SetLanguage(langu string) {
	l.langu = langu
}

func (l localizer) Localize(id string, t ...string) string {
	mess := l.messages[id][l.langu]
	if mess != "" {
		mess = fmt.Sprintf(mess, t)
	} else {
		mess = id
	}

	return mess
}

type localizer struct {
	messages map[string]map[string]string
	langu    string
}
