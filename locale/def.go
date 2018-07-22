package locale

var (
	// AmericanEnglish represents "en_US" - (i.e., American English)
	AmericanEnglish Locale = Locale{dialect: 12, language: 1}

	// BrazilianPortuguese represents "pt_BR" - (i.e., Brazilian Portuguese)
	BrazilianPortuguese Locale = Locale{dialect: 0, language: 6}

	// BritishEnglish represents "en_GB" - (i.e., British English)
	BritishEnglish Locale = Locale{dialect: 5, language: 1}

	// CastilianSpanish represents "es_ES" - (i.e., Castilian Spanish)
	CastilianSpanish Locale = Locale{dialect: 3, language: 2}

	// EuropeanPortuguese represents "pt_PT" - (i.e., European Portuguese)
	EuropeanPortuguese Locale = Locale{dialect: 9, language: 6}

	// MexicanSpanish represents "es_MX" - (i.e., Mexican Spanish)
	MexicanSpanish Locale = Locale{dialect: 8, language: 2}

	// SimplifiedChinese represents "zh_CN" - (i.e., Simplified Chinese)
	SimplifiedChinese Locale = Locale{dialect: 1, language: 8}

	// StandardFrench represents "fr_FR" - (i.e., Standard French)
	StandardFrench Locale = Locale{dialect: 4, language: 3}

	// StandardGerman represents "de_DE" - (i.e., Standard German)
	StandardGerman Locale = Locale{dialect: 2, language: 0}

	// StandardItalian represents "it_IT" - (i.e., Standard Italian)
	StandardItalian Locale = Locale{dialect: 6, language: 4}

	// StandardKorean represents "ko_KR" - (i.e., Standard Korean)
	StandardKorean Locale = Locale{dialect: 7, language: 5}

	// StandardRussian represents "ru_RU" - (i.e., Standard Russian)
	StandardRussian Locale = Locale{dialect: 10, language: 7}

	// TraditionalChinese represents "zh_TW" - (i.e., Traditional Chinese)
	TraditionalChinese Locale = Locale{dialect: 11, language: 8}
)
