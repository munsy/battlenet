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
		fmt.Println("Invalid type. Possible types are:")
		fmt.Println("\t- en_US (American English)")
		fmt.Println("\t- pt_BR (Brazilian Portugues)")
		fmt.Println("\t- en_GB (British English)")
		fmt.Println("\t- es_ES (Castilian Spanish)")
		fmt.Println("\t- pt_PT (European Portuguese)")
		fmt.Println("\t- es_MX (Mexican Spanish)")
		fmt.Println("\t- zh_CN (Simplified Chinese)")
		fmt.Println("\t- fr_FR (Standard French)")
		fmt.Println("\t- de_DE (Standard German)")
		fmt.Println("\t- it_IT (Standard Italian)")
		fmt.Println("\t- ko_KR (Standard Korean)")
		fmt.Println("\t- ru_RU (Standard Russian)")
		fmt.Println("\t- zh_TW (Traditional Chinese)")
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
		fmt.Println("Invalid type. Possible types are:")
		fmt.Println("\t- us (United States)")
		fmt.Println("\t- eu (Europe)")
		fmt.Println("\t- kr (Korea)")
		fmt.Println("\t- tw (Taiwan)")
		fmt.Println("\t- sea (Oceanic)")
		fmt.Println("\t- cn (China)")
		os.Exit(1)
	}

	return s
}
