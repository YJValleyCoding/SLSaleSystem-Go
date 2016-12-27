package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//首字母大写的表才能被外部引用
type Account struct {
	Id       int64
	User     *User `orm:"rel(fk)"`
	Money    float64
	MoneyBak float64
}

type AccountDetail struct {
	Id int64

	User           *User `orm:"rel(fk)"`
	DetailType     int64
	DetailTypeName string
	Money          float64
	AccountMoney   float64
	Memo           string
	DetailDateTime time.Time
}
type Admin struct {
	Id       int64
	UserName string
}

type Contacts struct {
	Id int64

	Customs      *Customs `orm:"rel(fk)"`
	ContactName  string
	ContactTel   string
	ContactFax   string
	ContactEmail string
	ContactRole  string
}
type Customs struct {
	Id int64

	Agent          *User `orm:"rel(fk)"`
	AgentName      string
	CustomName     string
	CustomType     int64
	CustomTypeName string
	SiteUrl        string
	CustomStatus   int64
	BossName       string
	CardType       int64
	CardTypeName   string
	CardNum        string
	CompanyTel     string
	CompanyFax     string
	RegDatetime    time.Time
	Country        string
	Province       string
	City           string
	Area           string
	CompanyAddress string
	Memo           string
	AgentCode      string
}
type Function struct {
	Id             int64
	FunctionCode   string
	FunctionName   string
	CreationTime   time.Time
	CreatedBy      string
	LastUpdateTime time.Time
	FuncUrl        string
	IsStart        int64
	ParentId       int64
}

type Keywords struct {
	Id        int64
	Keywords  string
	Agent     *User `orm:"rel(fk)"`
	AgentName string

	Custom             *Customs `orm:"rel(fk)"`
	CustomName         string
	PreRegFrozenMoney  float64
	Price              float64
	ProductType        int64
	ServiceYears       int64
	OpenApp            int64
	AppUserName        string
	AppPassword        string
	LoginUrl           string
	IosDownloadUrl     string
	AndroidDownloadUrl string
	CodeIosUrl         string
	CodeAndroidUrl     string
	PreRegDatetime     time.Time
	PreRegPassDatetime time.Time
	RegDatetime        time.Time
	RegPassDatetime    time.Time
	IsPass             int64
	CheckStatus        int64
	IsUse              int64
}

type Logs struct {
	Id              int64
	User            *User `orm:"rel(fk)"`
	UserName        string
	OperateInfo     string
	OperateDatetime time.Time
}

type Role struct {
	Id             int64
	RoleName       string
	CreationTime   time.Time
	CreatedBy      string
	LastUpdateTime time.Time
	IsStart        int64
}

type RolePermission struct {
	Id             int64
	Role           *Role `orm:"rel(fk)"`
	FunctionId     int64
	CreationTime   time.Time
	CreatedBy      string
	LastUpdateTime time.Time
	IsStart        int64
}

type SiteConfig struct {
	Id       int64
	SiteName string
}

type SystemConfig struct {
	Id              int64
	ConfigType      int64
	ConfigTypeName  string
	ConfigTypeValue int64
	ConfigValue     string
	IsStart         int64
}

type User struct {
	Id             int64
	UserCode       string
	UserName       string
	UserPassword   string
	CreationTime   time.Time
	LastLoginTime  time.Time
	CreatedBy      string
	LastUpdateTime time.Time
	IsStart        int64
	Role           *Role `orm:"rel(fk)"`
}

type HatArea struct {
	Are  string
	Id   int64
	Area string
	City *HatCity `orm:"rel(fk)"`
}

type HatCity struct {
	Cit string
	Id  int64

	City     string
	Province *HatProvince `orm:"rel(fk)"`
}

type HatProvince struct {
	Provi    string
	Id       int64
	Province string
}
