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

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A ConversationsApiController binds http requests to an api service and writes the service results to the http response
type ConversationsApiController struct {
	service ConversationsApiServicer
}

// NewConversationsApiController creates a default api controller
func NewConversationsApiController(s ConversationsApiServicer) Router {
	return &ConversationsApiController{ service: s }
}

// Routes returns all of the api route for the ConversationsApiController
func (c *ConversationsApiController) Routes() Routes {
	return Routes{ 
		{
			"ConversationsCreateConversation",
			strings.ToUpper("Post"),
			"/v3/conversations",
			c.ConversationsCreateConversation,
		},
		{
			"ConversationsDeleteActivity",
			strings.ToUpper("Delete"),
			"/v3/conversations/{conversationId}/activities/{activityId}",
			c.ConversationsDeleteActivity,
		},
		{
			"ConversationsDeleteConversationMember",
			strings.ToUpper("Delete"),
			"/v3/conversations/{conversationId}/members/{memberId}",
			c.ConversationsDeleteConversationMember,
		},
		{
			"ConversationsGetActivityMembers",
			strings.ToUpper("Get"),
			"/v3/conversations/{conversationId}/activities/{activityId}/members",
			c.ConversationsGetActivityMembers,
		},
		{
			"ConversationsGetConversationMembers",
			strings.ToUpper("Get"),
			"/v3/conversations/{conversationId}/members",
			c.ConversationsGetConversationMembers,
		},
		{
			"ConversationsGetConversationPagedMembers",
			strings.ToUpper("Get"),
			"/v3/conversations/{conversationId}/pagedmembers",
			c.ConversationsGetConversationPagedMembers,
		},
		{
			"ConversationsGetConversations",
			strings.ToUpper("Get"),
			"/v3/conversations",
			c.ConversationsGetConversations,
		},
		{
			"ConversationsReplyToActivity",
			strings.ToUpper("Post"),
			"/v3/conversations/{conversationId}/activities/{activityId}",
			c.ConversationsReplyToActivity,
		},
		{
			"ConversationsSendConversationHistory",
			strings.ToUpper("Post"),
			"/v3/conversations/{conversationId}/activities/history",
			c.ConversationsSendConversationHistory,
		},
		{
			"ConversationsSendToConversation",
			strings.ToUpper("Post"),
			"/v3/conversations/{conversationId}/activities",
			c.ConversationsSendToConversation,
		},
		{
			"ConversationsUpdateActivity",
			strings.ToUpper("Put"),
			"/v3/conversations/{conversationId}/activities/{activityId}",
			c.ConversationsUpdateActivity,
		},
		{
			"ConversationsUploadAttachment",
			strings.ToUpper("Post"),
			"/v3/conversations/{conversationId}/attachments",
			c.ConversationsUploadAttachment,
		},
	}
}

// ConversationsCreateConversation - CreateConversation
func (c *ConversationsApiController) ConversationsCreateConversation(w http.ResponseWriter, r *http.Request) { 
	parameters := &ConversationParameters{}
	if err := json.NewDecoder(r.Body).Decode(&parameters); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConversationsCreateConversation(*parameters)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsDeleteActivity - DeleteActivity
func (c *ConversationsApiController) ConversationsDeleteActivity(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	activityId := params["activityId"]
	result, err := c.service.ConversationsDeleteActivity(conversationId, activityId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsDeleteConversationMember - DeleteConversationMember
func (c *ConversationsApiController) ConversationsDeleteConversationMember(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	memberId := params["memberId"]
	result, err := c.service.ConversationsDeleteConversationMember(conversationId, memberId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsGetActivityMembers - GetActivityMembers
func (c *ConversationsApiController) ConversationsGetActivityMembers(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	activityId := params["activityId"]
	result, err := c.service.ConversationsGetActivityMembers(conversationId, activityId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsGetConversationMembers - GetConversationMembers
func (c *ConversationsApiController) ConversationsGetConversationMembers(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	result, err := c.service.ConversationsGetConversationMembers(conversationId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsGetConversationPagedMembers - GetConversationPagedMembers
func (c *ConversationsApiController) ConversationsGetConversationPagedMembers(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	query := r.URL.Query()
	conversationId := params["conversationId"]
	pageSize := query.Get("pageSize")
	continuationToken := query.Get("continuationToken")
	result, err := c.service.ConversationsGetConversationPagedMembers(conversationId, pageSize, continuationToken)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsGetConversations - GetConversations
func (c *ConversationsApiController) ConversationsGetConversations(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	continuationToken := query.Get("continuationToken")
	result, err := c.service.ConversationsGetConversations(continuationToken)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsReplyToActivity - ReplyToActivity
func (c *ConversationsApiController) ConversationsReplyToActivity(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	activityId := params["activityId"]
	activity := &Activity{}
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConversationsReplyToActivity(conversationId, activityId, *activity)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsSendConversationHistory - SendConversationHistory
func (c *ConversationsApiController) ConversationsSendConversationHistory(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	history := &Transcript{}
	if err := json.NewDecoder(r.Body).Decode(&history); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConversationsSendConversationHistory(conversationId, *history)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsSendToConversation - SendToConversation
func (c *ConversationsApiController) ConversationsSendToConversation(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	activity := &Activity{}
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConversationsSendToConversation(conversationId, *activity)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsUpdateActivity - UpdateActivity
func (c *ConversationsApiController) ConversationsUpdateActivity(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	activityId := params["activityId"]
	activity := &Activity{}
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConversationsUpdateActivity(conversationId, activityId, *activity)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ConversationsUploadAttachment - UploadAttachment
func (c *ConversationsApiController) ConversationsUploadAttachment(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	conversationId := params["conversationId"]
	attachmentUpload := &AttachmentData{}
	if err := json.NewDecoder(r.Body).Decode(&attachmentUpload); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConversationsUploadAttachment(conversationId, *attachmentUpload)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
