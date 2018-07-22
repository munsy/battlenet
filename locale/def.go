package locale

var (
	AmericanEnglish     Locale = Locale{dialect: 12, language: 1} // "en_US" - American English
	BrazilianPortuguese Locale = Locale{dialect: 0, language: 6}  // "pt_BR" - Brazilian Portuguese
	BritishEnglish      Locale = Locale{dialect: 5, language: 1}  // "en_GB" - British English
	CastilianSpanish    Locale = Locale{dialect: 3, language: 2}  // "es_ES" - Castilian Spanish
	EuropeanPortuguese  Locale = Locale{dialect: 9, language: 6}  // "pt_PT" - European Portuguese
	MexicanSpanish      Locale = Locale{dialect: 8, language: 2}  // "es_MX" - Mexican Spanish
	SimplifiedChinese   Locale = Locale{dialect: 1, language: 8}  // "zh_CN" - Simplified Chinese
	StandardFrench      Locale = Locale{dialect: 4, language: 3}  // "fr_FR" - Standard French
	StandardGerman      Locale = Locale{dialect: 2, language: 0}  // "de_DE" - Standard German
	StandardItalian     Locale = Locale{dialect: 6, language: 4}  // "it_IT" - Standard Italian
	StandardKorean      Locale = Locale{dialect: 7, language: 5}  // "ko_KR" - Standard Korean
	StandardRussian     Locale = Locale{dialect: 10, language: 7} // "ru_RU" - Standard Russian
	TraditionalChinese  Locale = Locale{dialect: 11, language: 8} // "zh_TW" - Traditional Chinese
)
