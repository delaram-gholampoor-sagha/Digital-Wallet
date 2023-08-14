package middleware

import (
	"sort"
	"strconv"
	"strings"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/translation"
	"github.com/labstack/echo/v4"
)

type acceptedLanguage struct {
	name    string
	quality float64
}

func Locale() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			acceptedLanguages := make([]acceptedLanguage, 0)

			header := c.Request().Header.Get("Accept-Language")

			languages := strings.Split(header, ",")
			for _, language := range languages {
				trimmedLanguage := strings.TrimSpace(language)
				weightedLanguage := strings.Split(trimmedLanguage, ";")
				if len(weightedLanguage) == 1 {
					acceptedLanguages = append(acceptedLanguages, acceptedLanguage{name: weightedLanguage[0], quality: 1})
				} else {
					languageQuality := strings.Split(weightedLanguage[1], "=")
					quality, err := strconv.ParseFloat(languageQuality[1], 64)
					if err != nil {
						return derror.NewBadRequestError(message.InvalidRequest)
					}
					acceptedLanguages = append(acceptedLanguages, acceptedLanguage{name: weightedLanguage[0], quality: quality})
				}
			}

			sort.Slice(acceptedLanguages, func(i, j int) bool {
				return acceptedLanguages[i].quality > acceptedLanguages[j].quality
			})

			for _, language := range acceptedLanguages {
				switch language.name {
				case "fa", "fa-IR":
					c.Set(translation.Locale, translation.Farsi)
					return next(c)
				case "en", "en-US":
					c.Set(translation.Locale, translation.English)
					return next(c)
				}
			}

			c.Set(translation.Locale, translation.Farsi)
			return next(c)
		}
	}
}
