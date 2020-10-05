// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: shared.proto

package degitxpb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ObjectType int32

const (
	ObjectType_UNKNOWN ObjectType = 0
	ObjectType_COMMIT  ObjectType = 1
	ObjectType_BLOB    ObjectType = 2
	ObjectType_TREE    ObjectType = 3
	ObjectType_TAG     ObjectType = 4
)

// Enum value maps for ObjectType.
var (
	ObjectType_name = map[int32]string{
		0: "UNKNOWN",
		1: "COMMIT",
		2: "BLOB",
		3: "TREE",
		4: "TAG",
	}
	ObjectType_value = map[string]int32{
		"UNKNOWN": 0,
		"COMMIT":  1,
		"BLOB":    2,
		"TREE":    3,
		"TAG":     4,
	}
)

func (x ObjectType) Enum() *ObjectType {
	p := new(ObjectType)
	*p = x
	return p
}

func (x ObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_proto_enumTypes[0].Descriptor()
}

func (ObjectType) Type() protoreflect.EnumType {
	return &file_shared_proto_enumTypes[0]
}

func (x ObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectType.Descriptor instead.
func (ObjectType) EnumDescriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

type SignatureType int32

const (
	SignatureType_NONE SignatureType = 0
	SignatureType_PGP  SignatureType = 1
	SignatureType_X509 SignatureType = 2 // maybe add X509+TSA or other combinations at a later step
)

// Enum value maps for SignatureType.
var (
	SignatureType_name = map[int32]string{
		0: "NONE",
		1: "PGP",
		2: "X509",
	}
	SignatureType_value = map[string]int32{
		"NONE": 0,
		"PGP":  1,
		"X509": 2,
	}
)

func (x SignatureType) Enum() *SignatureType {
	p := new(SignatureType)
	*p = x
	return p
}

func (x SignatureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignatureType) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_proto_enumTypes[1].Descriptor()
}

func (SignatureType) Type() protoreflect.EnumType {
	return &file_shared_proto_enumTypes[1]
}

func (x SignatureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignatureType.Descriptor instead.
func (SignatureType) EnumDescriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageName  string `protobuf:"bytes,2,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	RelativePath string `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	// Sets the GIT_OBJECT_DIRECTORY envvar on git commands to the value of this field.
	// It influences the object storage directory the SHA1 directories are created underneath.
	GitObjectDirectory string `protobuf:"bytes,4,opt,name=git_object_directory,json=gitObjectDirectory,proto3" json:"git_object_directory,omitempty"`
	// Sets the GIT_ALTERNATE_OBJECT_DIRECTORIES envvar on git commands to the values of this field.
	// It influences the list of Git object directories which can be used to search for Git objects.
	GitAlternateObjectDirectories []string `protobuf:"bytes,5,rep,name=git_alternate_object_directories,json=gitAlternateObjectDirectories,proto3" json:"git_alternate_object_directories,omitempty"`
	// Used in callbacks to GitLab so that it knows what repository the event is
	// associated with. May be left empty on RPC's that do not perform callbacks.
	// During project creation, `gl_repository` may not be known.
	GlRepository string `protobuf:"bytes,6,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	// The human-readable GitLab project path (e.g. gitlab-org/gitlab-ce).
	// When hashed storage is use, this associates a project path with its
	// path on disk. The name can change over time (e.g. when a project is
	// renamed). This is primarily used for logging/debugging at the
	// moment.
	GlProjectPath string `protobuf:"bytes,8,opt,name=gl_project_path,json=glProjectPath,proto3" json:"gl_project_path,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Repository) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *Repository) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

func (x *Repository) GetGitObjectDirectory() string {
	if x != nil {
		return x.GitObjectDirectory
	}
	return ""
}

func (x *Repository) GetGitAlternateObjectDirectories() []string {
	if x != nil {
		return x.GitAlternateObjectDirectories
	}
	return nil
}

func (x *Repository) GetGlRepository() string {
	if x != nil {
		return x.GlRepository
	}
	return ""
}

func (x *Repository) GetGlProjectPath() string {
	if x != nil {
		return x.GlProjectPath
	}
	return ""
}

// Corresponds to Gitlab::Git::Commit
type GitCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject   []byte        `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,5,opt,name=committer,proto3" json:"committer,omitempty"`
	ParentIds []string      `protobuf:"bytes,6,rep,name=parent_ids,json=parentIds,proto3" json:"parent_ids,omitempty"`
	// If body exceeds a certain threshold, it will be nullified,
	// but its size will be set in body_size so we can know if
	// a commit had a body in the first place.
	BodySize      int64         `protobuf:"varint,7,opt,name=body_size,json=bodySize,proto3" json:"body_size,omitempty"`
	SignatureType SignatureType `protobuf:"varint,8,opt,name=signature_type,json=signatureType,proto3,enum=degitx.SignatureType" json:"signature_type,omitempty"`
}

func (x *GitCommit) Reset() {
	*x = GitCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitCommit) ProtoMessage() {}

func (x *GitCommit) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitCommit.ProtoReflect.Descriptor instead.
func (*GitCommit) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

func (x *GitCommit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GitCommit) GetSubject() []byte {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *GitCommit) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *GitCommit) GetAuthor() *CommitAuthor {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *GitCommit) GetCommitter() *CommitAuthor {
	if x != nil {
		return x.Committer
	}
	return nil
}

func (x *GitCommit) GetParentIds() []string {
	if x != nil {
		return x.ParentIds
	}
	return nil
}

func (x *GitCommit) GetBodySize() int64 {
	if x != nil {
		return x.BodySize
	}
	return 0
}

func (x *GitCommit) GetSignatureType() SignatureType {
	if x != nil {
		return x.SignatureType
	}
	return SignatureType_NONE
}

type CommitAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     []byte               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    []byte               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Timezone []byte               `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
}

func (x *CommitAuthor) Reset() {
	*x = CommitAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAuthor) ProtoMessage() {}

func (x *CommitAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAuthor.ProtoReflect.Descriptor instead.
func (*CommitAuthor) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{2}
}

func (x *CommitAuthor) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CommitAuthor) GetEmail() []byte {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *CommitAuthor) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *CommitAuthor) GetTimezone() []byte {
	if x != nil {
		return x.Timezone
	}
	return nil
}

type ExitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExitStatus) Reset() {
	*x = ExitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitStatus) ProtoMessage() {}

func (x *ExitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitStatus.ProtoReflect.Descriptor instead.
func (*ExitStatus) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{3}
}

func (x *ExitStatus) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Corresponds to Gitlab::Git::Branch
type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TargetCommit *GitCommit `protobuf:"bytes,2,opt,name=target_commit,json=targetCommit,proto3" json:"target_commit,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{4}
}

func (x *Branch) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Branch) GetTargetCommit() *GitCommit {
	if x != nil {
		return x.TargetCommit
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id           string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TargetCommit *GitCommit `protobuf:"bytes,3,opt,name=target_commit,json=targetCommit,proto3" json:"target_commit,omitempty"`
	// If message exceeds a certain threshold, it will be nullified,
	// but its size will be set in message_size so we can know if
	// a tag had a message in the first place.
	Message       []byte        `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	MessageSize   int64         `protobuf:"varint,5,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	Tagger        *CommitAuthor `protobuf:"bytes,6,opt,name=tagger,proto3" json:"tagger,omitempty"`
	SignatureType SignatureType `protobuf:"varint,7,opt,name=signature_type,json=signatureType,proto3,enum=degitx.SignatureType" json:"signature_type,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{5}
}

func (x *Tag) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetTargetCommit() *GitCommit {
	if x != nil {
		return x.TargetCommit
	}
	return nil
}

func (x *Tag) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Tag) GetMessageSize() int64 {
	if x != nil {
		return x.MessageSize
	}
	return 0
}

func (x *Tag) GetTagger() *CommitAuthor {
	if x != nil {
		return x.Tagger
	}
	return nil
}

func (x *Tag) GetSignatureType() SignatureType {
	if x != nil {
		return x.SignatureType
	}
	return SignatureType_NONE
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlId       string `protobuf:"bytes,1,opt,name=gl_id,json=glId,proto3" json:"gl_id,omitempty"`
	Name       []byte `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email      []byte `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	GlUsername string `protobuf:"bytes,4,opt,name=gl_username,json=glUsername,proto3" json:"gl_username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetGlId() string {
	if x != nil {
		return x.GlId
	}
	return ""
}

func (x *User) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *User) GetEmail() []byte {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *User) GetGlUsername() string {
	if x != nil {
		return x.GlUsername
	}
	return ""
}

type ObjectPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *ObjectPool) Reset() {
	*x = ObjectPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectPool) ProtoMessage() {}

func (x *ObjectPool) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectPool.ProtoReflect.Descriptor instead.
func (*ObjectPool) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{7}
}

func (x *ObjectPool) GetRepository() *Repository {
	if x != nil {
		return x.Repository
	}
	return nil
}

type PaginationParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instructs pagination to start sending results after the provided page
	// token appears. A page token allows for a generic pattern to uniquely
	// identify a result or 'page'. Each paginated RPC may interpret a page
	// token differently.
	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// When fully consuming the response the client will receive _at most_
	// `limit` number of resulting objects. Note that the number of response
	// messages might be much lower, as some response messages already send
	// multiple objects per message.
	// When the limit is smaller than 0, it will be normalized to 2147483647
	// on the server side. When limit is not set, it defaults to 0, and no
	// results are send in the response.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *PaginationParameter) Reset() {
	*x = PaginationParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationParameter) ProtoMessage() {}

func (x *PaginationParameter) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationParameter.ProtoReflect.Descriptor instead.
func (*PaginationParameter) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{8}
}

func (x *PaginationParameter) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *PaginationParameter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// https://git-scm.com/docs/git/#_options
type GlobalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Treat pathspecs literally (i.e. no globbing, no pathspec magic)
	LiteralPathspecs bool `protobuf:"varint,1,opt,name=literal_pathspecs,json=literalPathspecs,proto3" json:"literal_pathspecs,omitempty"`
}

func (x *GlobalOptions) Reset() {
	*x = GlobalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalOptions) ProtoMessage() {}

func (x *GlobalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalOptions.ProtoReflect.Descriptor instead.
func (*GlobalOptions) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{9}
}

func (x *GlobalOptions) GetLiteralPathspecs() bool {
	if x != nil {
		return x.LiteralPathspecs
	}
	return false
}

var File_shared_proto protoreflect.FileDescriptor

var file_shared_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x69,
	0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x69, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x20,
	0x67, 0x69, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x67, 0x69, 0x74, 0x41, 0x6c, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0xa5, 0x02, 0x0a, 0x09, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x2c, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x32,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3c,
	0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x45, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64,
	0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x8a, 0x02,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x0d, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x66, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x13, 0x0a, 0x05, 0x67, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x67, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x46, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x12, 0x38, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x04, 0x90, 0xc6, 0x2c, 0x01, 0x52, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x4a, 0x0a, 0x13, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3c, 0x0a, 0x0d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x2a, 0x42, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x4c, 0x4f, 0x42, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x52, 0x45, 0x45, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x41, 0x47, 0x10, 0x04, 0x2a, 0x2c, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x47, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x58, 0x35, 0x30, 0x39, 0x10, 0x02, 0x42, 0x23, 0x5a, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x71,
	0x66, 0x6e, 0x2f, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x64, 0x65, 0x67, 0x69, 0x74, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_shared_proto_rawDescOnce sync.Once
	file_shared_proto_rawDescData = file_shared_proto_rawDesc
)

func file_shared_proto_rawDescGZIP() []byte {
	file_shared_proto_rawDescOnce.Do(func() {
		file_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_rawDescData)
	})
	return file_shared_proto_rawDescData
}

var file_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_shared_proto_goTypes = []interface{}{
	(ObjectType)(0),             // 0: degitx.ObjectType
	(SignatureType)(0),          // 1: degitx.SignatureType
	(*Repository)(nil),          // 2: degitx.Repository
	(*GitCommit)(nil),           // 3: degitx.GitCommit
	(*CommitAuthor)(nil),        // 4: degitx.CommitAuthor
	(*ExitStatus)(nil),          // 5: degitx.ExitStatus
	(*Branch)(nil),              // 6: degitx.Branch
	(*Tag)(nil),                 // 7: degitx.Tag
	(*User)(nil),                // 8: degitx.User
	(*ObjectPool)(nil),          // 9: degitx.ObjectPool
	(*PaginationParameter)(nil), // 10: degitx.PaginationParameter
	(*GlobalOptions)(nil),       // 11: degitx.GlobalOptions
	(*timestamp.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_shared_proto_depIdxs = []int32{
	4,  // 0: degitx.GitCommit.author:type_name -> degitx.CommitAuthor
	4,  // 1: degitx.GitCommit.committer:type_name -> degitx.CommitAuthor
	1,  // 2: degitx.GitCommit.signature_type:type_name -> degitx.SignatureType
	12, // 3: degitx.CommitAuthor.date:type_name -> google.protobuf.Timestamp
	3,  // 4: degitx.Branch.target_commit:type_name -> degitx.GitCommit
	3,  // 5: degitx.Tag.target_commit:type_name -> degitx.GitCommit
	4,  // 6: degitx.Tag.tagger:type_name -> degitx.CommitAuthor
	1,  // 7: degitx.Tag.signature_type:type_name -> degitx.SignatureType
	2,  // 8: degitx.ObjectPool.repository:type_name -> degitx.Repository
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_shared_proto_init() }
func file_shared_proto_init() {
	if File_shared_proto != nil {
		return
	}
	file_lint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitCommit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAuthor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectPool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationParameter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shared_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_proto_goTypes,
		DependencyIndexes: file_shared_proto_depIdxs,
		EnumInfos:         file_shared_proto_enumTypes,
		MessageInfos:      file_shared_proto_msgTypes,
	}.Build()
	File_shared_proto = out.File
	file_shared_proto_rawDesc = nil
	file_shared_proto_goTypes = nil
	file_shared_proto_depIdxs = nil
}