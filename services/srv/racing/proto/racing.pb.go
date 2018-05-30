// Code generated by protoc-gen-go. DO NOT EDIT.
// source: racing.proto

package racing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListMeetingsByDateRequest struct {
	StartDate            int64    `protobuf:"varint,1,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	EndDate              int64    `protobuf:"varint,2,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMeetingsByDateRequest) Reset()         { *m = ListMeetingsByDateRequest{} }
func (m *ListMeetingsByDateRequest) String() string { return proto.CompactTextString(m) }
func (*ListMeetingsByDateRequest) ProtoMessage()    {}
func (*ListMeetingsByDateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{0}
}
func (m *ListMeetingsByDateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeetingsByDateRequest.Unmarshal(m, b)
}
func (m *ListMeetingsByDateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeetingsByDateRequest.Marshal(b, m, deterministic)
}
func (dst *ListMeetingsByDateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeetingsByDateRequest.Merge(dst, src)
}
func (m *ListMeetingsByDateRequest) XXX_Size() int {
	return xxx_messageInfo_ListMeetingsByDateRequest.Size(m)
}
func (m *ListMeetingsByDateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeetingsByDateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeetingsByDateRequest proto.InternalMessageInfo

func (m *ListMeetingsByDateRequest) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *ListMeetingsByDateRequest) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

type ListMeetingsByDateResponse struct {
	Meetings             []*Meeting `protobuf:"bytes,1,rep,name=meetings" json:"meetings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMeetingsByDateResponse) Reset()         { *m = ListMeetingsByDateResponse{} }
func (m *ListMeetingsByDateResponse) String() string { return proto.CompactTextString(m) }
func (*ListMeetingsByDateResponse) ProtoMessage()    {}
func (*ListMeetingsByDateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{1}
}
func (m *ListMeetingsByDateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeetingsByDateResponse.Unmarshal(m, b)
}
func (m *ListMeetingsByDateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeetingsByDateResponse.Marshal(b, m, deterministic)
}
func (dst *ListMeetingsByDateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeetingsByDateResponse.Merge(dst, src)
}
func (m *ListMeetingsByDateResponse) XXX_Size() int {
	return xxx_messageInfo_ListMeetingsByDateResponse.Size(m)
}
func (m *ListMeetingsByDateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeetingsByDateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeetingsByDateResponse proto.InternalMessageInfo

func (m *ListMeetingsByDateResponse) GetMeetings() []*Meeting {
	if m != nil {
		return m.Meetings
	}
	return nil
}

type ListRacesByMeetingDateRequest struct {
	StartDate            int64    `protobuf:"varint,1,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	EndDate              int64    `protobuf:"varint,2,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRacesByMeetingDateRequest) Reset()         { *m = ListRacesByMeetingDateRequest{} }
func (m *ListRacesByMeetingDateRequest) String() string { return proto.CompactTextString(m) }
func (*ListRacesByMeetingDateRequest) ProtoMessage()    {}
func (*ListRacesByMeetingDateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{2}
}
func (m *ListRacesByMeetingDateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRacesByMeetingDateRequest.Unmarshal(m, b)
}
func (m *ListRacesByMeetingDateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRacesByMeetingDateRequest.Marshal(b, m, deterministic)
}
func (dst *ListRacesByMeetingDateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRacesByMeetingDateRequest.Merge(dst, src)
}
func (m *ListRacesByMeetingDateRequest) XXX_Size() int {
	return xxx_messageInfo_ListRacesByMeetingDateRequest.Size(m)
}
func (m *ListRacesByMeetingDateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRacesByMeetingDateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRacesByMeetingDateRequest proto.InternalMessageInfo

func (m *ListRacesByMeetingDateRequest) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *ListRacesByMeetingDateRequest) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

type ListRacesByMeetingDateResponse struct {
	Races                []*Race  `protobuf:"bytes,1,rep,name=races" json:"races,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRacesByMeetingDateResponse) Reset()         { *m = ListRacesByMeetingDateResponse{} }
func (m *ListRacesByMeetingDateResponse) String() string { return proto.CompactTextString(m) }
func (*ListRacesByMeetingDateResponse) ProtoMessage()    {}
func (*ListRacesByMeetingDateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{3}
}
func (m *ListRacesByMeetingDateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRacesByMeetingDateResponse.Unmarshal(m, b)
}
func (m *ListRacesByMeetingDateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRacesByMeetingDateResponse.Marshal(b, m, deterministic)
}
func (dst *ListRacesByMeetingDateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRacesByMeetingDateResponse.Merge(dst, src)
}
func (m *ListRacesByMeetingDateResponse) XXX_Size() int {
	return xxx_messageInfo_ListRacesByMeetingDateResponse.Size(m)
}
func (m *ListRacesByMeetingDateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRacesByMeetingDateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRacesByMeetingDateResponse proto.InternalMessageInfo

func (m *ListRacesByMeetingDateResponse) GetRaces() []*Race {
	if m != nil {
		return m.Races
	}
	return nil
}

type AddRacesRequest struct {
	MeetingId            string   `protobuf:"bytes,1,opt,name=meeting_id,json=meetingId" json:"meeting_id,omitempty"`
	Races                []*Race  `protobuf:"bytes,2,rep,name=races" json:"races,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRacesRequest) Reset()         { *m = AddRacesRequest{} }
func (m *AddRacesRequest) String() string { return proto.CompactTextString(m) }
func (*AddRacesRequest) ProtoMessage()    {}
func (*AddRacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{4}
}
func (m *AddRacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRacesRequest.Unmarshal(m, b)
}
func (m *AddRacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRacesRequest.Marshal(b, m, deterministic)
}
func (dst *AddRacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRacesRequest.Merge(dst, src)
}
func (m *AddRacesRequest) XXX_Size() int {
	return xxx_messageInfo_AddRacesRequest.Size(m)
}
func (m *AddRacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRacesRequest proto.InternalMessageInfo

func (m *AddRacesRequest) GetMeetingId() string {
	if m != nil {
		return m.MeetingId
	}
	return ""
}

func (m *AddRacesRequest) GetRaces() []*Race {
	if m != nil {
		return m.Races
	}
	return nil
}

type AddRacesResponse struct {
	Created              bool     `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRacesResponse) Reset()         { *m = AddRacesResponse{} }
func (m *AddRacesResponse) String() string { return proto.CompactTextString(m) }
func (*AddRacesResponse) ProtoMessage()    {}
func (*AddRacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{5}
}
func (m *AddRacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRacesResponse.Unmarshal(m, b)
}
func (m *AddRacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRacesResponse.Marshal(b, m, deterministic)
}
func (dst *AddRacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRacesResponse.Merge(dst, src)
}
func (m *AddRacesResponse) XXX_Size() int {
	return xxx_messageInfo_AddRacesResponse.Size(m)
}
func (m *AddRacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRacesResponse proto.InternalMessageInfo

func (m *AddRacesResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type AddMeetingsRequest struct {
	Meetings             []*Meeting `protobuf:"bytes,1,rep,name=meetings" json:"meetings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddMeetingsRequest) Reset()         { *m = AddMeetingsRequest{} }
func (m *AddMeetingsRequest) String() string { return proto.CompactTextString(m) }
func (*AddMeetingsRequest) ProtoMessage()    {}
func (*AddMeetingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{6}
}
func (m *AddMeetingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMeetingsRequest.Unmarshal(m, b)
}
func (m *AddMeetingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMeetingsRequest.Marshal(b, m, deterministic)
}
func (dst *AddMeetingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMeetingsRequest.Merge(dst, src)
}
func (m *AddMeetingsRequest) XXX_Size() int {
	return xxx_messageInfo_AddMeetingsRequest.Size(m)
}
func (m *AddMeetingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMeetingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMeetingsRequest proto.InternalMessageInfo

func (m *AddMeetingsRequest) GetMeetings() []*Meeting {
	if m != nil {
		return m.Meetings
	}
	return nil
}

type AddMeetingsResponse struct {
	Created              bool     `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMeetingsResponse) Reset()         { *m = AddMeetingsResponse{} }
func (m *AddMeetingsResponse) String() string { return proto.CompactTextString(m) }
func (*AddMeetingsResponse) ProtoMessage()    {}
func (*AddMeetingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{7}
}
func (m *AddMeetingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMeetingsResponse.Unmarshal(m, b)
}
func (m *AddMeetingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMeetingsResponse.Marshal(b, m, deterministic)
}
func (dst *AddMeetingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMeetingsResponse.Merge(dst, src)
}
func (m *AddMeetingsResponse) XXX_Size() int {
	return xxx_messageInfo_AddMeetingsResponse.Size(m)
}
func (m *AddMeetingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMeetingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMeetingsResponse proto.InternalMessageInfo

func (m *AddMeetingsResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type UpdateRaceRequest struct {
	Race                 *Race        `protobuf:"bytes,1,opt,name=race" json:"race,omitempty"`
	Selections           []*Selection `protobuf:"bytes,2,rep,name=selections" json:"selections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateRaceRequest) Reset()         { *m = UpdateRaceRequest{} }
func (m *UpdateRaceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRaceRequest) ProtoMessage()    {}
func (*UpdateRaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{8}
}
func (m *UpdateRaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRaceRequest.Unmarshal(m, b)
}
func (m *UpdateRaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRaceRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRaceRequest.Merge(dst, src)
}
func (m *UpdateRaceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRaceRequest.Size(m)
}
func (m *UpdateRaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRaceRequest proto.InternalMessageInfo

func (m *UpdateRaceRequest) GetRace() *Race {
	if m != nil {
		return m.Race
	}
	return nil
}

func (m *UpdateRaceRequest) GetSelections() []*Selection {
	if m != nil {
		return m.Selections
	}
	return nil
}

type UpdateRaceResponse struct {
	Updated              bool     `protobuf:"varint,1,opt,name=updated" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRaceResponse) Reset()         { *m = UpdateRaceResponse{} }
func (m *UpdateRaceResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateRaceResponse) ProtoMessage()    {}
func (*UpdateRaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{9}
}
func (m *UpdateRaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRaceResponse.Unmarshal(m, b)
}
func (m *UpdateRaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRaceResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateRaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRaceResponse.Merge(dst, src)
}
func (m *UpdateRaceResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateRaceResponse.Size(m)
}
func (m *UpdateRaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRaceResponse proto.InternalMessageInfo

func (m *UpdateRaceResponse) GetUpdated() bool {
	if m != nil {
		return m.Updated
	}
	return false
}

type GetNextRaceRequest struct {
	MeetingId            string   `protobuf:"bytes,1,opt,name=meeting_id,json=meetingId" json:"meeting_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNextRaceRequest) Reset()         { *m = GetNextRaceRequest{} }
func (m *GetNextRaceRequest) String() string { return proto.CompactTextString(m) }
func (*GetNextRaceRequest) ProtoMessage()    {}
func (*GetNextRaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{10}
}
func (m *GetNextRaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNextRaceRequest.Unmarshal(m, b)
}
func (m *GetNextRaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNextRaceRequest.Marshal(b, m, deterministic)
}
func (dst *GetNextRaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextRaceRequest.Merge(dst, src)
}
func (m *GetNextRaceRequest) XXX_Size() int {
	return xxx_messageInfo_GetNextRaceRequest.Size(m)
}
func (m *GetNextRaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextRaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextRaceRequest proto.InternalMessageInfo

func (m *GetNextRaceRequest) GetMeetingId() string {
	if m != nil {
		return m.MeetingId
	}
	return ""
}

type GetNextRaceResponse struct {
	Race                 *Race    `protobuf:"bytes,1,opt,name=race" json:"race,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNextRaceResponse) Reset()         { *m = GetNextRaceResponse{} }
func (m *GetNextRaceResponse) String() string { return proto.CompactTextString(m) }
func (*GetNextRaceResponse) ProtoMessage()    {}
func (*GetNextRaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{11}
}
func (m *GetNextRaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNextRaceResponse.Unmarshal(m, b)
}
func (m *GetNextRaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNextRaceResponse.Marshal(b, m, deterministic)
}
func (dst *GetNextRaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextRaceResponse.Merge(dst, src)
}
func (m *GetNextRaceResponse) XXX_Size() int {
	return xxx_messageInfo_GetNextRaceResponse.Size(m)
}
func (m *GetNextRaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextRaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextRaceResponse proto.InternalMessageInfo

func (m *GetNextRaceResponse) GetRace() *Race {
	if m != nil {
		return m.Race
	}
	return nil
}

type Meeting struct {
	MeetingId            string   `protobuf:"bytes,1,opt,name=meeting_id,json=meetingId" json:"meeting_id,omitempty"`
	SourceId             string   `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
	RaceType             string   `protobuf:"bytes,5,opt,name=race_type,json=raceType" json:"race_type,omitempty"`
	ScheduledStart       int64    `protobuf:"varint,6,opt,name=scheduled_start,json=scheduledStart" json:"scheduled_start,omitempty"`
	RaceIds              []string `protobuf:"bytes,7,rep,name=race_ids,json=raceIds" json:"race_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meeting) Reset()         { *m = Meeting{} }
func (m *Meeting) String() string { return proto.CompactTextString(m) }
func (*Meeting) ProtoMessage()    {}
func (*Meeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{12}
}
func (m *Meeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meeting.Unmarshal(m, b)
}
func (m *Meeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meeting.Marshal(b, m, deterministic)
}
func (dst *Meeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meeting.Merge(dst, src)
}
func (m *Meeting) XXX_Size() int {
	return xxx_messageInfo_Meeting.Size(m)
}
func (m *Meeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Meeting.DiscardUnknown(m)
}

var xxx_messageInfo_Meeting proto.InternalMessageInfo

func (m *Meeting) GetMeetingId() string {
	if m != nil {
		return m.MeetingId
	}
	return ""
}

func (m *Meeting) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Meeting) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Meeting) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Meeting) GetRaceType() string {
	if m != nil {
		return m.RaceType
	}
	return ""
}

func (m *Meeting) GetScheduledStart() int64 {
	if m != nil {
		return m.ScheduledStart
	}
	return 0
}

func (m *Meeting) GetRaceIds() []string {
	if m != nil {
		return m.RaceIds
	}
	return nil
}

type Competitor struct {
	CompetitorId         string   `protobuf:"bytes,1,opt,name=competitor_id,json=competitorId" json:"competitor_id,omitempty"`
	SourceId             string   `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth" json:"date_of_birth,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	Trainer              string   `protobuf:"bytes,6,opt,name=trainer" json:"trainer,omitempty"`
	Gender               string   `protobuf:"bytes,7,opt,name=gender" json:"gender,omitempty"`
	ImageUrl             string   `protobuf:"bytes,8,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Competitor) Reset()         { *m = Competitor{} }
func (m *Competitor) String() string { return proto.CompactTextString(m) }
func (*Competitor) ProtoMessage()    {}
func (*Competitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{13}
}
func (m *Competitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Competitor.Unmarshal(m, b)
}
func (m *Competitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Competitor.Marshal(b, m, deterministic)
}
func (dst *Competitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Competitor.Merge(dst, src)
}
func (m *Competitor) XXX_Size() int {
	return xxx_messageInfo_Competitor.Size(m)
}
func (m *Competitor) XXX_DiscardUnknown() {
	xxx_messageInfo_Competitor.DiscardUnknown(m)
}

var xxx_messageInfo_Competitor proto.InternalMessageInfo

func (m *Competitor) GetCompetitorId() string {
	if m != nil {
		return m.CompetitorId
	}
	return ""
}

func (m *Competitor) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Competitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Competitor) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *Competitor) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Competitor) GetTrainer() string {
	if m != nil {
		return m.Trainer
	}
	return ""
}

func (m *Competitor) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Competitor) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type Selection struct {
	SelectionId          string   `protobuf:"bytes,1,opt,name=selection_id,json=selectionId" json:"selection_id,omitempty"`
	SourceId             string   `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	CompetitorId         string   `protobuf:"bytes,3,opt,name=competitor_id,json=competitorId" json:"competitor_id,omitempty"`
	SourceCompetitorId   string   `protobuf:"bytes,4,opt,name=source_competitor_id,json=sourceCompetitorId" json:"source_competitor_id,omitempty"`
	RaceId               string   `protobuf:"bytes,5,opt,name=race_id,json=raceId" json:"race_id,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Jockey               string   `protobuf:"bytes,7,opt,name=jockey" json:"jockey,omitempty"`
	Number               int32    `protobuf:"varint,8,opt,name=number" json:"number,omitempty"`
	BarrierNumber        int32    `protobuf:"varint,9,opt,name=barrier_number,json=barrierNumber" json:"barrier_number,omitempty"`
	LastUpdated          int64    `protobuf:"varint,10,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selection) Reset()         { *m = Selection{} }
func (m *Selection) String() string { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()    {}
func (*Selection) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{14}
}
func (m *Selection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selection.Unmarshal(m, b)
}
func (m *Selection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selection.Marshal(b, m, deterministic)
}
func (dst *Selection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selection.Merge(dst, src)
}
func (m *Selection) XXX_Size() int {
	return xxx_messageInfo_Selection.Size(m)
}
func (m *Selection) XXX_DiscardUnknown() {
	xxx_messageInfo_Selection.DiscardUnknown(m)
}

var xxx_messageInfo_Selection proto.InternalMessageInfo

func (m *Selection) GetSelectionId() string {
	if m != nil {
		return m.SelectionId
	}
	return ""
}

func (m *Selection) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Selection) GetCompetitorId() string {
	if m != nil {
		return m.CompetitorId
	}
	return ""
}

func (m *Selection) GetSourceCompetitorId() string {
	if m != nil {
		return m.SourceCompetitorId
	}
	return ""
}

func (m *Selection) GetRaceId() string {
	if m != nil {
		return m.RaceId
	}
	return ""
}

func (m *Selection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Selection) GetJockey() string {
	if m != nil {
		return m.Jockey
	}
	return ""
}

func (m *Selection) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Selection) GetBarrierNumber() int32 {
	if m != nil {
		return m.BarrierNumber
	}
	return 0
}

func (m *Selection) GetLastUpdated() int64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

type Race struct {
	RaceId               string   `protobuf:"bytes,1,opt,name=race_id,json=raceId" json:"race_id,omitempty"`
	SourceId             string   `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	MeetingId            string   `protobuf:"bytes,3,opt,name=meeting_id,json=meetingId" json:"meeting_id,omitempty"`
	Number               int32    `protobuf:"varint,4,opt,name=number" json:"number,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	ScheduledStart       int64    `protobuf:"varint,6,opt,name=scheduled_start,json=scheduledStart" json:"scheduled_start,omitempty"`
	ActualStart          int64    `protobuf:"varint,7,opt,name=actual_start,json=actualStart" json:"actual_start,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	Results              string   `protobuf:"bytes,9,opt,name=results" json:"results,omitempty"`
	MeetingStart         int64    `protobuf:"varint,10,opt,name=meeting_start,json=meetingStart" json:"meeting_start,omitempty"`
	DateCreated          int64    `protobuf:"varint,11,opt,name=date_created,json=dateCreated" json:"date_created,omitempty"`
	LastUpdated          int64    `protobuf:"varint,12,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Race) Reset()         { *m = Race{} }
func (m *Race) String() string { return proto.CompactTextString(m) }
func (*Race) ProtoMessage()    {}
func (*Race) Descriptor() ([]byte, []int) {
	return fileDescriptor_racing_869a59ac434a4990, []int{15}
}
func (m *Race) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Race.Unmarshal(m, b)
}
func (m *Race) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Race.Marshal(b, m, deterministic)
}
func (dst *Race) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Race.Merge(dst, src)
}
func (m *Race) XXX_Size() int {
	return xxx_messageInfo_Race.Size(m)
}
func (m *Race) XXX_DiscardUnknown() {
	xxx_messageInfo_Race.DiscardUnknown(m)
}

var xxx_messageInfo_Race proto.InternalMessageInfo

func (m *Race) GetRaceId() string {
	if m != nil {
		return m.RaceId
	}
	return ""
}

func (m *Race) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Race) GetMeetingId() string {
	if m != nil {
		return m.MeetingId
	}
	return ""
}

func (m *Race) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Race) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Race) GetScheduledStart() int64 {
	if m != nil {
		return m.ScheduledStart
	}
	return 0
}

func (m *Race) GetActualStart() int64 {
	if m != nil {
		return m.ActualStart
	}
	return 0
}

func (m *Race) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Race) GetResults() string {
	if m != nil {
		return m.Results
	}
	return ""
}

func (m *Race) GetMeetingStart() int64 {
	if m != nil {
		return m.MeetingStart
	}
	return 0
}

func (m *Race) GetDateCreated() int64 {
	if m != nil {
		return m.DateCreated
	}
	return 0
}

func (m *Race) GetLastUpdated() int64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

func init() {
	proto.RegisterType((*ListMeetingsByDateRequest)(nil), "racing.ListMeetingsByDateRequest")
	proto.RegisterType((*ListMeetingsByDateResponse)(nil), "racing.ListMeetingsByDateResponse")
	proto.RegisterType((*ListRacesByMeetingDateRequest)(nil), "racing.ListRacesByMeetingDateRequest")
	proto.RegisterType((*ListRacesByMeetingDateResponse)(nil), "racing.ListRacesByMeetingDateResponse")
	proto.RegisterType((*AddRacesRequest)(nil), "racing.AddRacesRequest")
	proto.RegisterType((*AddRacesResponse)(nil), "racing.AddRacesResponse")
	proto.RegisterType((*AddMeetingsRequest)(nil), "racing.AddMeetingsRequest")
	proto.RegisterType((*AddMeetingsResponse)(nil), "racing.AddMeetingsResponse")
	proto.RegisterType((*UpdateRaceRequest)(nil), "racing.UpdateRaceRequest")
	proto.RegisterType((*UpdateRaceResponse)(nil), "racing.UpdateRaceResponse")
	proto.RegisterType((*GetNextRaceRequest)(nil), "racing.GetNextRaceRequest")
	proto.RegisterType((*GetNextRaceResponse)(nil), "racing.GetNextRaceResponse")
	proto.RegisterType((*Meeting)(nil), "racing.Meeting")
	proto.RegisterType((*Competitor)(nil), "racing.Competitor")
	proto.RegisterType((*Selection)(nil), "racing.Selection")
	proto.RegisterType((*Race)(nil), "racing.Race")
}

func init() { proto.RegisterFile("racing.proto", fileDescriptor_racing_869a59ac434a4990) }

var fileDescriptor_racing_869a59ac434a4990 = []byte{
	// 871 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0x27, 0xff, 0xec, 0x64, 0xe2, 0xf4, 0xe8, 0x16, 0xa5, 0x3e, 0x57, 0x45, 0x89, 0x51, 0xa1,
	0x12, 0xe8, 0x80, 0xf6, 0x81, 0x27, 0x1e, 0xee, 0x52, 0x15, 0x45, 0x82, 0x22, 0xf9, 0x9a, 0x07,
	0x1e, 0x90, 0xe5, 0xd8, 0x73, 0x39, 0x43, 0x62, 0x87, 0xdd, 0x35, 0x22, 0x1f, 0x91, 0x77, 0xbe,
	0x06, 0x82, 0x8f, 0x80, 0xf6, 0x9f, 0xed, 0x9c, 0x93, 0x0b, 0xa0, 0xbe, 0x65, 0x7e, 0x33, 0x3b,
	0xf3, 0x9b, 0x9f, 0x77, 0x76, 0x02, 0x0e, 0x8d, 0xe2, 0x34, 0x5b, 0x5d, 0x6c, 0x69, 0xce, 0x73,
	0x62, 0x29, 0xcb, 0x5f, 0xc0, 0xf9, 0xb7, 0x29, 0xe3, 0xdf, 0x21, 0xf2, 0x34, 0x5b, 0xb1, 0xab,
	0xdd, 0xab, 0x88, 0x63, 0x80, 0xbf, 0x14, 0xc8, 0x38, 0x79, 0x0a, 0xc0, 0x78, 0x44, 0x79, 0x98,
	0x44, 0x1c, 0xdd, 0xd6, 0xa4, 0xf5, 0xbc, 0x13, 0x0c, 0x24, 0x22, 0xa2, 0xc8, 0x39, 0xf4, 0x31,
	0x4b, 0x94, 0xb3, 0x2d, 0x9d, 0x36, 0x66, 0x89, 0x70, 0xf9, 0x73, 0xf0, 0x0e, 0xa5, 0x65, 0xdb,
	0x3c, 0x63, 0x48, 0x3e, 0x85, 0xfe, 0x46, 0x7b, 0xdc, 0xd6, 0xa4, 0xf3, 0x7c, 0xf8, 0xe2, 0xec,
	0x42, 0xb3, 0xd3, 0x27, 0x82, 0x32, 0xc0, 0xff, 0x01, 0x9e, 0x8a, 0x54, 0x41, 0x14, 0x23, 0xbb,
	0xda, 0x69, 0xff, 0xbb, 0x61, 0xf9, 0x0a, 0x3e, 0x3c, 0x96, 0x5a, 0x33, 0xf5, 0xa1, 0x47, 0x85,
	0x57, 0xd3, 0x74, 0x0c, 0x4d, 0x71, 0x24, 0x50, 0x2e, 0xff, 0x2d, 0x9c, 0x5d, 0x26, 0x89, 0x4c,
	0x52, 0xa3, 0xa4, 0xf9, 0x87, 0x69, 0x22, 0x29, 0x0d, 0x82, 0x81, 0x46, 0xe6, 0x49, 0x95, 0xb5,
	0x7d, 0x3c, 0xeb, 0x67, 0xf0, 0x7e, 0x95, 0x55, 0xb3, 0x71, 0xc1, 0x8e, 0x29, 0x46, 0x1c, 0x55,
	0xce, 0x7e, 0x60, 0x4c, 0xff, 0x12, 0xc8, 0x65, 0x92, 0x18, 0xb9, 0x0d, 0x8d, 0xff, 0xa4, 0xf3,
	0xe7, 0xf0, 0x68, 0x2f, 0xc5, 0xc9, 0x9a, 0xb7, 0xf0, 0x70, 0xb1, 0x15, 0xb2, 0x4a, 0xda, 0xba,
	0xe4, 0x04, 0xba, 0x82, 0xbf, 0x8c, 0xbd, 0xdb, 0x99, 0xf4, 0x90, 0x2f, 0x01, 0x18, 0xae, 0x31,
	0xe6, 0x69, 0x9e, 0x19, 0x05, 0x1e, 0x9a, 0xb8, 0x6b, 0xe3, 0x09, 0x6a, 0x41, 0xfe, 0x05, 0x90,
	0x7a, 0xa5, 0x8a, 0x59, 0x21, 0xd1, 0x92, 0x99, 0x36, 0xfd, 0x97, 0x40, 0xbe, 0x41, 0xfe, 0x06,
	0x7f, 0xe3, 0x75, 0x6a, 0xf7, 0x7f, 0x14, 0xff, 0x2b, 0x78, 0xb4, 0x77, 0x48, 0x57, 0x39, 0xd9,
	0x90, 0xff, 0x47, 0x0b, 0x6c, 0x2d, 0xdb, 0xa9, 0x0f, 0xff, 0x04, 0x06, 0x2c, 0x2f, 0x68, 0x8c,
	0xc2, 0xdb, 0x96, 0xde, 0xbe, 0x02, 0xe6, 0x09, 0x21, 0xd0, 0xcd, 0xa2, 0x0d, 0xba, 0x1d, 0x89,
	0xcb, 0xdf, 0x52, 0xfd, 0xbc, 0xc8, 0x38, 0xdd, 0xb9, 0x5d, 0x09, 0x1b, 0x53, 0xa4, 0x12, 0xd5,
	0x43, 0xbe, 0xdb, 0xa2, 0xdb, 0x53, 0xa9, 0x04, 0xf0, 0x76, 0xb7, 0x45, 0xf2, 0x09, 0x9c, 0xb1,
	0xf8, 0x16, 0x93, 0x62, 0x8d, 0x49, 0x28, 0x47, 0xc1, 0xb5, 0xe4, 0xd5, 0x7f, 0x50, 0xc2, 0xd7,
	0x02, 0x15, 0xc3, 0x21, 0xb3, 0xa4, 0x09, 0x73, 0xed, 0x49, 0x47, 0x14, 0x10, 0xf6, 0x3c, 0x61,
	0xfe, 0x5f, 0x2d, 0x80, 0x59, 0xbe, 0xd9, 0x22, 0x4f, 0x79, 0x4e, 0xc9, 0x47, 0x30, 0x8a, 0x4b,
	0xab, 0x6a, 0xce, 0xa9, 0xc0, 0xff, 0xd3, 0x9f, 0x0f, 0x23, 0xf1, 0xc9, 0xc2, 0xfc, 0x26, 0x5c,
	0xa6, 0x94, 0xdf, 0xea, 0x2e, 0x87, 0x02, 0xfc, 0xfe, 0xe6, 0x4a, 0x40, 0x75, 0x0d, 0x7a, 0xfb,
	0x1a, 0xb8, 0x60, 0x73, 0x1a, 0xa5, 0x19, 0x52, 0xd9, 0xde, 0x20, 0x30, 0x26, 0x19, 0x83, 0xb5,
	0xc2, 0x2c, 0x41, 0xea, 0xda, 0xd2, 0xa1, 0x2d, 0x41, 0x30, 0xdd, 0x44, 0x2b, 0x0c, 0x0b, 0xba,
	0x76, 0xfb, 0x8a, 0xa0, 0x04, 0x16, 0x74, 0xed, 0xff, 0xde, 0x86, 0x41, 0x79, 0x01, 0xc9, 0x14,
	0x9c, 0xf2, 0x0a, 0x56, 0xfd, 0x0e, 0x4b, 0xec, 0x54, 0xbb, 0x0d, 0xc1, 0x3a, 0x07, 0x04, 0xfb,
	0x02, 0x3e, 0xd0, 0x19, 0xf6, 0x63, 0x95, 0x0c, 0x44, 0xf9, 0x66, 0xf5, 0x13, 0x8f, 0xc1, 0xd6,
	0x5f, 0x4c, 0xab, 0x61, 0xa9, 0x0f, 0x56, 0xca, 0x6b, 0xd5, 0xe4, 0x1d, 0x83, 0xf5, 0x53, 0x1e,
	0xff, 0x8c, 0x3b, 0x23, 0x83, 0xb2, 0x04, 0x9e, 0x15, 0x9b, 0x25, 0x52, 0xa9, 0x41, 0x2f, 0xd0,
	0x16, 0x79, 0x06, 0x0f, 0x96, 0x11, 0xa5, 0x29, 0xd2, 0x50, 0xfb, 0x07, 0xd2, 0x3f, 0xd2, 0xe8,
	0x1b, 0x15, 0x36, 0x05, 0x67, 0x1d, 0x31, 0x1e, 0x9a, 0xf1, 0x03, 0x79, 0xb7, 0x86, 0x02, 0x5b,
	0xe8, 0x11, 0xfc, 0xbb, 0x0d, 0x5d, 0x31, 0x23, 0x75, 0xbe, 0xad, 0x3d, 0xbe, 0xf7, 0x8a, 0xb7,
	0x3f, 0x47, 0x9d, 0xbb, 0x73, 0x54, 0xf1, 0xef, 0xee, 0xf1, 0x37, 0x1a, 0xf4, 0x6a, 0x1a, 0xfc,
	0xeb, 0x59, 0x98, 0x82, 0x13, 0xc5, 0xbc, 0x88, 0xd6, 0x3a, 0xca, 0x56, 0x5d, 0x29, 0x4c, 0x85,
	0x8c, 0xc1, 0x62, 0x3c, 0xe2, 0x05, 0xd3, 0x77, 0x47, 0x5b, 0xe2, 0x22, 0x52, 0x64, 0xc5, 0x9a,
	0x33, 0x29, 0x98, 0x98, 0x22, 0x65, 0x8a, 0x5b, 0x60, 0x1a, 0x51, 0x59, 0x95, 0x56, 0x8e, 0x06,
	0xcb, 0xca, 0x72, 0x0a, 0xcc, 0x43, 0x3b, 0x54, 0x95, 0x05, 0x36, 0x53, 0x50, 0x43, 0x72, 0xa7,
	0x21, 0xf9, 0x8b, 0x3f, 0x3b, 0x30, 0x0a, 0xe4, 0xeb, 0x74, 0x8d, 0xf4, 0xd7, 0x34, 0x46, 0xf2,
	0x23, 0x90, 0xe6, 0x16, 0x26, 0x53, 0xf3, 0x86, 0x1d, 0x5d, 0xfc, 0x9e, 0x7f, 0x5f, 0x88, 0x7a,
	0x18, 0xfd, 0xf7, 0x48, 0x0a, 0xe3, 0xc3, 0xeb, 0x93, 0x3c, 0xab, 0x9f, 0x3f, 0xba, 0xb9, 0xbd,
	0x8f, 0x4f, 0x85, 0x95, 0xa5, 0x5e, 0xc3, 0xb0, 0xb6, 0x9c, 0x88, 0x67, 0x0e, 0x36, 0x97, 0x9e,
	0xf7, 0xe4, 0xa0, 0x4f, 0xbf, 0xe6, 0x5f, 0x43, 0xdf, 0x6c, 0x55, 0xf2, 0xb8, 0x16, 0x58, 0xdf,
	0xde, 0x9e, 0xdb, 0x74, 0xe8, 0xe3, 0x33, 0x80, 0x6a, 0x11, 0x91, 0x73, 0x13, 0xd7, 0x58, 0x83,
	0x9e, 0x77, 0xc8, 0xa5, 0x93, 0xbc, 0x86, 0x61, 0x6d, 0xd1, 0x54, 0xbd, 0x34, 0x57, 0x56, 0xd5,
	0xcb, 0x81, 0xcd, 0xb4, 0xb4, 0xe4, 0x3f, 0xb9, 0x97, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x43, 0x51, 0x48, 0xd9, 0x09, 0x00, 0x00,
}
