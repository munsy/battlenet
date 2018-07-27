package locale

import "testing"

func TestLocaleConst(t *testing.T) {
	if AmericanEnglish.String() != "en_US" {
		t.Fatal("mismatch: 'AmericanEnglish' and 'en_US'")
	}

	if BrazilianPortuguese.String() != "pt_BR" {
		t.Fatal("mismatch: 'BrazilianPortuguese' and 'pt_BR'")
	}

	if BritishEnglish.String() != "en_GB" {
		t.Fatal("mismatch: 'BritishEnglish' and 'en_GB'")
	}

	if CastilianSpanish.String() != "es_ES" {
		t.Fatal("mismatch: 'CastilianSpanish' and 'es_ES'")
	}

	if EuropeanPortuguese.String() != "pt_PT" {
		t.Fatal("mismatch: 'EuropeanPortuguese' and 'pt_PT'")
	}

	if MexicanSpanish.String() != "es_MX" {
		t.Fatal("mismatch: 'MexicanSpanish' and 'es_MX'")
	}

	if SimplifiedChinese.String() != "zh_CN" {
		t.Fatal("mismatch: 'SimplifiedChinese' and 'zh_CN'")
	}

	if StandardFrench.String() != "fr_FR" {
		t.Fatal("mismatch: 'StandardFrench' and 'fr_FR'")
	}

	if StandardGerman.String() != "de_DE" {
		t.Fatal("mismatch: 'StandardGerman' and 'de_DE'")
	}

	if StandardItalian.String() != "it_IT" {
		t.Fatal("mismatch: 'StandardItalian' and 'it_IT'")
	}

	if StandardKorean.String() != "ko_KR" {
		t.Fatal("mismatch: 'StandardKorean' and 'ko_KR'")
	}

	if StandardRussian.String() != "ru_RU" {
		t.Fatal("mismatch: 'StandardRussian' and 'ru_RU'")
	}

	if TraditionalChinese.String() != "zh_TW" {
		t.Fatal("mismatch: 'TraditionalChinese' and 'zh_TW'")
	}
}
