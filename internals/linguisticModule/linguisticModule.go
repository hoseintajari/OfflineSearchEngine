package linguisticModule

type LinguisticModule struct {
	converters []Converter
}

func NewLinguisticModule(converters ...Converter) LinguisticModule {
	return LinguisticModule{converters: converters}
}

func (lm LinguisticModule) Convert(s []string) []string {
	var result []string
	for _, converter := range lm.converters {
		result = converter.Convert(s)
	}
	return result
}
