package logger

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

type hook struct {
	CallBack func(msg string, data map[string]interface{}) error
}

func (h hook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.TraceLevel}
}
func (h hook) Fire(entry *logrus.Entry) error {
	return h.CallBack(entry.Message, entry.Data)
}

var loginstance = logrus.New()

func Initlogger() {
	loginstance.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //"2006-01-02T15:04:05Z07:00"
		PrettyPrint:     true,
	})
}

func AddHook(callBack func(msg string, data map[string]interface{}) error) {
	loginstance.AddHook(hook{
		CallBack: callBack,
	})
}

func Info(msg string, detail ...logrus.Fields) {
	f := GenerateLogFields(detail...)
	loginstance.WithFields(f).Info(msg)
}

func Trace(msg string, detail ...logrus.Fields) {
	f := GenerateLogFields(detail...)
	loginstance.WithFields(f).Trace(msg)
}

func Warn(msg string, detail ...logrus.Fields) {
	loginstance.WithFields(GenerateLogFields(detail...)).Warn(msg)
}

func Error(msg string, detail ...logrus.Fields) {
	loginstance.WithFields(GenerateLogFields(detail...)).Error(msg)
}

func Fatal(msg string, detail ...logrus.Fields) {
	loginstance.WithFields(GenerateLogFields(detail...)).Fatal(msg)
}

func StringField(key string, value interface{}) logrus.Fields {
	return logrus.Fields{
		key: value,
	}
}

func ErrField(err error) logrus.Fields {
	return logrus.Fields{
		"error": err,
	}
}

func Field(key string, value interface{}) logrus.Fields {
	b, err := jsoniter.Marshal(value)
	if err != nil {
		return logrus.Fields{
			key: err,
		}
	}
	f := logrus.Fields{}
	jsoniter.Unmarshal(b, &f)
	return logrus.Fields{key: f}
}

func GenerateLogFields(detail ...logrus.Fields) logrus.Fields {

	field := logrus.Fields{}

	for _, v := range detail {
		maps.Copy(field, v)
	}
	return field
}
