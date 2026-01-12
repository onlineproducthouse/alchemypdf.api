package configmock

import (
	"strings"

	"1702tech.oph.api/lib/1702tech.oph.api.infrastructure/config"
)

type MockConfig struct{}

func NewMockConfig() config.IConfig {
	return MockConfig{}
}

func (c MockConfig) ProjectName() string {
	return "example"
}

func (c MockConfig) ProjectShortName() string {
	return "e.g."
}

func (c MockConfig) Protocol() string {
	return "API_PROTOCOL"
}

func (c MockConfig) Host() string {
	return "API_HOST"
}

func (c MockConfig) Port() string {
	return "API_PORT"
}

func (c MockConfig) EnvName() string {
	return ""
}

func (c MockConfig) ReqHeaderApiKey() string {
	return "x-api-key"
}

func (c MockConfig) ReqHeaderRequestID() string {
	return "x-request-id"
}

func (c MockConfig) ReqHeaderSessionExternalID() string {
	return "x-session-external-id"
}

func (c MockConfig) ReqHeaderOrganisationExternalID() string {
	return "x-organisation-external-id"
}

func (c MockConfig) ReqHeaderRepresentedOrganisationExternalID() string {
	return "x-rep-organisation-external-id"
}

func (c MockConfig) ReqHeaderProjectExternalID() string {
	return "x-project-external-id"
}

func (c MockConfig) ReqHeaderWebApp() string {
	return "x-web-app"
}

func (c MockConfig) UserExternalIDSessionKey() string {
	return "userExternalID"
}

func (c MockConfig) DbConnectionString() string {
	return "db://0.0.0.0:0000"
}

func (c MockConfig) RedisHost() string {
	return "0.0.0.0"
}

func (c MockConfig) RedisPort() string {
	return "0000"
}

func (c MockConfig) RedisPassword() string {
	return "$tr0ngP@$$W0rD"
}

func (c MockConfig) OtpLength() int {
	return 6
}

func (c MockConfig) NoReplyEmailAddr() string {
	return "no-reply@example.org"
}

func (c MockConfig) APIKeys() []string {
	return strings.Split("API_KEY", " ")
}

func (c MockConfig) RunSwagger() bool {
	return true
}

func (c MockConfig) SendGridAPIKey() string {
	return "SENDGRID_API_KEY"
}

func (c MockConfig) SendGridAddress() string {
	return "SENDGRID_SENDER_ADDRESS"
}

func (c MockConfig) SendGridCity() string {
	return "SENDGRID_SENDER_CITY"
}

func (c MockConfig) SendGridState() string {
	return "SENDGRID_SENDER_STATE"
}

func (c MockConfig) SendGridZip() string {
	return "SENDGRID_SENDER_ZIP"
}

func (c MockConfig) SendGridNewAccTemplID() string {
	return "SENDGRID_SENDER_NEW_ACCOUNT_TEMPL_ID"
}

func (c MockConfig) SendGridRecoverAccTemplID() string {
	return "SENDGRID_SENDER_RECOVER_ACCOUNT_TEMPL_ID"
}

func (c MockConfig) SendGridNewEmailAddrTemplID() string {
	return "SENDGRID_SENDER_NEW_EMAIL_ADDR_TEMPL_ID"
}

func (c MockConfig) SendGridLeadLinkInviteTemplID() string {
	return "SENDGRID_SENDER_LEAD_LINK_INVITE_TEMPL_ID"
}

func (c MockConfig) SendGridLeadSignUpClosedTemplID() string {
	return "SENDGRID_LEAD_SIGN_UP_CLOSED_TEMPL_ID"
}

func (c MockConfig) SendGridLeadNewClientTemplID() string {
	return "SENDGRID_LEAD_NEW_CLIENT_TEMPL_ID"
}

func (c MockConfig) PortalAppURL() string {
	return "http://portal.example.org"
}

func (c MockConfig) ConsoleAppURL() string {
	return "http://console.example.org"
}

func (c MockConfig) RegistrationAppURL() string {
	return "http://registration.example.org"
}

func (c MockConfig) OTPTimeToLiveInMinutes() int {
	return 5
}

func (c MockConfig) RequestIDKey() string {
	return "x-request-id"
}

func (c MockConfig) SendGridAgreementVersionPublishedTemplID() string {
	return "SENDGRID_SENDER_AGREEMENT_VERSION_PUBLISHED_TEMPL_ID"
}

func (c MockConfig) SendGridOrganisationMemberInviteTemplID() string {
	return "SENDGRID_SENDER_ORGANISATION_MEMBER_INVITE_TEMPL_ID"
}

func (c MockConfig) PortalUserEmailAdress() string {
	return "info@example.org"
}

func (c MockConfig) FileServiceS3Bucket() string {
	return "FILE_SERVICE_S3_BUCKET"
}

func (c MockConfig) PaystackPublicKey() string {
	return "PAYSTACK_PUBLIC_KEY"
}

func (c MockConfig) PaystackSecretKey() string {
	return "PAYSTACK_SECRET_KEY"
}

func (c MockConfig) SendGridBillingQuoteGeneratedTemplID() string {
	return "SENDGRID_BILLING_QUOTE_GENERATED_TEMPL_ID"
}

func (c MockConfig) SendGridBillingInvoiceGeneratedTemplID() string {
	return "SENDGRID_BILLING_INVOICE_GENERATED_TEMPL_ID"
}

func (c MockConfig) SendGridBillingPaymentConfirmedTemplID() string {
	return "SENDGRID_BILLING_PAYMENT_CONFIRMED_TEMPL_ID"
}

func (c MockConfig) SendGridNewChatMessageTemplID() string {
	return "SENDGRID_NEW_CHAT_MESSAGE_TEMPL_ID"
}

func (c MockConfig) AlcheMyPDFAPIURL() string {
	return "ALCHEMYPDF_API_URL"
}

func (c MockConfig) AlcheMyPDFAPIKey() string {
	return "ALCHEMYPDF_API_KEY"
}
