package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/regions"
	"github.com/munsy/gobattlenet/settings"
)

type Config struct {
	Token  string
	Key    string
	Region string
	Locale string
}

func (c Config) Settings(cmdType string) settings.BNetSettings {
	var s settings.BNetSettings

	s.Client = &http.Client{Timeout: (10 * time.Second)}

	switch cmdType {
	case "key":
		s.Key = c.Key
		break
	case "token":
		s.Key = c.Token
		break
	default:
		fmt.Println("Bad setting input.")
		os.Exit(1)
	}

	switch c.Locale {
	case "en_US":
		s.Locale = locale.AmericanEnglish
		break
	case "pt_BR":
		s.Locale = locale.BrazilianPortugues
		break
	case "en_GB":
		s.Locale = locale.BritishEnglish
		break
	case "es_ES":
		s.Locale = locale.CastilianSpanish
		break
	case "pt_PT":
		s.Locale = locale.EuropeanPortuguese
		break
	case "es_MX":
		s.Locale = locale.MexicanSpanish
		break
	case "zh_CN":
		s.Locale = locale.SimplifiedChinese
		break
	case "fr_FR":
		s.Locale = locale.StandardFrench
		break
	case "de_DE":
		s.Locale = locale.StandardGerman
		break
	case "it_IT":
		s.Locale = locale.StandardItalian
		break
	case "ko_KR":
		s.Locale = locale.StandardKorean
		break
	case "ru_RU":
		s.Locale = locale.StandardRussian
		break
	case "zh_TW":
		s.Locale = locale.TraditionalChinese
		break
	default:
		fmt.Println("\tInvalid type. Possible types are:")
		fmt.Println("\t\ten_US")
		fmt.Println("\t\tpt_BR")
		fmt.Println("\t\ten_GB")
		fmt.Println("\t\tes_ES")
		fmt.Println("\t\tpt_PT")
		fmt.Println("\t\tes_MX")
		fmt.Println("\t\tzh_CN")
		fmt.Println("\t\tfr_FR")
		fmt.Println("\t\tde_DE")
		fmt.Println("\t\tit_IT")
		fmt.Println("\t\tko_KR")
		fmt.Println("\t\tru_RU")
		fmt.Println("\t\tzh_TW")
		os.Exit(1)
	}

	switch c.Region {
	case "us":
		s.Region = regions.US
		break
	case "eu":
		s.Region = regions.EU
		break
	case "kr":
		s.Region = regions.KR
		break
	case "tw":
		s.Region = regions.TW
		break
	case "sea":
		s.Region = regions.SEA
		break
	case "cn":
		s.Region = regions.CN
		break
	default:
		fmt.Println("\tInvalid type. Possible types are:")
		fmt.Println("\t\tus")
		fmt.Println("\t\teu")
		fmt.Println("\t\tkr")
		fmt.Println("\t\ttw")
		fmt.Println("\t\tsea")
		fmt.Println("\t\tcn")
		os.Exit(1)
	}

	return s
}
