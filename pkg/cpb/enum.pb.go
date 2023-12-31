// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: enum.proto

// proto 包名

package cpb

import (
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

type EitherStatusEnum int32

const (
	EitherStatusEnum_False EitherStatusEnum = 0
	EitherStatusEnum_True  EitherStatusEnum = 1
)

// Enum value maps for EitherStatusEnum.
var (
	EitherStatusEnum_name = map[int32]string{
		0: "False",
		1: "True",
	}
	EitherStatusEnum_value = map[string]int32{
		"False": 0,
		"True":  1,
	}
)

func (x EitherStatusEnum) Enum() *EitherStatusEnum {
	p := new(EitherStatusEnum)
	*p = x
	return p
}

func (x EitherStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EitherStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (EitherStatusEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x EitherStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EitherStatusEnum.Descriptor instead.
func (EitherStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

type CurrencyEnum int32

const (
	CurrencyEnum_CurrencyUnknow CurrencyEnum = 0
	// 云豆
	CurrencyEnum_CurrencyBean CurrencyEnum = 1
	// 银币
	CurrencyEnum_CurrencySilver CurrencyEnum = 2
	// 人民币
	CurrencyEnum_CurrencyMoney CurrencyEnum = 3
)

// Enum value maps for CurrencyEnum.
var (
	CurrencyEnum_name = map[int32]string{
		0: "CurrencyUnknow",
		1: "CurrencyBean",
		2: "CurrencySilver",
		3: "CurrencyMoney",
	}
	CurrencyEnum_value = map[string]int32{
		"CurrencyUnknow": 0,
		"CurrencyBean":   1,
		"CurrencySilver": 2,
		"CurrencyMoney":  3,
	}
)

func (x CurrencyEnum) Enum() *CurrencyEnum {
	p := new(CurrencyEnum)
	*p = x
	return p
}

func (x CurrencyEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CurrencyEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[1].Descriptor()
}

func (CurrencyEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[1]
}

func (x CurrencyEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CurrencyEnum.Descriptor instead.
func (CurrencyEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{1}
}

// 资产操作类型
type OperatingCurrencyEnum int32

const (
	OperatingCurrencyEnum_OperatingCurrencyUnknow OperatingCurrencyEnum = 0
	OperatingCurrencyEnum_OperatingCurrencyIn     OperatingCurrencyEnum = 1
	OperatingCurrencyEnum_OperatingCurrencyOut    OperatingCurrencyEnum = 2
)

// Enum value maps for OperatingCurrencyEnum.
var (
	OperatingCurrencyEnum_name = map[int32]string{
		0: "OperatingCurrencyUnknow",
		1: "OperatingCurrencyIn",
		2: "OperatingCurrencyOut",
	}
	OperatingCurrencyEnum_value = map[string]int32{
		"OperatingCurrencyUnknow": 0,
		"OperatingCurrencyIn":     1,
		"OperatingCurrencyOut":    2,
	}
)

func (x OperatingCurrencyEnum) Enum() *OperatingCurrencyEnum {
	p := new(OperatingCurrencyEnum)
	*p = x
	return p
}

func (x OperatingCurrencyEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatingCurrencyEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[2].Descriptor()
}

func (OperatingCurrencyEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[2]
}

func (x OperatingCurrencyEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatingCurrencyEnum.Descriptor instead.
func (OperatingCurrencyEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{2}
}

type ActiveLevelEnum int32

const (
	ActiveLevelEnum_unknow         ActiveLevelEnum = 0
	ActiveLevelEnum_reward_source  ActiveLevelEnum = 1
	ActiveLevelEnum_bind_new_user  ActiveLevelEnum = 2
	ActiveLevelEnum_watch_advert   ActiveLevelEnum = 3
	ActiveLevelEnum_buy_vip        ActiveLevelEnum = 4
	ActiveLevelEnum_compound_build ActiveLevelEnum = 5
	ActiveLevelEnum_login          ActiveLevelEnum = 6
)

// Enum value maps for ActiveLevelEnum.
var (
	ActiveLevelEnum_name = map[int32]string{
		0: "unknow",
		1: "reward_source",
		2: "bind_new_user",
		3: "watch_advert",
		4: "buy_vip",
		5: "compound_build",
		6: "login",
	}
	ActiveLevelEnum_value = map[string]int32{
		"unknow":         0,
		"reward_source":  1,
		"bind_new_user":  2,
		"watch_advert":   3,
		"buy_vip":        4,
		"compound_build": 5,
		"login":          6,
	}
)

func (x ActiveLevelEnum) Enum() *ActiveLevelEnum {
	p := new(ActiveLevelEnum)
	*p = x
	return p
}

func (x ActiveLevelEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActiveLevelEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[3].Descriptor()
}

func (ActiveLevelEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[3]
}

func (x ActiveLevelEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActiveLevelEnum.Descriptor instead.
func (ActiveLevelEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{3}
}

type WinPrizeRewardSourceTypeEnum int32

const (
	WinPrizeRewardSourceTypeEnum_WinPrizeRewardSourceTypeEnumUnknow   WinPrizeRewardSourceTypeEnum = 0
	WinPrizeRewardSourceTypeEnum_WinPrizeRewardSourceTypeEnumBean     WinPrizeRewardSourceTypeEnum = 1
	WinPrizeRewardSourceTypeEnum_WinPrizeRewardSourceTypeEnumFragment WinPrizeRewardSourceTypeEnum = 2
)

// Enum value maps for WinPrizeRewardSourceTypeEnum.
var (
	WinPrizeRewardSourceTypeEnum_name = map[int32]string{
		0: "WinPrizeRewardSourceTypeEnumUnknow",
		1: "WinPrizeRewardSourceTypeEnumBean",
		2: "WinPrizeRewardSourceTypeEnumFragment",
	}
	WinPrizeRewardSourceTypeEnum_value = map[string]int32{
		"WinPrizeRewardSourceTypeEnumUnknow":   0,
		"WinPrizeRewardSourceTypeEnumBean":     1,
		"WinPrizeRewardSourceTypeEnumFragment": 2,
	}
)

func (x WinPrizeRewardSourceTypeEnum) Enum() *WinPrizeRewardSourceTypeEnum {
	p := new(WinPrizeRewardSourceTypeEnum)
	*p = x
	return p
}

func (x WinPrizeRewardSourceTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WinPrizeRewardSourceTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[4].Descriptor()
}

func (WinPrizeRewardSourceTypeEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[4]
}

func (x WinPrizeRewardSourceTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WinPrizeRewardSourceTypeEnum.Descriptor instead.
func (WinPrizeRewardSourceTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{4}
}

// 用户删除状态
type UserDeleteStatusEnum int32

const (
	UserDeleteStatusEnum_UserDeleteStatusEnumNormal UserDeleteStatusEnum = 0 // 正常
	UserDeleteStatusEnum_UserDeleteStatusEnumDelete UserDeleteStatusEnum = 1 // 删除
)

// Enum value maps for UserDeleteStatusEnum.
var (
	UserDeleteStatusEnum_name = map[int32]string{
		0: "UserDeleteStatusEnumNormal",
		1: "UserDeleteStatusEnumDelete",
	}
	UserDeleteStatusEnum_value = map[string]int32{
		"UserDeleteStatusEnumNormal": 0,
		"UserDeleteStatusEnumDelete": 1,
	}
)

func (x UserDeleteStatusEnum) Enum() *UserDeleteStatusEnum {
	p := new(UserDeleteStatusEnum)
	*p = x
	return p
}

func (x UserDeleteStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserDeleteStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[5].Descriptor()
}

func (UserDeleteStatusEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[5]
}

func (x UserDeleteStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserDeleteStatusEnum.Descriptor instead.
func (UserDeleteStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{5}
}

type AccountbookEnum int32

const (
	AccountbookEnum_AccountbookEnumUnknow      AccountbookEnum = 0
	AccountbookEnum_AccountbookEnumLotteryType AccountbookEnum = 1
)

// Enum value maps for AccountbookEnum.
var (
	AccountbookEnum_name = map[int32]string{
		0: "AccountbookEnumUnknow",
		1: "AccountbookEnumLotteryType",
	}
	AccountbookEnum_value = map[string]int32{
		"AccountbookEnumUnknow":      0,
		"AccountbookEnumLotteryType": 1,
	}
)

func (x AccountbookEnum) Enum() *AccountbookEnum {
	p := new(AccountbookEnum)
	*p = x
	return p
}

func (x AccountbookEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountbookEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[6].Descriptor()
}

func (AccountbookEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[6]
}

func (x AccountbookEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountbookEnum.Descriptor instead.
func (AccountbookEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{6}
}

type SysUserStatusEnum int32

const (
	SysUserStatusEnum_SysUserStatusEnumUnknow SysUserStatusEnum = 0
	SysUserStatusEnum_SysUserStatusEnumNormal SysUserStatusEnum = 1
	SysUserStatusEnum_SysUserStatusEnumLock   SysUserStatusEnum = 2
)

// Enum value maps for SysUserStatusEnum.
var (
	SysUserStatusEnum_name = map[int32]string{
		0: "SysUserStatusEnumUnknow",
		1: "SysUserStatusEnumNormal",
		2: "SysUserStatusEnumLock",
	}
	SysUserStatusEnum_value = map[string]int32{
		"SysUserStatusEnumUnknow": 0,
		"SysUserStatusEnumNormal": 1,
		"SysUserStatusEnumLock":   2,
	}
)

func (x SysUserStatusEnum) Enum() *SysUserStatusEnum {
	p := new(SysUserStatusEnum)
	*p = x
	return p
}

func (x SysUserStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SysUserStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[7].Descriptor()
}

func (SysUserStatusEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[7]
}

func (x SysUserStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SysUserStatusEnum.Descriptor instead.
func (SysUserStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{7}
}

// 通用开关状态枚举
type StatusSwitchEnum int32

const (
	StatusSwitchEnum_StatusSwitch_Close StatusSwitchEnum = 0 //关闭
	StatusSwitchEnum_StatusSwitch_Open  StatusSwitchEnum = 1 // 开启
)

// Enum value maps for StatusSwitchEnum.
var (
	StatusSwitchEnum_name = map[int32]string{
		0: "StatusSwitch_Close",
		1: "StatusSwitch_Open",
	}
	StatusSwitchEnum_value = map[string]int32{
		"StatusSwitch_Close": 0,
		"StatusSwitch_Open":  1,
	}
)

func (x StatusSwitchEnum) Enum() *StatusSwitchEnum {
	p := new(StatusSwitchEnum)
	*p = x
	return p
}

func (x StatusSwitchEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusSwitchEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[8].Descriptor()
}

func (StatusSwitchEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[8]
}

func (x StatusSwitchEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusSwitchEnum.Descriptor instead.
func (StatusSwitchEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{8}
}

// 通用是否状态枚举
type StatusOrEnum int32

const (
	StatusOrEnum_StatusOr_Flase StatusOrEnum = 0 //否
	StatusOrEnum_StatusOr_True  StatusOrEnum = 1 // 是
)

// Enum value maps for StatusOrEnum.
var (
	StatusOrEnum_name = map[int32]string{
		0: "StatusOr_Flase",
		1: "StatusOr_True",
	}
	StatusOrEnum_value = map[string]int32{
		"StatusOr_Flase": 0,
		"StatusOr_True":  1,
	}
)

func (x StatusOrEnum) Enum() *StatusOrEnum {
	p := new(StatusOrEnum)
	*p = x
	return p
}

func (x StatusOrEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusOrEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[9].Descriptor()
}

func (StatusOrEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[9]
}

func (x StatusOrEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusOrEnum.Descriptor instead.
func (StatusOrEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{9}
}

// 系统菜单类型
type SysMenuTypeEnum int32

const (
	SysMenuTypeEnum_SysMenuType_Unknow SysMenuTypeEnum = 0
	SysMenuTypeEnum_SysMenuType_Dir    SysMenuTypeEnum = 1 // 目录
	SysMenuTypeEnum_SysMenuType_Menu   SysMenuTypeEnum = 2 // 菜单
	SysMenuTypeEnum_SysMenutype_Button SysMenuTypeEnum = 3 // 按钮
)

// Enum value maps for SysMenuTypeEnum.
var (
	SysMenuTypeEnum_name = map[int32]string{
		0: "SysMenuType_Unknow",
		1: "SysMenuType_Dir",
		2: "SysMenuType_Menu",
		3: "SysMenutype_Button",
	}
	SysMenuTypeEnum_value = map[string]int32{
		"SysMenuType_Unknow": 0,
		"SysMenuType_Dir":    1,
		"SysMenuType_Menu":   2,
		"SysMenutype_Button": 3,
	}
)

func (x SysMenuTypeEnum) Enum() *SysMenuTypeEnum {
	p := new(SysMenuTypeEnum)
	*p = x
	return p
}

func (x SysMenuTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SysMenuTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[10].Descriptor()
}

func (SysMenuTypeEnum) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[10]
}

func (x SysMenuTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SysMenuTypeEnum.Descriptor instead.
func (SysMenuTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{10}
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x70,
	0x62, 0x2a, 0x27, 0x0a, 0x10, 0x45, 0x69, 0x74, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x72, 0x75, 0x65, 0x10, 0x01, 0x2a, 0x5b, 0x0a, 0x0c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x65, 0x61, 0x6e, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x69, 0x6c, 0x76,
	0x65, 0x72, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x10, 0x03, 0x2a, 0x67, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x49, 0x6e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4f, 0x75, 0x74, 0x10, 0x02,
	0x2a, 0x81, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0a, 0x0a, 0x06, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x6e, 0x65, 0x77, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x62, 0x75, 0x79, 0x5f,
	0x76, 0x69, 0x70, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x10, 0x06, 0x2a, 0x96, 0x01, 0x0a, 0x1c, 0x57, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x22, 0x57, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x24, 0x0a,
	0x20, 0x57, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x61,
	0x6e, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x2a, 0x56, 0x0a,
	0x14, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x10, 0x01, 0x2a, 0x4c, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x75, 0x6d, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x6f,
	0x6f, 0x6b, 0x45, 0x6e, 0x75, 0x6d, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x01, 0x2a, 0x68, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x79, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x02, 0x2a, 0x41, 0x0a,
	0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x5f, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x4f, 0x70, 0x65, 0x6e, 0x10, 0x01,
	0x2a, 0x35, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x5f, 0x46, 0x6c, 0x61,
	0x73, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72,
	0x5f, 0x54, 0x72, 0x75, 0x65, 0x10, 0x01, 0x2a, 0x6c, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x4d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x79,
	0x73, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x44, 0x69, 0x72, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x4d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4d, 0x65, 0x6e, 0x75, 0x10, 0x02, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x10, 0x03, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x63, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_proto_rawDescOnce sync.Once
	file_enum_proto_rawDescData = file_enum_proto_rawDesc
)

func file_enum_proto_rawDescGZIP() []byte {
	file_enum_proto_rawDescOnce.Do(func() {
		file_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_proto_rawDescData)
	})
	return file_enum_proto_rawDescData
}

var file_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 11)
var file_enum_proto_goTypes = []interface{}{
	(EitherStatusEnum)(0),             // 0: cpb.EitherStatusEnum
	(CurrencyEnum)(0),                 // 1: cpb.CurrencyEnum
	(OperatingCurrencyEnum)(0),        // 2: cpb.OperatingCurrencyEnum
	(ActiveLevelEnum)(0),              // 3: cpb.ActiveLevelEnum
	(WinPrizeRewardSourceTypeEnum)(0), // 4: cpb.WinPrizeRewardSourceTypeEnum
	(UserDeleteStatusEnum)(0),         // 5: cpb.UserDeleteStatusEnum
	(AccountbookEnum)(0),              // 6: cpb.AccountbookEnum
	(SysUserStatusEnum)(0),            // 7: cpb.SysUserStatusEnum
	(StatusSwitchEnum)(0),             // 8: cpb.StatusSwitchEnum
	(StatusOrEnum)(0),                 // 9: cpb.StatusOrEnum
	(SysMenuTypeEnum)(0),              // 10: cpb.SysMenuTypeEnum
}
var file_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      11,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}
