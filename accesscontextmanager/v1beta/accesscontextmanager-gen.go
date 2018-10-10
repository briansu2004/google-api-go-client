// Package accesscontextmanager provides access to the Access Context Manager API.
//
// See https://cloud.google.com/access-context-manager/docs/reference/rest/
//
// Usage example:
//
//   import "google.golang.org/api/accesscontextmanager/v1beta"
//   ...
//   accesscontextmanagerService, err := accesscontextmanager.New(oauthHttpClient)
package accesscontextmanager // import "google.golang.org/api/accesscontextmanager/v1beta"

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

const apiId = "accesscontextmanager:v1beta"
const apiName = "accesscontextmanager"
const apiVersion = "v1beta"
const basePath = "https://accesscontextmanager.googleapis.com/"

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
	s.AccessPolicies = NewAccessPoliciesService(s)
	s.Operations = NewOperationsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	AccessPolicies *AccessPoliciesService

	Operations *OperationsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAccessPoliciesService(s *Service) *AccessPoliciesService {
	rs := &AccessPoliciesService{s: s}
	rs.AccessLevels = NewAccessPoliciesAccessLevelsService(s)
	rs.ServicePerimeters = NewAccessPoliciesServicePerimetersService(s)
	return rs
}

type AccessPoliciesService struct {
	s *Service

	AccessLevels *AccessPoliciesAccessLevelsService

	ServicePerimeters *AccessPoliciesServicePerimetersService
}

func NewAccessPoliciesAccessLevelsService(s *Service) *AccessPoliciesAccessLevelsService {
	rs := &AccessPoliciesAccessLevelsService{s: s}
	return rs
}

type AccessPoliciesAccessLevelsService struct {
	s *Service
}

func NewAccessPoliciesServicePerimetersService(s *Service) *AccessPoliciesServicePerimetersService {
	rs := &AccessPoliciesServicePerimetersService{s: s}
	return rs
}

type AccessPoliciesServicePerimetersService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

// AccessLevel: An `AccessLevel` is a label that can be applied to
// requests to GCP services,
// along with a list of requirements necessary for the label to be
// applied.
// `AccessLevels` can be referenced in `AccessZones` and in the `Cloud
// Org
// Policy` API.
type AccessLevel struct {
	// Basic: A `BasicLevel` composed of `Conditions`.
	Basic *BasicLevel `json:"basic,omitempty"`

	// CreateTime: Output only. Time the `AccessLevel` was created in UTC.
	CreateTime string `json:"createTime,omitempty"`

	// Description: Description of the `AccessLevel` and its use. Does not
	// affect behavior.
	Description string `json:"description,omitempty"`

	// Name: Required. Resource name for the Access Level. The `short_name`
	// component
	// must begin with a letter and only include alphanumeric and '_'.
	// Format:
	// `accessPolicies/{policy_id}/accessLevels/{short_name}`
	Name string `json:"name,omitempty"`

	// Title: Human readable title. Must be unique within the Policy.
	Title string `json:"title,omitempty"`

	// UpdateTime: Output only. Time the `AccessLevel` was updated in UTC.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Basic") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Basic") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AccessLevel) MarshalJSON() ([]byte, error) {
	type NoMethod AccessLevel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AccessPolicy: `AccessPolicy` is a container for `AccessLevels` (which
// define the necessary
// attributes to use GCP services) and `ServicePerimeters` (which define
// regions
// of services able to freely pass data within a perimeter). An access
// policy is
// globally visible within an organization, and the restrictions it
// specifies
// apply to all projects within an organization.
type AccessPolicy struct {
	// CreateTime: Output only. Time the `AccessPolicy` was created in UTC.
	CreateTime string `json:"createTime,omitempty"`

	// Name: Output only. Resource name of the `AccessPolicy`.
	// Format:
	// `accessPolicies/{policy_id}`
	Name string `json:"name,omitempty"`

	// Parent: Required. The parent of this `AccessPolicy` in the Cloud
	// Resource
	// Hierarchy. Currently immutable once created.
	// Format:
	// `organizations/{organization_id}`
	Parent string `json:"parent,omitempty"`

	// Title: Human readable title. Does not affect behavior.
	Title string `json:"title,omitempty"`

	// UpdateTime: Output only. Time the `AccessPolicy` was updated in UTC.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AccessPolicy) MarshalJSON() ([]byte, error) {
	type NoMethod AccessPolicy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BasicLevel: `BasicLevel` is an `AccessLevel` using a set of
// recommended features.
type BasicLevel struct {
	// CombiningFunction: How the `conditions` list should be combined to
	// determine if a request is
	// granted this `AccessLevel`. If AND is used, each `Condition`
	// in
	// `conditions` must be satisfied for the `AccessLevel` to be applied.
	// If OR
	// is used, at least one `Condition` in `conditions` must be satisfied
	// for the
	// `AccessLevel` to be applied. Default behavior is AND.
	//
	// Possible values:
	//   "AND" - All `Conditions` must be true for the `BasicLevel` to be
	// true.
	//   "OR" - If at least one `Condition` is true, then the `BasicLevel`
	// is true.
	CombiningFunction string `json:"combiningFunction,omitempty"`

	// Conditions: Required. A list of requirements for the `AccessLevel` to
	// be granted.
	Conditions []*Condition `json:"conditions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CombiningFunction")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CombiningFunction") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BasicLevel) MarshalJSON() ([]byte, error) {
	type NoMethod BasicLevel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Condition: A condition necessary for an `AccessLevel` to be granted.
// The Condition is an
// AND over its fields. So a Condition is true if: 1) the request IP is
// from one
// of the listed subnetworks AND 2) the originating device complies with
// the
// listed device policy AND 3) all listed access levels are granted AND
// 4) the
// request was sent at a time allowed by the DateTimeRestriction.
type Condition struct {
	// DevicePolicy: Device specific restrictions, all restrictions must
	// hold for the
	// Condition to be true. If not specified, all devices are allowed.
	DevicePolicy *DevicePolicy `json:"devicePolicy,omitempty"`

	// IpSubnetworks: CIDR block IP subnetwork specification. May be IPv4 or
	// IPv6. Note that for
	// a CIDR IP address block, the specified IP address portion must be
	// properly
	// truncated (i.e. all the host bits must be zero) or the input is
	// considered
	// malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24"
	// is
	// not. Similarly, for IPv6, "2001:db8::/32" is accepted
	// whereas
	// "2001:db8::1/32" is not. The originating IP of a request must be in
	// one of
	// the listed subnets in order for this Condition to be true. If empty,
	// all IP
	// addresses are allowed.
	IpSubnetworks []string `json:"ipSubnetworks,omitempty"`

	// Members: The signed-in user originating the request must be a part of
	// one of the
	// provided
	// members.
	// Syntax:
	// `user:{emailid}`
	// `group:{emailid}`
	// `serviceAccount:{e
	// mailid}`
	// If not specified, a request may come from any user (logged in/not
	// logged
	// in, not present in any groups, etc.).
	Members []string `json:"members,omitempty"`

	// Negate: Whether to negate the Condition. If true, the Condition
	// becomes a NAND over
	// its non-empty fields, each field must be false for the Condition
	// overall to
	// be satisfied. Defaults to false.
	Negate bool `json:"negate,omitempty"`

	// RequiredAccessLevels: A list of other access levels defined in the
	// same `Policy`, referenced by
	// resource name. Referencing an `AccessLevel` which does not exist is
	// an
	// error. All access levels listed must be granted for the Condition
	// to be true.
	// Example:
	// "accessPolicies/MY_POLICY/accessLevels/LEVEL_NAME"
	RequiredAccessLevels []string `json:"requiredAccessLevels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DevicePolicy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DevicePolicy") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Condition) MarshalJSON() ([]byte, error) {
	type NoMethod Condition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DevicePolicy: `DevicePolicy` specifies device specific restrictions
// necessary to acquire a
// given access level. A `DevicePolicy` specifies requirements for
// requests from
// devices to be granted access levels, it does not do any enforcement
// on the
// device. `DevicePolicy` acts as an AND over all specified fields, and
// each
// repeated field is an OR over its elements. Any unset fields are
// ignored. For
// example, if the proto is { os_type : DESKTOP_WINDOWS, os_type
// :
// DESKTOP_LINUX, encryption_status: ENCRYPTED}, then the DevicePolicy
// will be
// true for requests originating from encrypted Linux desktops and
// encrypted
// Windows desktops.
type DevicePolicy struct {
	// AllowedDeviceManagementLevels: Allowed device management levels, an
	// empty list allows all management
	// levels.
	//
	// Possible values:
	//   "MANAGEMENT_UNSPECIFIED" - The device's management level is not
	// specified or not known.
	//   "NONE" - The device is not managed.
	//   "BASIC" - Basic management is enabled, which is generally limited
	// to monitoring and
	// wiping the corporate account.
	//   "COMPLETE" - Complete device management. This includes more
	// thorough monitoring and the
	// ability to directly manage the device (such as remote wiping). This
	// can be
	// enabled through the Android Enterprise Platform.
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty"`

	// AllowedEncryptionStatuses: Allowed encryptions statuses, an empty
	// list allows all statuses.
	//
	// Possible values:
	//   "ENCRYPTION_UNSPECIFIED" - The encryption status of the device is
	// not specified or not known.
	//   "ENCRYPTION_UNSUPPORTED" - The device does not support encryption.
	//   "UNENCRYPTED" - The device supports encryption, but is currently
	// unencrypted.
	//   "ENCRYPTED" - The device is encrypted.
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty"`

	// OsConstraints: Allowed OS versions, an empty list allows all types
	// and all versions.
	OsConstraints []*OsConstraint `json:"osConstraints,omitempty"`

	// RequireScreenlock: Whether or not screenlock is required for the
	// DevicePolicy to be true.
	// Defaults to `false`.
	RequireScreenlock bool `json:"requireScreenlock,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AllowedDeviceManagementLevels") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AllowedDeviceManagementLevels") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DevicePolicy) MarshalJSON() ([]byte, error) {
	type NoMethod DevicePolicy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAccessLevelsResponse: A response to `ListAccessLevelsRequest`.
type ListAccessLevelsResponse struct {
	// AccessLevels: List of the Access Level instances.
	AccessLevels []*AccessLevel `json:"accessLevels,omitempty"`

	// NextPageToken: The pagination token to retrieve the next page of
	// results. If the value is
	// empty, no further results remain.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccessLevels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessLevels") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAccessLevelsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAccessLevelsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAccessPoliciesResponse: A response to
// `ListAccessPoliciesRequest`.
type ListAccessPoliciesResponse struct {
	// AccessPolicies: List of the AccessPolicy instances.
	AccessPolicies []*AccessPolicy `json:"accessPolicies,omitempty"`

	// NextPageToken: The pagination token to retrieve the next page of
	// results. If the value is
	// empty, no further results remain.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccessPolicies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessPolicies") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListAccessPoliciesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAccessPoliciesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListServicePerimetersResponse: A response to
// `ListServicePerimetersRequest`.
type ListServicePerimetersResponse struct {
	// NextPageToken: The pagination token to retrieve the next page of
	// results. If the value is
	// empty, no further results remain.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServicePerimeters: List of the Service Perimeter instances.
	ServicePerimeters []*ServicePerimeter `json:"servicePerimeters,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListServicePerimetersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListServicePerimetersResponse
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

// OsConstraint: A restriction on the OS type and version of devices
// making requests.
type OsConstraint struct {
	// MinimumVersion: The minimum allowed OS version. If not set, any
	// version of this OS
	// satisfies the constraint. Format: "major.minor.patch".
	// Examples: "10.5.301", "9.2.1".
	MinimumVersion string `json:"minimumVersion,omitempty"`

	// OsType: Required. The allowed OS type.
	//
	// Possible values:
	//   "OS_UNSPECIFIED" - The operating system of the device is not
	// specified or not known.
	//   "DESKTOP_MAC" - A desktop Mac operating system.
	//   "DESKTOP_WINDOWS" - A desktop Windows operating system.
	//   "DESKTOP_LINUX" - A desktop Linux operating system.
	//   "DESKTOP_CHROME_OS" - A desktop ChromeOS operating system.
	//   "ANDROID" - An Android operating system.
	//   "IOS" - An iOS operating system.
	OsType string `json:"osType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MinimumVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MinimumVersion") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *OsConstraint) MarshalJSON() ([]byte, error) {
	type NoMethod OsConstraint
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServicePerimeter: `ServicePerimeter` describes a set of GCP resources
// which can freely import
// and export data amongst themselves, but not export outside of
// the
// `ServicePerimeter`. If a request with a source within this
// `ServicePerimeter`
// has a target outside of the `ServicePerimeter`, the request will be
// blocked.
// Otherwise the request is allowed. There are two types of Service
// Perimeter -
// Regular and Bridge. Regular Service Perimeters cannot overlap, a
// single GCP
// project can only belong to a single regular Service Perimeter.
// Service
// Perimeter Bridges can contain only GCP projects as members, a single
// GCP
// project may belong to multiple Service Perimeter Bridges.
type ServicePerimeter struct {
	// CreateTime: Output only. Time the `ServicePerimeter` was created in
	// UTC.
	CreateTime string `json:"createTime,omitempty"`

	// Description: Description of the `ServicePerimeter` and its use. Does
	// not affect
	// behavior.
	Description string `json:"description,omitempty"`

	// Name: Required. Resource name for the ServicePerimeter.  The
	// `short_name`
	// component must begin with a letter and only include alphanumeric and
	// '_'.
	// Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
	Name string `json:"name,omitempty"`

	// PerimeterType: Perimeter type indicator. A single project is
	// allowed to be a member of single regular perimeter, but multiple
	// service
	// perimeter bridges. A project cannot be a included in a perimeter
	// bridge
	// without being included in regular perimeter. For perimeter
	// bridges,
	// restricted/unrestricted service lists as well as access lists must
	// be
	// empty.
	//
	// Possible values:
	//   "PERIMETER_TYPE_REGULAR" - Regular Perimeter.
	//   "PERIMETER_TYPE_BRIDGE" - Perimeter Bridge.
	PerimeterType string `json:"perimeterType,omitempty"`

	// Status: Current ServicePerimeter configuration. Specifies sets of
	// resources,
	// restricted/unrestricted services and access levels that determine
	// perimeter
	// content and boundaries.
	Status *ServicePerimeterConfig `json:"status,omitempty"`

	// Title: Human readable title. Must be unique within the Policy.
	Title string `json:"title,omitempty"`

	// UpdateTime: Output only. Time the `ServicePerimeter` was updated in
	// UTC.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServicePerimeter) MarshalJSON() ([]byte, error) {
	type NoMethod ServicePerimeter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServicePerimeterConfig: `ServicePerimeterConfig` specifies a set of
// GCP resources that describe
// specific Service Perimeter configuration.
type ServicePerimeterConfig struct {
	// AccessLevels: A list of `AccessLevel` resource names that allow
	// resources within the
	// `ServicePerimeter` to be accessed from the internet. `AccessLevels`
	// listed
	// must be in the same policy as this `ServicePerimeter`. Referencing
	// a
	// nonexistent `AccessLevel` is a syntax error. If no `AccessLevel`
	// names are
	// listed, resources within the perimeter can only be accessed via GCP
	// calls with
	// request origins within the perimeter.
	// Example:
	// "accessPolicies/MY_POLICY/accessLevels/MY_LEVEL".
	// For Service Perimeter Bridge, must be empty.
	AccessLevels []string `json:"accessLevels,omitempty"`

	// Resources: A list of GCP resources that are inside of the service
	// perimeter.
	// Currently only projects are allowed. Format:
	// `projects/{project_number}`
	Resources []string `json:"resources,omitempty"`

	// RestrictedServices: GCP services that are subject to the Service
	// Perimeter restrictions. May
	// contain a list of services or a single wildcard "*". For example,
	// if
	// `storage.googleapis.com` is specified, access to the storage
	// buckets
	// inside the perimeter must meet the perimeter's access
	// restrictions.
	//
	// Wildcard means that unless explicitly specified by
	// "unrestricted_services"
	// list, any service is treated as restricted. One of the
	// fields
	// "restricted_services", "unrestricted_services" must contain a
	// wildcard "*",
	// otherwise the Service Perimeter specification is invalid. It also
	// means
	// that both field being empty is invalid as well. "restricted_services"
	// can
	// be empty if and only if "unrestricted_services" list contains a
	// "*"
	// wildcard.
	RestrictedServices []string `json:"restrictedServices,omitempty"`

	// UnrestrictedServices: GCP services that are not subject to the
	// Service Perimeter restrictions.
	// May contain a list of services or a single wildcard "*". For example,
	// if
	// `logging.googleapis.com` is unrestricted, users can access logs
	// inside the
	// perimeter as if the perimeter doesn't exist, and it also means VMs
	// inside the perimeter
	// can access logs outside the perimeter.
	//
	// The wildcard means that unless explicitly specified
	// by
	// "restricted_services" list, any service is treated as unrestricted.
	// One of
	// the fields "restricted_services", "unrestricted_services" must
	// contain a
	// wildcard "*", otherwise the Service Perimeter specification is
	// invalid. It
	// also means that both field being empty is invalid as
	// well.
	// "unrestricted_services" can be empty if and only if
	// "restricted_services"
	// list contains a "*" wildcard.
	UnrestrictedServices []string `json:"unrestrictedServices,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccessLevels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessLevels") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServicePerimeterConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ServicePerimeterConfig
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

// method id "accesscontextmanager.accessPolicies.create":

type AccessPoliciesCreateCall struct {
	s            *Service
	accesspolicy *AccessPolicy
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Create: Create an `AccessPolicy`. Fails if this organization already
// has a
// `AccessPolicy`. The longrunning Operation will have a successful
// status
// once the `AccessPolicy` has propagated to long-lasting
// storage.
// Syntactic and basic semantic errors will be returned in `metadata` as
// a
// BadRequest proto.
func (r *AccessPoliciesService) Create(accesspolicy *AccessPolicy) *AccessPoliciesCreateCall {
	c := &AccessPoliciesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accesspolicy = accesspolicy
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesCreateCall) Fields(s ...googleapi.Field) *AccessPoliciesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesCreateCall) Context(ctx context.Context) *AccessPoliciesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.accesspolicy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/accessPolicies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Create an `AccessPolicy`. Fails if this organization already has a\n`AccessPolicy`. The longrunning Operation will have a successful status\nonce the `AccessPolicy` has propagated to long-lasting storage.\nSyntactic and basic semantic errors will be returned in `metadata` as a\nBadRequest proto.",
	//   "flatPath": "v1beta/accessPolicies",
	//   "httpMethod": "POST",
	//   "id": "accesscontextmanager.accessPolicies.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta/accessPolicies",
	//   "request": {
	//     "$ref": "AccessPolicy"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.delete":

type AccessPoliciesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an AccessPolicy by resource
// name. The longrunning Operation will have a successful status once
// the
// AccessPolicy
// has been removed from long-lasting storage.
func (r *AccessPoliciesService) Delete(name string) *AccessPoliciesDeleteCall {
	c := &AccessPoliciesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesDeleteCall) Fields(s ...googleapi.Field) *AccessPoliciesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesDeleteCall) Context(ctx context.Context) *AccessPoliciesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Delete an AccessPolicy by resource\nname. The longrunning Operation will have a successful status once the\nAccessPolicy\nhas been removed from long-lasting storage.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}",
	//   "httpMethod": "DELETE",
	//   "id": "accesscontextmanager.accessPolicies.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the access policy to delete.\n\nFormat `accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.get":

type AccessPoliciesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get an AccessPolicy by name.
func (r *AccessPoliciesService) Get(name string) *AccessPoliciesGetCall {
	c := &AccessPoliciesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesGetCall) Fields(s ...googleapi.Field) *AccessPoliciesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccessPoliciesGetCall) IfNoneMatch(entityTag string) *AccessPoliciesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesGetCall) Context(ctx context.Context) *AccessPoliciesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.get" call.
// Exactly one of *AccessPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AccessPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesGetCall) Do(opts ...googleapi.CallOption) (*AccessPolicy, error) {
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
	ret := &AccessPolicy{
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
	//   "description": "Get an AccessPolicy by name.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.accessPolicies.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the access policy to get.\n\nFormat `accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "AccessPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.list":

type AccessPoliciesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all AccessPolicies under a
// container.
func (r *AccessPoliciesService) List() *AccessPoliciesListCall {
	c := &AccessPoliciesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": Number of
// AccessPolicy instances to include in the list. Default 100.
func (c *AccessPoliciesListCall) PageSize(pageSize int64) *AccessPoliciesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Next page token
// for the next batch of AccessPolicy instances. Defaults to
// the first page of results.
func (c *AccessPoliciesListCall) PageToken(pageToken string) *AccessPoliciesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Parent sets the optional parameter "parent": Required. Resource name
// for the container to list AccessPolicy
// instances
// from.
//
// Format:
// `organizations/{org_id}`
func (c *AccessPoliciesListCall) Parent(parent string) *AccessPoliciesListCall {
	c.urlParams_.Set("parent", parent)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesListCall) Fields(s ...googleapi.Field) *AccessPoliciesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccessPoliciesListCall) IfNoneMatch(entityTag string) *AccessPoliciesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesListCall) Context(ctx context.Context) *AccessPoliciesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/accessPolicies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.list" call.
// Exactly one of *ListAccessPoliciesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListAccessPoliciesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccessPoliciesListCall) Do(opts ...googleapi.CallOption) (*ListAccessPoliciesResponse, error) {
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
	ret := &ListAccessPoliciesResponse{
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
	//   "description": "List all AccessPolicies under a\ncontainer.",
	//   "flatPath": "v1beta/accessPolicies",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.accessPolicies.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Number of AccessPolicy instances to include in the list. Default 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Next page token for the next batch of AccessPolicy instances. Defaults to\nthe first page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. Resource name for the container to list AccessPolicy instances\nfrom.\n\nFormat:\n`organizations/{org_id}`",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/accessPolicies",
	//   "response": {
	//     "$ref": "ListAccessPoliciesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccessPoliciesListCall) Pages(ctx context.Context, f func(*ListAccessPoliciesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "accesscontextmanager.accessPolicies.patch":

type AccessPoliciesPatchCall struct {
	s            *Service
	name         string
	accesspolicy *AccessPolicy
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Patch: Update an AccessPolicy. The
// longrunning Operation from this RPC will have a successful status
// once the
// changes to the AccessPolicy have propagated
// to long-lasting storage. Syntactic and basic semantic errors will
// be
// returned in `metadata` as a BadRequest proto.
func (r *AccessPoliciesService) Patch(name string, accesspolicy *AccessPolicy) *AccessPoliciesPatchCall {
	c := &AccessPoliciesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.accesspolicy = accesspolicy
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. Mask
// to control which fields get updated. Must be non-empty.
func (c *AccessPoliciesPatchCall) UpdateMask(updateMask string) *AccessPoliciesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesPatchCall) Fields(s ...googleapi.Field) *AccessPoliciesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesPatchCall) Context(ctx context.Context) *AccessPoliciesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.accesspolicy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Update an AccessPolicy. The\nlongrunning Operation from this RPC will have a successful status once the\nchanges to the AccessPolicy have propagated\nto long-lasting storage. Syntactic and basic semantic errors will be\nreturned in `metadata` as a BadRequest proto.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}",
	//   "httpMethod": "PATCH",
	//   "id": "accesscontextmanager.accessPolicies.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output only. Resource name of the `AccessPolicy`. Format:\n`accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Mask to control which fields get updated. Must be non-empty.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "request": {
	//     "$ref": "AccessPolicy"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.accessLevels.create":

type AccessPoliciesAccessLevelsCreateCall struct {
	s           *Service
	parent      string
	accesslevel *AccessLevel
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Create: Create an Access Level. The longrunning
// operation from this RPC will have a successful status once the
// Access
// Level has
// propagated to long-lasting storage. Access Levels containing
// errors will result in an error response for the first error
// encountered.
func (r *AccessPoliciesAccessLevelsService) Create(parent string, accesslevel *AccessLevel) *AccessPoliciesAccessLevelsCreateCall {
	c := &AccessPoliciesAccessLevelsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.accesslevel = accesslevel
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesAccessLevelsCreateCall) Fields(s ...googleapi.Field) *AccessPoliciesAccessLevelsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesAccessLevelsCreateCall) Context(ctx context.Context) *AccessPoliciesAccessLevelsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesAccessLevelsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesAccessLevelsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.accesslevel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/accessLevels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.accessLevels.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesAccessLevelsCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Create an Access Level. The longrunning\noperation from this RPC will have a successful status once the Access\nLevel has\npropagated to long-lasting storage. Access Levels containing\nerrors will result in an error response for the first error encountered.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/accessLevels",
	//   "httpMethod": "POST",
	//   "id": "accesscontextmanager.accessPolicies.accessLevels.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Resource name for the access policy which owns this Access\nLevel.\n\nFormat: `accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+parent}/accessLevels",
	//   "request": {
	//     "$ref": "AccessLevel"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.accessLevels.delete":

type AccessPoliciesAccessLevelsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an Access Level by resource
// name. The longrunning operation from this RPC will have a successful
// status
// once the Access Level has been removed
// from long-lasting storage.
func (r *AccessPoliciesAccessLevelsService) Delete(name string) *AccessPoliciesAccessLevelsDeleteCall {
	c := &AccessPoliciesAccessLevelsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesAccessLevelsDeleteCall) Fields(s ...googleapi.Field) *AccessPoliciesAccessLevelsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesAccessLevelsDeleteCall) Context(ctx context.Context) *AccessPoliciesAccessLevelsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesAccessLevelsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesAccessLevelsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.accessLevels.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesAccessLevelsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Delete an Access Level by resource\nname. The longrunning operation from this RPC will have a successful status\nonce the Access Level has been removed\nfrom long-lasting storage.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/accessLevels/{accessLevelsId}",
	//   "httpMethod": "DELETE",
	//   "id": "accesscontextmanager.accessPolicies.accessLevels.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the Access Level.\n\nFormat:\n`accessPolicies/{policy_id}/accessLevels/{access_level_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+/accessLevels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.accessLevels.get":

type AccessPoliciesAccessLevelsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get an Access Level by resource
// name.
func (r *AccessPoliciesAccessLevelsService) Get(name string) *AccessPoliciesAccessLevelsGetCall {
	c := &AccessPoliciesAccessLevelsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// AccessLevelFormat sets the optional parameter "accessLevelFormat":
// Whether to return `BasicLevels` in the Cloud Common
// Expression
// Language rather than as `BasicLevels`. Defaults to AS_DEFINED,
// where
// Access Levels
// are returned as `BasicLevels` or `CustomLevels` based on how they
// were
// created. If set to CEL, all Access Levels are returned
// as
// `CustomLevels`. In the CEL case, `BasicLevels` are translated to
// equivalent
// `CustomLevels`.
//
// Possible values:
//   "LEVEL_FORMAT_UNSPECIFIED"
//   "AS_DEFINED"
//   "CEL"
func (c *AccessPoliciesAccessLevelsGetCall) AccessLevelFormat(accessLevelFormat string) *AccessPoliciesAccessLevelsGetCall {
	c.urlParams_.Set("accessLevelFormat", accessLevelFormat)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesAccessLevelsGetCall) Fields(s ...googleapi.Field) *AccessPoliciesAccessLevelsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccessPoliciesAccessLevelsGetCall) IfNoneMatch(entityTag string) *AccessPoliciesAccessLevelsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesAccessLevelsGetCall) Context(ctx context.Context) *AccessPoliciesAccessLevelsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesAccessLevelsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesAccessLevelsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.accessLevels.get" call.
// Exactly one of *AccessLevel or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AccessLevel.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesAccessLevelsGetCall) Do(opts ...googleapi.CallOption) (*AccessLevel, error) {
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
	ret := &AccessLevel{
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
	//   "description": "Get an Access Level by resource\nname.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/accessLevels/{accessLevelsId}",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.accessPolicies.accessLevels.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "accessLevelFormat": {
	//       "description": "Whether to return `BasicLevels` in the Cloud Common Expression\nLanguage rather than as `BasicLevels`. Defaults to AS_DEFINED, where\nAccess Levels\nare returned as `BasicLevels` or `CustomLevels` based on how they were\ncreated. If set to CEL, all Access Levels are returned as\n`CustomLevels`. In the CEL case, `BasicLevels` are translated to equivalent\n`CustomLevels`.",
	//       "enum": [
	//         "LEVEL_FORMAT_UNSPECIFIED",
	//         "AS_DEFINED",
	//         "CEL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required. Resource name for the Access Level.\n\nFormat:\n`accessPolicies/{policy_id}/accessLevels/{access_level_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+/accessLevels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "AccessLevel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.accessLevels.list":

type AccessPoliciesAccessLevelsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all Access Levels for an access
// policy.
func (r *AccessPoliciesAccessLevelsService) List(parent string) *AccessPoliciesAccessLevelsListCall {
	c := &AccessPoliciesAccessLevelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// AccessLevelFormat sets the optional parameter "accessLevelFormat":
// Whether to return `BasicLevels` in the Cloud Common Expression
// language, as
// `CustomLevels`, rather than as `BasicLevels`. Defaults to
// returning
// `AccessLevels` in the format they were defined.
//
// Possible values:
//   "LEVEL_FORMAT_UNSPECIFIED"
//   "AS_DEFINED"
//   "CEL"
func (c *AccessPoliciesAccessLevelsListCall) AccessLevelFormat(accessLevelFormat string) *AccessPoliciesAccessLevelsListCall {
	c.urlParams_.Set("accessLevelFormat", accessLevelFormat)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of Access
// Levels to include in
// the list. Default 100.
func (c *AccessPoliciesAccessLevelsListCall) PageSize(pageSize int64) *AccessPoliciesAccessLevelsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Next page token
// for the next batch of Access Level instances.
// Defaults to the first page of results.
func (c *AccessPoliciesAccessLevelsListCall) PageToken(pageToken string) *AccessPoliciesAccessLevelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesAccessLevelsListCall) Fields(s ...googleapi.Field) *AccessPoliciesAccessLevelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccessPoliciesAccessLevelsListCall) IfNoneMatch(entityTag string) *AccessPoliciesAccessLevelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesAccessLevelsListCall) Context(ctx context.Context) *AccessPoliciesAccessLevelsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesAccessLevelsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesAccessLevelsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/accessLevels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.accessLevels.list" call.
// Exactly one of *ListAccessLevelsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListAccessLevelsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccessPoliciesAccessLevelsListCall) Do(opts ...googleapi.CallOption) (*ListAccessLevelsResponse, error) {
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
	ret := &ListAccessLevelsResponse{
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
	//   "description": "List all Access Levels for an access\npolicy.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/accessLevels",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.accessPolicies.accessLevels.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "accessLevelFormat": {
	//       "description": "Whether to return `BasicLevels` in the Cloud Common Expression language, as\n`CustomLevels`, rather than as `BasicLevels`. Defaults to returning\n`AccessLevels` in the format they were defined.",
	//       "enum": [
	//         "LEVEL_FORMAT_UNSPECIFIED",
	//         "AS_DEFINED",
	//         "CEL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of Access Levels to include in\nthe list. Default 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Next page token for the next batch of Access Level instances.\nDefaults to the first page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. Resource name for the access policy to list Access Levels from.\n\nFormat:\n`accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+parent}/accessLevels",
	//   "response": {
	//     "$ref": "ListAccessLevelsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccessPoliciesAccessLevelsListCall) Pages(ctx context.Context, f func(*ListAccessLevelsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "accesscontextmanager.accessPolicies.accessLevels.patch":

type AccessPoliciesAccessLevelsPatchCall struct {
	s           *Service
	name        string
	accesslevel *AccessLevel
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Update an Access Level. The longrunning
// operation from this RPC will have a successful status once the
// changes to
// the Access Level have propagated
// to long-lasting storage. Access Levels containing
// errors will result in an error response for the first error
// encountered.
func (r *AccessPoliciesAccessLevelsService) Patch(name string, accesslevel *AccessLevel) *AccessPoliciesAccessLevelsPatchCall {
	c := &AccessPoliciesAccessLevelsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.accesslevel = accesslevel
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.  Mask
// to control which fields get updated. Must be non-empty.
func (c *AccessPoliciesAccessLevelsPatchCall) UpdateMask(updateMask string) *AccessPoliciesAccessLevelsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesAccessLevelsPatchCall) Fields(s ...googleapi.Field) *AccessPoliciesAccessLevelsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesAccessLevelsPatchCall) Context(ctx context.Context) *AccessPoliciesAccessLevelsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesAccessLevelsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesAccessLevelsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.accesslevel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.accessLevels.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesAccessLevelsPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Update an Access Level. The longrunning\noperation from this RPC will have a successful status once the changes to\nthe Access Level have propagated\nto long-lasting storage. Access Levels containing\nerrors will result in an error response for the first error encountered.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/accessLevels/{accessLevelsId}",
	//   "httpMethod": "PATCH",
	//   "id": "accesscontextmanager.accessPolicies.accessLevels.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the Access Level. The `short_name` component\nmust begin with a letter and only include alphanumeric and '_'. Format:\n`accessPolicies/{policy_id}/accessLevels/{short_name}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+/accessLevels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required.  Mask to control which fields get updated. Must be non-empty.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "request": {
	//     "$ref": "AccessLevel"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.servicePerimeters.create":

type AccessPoliciesServicePerimetersCreateCall struct {
	s                *Service
	parent           string
	serviceperimeter *ServicePerimeter
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Create: Create an Service Perimeter. The
// longrunning operation from this RPC will have a successful status
// once the
// Service Perimeter has
// propagated to long-lasting storage. Service Perimeters
// containing
// errors will result in an error response for the first error
// encountered.
func (r *AccessPoliciesServicePerimetersService) Create(parent string, serviceperimeter *ServicePerimeter) *AccessPoliciesServicePerimetersCreateCall {
	c := &AccessPoliciesServicePerimetersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.serviceperimeter = serviceperimeter
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesServicePerimetersCreateCall) Fields(s ...googleapi.Field) *AccessPoliciesServicePerimetersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesServicePerimetersCreateCall) Context(ctx context.Context) *AccessPoliciesServicePerimetersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesServicePerimetersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesServicePerimetersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.serviceperimeter)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/servicePerimeters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.servicePerimeters.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesServicePerimetersCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Create an Service Perimeter. The\nlongrunning operation from this RPC will have a successful status once the\nService Perimeter has\npropagated to long-lasting storage. Service Perimeters containing\nerrors will result in an error response for the first error encountered.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/servicePerimeters",
	//   "httpMethod": "POST",
	//   "id": "accesscontextmanager.accessPolicies.servicePerimeters.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Resource name for the access policy which owns this Service\nPerimeter.\n\nFormat: `accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+parent}/servicePerimeters",
	//   "request": {
	//     "$ref": "ServicePerimeter"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.servicePerimeters.delete":

type AccessPoliciesServicePerimetersDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an Service Perimeter by resource
// name. The longrunning operation from this RPC will have a successful
// status
// once the Service Perimeter has been
// removed from long-lasting storage.
func (r *AccessPoliciesServicePerimetersService) Delete(name string) *AccessPoliciesServicePerimetersDeleteCall {
	c := &AccessPoliciesServicePerimetersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesServicePerimetersDeleteCall) Fields(s ...googleapi.Field) *AccessPoliciesServicePerimetersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesServicePerimetersDeleteCall) Context(ctx context.Context) *AccessPoliciesServicePerimetersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesServicePerimetersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesServicePerimetersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.servicePerimeters.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesServicePerimetersDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Delete an Service Perimeter by resource\nname. The longrunning operation from this RPC will have a successful status\nonce the Service Perimeter has been\nremoved from long-lasting storage.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/servicePerimeters/{servicePerimetersId}",
	//   "httpMethod": "DELETE",
	//   "id": "accesscontextmanager.accessPolicies.servicePerimeters.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the Service Perimeter.\n\nFormat:\n`accessPolicies/{policy_id}/servicePerimeters/{service_perimeter_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+/servicePerimeters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.servicePerimeters.get":

type AccessPoliciesServicePerimetersGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get an Service Perimeter by resource
// name.
func (r *AccessPoliciesServicePerimetersService) Get(name string) *AccessPoliciesServicePerimetersGetCall {
	c := &AccessPoliciesServicePerimetersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesServicePerimetersGetCall) Fields(s ...googleapi.Field) *AccessPoliciesServicePerimetersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccessPoliciesServicePerimetersGetCall) IfNoneMatch(entityTag string) *AccessPoliciesServicePerimetersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesServicePerimetersGetCall) Context(ctx context.Context) *AccessPoliciesServicePerimetersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesServicePerimetersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesServicePerimetersGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.servicePerimeters.get" call.
// Exactly one of *ServicePerimeter or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ServicePerimeter.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccessPoliciesServicePerimetersGetCall) Do(opts ...googleapi.CallOption) (*ServicePerimeter, error) {
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
	ret := &ServicePerimeter{
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
	//   "description": "Get an Service Perimeter by resource\nname.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/servicePerimeters/{servicePerimetersId}",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.accessPolicies.servicePerimeters.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the Service Perimeter.\n\nFormat:\n`accessPolicies/{policy_id}/servicePerimeters/{service_perimeters_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+/servicePerimeters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "ServicePerimeter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.accessPolicies.servicePerimeters.list":

type AccessPoliciesServicePerimetersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all Service Perimeters for an
// access policy.
func (r *AccessPoliciesServicePerimetersService) List(parent string) *AccessPoliciesServicePerimetersListCall {
	c := &AccessPoliciesServicePerimetersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Number of Service
// Perimeters to include
// in the list. Default 100.
func (c *AccessPoliciesServicePerimetersListCall) PageSize(pageSize int64) *AccessPoliciesServicePerimetersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Next page token
// for the next batch of Service Perimeter instances.
// Defaults to the first page of results.
func (c *AccessPoliciesServicePerimetersListCall) PageToken(pageToken string) *AccessPoliciesServicePerimetersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesServicePerimetersListCall) Fields(s ...googleapi.Field) *AccessPoliciesServicePerimetersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccessPoliciesServicePerimetersListCall) IfNoneMatch(entityTag string) *AccessPoliciesServicePerimetersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesServicePerimetersListCall) Context(ctx context.Context) *AccessPoliciesServicePerimetersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesServicePerimetersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesServicePerimetersListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/servicePerimeters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.servicePerimeters.list" call.
// Exactly one of *ListServicePerimetersResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListServicePerimetersResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccessPoliciesServicePerimetersListCall) Do(opts ...googleapi.CallOption) (*ListServicePerimetersResponse, error) {
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
	ret := &ListServicePerimetersResponse{
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
	//   "description": "List all Service Perimeters for an\naccess policy.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/servicePerimeters",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.accessPolicies.servicePerimeters.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Number of Service Perimeters to include\nin the list. Default 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Next page token for the next batch of Service Perimeter instances.\nDefaults to the first page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. Resource name for the access policy to list Service Perimeters from.\n\nFormat:\n`accessPolicies/{policy_id}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+parent}/servicePerimeters",
	//   "response": {
	//     "$ref": "ListServicePerimetersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccessPoliciesServicePerimetersListCall) Pages(ctx context.Context, f func(*ListServicePerimetersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "accesscontextmanager.accessPolicies.servicePerimeters.patch":

type AccessPoliciesServicePerimetersPatchCall struct {
	s                *Service
	name             string
	serviceperimeter *ServicePerimeter
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Patch: Update an Service Perimeter. The
// longrunning operation from this RPC will have a successful status
// once the
// changes to the Service Perimeter have
// propagated to long-lasting storage. Service Perimeter
// containing
// errors will result in an error response for the first error
// encountered.
func (r *AccessPoliciesServicePerimetersService) Patch(name string, serviceperimeter *ServicePerimeter) *AccessPoliciesServicePerimetersPatchCall {
	c := &AccessPoliciesServicePerimetersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.serviceperimeter = serviceperimeter
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. Mask
// to control which fields get updated. Must be non-empty.
func (c *AccessPoliciesServicePerimetersPatchCall) UpdateMask(updateMask string) *AccessPoliciesServicePerimetersPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccessPoliciesServicePerimetersPatchCall) Fields(s ...googleapi.Field) *AccessPoliciesServicePerimetersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccessPoliciesServicePerimetersPatchCall) Context(ctx context.Context) *AccessPoliciesServicePerimetersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccessPoliciesServicePerimetersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccessPoliciesServicePerimetersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.serviceperimeter)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.accessPolicies.servicePerimeters.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccessPoliciesServicePerimetersPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Update an Service Perimeter. The\nlongrunning operation from this RPC will have a successful status once the\nchanges to the Service Perimeter have\npropagated to long-lasting storage. Service Perimeter containing\nerrors will result in an error response for the first error encountered.",
	//   "flatPath": "v1beta/accessPolicies/{accessPoliciesId}/servicePerimeters/{servicePerimetersId}",
	//   "httpMethod": "PATCH",
	//   "id": "accesscontextmanager.accessPolicies.servicePerimeters.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name for the ServicePerimeter.  The `short_name`\ncomponent must begin with a letter and only include alphanumeric and '_'.\nFormat: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`",
	//       "location": "path",
	//       "pattern": "^accessPolicies/[^/]+/servicePerimeters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Mask to control which fields get updated. Must be non-empty.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "request": {
	//     "$ref": "ServicePerimeter"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "accesscontextmanager.operations.get":

type OperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsGetCall) IfNoneMatch(entityTag string) *OperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "accesscontextmanager.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1beta/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "accesscontextmanager.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
