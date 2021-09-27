package secret

import (
	"bytes"
	"regexp"
	"text/template"
)

const DefaultTemplateName = "secret-template"
const TokenPattern = "{{(.*?)}}"

func HasToken(str string) bool {
	regex, _ := regexp.Compile(TokenPattern)
	return regex.MatchString(str)
}

func ReplaceToken(placeholder string, secretValue string, str string) (string, error) {
	t := template.Must(template.New(DefaultTemplateName).Parse(str))
	secretMap := map[string]string{placeholder: secretValue}
	var strBuffer bytes.Buffer
	err := t.Execute(&strBuffer, secretMap)
	if err != nil {
		return "", err
	}
	return strBuffer.String(), nil
}

func ReplaceTokens(secretMap map[string]string, str string) (string, error) {
	t := template.Must(template.New(DefaultTemplateName).Parse(str))
	var strBuffer bytes.Buffer
	err := t.Execute(&strBuffer, secretMap)
	if err != nil {
		return "", err
	}
	return strBuffer.String(), nil
}
