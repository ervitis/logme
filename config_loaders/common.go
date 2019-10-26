package config_loaders

import "github.com/sirupsen/logrus"

func CommonFormatter(typef string) logrus.Formatter {
	switch typef {
	case "json":
		return &logrus.JSONFormatter{
			PrettyPrint: false,
			FieldMap:    logrus.FieldMap{logrus.FieldKeyTime: "time"},
		}
	default:
		return &logrus.TextFormatter{
			DisableColors:          true,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			FieldMap:               logrus.FieldMap{logrus.FieldKeyTime: "time"},
		}
	}
}
