// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/credentials/v1/common.proto

package credentials // import "google.golang.org/genproto/googleapis/iam/credentials/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GenerateAccessTokenRequest struct {
	// The resource name of the service account for which the credentials
	// are requested, in the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The sequence of service accounts in a delegation chain. Each service
	// account must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on its next service account in the chain. The last service account in the
	// chain must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on the service account that is specified in the `name` field of the
	// request.
	//
	// The delegates must have the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`
	Delegates []string `protobuf:"bytes,2,rep,name=delegates,proto3" json:"delegates,omitempty"`
	// Code to identify the scopes to be included in the OAuth 2.0 access token.
	// See https://developers.google.com/identity/protocols/googlescopes for more
	// information.
	// At least one value required.
	Scope []string `protobuf:"bytes,4,rep,name=scope,proto3" json:"scope,omitempty"`
	// The desired lifetime duration of the access token in seconds.
	// Must be set to a value less than or equal to 3600 (1 hour). If a value is
	// not specified, the token's lifetime will be set to a default value of one
	// hour.
	Lifetime             *duration.Duration `protobuf:"bytes,7,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GenerateAccessTokenRequest) Reset()         { *m = GenerateAccessTokenRequest{} }
func (m *GenerateAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateAccessTokenRequest) ProtoMessage()    {}
func (*GenerateAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{0}
}
func (m *GenerateAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAccessTokenRequest.Unmarshal(m, b)
}
func (m *GenerateAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAccessTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAccessTokenRequest.Merge(dst, src)
}
func (m *GenerateAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateAccessTokenRequest.Size(m)
}
func (m *GenerateAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAccessTokenRequest proto.InternalMessageInfo

func (m *GenerateAccessTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenerateAccessTokenRequest) GetDelegates() []string {
	if m != nil {
		return m.Delegates
	}
	return nil
}

func (m *GenerateAccessTokenRequest) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *GenerateAccessTokenRequest) GetLifetime() *duration.Duration {
	if m != nil {
		return m.Lifetime
	}
	return nil
}

type GenerateAccessTokenResponse struct {
	// The OAuth 2.0 access token.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Token expiration time.
	// The expiration time is always set.
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GenerateAccessTokenResponse) Reset()         { *m = GenerateAccessTokenResponse{} }
func (m *GenerateAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateAccessTokenResponse) ProtoMessage()    {}
func (*GenerateAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{1}
}
func (m *GenerateAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAccessTokenResponse.Unmarshal(m, b)
}
func (m *GenerateAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAccessTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAccessTokenResponse.Merge(dst, src)
}
func (m *GenerateAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateAccessTokenResponse.Size(m)
}
func (m *GenerateAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAccessTokenResponse proto.InternalMessageInfo

func (m *GenerateAccessTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GenerateAccessTokenResponse) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

type SignBlobRequest struct {
	// The resource name of the service account for which the credentials
	// are requested, in the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The sequence of service accounts in a delegation chain. Each service
	// account must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on its next service account in the chain. The last service account in the
	// chain must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on the service account that is specified in the `name` field of the
	// request.
	//
	// The delegates must have the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`
	Delegates []string `protobuf:"bytes,3,rep,name=delegates,proto3" json:"delegates,omitempty"`
	// The bytes to sign.
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignBlobRequest) Reset()         { *m = SignBlobRequest{} }
func (m *SignBlobRequest) String() string { return proto.CompactTextString(m) }
func (*SignBlobRequest) ProtoMessage()    {}
func (*SignBlobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{2}
}
func (m *SignBlobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignBlobRequest.Unmarshal(m, b)
}
func (m *SignBlobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignBlobRequest.Marshal(b, m, deterministic)
}
func (dst *SignBlobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignBlobRequest.Merge(dst, src)
}
func (m *SignBlobRequest) XXX_Size() int {
	return xxx_messageInfo_SignBlobRequest.Size(m)
}
func (m *SignBlobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignBlobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignBlobRequest proto.InternalMessageInfo

func (m *SignBlobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignBlobRequest) GetDelegates() []string {
	if m != nil {
		return m.Delegates
	}
	return nil
}

func (m *SignBlobRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SignBlobResponse struct {
	// The ID of the key used to sign the blob.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// The signed blob.
	SignedBlob           []byte   `protobuf:"bytes,4,opt,name=signed_blob,json=signedBlob,proto3" json:"signed_blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignBlobResponse) Reset()         { *m = SignBlobResponse{} }
func (m *SignBlobResponse) String() string { return proto.CompactTextString(m) }
func (*SignBlobResponse) ProtoMessage()    {}
func (*SignBlobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{3}
}
func (m *SignBlobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignBlobResponse.Unmarshal(m, b)
}
func (m *SignBlobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignBlobResponse.Marshal(b, m, deterministic)
}
func (dst *SignBlobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignBlobResponse.Merge(dst, src)
}
func (m *SignBlobResponse) XXX_Size() int {
	return xxx_messageInfo_SignBlobResponse.Size(m)
}
func (m *SignBlobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignBlobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignBlobResponse proto.InternalMessageInfo

func (m *SignBlobResponse) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SignBlobResponse) GetSignedBlob() []byte {
	if m != nil {
		return m.SignedBlob
	}
	return nil
}

type SignJwtRequest struct {
	// The resource name of the service account for which the credentials
	// are requested, in the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The sequence of service accounts in a delegation chain. Each service
	// account must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on its next service account in the chain. The last service account in the
	// chain must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on the service account that is specified in the `name` field of the
	// request.
	//
	// The delegates must have the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`
	Delegates []string `protobuf:"bytes,3,rep,name=delegates,proto3" json:"delegates,omitempty"`
	// The JWT payload to sign: a JSON object that contains a JWT Claims Set.
	Payload              string   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignJwtRequest) Reset()         { *m = SignJwtRequest{} }
func (m *SignJwtRequest) String() string { return proto.CompactTextString(m) }
func (*SignJwtRequest) ProtoMessage()    {}
func (*SignJwtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{4}
}
func (m *SignJwtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignJwtRequest.Unmarshal(m, b)
}
func (m *SignJwtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignJwtRequest.Marshal(b, m, deterministic)
}
func (dst *SignJwtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignJwtRequest.Merge(dst, src)
}
func (m *SignJwtRequest) XXX_Size() int {
	return xxx_messageInfo_SignJwtRequest.Size(m)
}
func (m *SignJwtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignJwtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignJwtRequest proto.InternalMessageInfo

func (m *SignJwtRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignJwtRequest) GetDelegates() []string {
	if m != nil {
		return m.Delegates
	}
	return nil
}

func (m *SignJwtRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type SignJwtResponse struct {
	// The ID of the key used to sign the JWT.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// The signed JWT.
	SignedJwt            string   `protobuf:"bytes,2,opt,name=signed_jwt,json=signedJwt,proto3" json:"signed_jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignJwtResponse) Reset()         { *m = SignJwtResponse{} }
func (m *SignJwtResponse) String() string { return proto.CompactTextString(m) }
func (*SignJwtResponse) ProtoMessage()    {}
func (*SignJwtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{5}
}
func (m *SignJwtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignJwtResponse.Unmarshal(m, b)
}
func (m *SignJwtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignJwtResponse.Marshal(b, m, deterministic)
}
func (dst *SignJwtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignJwtResponse.Merge(dst, src)
}
func (m *SignJwtResponse) XXX_Size() int {
	return xxx_messageInfo_SignJwtResponse.Size(m)
}
func (m *SignJwtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignJwtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignJwtResponse proto.InternalMessageInfo

func (m *SignJwtResponse) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SignJwtResponse) GetSignedJwt() string {
	if m != nil {
		return m.SignedJwt
	}
	return ""
}

type GenerateIdTokenRequest struct {
	// The resource name of the service account for which the credentials
	// are requested, in the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The sequence of service accounts in a delegation chain. Each service
	// account must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on its next service account in the chain. The last service account in the
	// chain must be granted the `roles/iam.serviceAccountTokenCreator` role
	// on the service account that is specified in the `name` field of the
	// request.
	//
	// The delegates must have the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`
	Delegates []string `protobuf:"bytes,2,rep,name=delegates,proto3" json:"delegates,omitempty"`
	// The audience for the token, such as the API or account that this token
	// grants access to.
	Audience string `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	// Include the service account email in the token. If set to `true`, the
	// token will contain `email` and `email_verified` claims.
	IncludeEmail         bool     `protobuf:"varint,4,opt,name=include_email,json=includeEmail,proto3" json:"include_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateIdTokenRequest) Reset()         { *m = GenerateIdTokenRequest{} }
func (m *GenerateIdTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateIdTokenRequest) ProtoMessage()    {}
func (*GenerateIdTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{6}
}
func (m *GenerateIdTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateIdTokenRequest.Unmarshal(m, b)
}
func (m *GenerateIdTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateIdTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateIdTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateIdTokenRequest.Merge(dst, src)
}
func (m *GenerateIdTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateIdTokenRequest.Size(m)
}
func (m *GenerateIdTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateIdTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateIdTokenRequest proto.InternalMessageInfo

func (m *GenerateIdTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenerateIdTokenRequest) GetDelegates() []string {
	if m != nil {
		return m.Delegates
	}
	return nil
}

func (m *GenerateIdTokenRequest) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func (m *GenerateIdTokenRequest) GetIncludeEmail() bool {
	if m != nil {
		return m.IncludeEmail
	}
	return false
}

type GenerateIdTokenResponse struct {
	// The OpenId Connect ID token.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateIdTokenResponse) Reset()         { *m = GenerateIdTokenResponse{} }
func (m *GenerateIdTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateIdTokenResponse) ProtoMessage()    {}
func (*GenerateIdTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{7}
}
func (m *GenerateIdTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateIdTokenResponse.Unmarshal(m, b)
}
func (m *GenerateIdTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateIdTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateIdTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateIdTokenResponse.Merge(dst, src)
}
func (m *GenerateIdTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateIdTokenResponse.Size(m)
}
func (m *GenerateIdTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateIdTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateIdTokenResponse proto.InternalMessageInfo

func (m *GenerateIdTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GenerateIdentityBindingAccessTokenRequest struct {
	// The resource name of the service account for which the credentials
	// are requested, in the following format:
	// `projects/-/serviceAccounts/{ACCOUNT_EMAIL_OR_UNIQUEID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Code to identify the scopes to be included in the OAuth 2.0 access token.
	// See https://developers.google.com/identity/protocols/googlescopes for more
	// information.
	// At least one value required.
	Scope []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	// Required. Input token.
	// Must be in JWT format according to
	// RFC7523 (https://tools.ietf.org/html/rfc7523)
	// and must have 'kid' field in the header.
	// Supported signing algorithms: RS256 (RS512, ES256, ES512 coming soon).
	// Mandatory payload fields (along the lines of RFC 7523, section 3):
	// - iss: issuer of the token. Must provide a discovery document at
	//        $iss/.well-known/openid-configuration . The document needs to be
	//        formatted according to section 4.2 of the OpenID Connect Discovery
	//        1.0 specification.
	// - iat: Issue time in seconds since epoch. Must be in the past.
	// - exp: Expiration time in seconds since epoch. Must be less than 48 hours
	//        after iat. We recommend to create tokens that last shorter than 6
	//        hours to improve security unless business reasons mandate longer
	//        expiration times. Shorter token lifetimes are generally more secure
	//        since tokens that have been exfiltrated by attackers can be used for
	//        a shorter time. you can configure the maximum lifetime of the
	//        incoming token in the configuration of the mapper.
	//        The resulting Google token will expire within an hour or at "exp",
	//        whichever is earlier.
	// - sub: JWT subject, identity asserted in the JWT.
	// - aud: Configured in the mapper policy. By default the service account
	//        email.
	//
	// Claims from the incoming token can be transferred into the output token
	// accoding to the mapper configuration. The outgoing claim size is limited.
	// Outgoing claims size must be less than 4kB serialized as JSON without
	// whitespace.
	//
	// Example header:
	// {
	//   "alg": "RS256",
	//   "kid": "92a4265e14ab04d4d228a48d10d4ca31610936f8"
	// }
	// Example payload:
	// {
	//   "iss": "https://accounts.google.com",
	//   "iat": 1517963104,
	//   "exp": 1517966704,
	//   "aud": "https://iamcredentials.googleapis.com/",
	//   "sub": "113475438248934895348",
	//   "my_claims": {
	//     "additional_claim": "value"
	//   }
	// }
	Jwt                  string   `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateIdentityBindingAccessTokenRequest) Reset() {
	*m = GenerateIdentityBindingAccessTokenRequest{}
}
func (m *GenerateIdentityBindingAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateIdentityBindingAccessTokenRequest) ProtoMessage()    {}
func (*GenerateIdentityBindingAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{8}
}
func (m *GenerateIdentityBindingAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateIdentityBindingAccessTokenRequest.Unmarshal(m, b)
}
func (m *GenerateIdentityBindingAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateIdentityBindingAccessTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateIdentityBindingAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateIdentityBindingAccessTokenRequest.Merge(dst, src)
}
func (m *GenerateIdentityBindingAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateIdentityBindingAccessTokenRequest.Size(m)
}
func (m *GenerateIdentityBindingAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateIdentityBindingAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateIdentityBindingAccessTokenRequest proto.InternalMessageInfo

func (m *GenerateIdentityBindingAccessTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenerateIdentityBindingAccessTokenRequest) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *GenerateIdentityBindingAccessTokenRequest) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type GenerateIdentityBindingAccessTokenResponse struct {
	// The OAuth 2.0 access token.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Token expiration time.
	// The expiration time is always set.
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GenerateIdentityBindingAccessTokenResponse) Reset() {
	*m = GenerateIdentityBindingAccessTokenResponse{}
}
func (m *GenerateIdentityBindingAccessTokenResponse) String() string {
	return proto.CompactTextString(m)
}
func (*GenerateIdentityBindingAccessTokenResponse) ProtoMessage() {}
func (*GenerateIdentityBindingAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_7e5e13abca9c147a, []int{9}
}
func (m *GenerateIdentityBindingAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateIdentityBindingAccessTokenResponse.Unmarshal(m, b)
}
func (m *GenerateIdentityBindingAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateIdentityBindingAccessTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateIdentityBindingAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateIdentityBindingAccessTokenResponse.Merge(dst, src)
}
func (m *GenerateIdentityBindingAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateIdentityBindingAccessTokenResponse.Size(m)
}
func (m *GenerateIdentityBindingAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateIdentityBindingAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateIdentityBindingAccessTokenResponse proto.InternalMessageInfo

func (m *GenerateIdentityBindingAccessTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GenerateIdentityBindingAccessTokenResponse) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateAccessTokenRequest)(nil), "google.iam.credentials.v1.GenerateAccessTokenRequest")
	proto.RegisterType((*GenerateAccessTokenResponse)(nil), "google.iam.credentials.v1.GenerateAccessTokenResponse")
	proto.RegisterType((*SignBlobRequest)(nil), "google.iam.credentials.v1.SignBlobRequest")
	proto.RegisterType((*SignBlobResponse)(nil), "google.iam.credentials.v1.SignBlobResponse")
	proto.RegisterType((*SignJwtRequest)(nil), "google.iam.credentials.v1.SignJwtRequest")
	proto.RegisterType((*SignJwtResponse)(nil), "google.iam.credentials.v1.SignJwtResponse")
	proto.RegisterType((*GenerateIdTokenRequest)(nil), "google.iam.credentials.v1.GenerateIdTokenRequest")
	proto.RegisterType((*GenerateIdTokenResponse)(nil), "google.iam.credentials.v1.GenerateIdTokenResponse")
	proto.RegisterType((*GenerateIdentityBindingAccessTokenRequest)(nil), "google.iam.credentials.v1.GenerateIdentityBindingAccessTokenRequest")
	proto.RegisterType((*GenerateIdentityBindingAccessTokenResponse)(nil), "google.iam.credentials.v1.GenerateIdentityBindingAccessTokenResponse")
}

func init() {
	proto.RegisterFile("google/iam/credentials/v1/common.proto", fileDescriptor_common_7e5e13abca9c147a)
}

var fileDescriptor_common_7e5e13abca9c147a = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0xa6, 0x8d, 0x27, 0x01, 0x2a, 0xab, 0x80, 0x13, 0x7e, 0x1a, 0x5c, 0x09, 0x05,
	0x0e, 0xb6, 0x0a, 0xe2, 0xd4, 0x53, 0xd3, 0xa2, 0x2a, 0x91, 0x90, 0x2a, 0xd3, 0x13, 0x02, 0x59,
	0x1b, 0xef, 0xd4, 0x5a, 0x62, 0xef, 0x1a, 0x7b, 0xdd, 0x90, 0x03, 0x4f, 0x50, 0xde, 0x80, 0x17,
	0xe5, 0x88, 0xbc, 0xb6, 0xe3, 0xa8, 0x8d, 0x50, 0xf8, 0xb9, 0xed, 0x7c, 0xf3, 0xcd, 0xcc, 0x37,
	0xb3, 0x3b, 0x0b, 0xcf, 0x03, 0x21, 0x82, 0x10, 0x1d, 0x46, 0x22, 0xc7, 0x4f, 0x90, 0x22, 0x97,
	0x8c, 0x84, 0xa9, 0x73, 0x75, 0xe8, 0xf8, 0x22, 0x8a, 0x04, 0xb7, 0xe3, 0x44, 0x48, 0x61, 0xf4,
	0x0a, 0x9e, 0xcd, 0x48, 0x64, 0xaf, 0xf0, 0xec, 0xab, 0xc3, 0xfe, 0xd3, 0x32, 0x85, 0x22, 0x4e,
	0xb3, 0x4b, 0x87, 0x66, 0x09, 0x91, 0xac, 0x0a, 0xed, 0xef, 0xdf, 0xf4, 0x4b, 0x16, 0x61, 0x2a,
	0x49, 0x14, 0x17, 0x04, 0xeb, 0x87, 0x06, 0xfd, 0x33, 0xe4, 0x98, 0x10, 0x89, 0xc7, 0xbe, 0x8f,
	0x69, 0x7a, 0x21, 0x66, 0xc8, 0x5d, 0xfc, 0x92, 0x61, 0x2a, 0x0d, 0x03, 0xb6, 0x38, 0x89, 0xd0,
	0xd4, 0x06, 0xda, 0x50, 0x77, 0xd5, 0xd9, 0x78, 0x0c, 0x3a, 0xc5, 0x10, 0x03, 0x22, 0x31, 0x35,
	0x1b, 0x83, 0xe6, 0x50, 0x77, 0x6b, 0xc0, 0xd8, 0x83, 0x56, 0xea, 0x8b, 0x18, 0xcd, 0x2d, 0xe5,
	0x29, 0x0c, 0xe3, 0x0d, 0xb4, 0x43, 0x76, 0x89, 0x79, 0x75, 0x73, 0x67, 0xa0, 0x0d, 0x3b, 0xaf,
	0x7a, 0x76, 0xd9, 0x55, 0x25, 0xcd, 0x3e, 0x2d, 0xa5, 0xbb, 0x4b, 0xaa, 0xf5, 0x0d, 0x1e, 0xad,
	0x15, 0x97, 0xc6, 0x82, 0xa7, 0x68, 0x3c, 0x83, 0x2e, 0x51, 0xb0, 0x27, 0x73, 0xbc, 0x54, 0xd9,
	0x21, 0x35, 0xd5, 0x38, 0x82, 0x0e, 0x7e, 0x8d, 0x59, 0x82, 0x9e, 0xaa, 0xdd, 0x54, 0xb5, 0xfb,
	0xb7, 0x6a, 0x5f, 0x54, 0x63, 0x71, 0xa1, 0xa0, 0xe7, 0x80, 0xf5, 0x09, 0xee, 0xbd, 0x67, 0x01,
	0x1f, 0x85, 0x62, 0xba, 0xf1, 0x40, 0x9a, 0x37, 0x07, 0x62, 0xc2, 0x4e, 0x4c, 0x16, 0xa1, 0x20,
	0xd4, 0x6c, 0x0d, 0xb4, 0x61, 0xd7, 0xad, 0x4c, 0x6b, 0x02, 0xbb, 0x75, 0xfa, 0xb2, 0xa5, 0xfb,
	0xb0, 0x3d, 0xc3, 0x85, 0xc7, 0x68, 0x59, 0xa1, 0x35, 0xc3, 0xc5, 0x98, 0x1a, 0xfb, 0xd0, 0x49,
	0x59, 0xc0, 0x91, 0x7a, 0xd3, 0x50, 0x4c, 0xcd, 0x2d, 0x95, 0x08, 0x0a, 0x28, 0x8f, 0xb7, 0x3e,
	0xc2, 0xdd, 0x3c, 0xd7, 0x64, 0x2e, 0xff, 0x9b, 0x52, 0xbd, 0x56, 0x7a, 0x56, 0x0c, 0x42, 0x65,
	0xff, 0xbd, 0xd0, 0x27, 0x50, 0xaa, 0xf2, 0x3e, 0xcf, 0xa5, 0xd9, 0x50, 0x2e, 0xbd, 0x40, 0x26,
	0x73, 0x69, 0x5d, 0x6b, 0xf0, 0xa0, 0xba, 0xd1, 0x31, 0xfd, 0xc7, 0xa7, 0xd6, 0x87, 0x36, 0xc9,
	0x28, 0x43, 0xee, 0x17, 0x17, 0xab, 0xbb, 0x4b, 0xdb, 0x38, 0x80, 0x3b, 0x8c, 0xfb, 0x61, 0x46,
	0xd1, 0xc3, 0x88, 0xb0, 0x50, 0x8d, 0xac, 0xed, 0x76, 0x4b, 0xf0, 0x6d, 0x8e, 0x59, 0x0e, 0x3c,
	0xbc, 0x25, 0xa6, 0x6c, 0x6f, 0x0f, 0x5a, 0xab, 0x6f, 0xaa, 0x30, 0xac, 0x00, 0x5e, 0xd4, 0x01,
	0xf9, 0x1a, 0xca, 0xc5, 0x88, 0x71, 0xca, 0x78, 0xb0, 0xe1, 0xee, 0x2c, 0xb7, 0xa3, 0xb1, 0xba,
	0x1d, 0xbb, 0xd0, 0xcc, 0xa7, 0x55, 0xf4, 0x90, 0x1f, 0xad, 0xef, 0x1a, 0xbc, 0xdc, 0xa4, 0xd2,
	0x5f, 0x2f, 0x42, 0xe3, 0x4f, 0x16, 0x61, 0x74, 0xad, 0xc1, 0x81, 0x2f, 0xa2, 0x8a, 0xed, 0x87,
	0x22, 0xa3, 0x6b, 0xbe, 0xa3, 0x51, 0x6f, 0x7c, 0xfc, 0xee, 0xa4, 0x86, 0x4e, 0xd4, 0x2f, 0x76,
	0x9e, 0xe7, 0x3e, 0xd7, 0x3e, 0x9c, 0x96, 0xd1, 0x81, 0x08, 0x09, 0x0f, 0x6c, 0x91, 0x04, 0x4e,
	0x80, 0x5c, 0x55, 0x76, 0x0a, 0x17, 0x89, 0x59, 0xba, 0xe6, 0x37, 0x3c, 0x5a, 0x31, 0x7f, 0x6a,
	0xda, 0x74, 0x5b, 0xc5, 0xbc, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x97, 0xa7, 0x4c, 0xe2, 0x40,
	0x05, 0x00, 0x00,
}
