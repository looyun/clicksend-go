/*
 * ClickSend v3 API
 *
 *  This is an official SDK for [ClickSend](https://clicksend.com)  Below you will find a current list of the available methods for clicksend.  *NOTE: You will need to create a free account to use the API. You can register [here](https://dashboard.clicksend.com/#/signup/step1/)..* 
 *
 * API version: 3.1
 * Contact: support@clicksend.com
 * Generated by: Swagger Codegen (https://github.com/clicksend-api/clicksend-codegen.git)
 */

package clicksend

// Model for Inbound SMS Rules
type InboundSmsRule struct {
	// Dedicated Number. Can be '*' to apply to all numbers.
	DedicatedNumber string `json:"dedicated_number"`
	// Rule Name.
	RuleName string `json:"rule_name"`
	// Message Search Type: 0=Any message, 1=starts with, 2=contains, 3=does not contain.
	MessageSearchType float32 `json:"message_search_type"`
	// Message search term.
	MessageSearchTerm string `json:"message_search_term"`
	// Action to be taken (AUTO_REPLY, EMAIL_USER, EMAIL_FIXED, URL, SMS, POLL, GROUP_SMS, MOVE_CONTACT, CREATE_CONTACT, CREATE_CONTACT_PLUS_EMAIL, CREATE_CONTACT_PLUS_NAME_EMAIL CREATE_CONTACT_PLUS_NAME, SMPP, NONE).
	Action string `json:"action"`
	// Action address.
	ActionAddress string `json:"action_address"`
	// Enabled: Disabled=0 or Enabled=1.
	Enabled float32 `json:"enabled"`
}
