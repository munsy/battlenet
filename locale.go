package battlenet

import (
	"github.com/munsy/battlenet/locale"
)

// Locale holds all of the possible Battle.net locale definitions.
var Locale = struct {
	AmericanEnglish     locale.Locale
	BrazilianPortuguese locale.Locale
	BritishEnglish      locale.Locale
	CastilianSpanish    locale.Locale
	EuropeanPortuguese  locale.Locale
	MexicanSpanish      locale.Locale
	SimplifiedChinese   locale.Locale
	StandardFrench      locale.Locale
	StandardGerman      locale.Locale
	StandardItalian     locale.Locale
	StandardKorean      locale.Locale
	StandardRussian     locale.Locale
	TraditionalChinese  locale.Locale
}{
	locale.AmericanEnglish,
	locale.BrazilianPortuguese,
	locale.BritishEnglish,
	locale.CastilianSpanish,
	locale.EuropeanPortuguese,
	locale.MexicanSpanish,
	locale.SimplifiedChinese,
	locale.StandardFrench,
	locale.StandardGerman,
	locale.StandardItalian,
	locale.StandardKorean,
	locale.StandardRussian,
	locale.TraditionalChinese,
}
