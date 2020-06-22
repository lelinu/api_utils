package models

//EmailSettingsModel struct
type EmailSettingsModel struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}

//SendEmailModel model for send
type SendEmailModel struct {
	Settings     *EmailSettingsModel `json:"emailSettings"`
	From         string              `json:"from"`
	FromTag      string              `json:"fromTag"`
	ToList       string              `json:"toList"`
	BccList      string              `json:"bccList"`
	CcList       string              `json:"ccList"`
	Subject      string              `json:"subject"`
	HTMLContent  string              `json:"htmlContent"`
	PlainContent string              `json:"plainContent"`
}

//NewSendEmailModel will return a send email model
func NewSendEmailModel(host string, port int, username string, password, from, fromTag, toList, bccList,
	ccList, subject, htmlContent, plainContent string) *SendEmailModel {

	return &SendEmailModel{Settings: &EmailSettingsModel{Host: host, Port: port, Username: username, Password: password},
		From: from, FromTag: fromTag, ToList: toList, BccList: bccList,
		CcList: ccList, Subject: subject, HTMLContent: htmlContent, PlainContent: plainContent}
}
