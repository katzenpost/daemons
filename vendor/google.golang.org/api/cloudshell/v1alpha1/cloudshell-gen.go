// Package cloudshell provides access to the Cloud Shell API.
//
// See https://cloud.google.com/shell/docs/
//
// Usage example:
//
//   import "google.golang.org/api/cloudshell/v1alpha1"
//   ...
//   cloudshellService, err := cloudshell.New(oauthHttpClient)
package cloudshell // import "google.golang.org/api/cloudshell/v1alpha1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "cloudshell:v1alpha1"
const apiName = "cloudshell"
const apiVersion = "v1alpha1"
const basePath = "https://cloudshell.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	rs.Environments = NewUsersEnvironmentsService(s)
	return rs
}

type UsersService struct {
	s *Service

	Environments *UsersEnvironmentsService
}

func NewUsersEnvironmentsService(s *Service) *UsersEnvironmentsService {
	rs := &UsersEnvironmentsService{s: s}
	rs.PublicKeys = NewUsersEnvironmentsPublicKeysService(s)
	return rs
}

type UsersEnvironmentsService struct {
	s *Service

	PublicKeys *UsersEnvironmentsPublicKeysService
}

func NewUsersEnvironmentsPublicKeysService(s *Service) *UsersEnvironmentsPublicKeysService {
	rs := &UsersEnvironmentsPublicKeysService{s: s}
	return rs
}

type UsersEnvironmentsPublicKeysService struct {
	s *Service
}

// AuthorizeEnvironmentRequest: Request message for
// AuthorizeEnvironment.
type AuthorizeEnvironmentRequest struct {
	// AccessToken: The OAuth access token that should be sent to the
	// environment.
	AccessToken string `json:"accessToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccessToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AuthorizeEnvironmentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AuthorizeEnvironmentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreatePublicKeyRequest: Request message for CreatePublicKey.
type CreatePublicKeyRequest struct {
	// Key: Key that should be added to the environment.
	Key *PublicKey `json:"key,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreatePublicKeyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreatePublicKeyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Environment: A Cloud Shell environment, which is defined as the
// combination of a Docker
// image specifying what is installed on the environment and a home
// directory
// containing the user's data that will remain across sessions. Each
// user has a
// single environment with the ID "default".
type Environment struct {
	// DockerImage: Required. Full path to the Docker image used to run this
	// environment, e.g.
	// "gcr.io/dev-con/cloud-devshell:latest".
	DockerImage string `json:"dockerImage,omitempty"`

	// Id: Output only. The environment's identifier, which is always
	// "default".
	Id string `json:"id,omitempty"`

	// Name: Output only. Full name of this resource, in the
	// format
	// `users/{owner_email}/environments/{environment_id}`. `{owner_email}`
	// is the
	// email address of the user to whom this environment belongs,
	// and
	// `{environment_id}` is the identifier of this environment. For
	// example,
	// `users/someone@example.com/environments/default`.
	Name string `json:"name,omitempty"`

	// PublicKeys: Output only. Public keys associated with the environment.
	// Clients can
	// connect to this environment via SSH only if they possess a private
	// key
	// corresponding to at least one of these public keys. Keys can be added
	// to or
	// removed from the environment using the CreatePublicKey and
	// DeletePublicKey
	// methods.
	PublicKeys []*PublicKey `json:"publicKeys,omitempty"`

	// SshHost: Output only. Host to which clients can connect to initiate
	// SSH sessions
	// with the environment.
	SshHost string `json:"sshHost,omitempty"`

	// SshPort: Output only. Port to which clients can connect to initiate
	// SSH sessions
	// with the environment.
	SshPort int64 `json:"sshPort,omitempty"`

	// SshUsername: Output only. Username that clients should use when
	// initiating SSH sessions
	// with the environment.
	SshUsername string `json:"sshUsername,omitempty"`

	// State: Output only. Current execution state of this environment.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The environment's states is unknown.
	//   "DISABLED" - The environment is not running and can't be connected
	// to. Starting the
	// environment will transition it to the STARTING state.
	//   "STARTING" - The environment is being started but is not yet ready
	// to accept
	// connections.
	//   "RUNNING" - The environment is running and ready to accept
	// connections. It will
	// automatically transition back to DISABLED after a period of
	// inactivity or
	// if another environment is started.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DockerImage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DockerImage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Environment) MarshalJSON() ([]byte, error) {
	type NoMethod Environment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PublicKey: A public SSH key, corresponding to a private SSH key held
// by the client.
type PublicKey struct {
	// Format: Required. Format of this key's content.
	//
	// Possible values:
	//   "FORMAT_UNSPECIFIED" - Unknown format. Do not use.
	//   "SSH_DSS" - `ssh-dss` key format (see RFC4253).
	//   "SSH_RSA" - `ssh-rsa` key format (see RFC4253).
	//   "ECDSA_SHA2_NISTP256" - `ecdsa-sha2-nistp256` key format (see
	// RFC5656).
	//   "ECDSA_SHA2_NISTP384" - `ecdsa-sha2-nistp384` key format (see
	// RFC5656).
	//   "ECDSA_SHA2_NISTP521" - `ecdsa-sha2-nistp521` key format (see
	// RFC5656).
	Format string `json:"format,omitempty"`

	// Key: Required. Content of this key.
	Key string `json:"key,omitempty"`

	// Name: Output only. Full name of this resource, in the
	// format
	// `users/{owner_email}/environments/{environment_id}/publicKeys/{
	// key_id}`.
	// `{owner_email}` is the email address of the user to whom the key
	// belongs.
	// `{environment_id}` is the identifier of the environment to which the
	// key
	// grants access. `{key_id}` is the unique identifier of the key. For
	// example,
	// `users/someone@example.com/environments/default/publicKeys/my
	// Key`.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Format") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Format") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PublicKey) MarshalJSON() ([]byte, error) {
	type NoMethod PublicKey
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartEnvironmentMetadata: Message included in the metadata field of
// operations returned from
// StartEnvironment.
type StartEnvironmentMetadata struct {
	// State: Current state of the environment being started.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The environment's start state is unknown.
	//   "STARTING" - The environment is in the process of being started,
	// but no additional
	// details are available.
	//   "UNARCHIVING_DISK" - Startup is waiting for the user's disk to be
	// unarchived. This can happen
	// when the user returns to Cloud Shell after not having used it for
	// a
	// while, and suggests that startup will take longer than normal.
	//   "FINISHED" - Startup is complete and the user should be able to
	// establish an SSH
	// connection to their environment.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "State") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "State") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartEnvironmentMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod StartEnvironmentMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartEnvironmentRequest: Request message for StartEnvironment.
type StartEnvironmentRequest struct {
	// AccessToken: The initial access token passed to the environment. If
	// this is present and
	// valid, the environment will be pre-authenticated with gcloud so that
	// the
	// user can run gcloud commands in Cloud Shell without having to log in.
	// This
	// code can be updated later by calling AuthorizeEnvironment.
	AccessToken string `json:"accessToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccessToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartEnvironmentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod StartEnvironmentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartEnvironmentResponse: Message included in the response field of
// operations returned from
// StartEnvironment once the
// operation is complete.
type StartEnvironmentResponse struct {
	// Environment: Environment that was started.
	Environment *Environment `json:"environment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Environment") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Environment") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartEnvironmentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod StartEnvironmentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudshell.users.environments.authorize":

type UsersEnvironmentsAuthorizeCall struct {
	s                           *Service
	name                        string
	authorizeenvironmentrequest *AuthorizeEnvironmentRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// Authorize: Sends an access token to a running environment on behalf
// of a user. When
// this completes, the environment will be authorized to run gcloud
// commands
// without requiring the user to manually authenticate.
func (r *UsersEnvironmentsService) Authorize(name string, authorizeenvironmentrequest *AuthorizeEnvironmentRequest) *UsersEnvironmentsAuthorizeCall {
	c := &UsersEnvironmentsAuthorizeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.authorizeenvironmentrequest = authorizeenvironmentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersEnvironmentsAuthorizeCall) Fields(s ...googleapi.Field) *UsersEnvironmentsAuthorizeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersEnvironmentsAuthorizeCall) Context(ctx context.Context) *UsersEnvironmentsAuthorizeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersEnvironmentsAuthorizeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersEnvironmentsAuthorizeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.authorizeenvironmentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}:authorize")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudshell.users.environments.authorize" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *UsersEnvironmentsAuthorizeCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sends an access token to a running environment on behalf of a user. When\nthis completes, the environment will be authorized to run gcloud commands\nwithout requiring the user to manually authenticate.",
	//   "flatPath": "v1alpha1/users/{usersId}/environments/{environmentsId}:authorize",
	//   "httpMethod": "POST",
	//   "id": "cloudshell.users.environments.authorize",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the resource that should receive the token, for example\n`users/me/environments/default` or\n`users/someone@example.com/environments/default`.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}:authorize",
	//   "request": {
	//     "$ref": "AuthorizeEnvironmentRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudshell.users.environments.get":

type UsersEnvironmentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets an environment. Returns NOT_FOUND if the environment does
// not exist.
func (r *UsersEnvironmentsService) Get(name string) *UsersEnvironmentsGetCall {
	c := &UsersEnvironmentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersEnvironmentsGetCall) Fields(s ...googleapi.Field) *UsersEnvironmentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UsersEnvironmentsGetCall) IfNoneMatch(entityTag string) *UsersEnvironmentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersEnvironmentsGetCall) Context(ctx context.Context) *UsersEnvironmentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersEnvironmentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersEnvironmentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudshell.users.environments.get" call.
// Exactly one of *Environment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Environment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UsersEnvironmentsGetCall) Do(opts ...googleapi.CallOption) (*Environment, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Environment{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets an environment. Returns NOT_FOUND if the environment does not exist.",
	//   "flatPath": "v1alpha1/users/{usersId}/environments/{environmentsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudshell.users.environments.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the requested resource, for example `users/me/environments/default`\nor `users/someone@example.com/environments/default`.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Environment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudshell.users.environments.patch":

type UsersEnvironmentsPatchCall struct {
	s           *Service
	name        string
	environment *Environment
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Updates an existing environment.
func (r *UsersEnvironmentsService) Patch(name string, environment *Environment) *UsersEnvironmentsPatchCall {
	c := &UsersEnvironmentsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.environment = environment
	return c
}

// UpdateMask sets the optional parameter "updateMask": Mask specifying
// which fields in the environment should be updated.
func (c *UsersEnvironmentsPatchCall) UpdateMask(updateMask string) *UsersEnvironmentsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersEnvironmentsPatchCall) Fields(s ...googleapi.Field) *UsersEnvironmentsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersEnvironmentsPatchCall) Context(ctx context.Context) *UsersEnvironmentsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersEnvironmentsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersEnvironmentsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.environment)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudshell.users.environments.patch" call.
// Exactly one of *Environment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Environment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UsersEnvironmentsPatchCall) Do(opts ...googleapi.CallOption) (*Environment, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Environment{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing environment.",
	//   "flatPath": "v1alpha1/users/{usersId}/environments/{environmentsId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudshell.users.environments.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the resource to be updated, for example\n`users/me/environments/default` or\n`users/someone@example.com/environments/default`.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Mask specifying which fields in the environment should be updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "request": {
	//     "$ref": "Environment"
	//   },
	//   "response": {
	//     "$ref": "Environment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudshell.users.environments.start":

type UsersEnvironmentsStartCall struct {
	s                       *Service
	name                    string
	startenvironmentrequest *StartEnvironmentRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// Start: Starts an existing environment, allowing clients to connect to
// it. The
// returned operation will contain an instance of
// StartEnvironmentMetadata in
// its metadata field. Users can wait for the environment to start by
// polling
// this operation via GetOperation. Once the environment has finished
// starting
// and is ready to accept connections, the operation will contain
// a
// StartEnvironmentResponse in its response field.
func (r *UsersEnvironmentsService) Start(name string, startenvironmentrequest *StartEnvironmentRequest) *UsersEnvironmentsStartCall {
	c := &UsersEnvironmentsStartCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.startenvironmentrequest = startenvironmentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersEnvironmentsStartCall) Fields(s ...googleapi.Field) *UsersEnvironmentsStartCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersEnvironmentsStartCall) Context(ctx context.Context) *UsersEnvironmentsStartCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersEnvironmentsStartCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersEnvironmentsStartCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.startenvironmentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}:start")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudshell.users.environments.start" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UsersEnvironmentsStartCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts an existing environment, allowing clients to connect to it. The\nreturned operation will contain an instance of StartEnvironmentMetadata in\nits metadata field. Users can wait for the environment to start by polling\nthis operation via GetOperation. Once the environment has finished starting\nand is ready to accept connections, the operation will contain a\nStartEnvironmentResponse in its response field.",
	//   "flatPath": "v1alpha1/users/{usersId}/environments/{environmentsId}:start",
	//   "httpMethod": "POST",
	//   "id": "cloudshell.users.environments.start",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the resource that should be started, for example\n`users/me/environments/default` or\n`users/someone@example.com/environments/default`.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}:start",
	//   "request": {
	//     "$ref": "StartEnvironmentRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudshell.users.environments.publicKeys.create":

type UsersEnvironmentsPublicKeysCreateCall struct {
	s                      *Service
	parent                 string
	createpublickeyrequest *CreatePublicKeyRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Create: Adds a public SSH key to an environment, allowing clients
// with the
// corresponding private key to connect to that environment via SSH. If
// a key
// with the same format and content already exists, this will return
// the
// existing key.
func (r *UsersEnvironmentsPublicKeysService) Create(parent string, createpublickeyrequest *CreatePublicKeyRequest) *UsersEnvironmentsPublicKeysCreateCall {
	c := &UsersEnvironmentsPublicKeysCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createpublickeyrequest = createpublickeyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersEnvironmentsPublicKeysCreateCall) Fields(s ...googleapi.Field) *UsersEnvironmentsPublicKeysCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersEnvironmentsPublicKeysCreateCall) Context(ctx context.Context) *UsersEnvironmentsPublicKeysCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersEnvironmentsPublicKeysCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersEnvironmentsPublicKeysCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createpublickeyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+parent}/publicKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudshell.users.environments.publicKeys.create" call.
// Exactly one of *PublicKey or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *PublicKey.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UsersEnvironmentsPublicKeysCreateCall) Do(opts ...googleapi.CallOption) (*PublicKey, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PublicKey{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a public SSH key to an environment, allowing clients with the\ncorresponding private key to connect to that environment via SSH. If a key\nwith the same format and content already exists, this will return the\nexisting key.",
	//   "flatPath": "v1alpha1/users/{usersId}/environments/{environmentsId}/publicKeys",
	//   "httpMethod": "POST",
	//   "id": "cloudshell.users.environments.publicKeys.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Parent resource name, e.g. `users/me/environments/default`.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+parent}/publicKeys",
	//   "request": {
	//     "$ref": "CreatePublicKeyRequest"
	//   },
	//   "response": {
	//     "$ref": "PublicKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudshell.users.environments.publicKeys.delete":

type UsersEnvironmentsPublicKeysDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Removes a public SSH key from an environment. Clients will no
// longer be
// able to connect to the environment using the corresponding private
// key.
func (r *UsersEnvironmentsPublicKeysService) Delete(name string) *UsersEnvironmentsPublicKeysDeleteCall {
	c := &UsersEnvironmentsPublicKeysDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersEnvironmentsPublicKeysDeleteCall) Fields(s ...googleapi.Field) *UsersEnvironmentsPublicKeysDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersEnvironmentsPublicKeysDeleteCall) Context(ctx context.Context) *UsersEnvironmentsPublicKeysDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersEnvironmentsPublicKeysDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersEnvironmentsPublicKeysDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudshell.users.environments.publicKeys.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *UsersEnvironmentsPublicKeysDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes a public SSH key from an environment. Clients will no longer be\nable to connect to the environment using the corresponding private key.",
	//   "flatPath": "v1alpha1/users/{usersId}/environments/{environmentsId}/publicKeys/{publicKeysId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudshell.users.environments.publicKeys.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the resource to be deleted, e.g.\n`users/me/environments/default/publicKeys/my-key`.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+/environments/[^/]+/publicKeys/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
