/*
 * Microsoft Bot Connector API - v3.0
 *
 * The Bot Connector REST API allows your bot to send and receive messages to channels configured in the  [Bot Framework Developer Portal](https://dev.botframework.com). The Connector service uses industry-standard REST  and JSON over HTTPS.    Client libraries for this REST API are available. See below for a list.    Many bots will use both the Bot Connector REST API and the associated [Bot State REST API](/en-us/restapi/state). The  Bot State REST API allows a bot to store and retrieve state associated with users and conversations.    Authentication for both the Bot Connector and Bot State REST APIs is accomplished with JWT Bearer tokens, and is  described in detail in the [Connector Authentication](/en-us/restapi/authentication) document.    # Client Libraries for the Bot Connector REST API    * [Bot Builder for C#](/en-us/csharp/builder/sdkreference/)  * [Bot Builder for Node.js](/en-us/node/builder/overview/)  * Generate your own from the [Connector API Swagger file](https://raw.githubusercontent.com/Microsoft/BotBuilder/master/CSharp/Library/Microsoft.Bot.Connector.Shared/Swagger/ConnectorAPI.json)    © 2016 Microsoft
 *
 * API version: v3
 * Contact: botframework@microsoft.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package botbuilder

// ReceiptItem - An item on a receipt card
type ReceiptItem struct {

	// Title of the Card
	Title string `json:"title,omitempty"`

	// Subtitle appears just below Title field, differs from Title in font styling only
	Subtitle string `json:"subtitle,omitempty"`

	// Text field appears just below subtitle, differs from Subtitle in font styling only
	Text string `json:"text,omitempty"`

	Image CardImage `json:"image,omitempty"`

	// Amount with currency
	Price string `json:"price,omitempty"`

	// Number of items of given kind
	Quantity string `json:"quantity,omitempty"`

	Tap CardAction `json:"tap,omitempty"`
}
