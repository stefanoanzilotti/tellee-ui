// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Source
	Dashboard
	DashboardCell
	Axis
	Template
	TemplateValue
	TemplateQuery
	Server
	Layout
	Cell
	Query
	Range
	AlertRule
	User
*/
package internal

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	ID                 int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type               string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Username           string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	URL                string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Default            bool   `protobuf:"varint,7,opt,name=Default,proto3" json:"Default,omitempty"`
	Telegraf           string `protobuf:"bytes,8,opt,name=Telegraf,proto3" json:"Telegraf,omitempty"`
	InsecureSkipVerify bool   `protobuf:"varint,9,opt,name=InsecureSkipVerify,proto3" json:"InsecureSkipVerify,omitempty"`
	MetaURL            string `protobuf:"bytes,10,opt,name=MetaURL,proto3" json:"MetaURL,omitempty"`
	SharedSecret       string `protobuf:"bytes,11,opt,name=SharedSecret,proto3" json:"SharedSecret,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

type Dashboard struct {
	ID        int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string           `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Cells     []*DashboardCell `protobuf:"bytes,3,rep,name=cells" json:"cells,omitempty"`
	Templates []*Template      `protobuf:"bytes,4,rep,name=templates" json:"templates,omitempty"`
}

func (m *Dashboard) Reset()                    { *m = Dashboard{} }
func (m *Dashboard) String() string            { return proto.CompactTextString(m) }
func (*Dashboard) ProtoMessage()               {}
func (*Dashboard) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Dashboard) GetCells() []*DashboardCell {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Dashboard) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

type DashboardCell struct {
	X       int32            `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32            `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32            `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32            `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query         `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	Name    string           `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Type    string           `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	ID      string           `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
	Axes    map[string]*Axis `protobuf:"bytes,9,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DashboardCell) Reset()                    { *m = DashboardCell{} }
func (m *DashboardCell) String() string            { return proto.CompactTextString(m) }
func (*DashboardCell) ProtoMessage()               {}
func (*DashboardCell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *DashboardCell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *DashboardCell) GetAxes() map[string]*Axis {
	if m != nil {
		return m.Axes
	}
	return nil
}

type Axis struct {
	LegacyBounds []int64  `protobuf:"varint,1,rep,name=legacyBounds" json:"legacyBounds,omitempty"`
	Bounds       []string `protobuf:"bytes,2,rep,name=bounds" json:"bounds,omitempty"`
	Label        string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Prefix       string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix       string   `protobuf:"bytes,5,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Base         string   `protobuf:"bytes,6,opt,name=base,proto3" json:"base,omitempty"`
	Scale        string   `protobuf:"bytes,7,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (m *Axis) Reset()                    { *m = Axis{} }
func (m *Axis) String() string            { return proto.CompactTextString(m) }
func (*Axis) ProtoMessage()               {}
func (*Axis) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

type Template struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TempVar string           `protobuf:"bytes,2,opt,name=temp_var,json=tempVar,proto3" json:"temp_var,omitempty"`
	Values  []*TemplateValue `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Type    string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Label   string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Query   *TemplateQuery   `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Template) GetValues() []*TemplateValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Template) GetQuery() *TemplateQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type TemplateValue struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (m *TemplateValue) Reset()                    { *m = TemplateValue{} }
func (m *TemplateValue) String() string            { return proto.CompactTextString(m) }
func (*TemplateValue) ProtoMessage()               {}
func (*TemplateValue) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

type TemplateQuery struct {
	Command     string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Db          string `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
	Rp          string `protobuf:"bytes,3,opt,name=rp,proto3" json:"rp,omitempty"`
	Measurement string `protobuf:"bytes,4,opt,name=measurement,proto3" json:"measurement,omitempty"`
	TagKey      string `protobuf:"bytes,5,opt,name=tag_key,json=tagKey,proto3" json:"tag_key,omitempty"`
	FieldKey    string `protobuf:"bytes,6,opt,name=field_key,json=fieldKey,proto3" json:"field_key,omitempty"`
}

func (m *TemplateQuery) Reset()                    { *m = TemplateQuery{} }
func (m *TemplateQuery) String() string            { return proto.CompactTextString(m) }
func (*TemplateQuery) ProtoMessage()               {}
func (*TemplateQuery) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

type Server struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	URL      string `protobuf:"bytes,5,opt,name=URL,proto3" json:"URL,omitempty"`
	SrcID    int64  `protobuf:"varint,6,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	Active   bool   `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

type Layout struct {
	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Application string  `protobuf:"bytes,2,opt,name=Application,proto3" json:"Application,omitempty"`
	Measurement string  `protobuf:"bytes,3,opt,name=Measurement,proto3" json:"Measurement,omitempty"`
	Cells       []*Cell `protobuf:"bytes,4,rep,name=Cells" json:"Cells,omitempty"`
	Autoflow    bool    `protobuf:"varint,5,opt,name=Autoflow,proto3" json:"Autoflow,omitempty"`
}

func (m *Layout) Reset()                    { *m = Layout{} }
func (m *Layout) String() string            { return proto.CompactTextString(m) }
func (*Layout) ProtoMessage()               {}
func (*Layout) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{8} }

func (m *Layout) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Cell struct {
	X       int32            `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32            `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32            `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32            `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query         `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	I       string           `protobuf:"bytes,6,opt,name=i,proto3" json:"i,omitempty"`
	Name    string           `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Yranges []int64          `protobuf:"varint,8,rep,name=yranges" json:"yranges,omitempty"`
	Ylabels []string         `protobuf:"bytes,9,rep,name=ylabels" json:"ylabels,omitempty"`
	Type    string           `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Axes    map[string]*Axis `protobuf:"bytes,11,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{9} }

func (m *Cell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *Cell) GetAxes() map[string]*Axis {
	if m != nil {
		return m.Axes
	}
	return nil
}

type Query struct {
	Command  string   `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	DB       string   `protobuf:"bytes,2,opt,name=DB,proto3" json:"DB,omitempty"`
	RP       string   `protobuf:"bytes,3,opt,name=RP,proto3" json:"RP,omitempty"`
	GroupBys []string `protobuf:"bytes,4,rep,name=GroupBys" json:"GroupBys,omitempty"`
	Wheres   []string `protobuf:"bytes,5,rep,name=Wheres" json:"Wheres,omitempty"`
	Label    string   `protobuf:"bytes,6,opt,name=Label,proto3" json:"Label,omitempty"`
	Range    *Range   `protobuf:"bytes,7,opt,name=Range" json:"Range,omitempty"`
	Source   string   `protobuf:"bytes,8,opt,name=Source,proto3" json:"Source,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{10} }

func (m *Query) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type Range struct {
	Upper int64 `protobuf:"varint,1,opt,name=Upper,proto3" json:"Upper,omitempty"`
	Lower int64 `protobuf:"varint,2,opt,name=Lower,proto3" json:"Lower,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{11} }

type AlertRule struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JSON   string `protobuf:"bytes,2,opt,name=JSON,proto3" json:"JSON,omitempty"`
	SrcID  int64  `protobuf:"varint,3,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	KapaID int64  `protobuf:"varint,4,opt,name=KapaID,proto3" json:"KapaID,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{12} }

type User struct {
	ID       uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Provider string `protobuf:"bytes,3,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Scheme   string `protobuf:"bytes,4,opt,name=Scheme,proto3" json:"Scheme,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{13} }

func init() {
	proto.RegisterType((*Source)(nil), "internal.Source")
	proto.RegisterType((*Dashboard)(nil), "internal.Dashboard")
	proto.RegisterType((*DashboardCell)(nil), "internal.DashboardCell")
	proto.RegisterType((*Axis)(nil), "internal.Axis")
	proto.RegisterType((*Template)(nil), "internal.Template")
	proto.RegisterType((*TemplateValue)(nil), "internal.TemplateValue")
	proto.RegisterType((*TemplateQuery)(nil), "internal.TemplateQuery")
	proto.RegisterType((*Server)(nil), "internal.Server")
	proto.RegisterType((*Layout)(nil), "internal.Layout")
	proto.RegisterType((*Cell)(nil), "internal.Cell")
	proto.RegisterType((*Query)(nil), "internal.Query")
	proto.RegisterType((*Range)(nil), "internal.Range")
	proto.RegisterType((*AlertRule)(nil), "internal.AlertRule")
	proto.RegisterType((*User)(nil), "internal.User")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0xd6, 0xc4, 0x76, 0x12, 0xbf, 0x6c, 0x17, 0x34, 0xaa, 0xa8, 0x29, 0x97, 0x60, 0x81, 0x14,
	0x24, 0xba, 0xa0, 0x56, 0x48, 0x88, 0x5b, 0x76, 0x83, 0xaa, 0x65, 0xb7, 0x65, 0x99, 0xec, 0x2e,
	0x27, 0xa8, 0x26, 0xce, 0xcb, 0xc6, 0xaa, 0x13, 0x9b, 0xb1, 0xbd, 0x89, 0xff, 0x05, 0xbf, 0x00,
	0x09, 0x89, 0x13, 0x07, 0x0e, 0xfc, 0x01, 0xee, 0xfd, 0x55, 0xe8, 0xcd, 0x8c, 0x1d, 0x87, 0x2e,
	0x55, 0x2f, 0x70, 0x9b, 0xef, 0xbd, 0xc9, 0x9b, 0x99, 0xf7, 0xbd, 0xef, 0x8b, 0xe1, 0x30, 0x5e,
	0x17, 0xa8, 0xd6, 0x32, 0x39, 0xca, 0x54, 0x5a, 0xa4, 0xbc, 0x5f, 0xe3, 0xf0, 0x8f, 0x0e, 0x74,
	0xa7, 0x69, 0xa9, 0x22, 0xe4, 0x87, 0xd0, 0x39, 0x9d, 0x04, 0x6c, 0xc8, 0x46, 0x8e, 0xe8, 0x9c,
	0x4e, 0x38, 0x07, 0xf7, 0xb9, 0x5c, 0x61, 0xd0, 0x19, 0xb2, 0x91, 0x2f, 0xf4, 0x9a, 0x62, 0x97,
	0x55, 0x86, 0x81, 0x63, 0x62, 0xb4, 0xe6, 0x0f, 0xa1, 0x7f, 0x95, 0x53, 0xb5, 0x15, 0x06, 0xae,
	0x8e, 0x37, 0x98, 0x72, 0x17, 0x32, 0xcf, 0x37, 0xa9, 0x9a, 0x07, 0x9e, 0xc9, 0xd5, 0x98, 0xbf,
	0x0b, 0xce, 0x95, 0x38, 0x0f, 0xba, 0x3a, 0x4c, 0x4b, 0x1e, 0x40, 0x6f, 0x82, 0x0b, 0x59, 0x26,
	0x45, 0xd0, 0x1b, 0xb2, 0x51, 0x5f, 0xd4, 0x90, 0xea, 0x5c, 0x62, 0x82, 0x37, 0x4a, 0x2e, 0x82,
	0xbe, 0xa9, 0x53, 0x63, 0x7e, 0x04, 0xfc, 0x74, 0x9d, 0x63, 0x54, 0x2a, 0x9c, 0xbe, 0x8c, 0xb3,
	0x6b, 0x54, 0xf1, 0xa2, 0x0a, 0x7c, 0x5d, 0xe0, 0x8e, 0x0c, 0x9d, 0xf2, 0x0c, 0x0b, 0x49, 0x67,
	0x83, 0x2e, 0x55, 0x43, 0x1e, 0xc2, 0xc1, 0x74, 0x29, 0x15, 0xce, 0xa7, 0x18, 0x29, 0x2c, 0x82,
	0x81, 0x4e, 0xef, 0xc5, 0xc2, 0x9f, 0x19, 0xf8, 0x13, 0x99, 0x2f, 0x67, 0xa9, 0x54, 0xf3, 0xb7,
	0xea, 0xd9, 0x23, 0xf0, 0x22, 0x4c, 0x92, 0x3c, 0x70, 0x86, 0xce, 0x68, 0xf0, 0xf8, 0xc1, 0x51,
	0x43, 0x46, 0x53, 0xe7, 0x04, 0x93, 0x44, 0x98, 0x5d, 0xfc, 0x73, 0xf0, 0x0b, 0x5c, 0x65, 0x89,
	0x2c, 0x30, 0x0f, 0x5c, 0xfd, 0x13, 0xbe, 0xfb, 0xc9, 0xa5, 0x4d, 0x89, 0xdd, 0xa6, 0xf0, 0xf7,
	0x0e, 0xdc, 0xdb, 0x2b, 0xc5, 0x0f, 0x80, 0x6d, 0xf5, 0xad, 0x3c, 0xc1, 0xb6, 0x84, 0x2a, 0x7d,
	0x23, 0x4f, 0xb0, 0x8a, 0xd0, 0x46, 0xf3, 0xe7, 0x09, 0xb6, 0x21, 0xb4, 0xd4, 0xac, 0x79, 0x82,
	0x2d, 0xf9, 0x27, 0xd0, 0xfb, 0xa9, 0x44, 0x15, 0x63, 0x1e, 0x78, 0xfa, 0xe4, 0x77, 0x76, 0x27,
	0x7f, 0x57, 0xa2, 0xaa, 0x44, 0x9d, 0xa7, 0x97, 0x6a, 0xc6, 0x0d, 0x7d, 0x7a, 0x4d, 0xb1, 0x82,
	0xa6, 0xa3, 0x67, 0x62, 0xb4, 0xb6, 0x1d, 0x32, 0x9c, 0x51, 0x87, 0xbe, 0x00, 0x57, 0x6e, 0x31,
	0x0f, 0x7c, 0x5d, 0xff, 0xc3, 0x7f, 0x69, 0xc6, 0xd1, 0x78, 0x8b, 0xf9, 0xd7, 0xeb, 0x42, 0x55,
	0x42, 0x6f, 0x7f, 0xf8, 0x14, 0xfc, 0x26, 0x44, 0x93, 0xf3, 0x12, 0x2b, 0xfd, 0x40, 0x5f, 0xd0,
	0x92, 0x7f, 0x04, 0xde, 0xad, 0x4c, 0x4a, 0xd3, 0xf8, 0xc1, 0xe3, 0xc3, 0x5d, 0xd9, 0xf1, 0x36,
	0xce, 0x85, 0x49, 0x7e, 0xd5, 0xf9, 0x92, 0x85, 0x7f, 0x32, 0x70, 0x29, 0x46, 0x64, 0x27, 0x78,
	0x23, 0xa3, 0xea, 0x38, 0x2d, 0xd7, 0xf3, 0x3c, 0x60, 0x43, 0x67, 0xe4, 0x88, 0xbd, 0x18, 0x7f,
	0x0f, 0xba, 0x33, 0x93, 0xed, 0x0c, 0x9d, 0x91, 0x2f, 0x2c, 0xe2, 0xf7, 0xc1, 0x4b, 0xe4, 0x0c,
	0x13, 0xab, 0x03, 0x03, 0x68, 0x77, 0xa6, 0x70, 0x11, 0x6f, 0xad, 0x0c, 0x2c, 0xa2, 0x78, 0x5e,
	0x2e, 0x28, 0x6e, 0x24, 0x60, 0x11, 0xb5, 0x6b, 0x26, 0xf3, 0xa6, 0x85, 0xb4, 0xa6, 0xca, 0x79,
	0x24, 0x93, 0xba, 0x87, 0x06, 0x84, 0x7f, 0x31, 0x9a, 0x7f, 0xc3, 0x77, 0x6b, 0xe6, 0x4c, 0x47,
	0xdf, 0x87, 0x3e, 0xcd, 0xc2, 0x8b, 0x5b, 0xa9, 0xec, 0xdc, 0xf5, 0x08, 0x5f, 0x4b, 0xc5, 0x3f,
	0x83, 0xae, 0x7e, 0xf9, 0x1d, 0xb3, 0x57, 0x97, 0xbb, 0xa6, 0xbc, 0xb0, 0xdb, 0x1a, 0x06, 0xdd,
	0x16, 0x83, 0xcd, 0x63, 0xbd, 0xf6, 0x63, 0x1f, 0x81, 0x47, 0xa3, 0x50, 0xe9, 0xdb, 0xdf, 0x59,
	0xd9, 0x0c, 0x8c, 0xd9, 0x15, 0x5e, 0xc1, 0xbd, 0xbd, 0x13, 0x9b, 0x93, 0xd8, 0xfe, 0x49, 0x3b,
	0x16, 0x7d, 0xcb, 0x1a, 0x69, 0x3f, 0xc7, 0x04, 0xa3, 0x02, 0xe7, 0xba, 0xdf, 0x7d, 0xd1, 0xe0,
	0xf0, 0x57, 0xb6, 0xab, 0xab, 0xcf, 0x23, 0x75, 0x47, 0xe9, 0x6a, 0x25, 0xd7, 0x73, 0x5b, 0xba,
	0x86, 0xd4, 0xb7, 0xf9, 0xcc, 0x96, 0xee, 0xcc, 0x67, 0x84, 0x55, 0x66, 0x19, 0xec, 0xa8, 0x8c,
	0x0f, 0x61, 0xb0, 0x42, 0x99, 0x97, 0x0a, 0x57, 0xb8, 0x2e, 0x6c, 0x0b, 0xda, 0x21, 0xfe, 0x00,
	0x7a, 0x85, 0xbc, 0x79, 0x41, 0xb3, 0x67, 0x99, 0x2c, 0xe4, 0xcd, 0x19, 0x56, 0xfc, 0x03, 0xf0,
	0x17, 0x31, 0x26, 0x73, 0x9d, 0x32, 0x74, 0xf6, 0x75, 0xe0, 0x0c, 0xab, 0xf0, 0x37, 0x06, 0xdd,
	0x29, 0xaa, 0x5b, 0x54, 0x6f, 0x65, 0x17, 0x6d, 0x3b, 0x75, 0xde, 0x60, 0xa7, 0xee, 0xdd, 0x76,
	0xea, 0xed, 0xec, 0xf4, 0x3e, 0x78, 0x53, 0x15, 0x9d, 0x4e, 0xf4, 0x8d, 0x1c, 0x61, 0x00, 0x4d,
	0xe3, 0x38, 0x2a, 0xe2, 0x5b, 0xb4, 0x1e, 0x6b, 0x51, 0xf8, 0x0b, 0x83, 0xee, 0xb9, 0xac, 0xd2,
	0xb2, 0x78, 0x6d, 0xc2, 0x86, 0x30, 0x18, 0x67, 0x59, 0x12, 0x47, 0xb2, 0x88, 0xd3, 0xb5, 0xbd,
	0x6d, 0x3b, 0x44, 0x3b, 0x9e, 0xb5, 0x7a, 0x67, 0xee, 0xdd, 0x0e, 0x91, 0x42, 0x4f, 0xb4, 0x0b,
	0x1a, 0x4b, 0x6b, 0x29, 0xd4, 0x98, 0x9f, 0x4e, 0xd2, 0x03, 0xc7, 0x65, 0x91, 0x2e, 0x92, 0x74,
	0xa3, 0x5f, 0xd2, 0x17, 0x0d, 0x0e, 0x5f, 0x75, 0xc0, 0xfd, 0xbf, 0xdc, 0xed, 0x00, 0x58, 0x6c,
	0x89, 0x64, 0x71, 0xe3, 0x75, 0xbd, 0x96, 0xd7, 0x05, 0xd0, 0xab, 0x94, 0x5c, 0xdf, 0x60, 0x1e,
	0xf4, 0xb5, 0x73, 0xd4, 0x50, 0x67, 0xb4, 0x46, 0x8c, 0xc9, 0xf9, 0xa2, 0x86, 0xcd, 0xcc, 0x43,
	0x6b, 0xe6, 0x3f, 0xb5, 0x7e, 0x38, 0xd0, 0x37, 0x0a, 0xf6, 0xdb, 0xf2, 0xdf, 0xd9, 0xe0, 0x2b,
	0x06, 0x5e, 0x23, 0x98, 0x93, 0x7d, 0xc1, 0x9c, 0xec, 0x04, 0x33, 0x39, 0xae, 0x05, 0x33, 0x39,
	0x26, 0x2c, 0x2e, 0x6a, 0xc1, 0x88, 0x0b, 0x22, 0xeb, 0xa9, 0x4a, 0xcb, 0xec, 0xb8, 0x32, 0xac,
	0xfa, 0xa2, 0xc1, 0x34, 0x65, 0xdf, 0x2f, 0x51, 0xd9, 0x56, 0xfb, 0xc2, 0x22, 0x9a, 0xc9, 0x73,
	0x6d, 0x26, 0xa6, 0xb9, 0x06, 0xf0, 0x8f, 0xc1, 0x13, 0xd4, 0x3c, 0xdd, 0xe1, 0x3d, 0x5e, 0x74,
	0x58, 0x98, 0x2c, 0x15, 0x35, 0xdf, 0x2a, 0xf6, 0xff, 0xc4, 0xa2, 0xf0, 0x89, 0xfd, 0x39, 0x55,
	0xbf, 0xca, 0x32, 0x54, 0x56, 0x62, 0x06, 0xe8, 0x33, 0xd3, 0x0d, 0x1a, 0x77, 0x74, 0x84, 0x01,
	0xe1, 0x0f, 0xe0, 0x8f, 0x13, 0x54, 0x85, 0x28, 0x93, 0xd7, 0x3d, 0x95, 0x83, 0xfb, 0xcd, 0xf4,
	0xdb, 0xe7, 0xb5, 0x30, 0x69, 0xbd, 0x93, 0x93, 0xf3, 0x0f, 0x39, 0x9d, 0xc9, 0x4c, 0x9e, 0x4e,
	0xf4, 0x9c, 0x39, 0xc2, 0xa2, 0xf0, 0x47, 0x70, 0x49, 0xb6, 0xad, 0xca, 0xee, 0x9b, 0x24, 0x7f,
	0xa1, 0xd2, 0xdb, 0x78, 0x8e, 0xaa, 0x96, 0x7c, 0x8d, 0xf5, 0x9b, 0xa3, 0x25, 0x36, 0xdf, 0x56,
	0x16, 0xcd, 0xba, 0xfa, 0x4b, 0xee, 0xc9, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xcf, 0x99,
	0xa9, 0xdb, 0x09, 0x00, 0x00,
}
