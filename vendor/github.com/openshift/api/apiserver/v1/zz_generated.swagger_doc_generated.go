package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_DeprecatedAPIRequest = map[string]string{
	"":       "DeprecatedAPIRequest tracts requests made to a deprecated API. The instance name should be of the form `resource.version.group`, matching the deprecated resource.",
	"spec":   "spec defines the characteristics of the resource.",
	"status": "status contains the observed state of the resource.",
}

func (DeprecatedAPIRequest) SwaggerDoc() map[string]string {
	return map_DeprecatedAPIRequest
}

var map_DeprecatedAPIRequestList = map[string]string{
	"": "DeprecatedAPIRequestList is a list of DeprecatedAPIRequest resources.",
}

func (DeprecatedAPIRequestList) SwaggerDoc() map[string]string {
	return map_DeprecatedAPIRequestList
}

var map_DeprecatedAPIRequestSpec = map[string]string{
	"removedRelease": "removedRelease is when the API will be removed.",
}

func (DeprecatedAPIRequestSpec) SwaggerDoc() map[string]string {
	return map_DeprecatedAPIRequestSpec
}

var map_DeprecatedAPIRequestStatus = map[string]string{
	"conditions":       "conditions contains details of the current status of this API Resource.",
	"requestsLastHour": "requestsLastHour contains request history for the current hour. This is porcelain to make the API easier to read by humans seeing if they addressed a problem. This field is reset on the hour.",
	"requestsLast24h":  "requestsLast24h contains request history for the last 24 hours, indexed by the hour, so 12:00AM-12:59 is in index 0, 6am-6:59am is index 6, etc. The index of the current hour is updated live and then duplicated into the requestsLastHour field.",
}

func (DeprecatedAPIRequestStatus) SwaggerDoc() map[string]string {
	return map_DeprecatedAPIRequestStatus
}

var map_NodeRequestLog = map[string]string{
	"":           "NodeRequestLog contains logs of requests to a certain node.",
	"nodeName":   "nodeName where the request are being handled.",
	"lastUpdate": "lastUpdate should *always* being within the hour this is for.  This is a time indicating the last moment the server is recording for, not the actual update time.",
	"users":      "users contains request details by top 10 users. Note that because in the case of an apiserver restart the list of top 10 users is determined on a best-effort basis, the list might be imprecise.",
}

func (NodeRequestLog) SwaggerDoc() map[string]string {
	return map_NodeRequestLog
}

var map_RequestCount = map[string]string{
	"":      "RequestCount counts requests by API request verb.",
	"verb":  "verb of API request (get, list, create, etc...)",
	"count": "count of requests for verb.",
}

func (RequestCount) SwaggerDoc() map[string]string {
	return map_RequestCount
}

var map_RequestLog = map[string]string{
	"":      "RequestLog logs request for various nodes.",
	"nodes": "nodes contains logs of requests per node.",
}

func (RequestLog) SwaggerDoc() map[string]string {
	return map_RequestLog
}

var map_RequestUser = map[string]string{
	"":         "RequestUser contains logs of a user's requests.",
	"username": "userName that made the request.",
	"count":    "count of requests.",
	"requests": "requests details by verb.",
}

func (RequestUser) SwaggerDoc() map[string]string {
	return map_RequestUser
}

// AUTO-GENERATED FUNCTIONS END HERE