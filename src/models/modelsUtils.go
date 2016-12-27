package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func QueryUser() orm.QuerySeter {
	return orm.NewOrm().QueryTable("as_user").OrderBy("-Id")
}

//前面的括号里面的参数表示 这个方法附属在那个类上面
func (m *User) EditUser() error {

	err := m.Update("UserCode", "UserName", "Role", "UserPassword", "IsStart")

	return err

}
func GetAccountSystemConfig() ([]*SystemConfig, error) {
	list := make([]*SystemConfig, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("as_systemconfig").Filter("ConfigType", 1).All(&list)

	return list, err

}
func GetProvinceList() ([]*HatProvince, error) {
	list := make([]*HatProvince, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("hat_province").All(&list)

	return list, err

}
func GetCITYList(province *HatProvince) ([]*HatCity, error) {
	list := make([]*HatCity, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("hat_city").Filter("Province", province).All(&list)

	return list, err

}

func GetAreAist(city *HatCity) ([]*HatArea, error) {
	list := make([]*HatArea, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("hat_area").Filter("City", city).All(&list)

	return list, err

}

func Tx_operationAccount(oldAccount, newAccount *Account, AccountDetail *AccountDetail, logs Logs) bool {
	flag := false
	o := orm.NewOrm()
	o.Begin()
	money := oldAccount.Money + newAccount.Money
	oldAccount.Money = money
	oldAccount.MoneyBak = money

	err := oldAccount.Update("Money", "MoneyBak")
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return flag
	}

	err1 := AccountDetail.Insert()
	if err1 != nil {
		beego.Error(err1)
		o.Rollback()
		return flag

	}

	err2 := logs.Insert()
	if err2 != nil {
		beego.Error(err2)
		o.Rollback()
		return flag

	}

	err3 := o.Commit()
	if err3 != nil {
		beego.Error(err3)
		return flag
	}

	return true

}

func (m *Customs) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Keywords) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *SystemConfig) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *Account) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *Customs) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *Keywords) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *User) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *Account) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *User) Insert() bool {
	flagg := false
	o := orm.NewOrm()
	var nummm interface{}
	var list orm.ParamsList
	//自定义sql语句返回string类型的 list  查询id的最大值
	num, err := o.Raw("select max(id) from as_user").ValuesFlat(&list)
	if err != nil {
		beego.Error(err)
	}

	if num > 0 && err == nil {
		nummm = list[0]

	}
	var name string
	//类型转换
	if i, ok := nummm.(string); ok {
		name = i
	}

	num, err = strconv.ParseInt(name, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	num = num + 1

	m.Id = num
	o.Begin()

	_, err = o.Insert(m)

	if err != nil {
		beego.Error(err)
		o.Rollback()

	} else {

		account := Account{User: m, Money: 0, MoneyBak: 0}

		err := account.Insert()

		if err != nil {
			beego.Error(err)
			o.Rollback()

		} else {

			errr := o.Commit()
			if errr != nil {
				o.Rollback()
				beego.Error(errr)

			} else {
				flagg = true
			}

		}

	}

	return flagg

}
func (m *Customs) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *AccountDetail) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *Logs) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func GetFunctionById(id string) (*Function, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	function := new(Function)
	qs := o.QueryTable("Function")
	err = qs.Filter("id", cid).One(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}
func GetRoleList() (roles []*Role, err error) {

	o := orm.NewOrm()

	roles = make([]*Role, 0)
	qs := o.QueryTable("Role")
	//   查询list的时候  需要传 数组的 内置地址作为参数
	_, err = qs.Filter("is_start", 1).All(&roles)

	return roles, err

}

func GetRolePremissionList(isStart, roleId string) (role_premissions []*RolePermission, err error) {
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return nil, err
	}
	croleId, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	role_premissions = make([]*RolePermission, 0)
	qs := o.QueryTable("RolePermission")
	//   查询list的时候  需要传 数组的 内置地址作为参数
	_, err = qs.Filter("is_start", cisStart).Filter("role_id", croleId).All(&role_premissions)

	return role_premissions, err
}

func RegisterDB() {
	// 检查数据库文件
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterModel(new(Account), new(AccountDetail), new(Admin), new(Contacts), new(Customs), new(Function), new(Keywords), new(Logs), new(Role), new(RolePermission), new(SiteConfig), new(User), new(HatArea), new(HatCity), new(HatProvince), new(SystemConfig))

	orm.RegisterDataBase("default", "mysql", "root:bdqn@/agentsystemdb?charset=utf8")

}

func GetAllAs_functions(isStart string) (functions []*Function, err error) {
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	functions = make([]*Function, 0)
	qs := o.QueryTable("Function")
	//   查询list的时候  需要传 数组的 内置地址作为参数
	_, err = qs.Filter("parent_id", 0).Filter("is_start", cisStart).All(&functions)

	return functions, err
}

func AddUser(userCode, userName, userPassword, createdBy, isStart, roleId string) error {
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return err
	}

	croleId, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return err
	}

	var role Role
	role.Id = croleId

	o := orm.NewOrm()
	u := &User{UserCode: userCode, UserPassword: userPassword, UserName: userName, CreatedBy: createdBy, IsStart: cisStart, Role: &role, CreationTime: time.Now(), LastLoginTime: time.Now(), LastUpdateTime: time.Now()}
	_, err = o.Insert(u)
	if err != nil {
		return err

	}
	return err

}

func ModifyUserPassword(userPassword, id string) error {

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	u := &User{Id: cid}
	// read 从数据库看看有没有这条数据
	if o.Read(u) == nil {
		u.UserPassword = userPassword
		_, err := o.Update(u)
		if err != nil {
			return err

		}
	}

	return nil

}

func ModifyUserLast(id int64, last time.Time) error {

	o := orm.NewOrm()
	u := &User{Id: id}
	// read 从数据库看看有没有这条数据
	if o.Read(u) == nil {
		u.LastLoginTime = last
		_, err := o.Update(u)
		if err != nil {
			return err

		}
	}

	return nil

}

func GetUserRolename() ([]*User, error) {

	o := orm.NewOrm()
	user := make([]*User, 0)
	qs := o.QueryTable("User")
	_, err := qs.Filter("Role", 1).All(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func ModifyUser(userCode, userName, userPassword, isStart, roleId, id string) error {
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return err
	}

	croleId, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return err
	}
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	u := &User{Id: cid}

	if o.Read(u) == nil {
		u.UserPassword = userPassword
		u.UserName = userName
		u.UserCode = userCode
		u.IsStart = cisStart
		u.Role.Id = croleId
		_, err := o.Update(u)
		if err != nil {
			return err

		}
	}

	return nil

}
func GetAccountByUserId(userId string) (*Account, error) {
	cid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	account := new(Account)
	qs := o.QueryTable("Account")
	err = qs.Filter("user_id", cid).RelatedSel().One(account)
	if err != nil {
		return nil, err
	}
	return account, nil

}

func (m *User) Delete() bool {
	flagg := false

	o := orm.NewOrm()
	o.Begin()
	//需要new 对象来时候
	account := new(Account)

	account.User = m
	errrere := account.Read("User")
	if errrere != nil {
		beego.Error(errrere)
		o.Rollback()

	} else {
		_, err := o.Delete(account)

		if err != nil {
			beego.Error(err)
			o.Rollback()

		} else {

			_, errr := o.Delete(m)

			if errr != nil {
				beego.Error(errr)
				o.Rollback()

			} else {

				errree := o.Commit()
				if errree != nil {

					beego.Error(errree)
				} else {
					flagg = true
				}

			}

		}
	}

	return flagg
}

func (m *Keywords) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *Account) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *SystemConfig) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Account) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *SystemConfig) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func GetUserById(id string) (*User, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("User")
	err = qs.Filter("id", cid).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func GetAllUser() ([]*User, error) {
	o := orm.NewOrm()

	cates := make([]*User, 0)

	qs := o.QueryTable("User")
	_, err := qs.All(&cates)
	return cates, err
}

func IsCunZaiUserCode(userCode string) (int64, error) {
	o := orm.NewOrm()

	qs := o.QueryTable("User")
	a, err := qs.Filter("user_code", userCode).Count()
	return a, err
}
func GetUserByUserCode(userCode string) (*User, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("User")
	err := qs.Filter("user_code", userCode).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserBySearch(usercode string) ([]*User, error) {
	o := orm.NewOrm()

	cates := make([]*User, 0)

	qs := o.QueryTable("User")
	_, err := qs.Filter("user_code__contains", usercode).All(&cates)
	return cates, err
}

func GetUserList(userName, roleId, isStart, starNum, pageSize string) (int64, []*User, error) {

	o := orm.NewOrm()

	cates := make([]*User, 0)

	qs := o.QueryTable("User")
	qs = qs.Filter("user_name__contains", userName)

	if roleId != "" {

		croleId, err := strconv.ParseInt(roleId, 10, 64)
		if err != nil {
			return 0, nil, err
		}

		qs = qs.Filter("role_id", croleId)
	}

	if isStart != "" {

		cisStart, err := strconv.ParseInt(isStart, 10, 64)
		if err != nil {
			return 0, nil, err
		}

		qs = qs.Filter("is_start", cisStart)
	}

	bb, err := qs.Count()

	if starNum != "" && pageSize != "" {

		cstarNum, err := strconv.ParseInt(starNum, 10, 64)
		if err != nil {
			return 0, nil, err
		}
		cpageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return 0, nil, err
		}
		qs = qs.Limit(cpageSize, cstarNum)
	}

	//RelatedSel  自动关联查询

	_, err = qs.RelatedSel().All(&cates)

	return bb, cates, err
}

func GetSystemconfigByConfigType(ConfigType int64) ([]*SystemConfig, error) {

	o := orm.NewOrm()
	user := make([]*SystemConfig, 0)
	qs := o.QueryTable("as_systemconfig")
	_, err := qs.Filter("ConfigType", ConfigType).All(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func GetCustomByName(customName string) ([]*Customs, error) {

	o := orm.NewOrm()
	user := make([]*Customs, 0)
	qs := o.QueryTable("Customs")
	_, err := qs.Filter("CustomName__contains", customName).All(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func Tx_saveCustomContact(AccountList []*Contacts, customs *Customs) bool {

	flagg := false
	o := orm.NewOrm()
	var nummm interface{}
	var list orm.ParamsList
	//自定义sql语句返回list 查询id的最大值
	num, err := o.Raw("select max(id) from Customs").ValuesFlat(&list)
	if err != nil {
		beego.Error(err)
	}

	nummm = list[0]

	var name string
	//类型转换
	if i, ok := nummm.(string); ok {
		name = i
	}

	num, err = strconv.ParseInt(name, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	num = num + 1

	o.Begin()

	customs.Id = num

	_, err = o.Insert(customs)

	if err != nil {
		beego.Error(err)
		o.Rollback()

	} else {
		for i := 0; i < len(AccountList); i++ {
			AccountList[i].Customs = customs

			_, err := o.Insert(AccountList[i])

			if err != nil {
				beego.Error(err)
				o.Rollback()

			} else {

				errr := o.Commit()
				if errr != nil {
					o.Rollback()
					beego.Error(errr)

				} else {
					flagg = true
				}

			}
		}

	}

	return flagg

}

func IsCunZaicustem(customName string) (int64, error) {
	o := orm.NewOrm()

	qs := o.QueryTable("Customs")
	a, err := qs.Filter("CustomName", customName).Count()
	return a, err
}

func IsCunZaikeywords(customName string) (int64, error) {
	o := orm.NewOrm()

	qs := o.QueryTable("Keywords")
	a, err := qs.Filter("Keywords", customName).Count()
	return a, err
}
func IsCunZaiConfigTypeName(customName string) (int64, error) {
	o := orm.NewOrm()

	qs := o.QueryTable("SystemConfig")
	a, err := qs.Filter("ConfigTypeName", customName).Count()
	return a, err
}

func Tx_SaveKeywords(keywords *Keywords, user *User) bool {
	flag := false
	o := orm.NewOrm()
	o.Begin()
	_, errre := o.Insert(keywords)

	timeee := time.Now()

	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	} else {

		account := Account{User: user}

		account.Read("User")
		oney := account.Money - keywords.Price
		account.Money = oney

		account.MoneyBak = account.Money
		err := account.Update("Money", "MoneyBak")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		} else {
			ServiceYears := strconv.FormatInt(keywords.ServiceYears, 10)
			Price := strconv.FormatFloat(keywords.Price, 'f', 2, 64)

			memo := user.UserName + "对" + keywords.CustomName + "进行关键词申请操作,扣除预付款资金：" + ServiceYears + "年" + Price + "元"
			AccountDetail := AccountDetail{DetailType: 9999, User: user, DetailTypeName: "预注册冻结资金",
				DetailDateTime: timeee, Money: keywords.Price, AccountMoney: account.Money, Memo: memo}

			eeerr := AccountDetail.Insert()

			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			} else {

				log := Logs{User: user, UserName: user.UserCode,
					OperateDatetime: timeee, OperateInfo: memo}
				errrer := log.Insert()
				if errrer != nil {
					beego.Error(errrer)
					o.Rollback()
					return flag

				} else {

					errr := o.Commit()
					if errr != nil {
						beego.Error(errr)
					} else {
						flag = true
					}
				}
			}
		}

	}

	return flag

}

func GetCustomByCustomName(CustomName, starNum, pageSize string) (int64, []*Customs, error) {

	o := orm.NewOrm()

	cates := make([]*Customs, 0)

	qs := o.QueryTable("Customs")
	qs = qs.Filter("custom_name__contains", CustomName)

	bb, err := qs.Count()

	if starNum != "" && pageSize != "" {

		cstarNum, err := strconv.ParseInt(starNum, 10, 64)
		if err != nil {
			return 0, nil, err
		}
		cpageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return 0, nil, err
		}
		qs = qs.Limit(cpageSize, cstarNum)
	}

	//RelatedSel  自动关联查询

	_, err = qs.RelatedSel().All(&cates)

	return bb, cates, err
}

func GetConstactByCumstom(customs *Customs) ([]*Contacts, error) {

	o := orm.NewOrm()
	user := make([]*Contacts, 0)
	qs := o.QueryTable("Contacts")
	_, err := qs.Filter("Customs", customs).All(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func GetAreaById(areaId string) (*HatArea, error) {
	areaIDd, err := strconv.ParseInt(areaId, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()

	cates := new(HatArea)

	qs := o.QueryTable("hat_area")
	err = qs.Filter("Id", areaIDd).RelatedSel().One(cates)

	return cates, err
}

func Tx_ModifyCustomContact(AccountList []*Contacts, customs *Customs) bool {

	flagg := false
	o := orm.NewOrm()

	o.Begin()

	_, err := o.Update(customs)

	if err != nil {
		beego.Error(err)
		o.Rollback()

	} else {

		oldAccountList := make([]*Contacts, 0)

		_, err := o.QueryTable("Contacts").Filter("Customs", customs).All(&oldAccountList)
		if err != nil {
			beego.Error(err)
			o.Rollback()

		} else {

			for _, v := range oldAccountList {
				_, err := o.Delete(v)

				if err != nil {
					beego.Error(err)
					o.Rollback()

				}

			}

			for i := 0; i < len(AccountList); i++ {
				AccountList[i].Customs = customs

				_, err := o.Insert(AccountList[i])

				if err != nil {
					beego.Error(err)
					o.Rollback()

				} else {

					errr := o.Commit()
					if errr != nil {
						o.Rollback()
						beego.Error(errr)

					} else {
						flagg = true
					}
				}
			}
		}

	}

	return flagg

}

func GetKeywordBySearch(keyword, starNum, pageSize string) (int64, []*Keywords, error) {

	o := orm.NewOrm()

	cates := make([]*Keywords, 0)

	qs := o.QueryTable("Keywords")
	qs = qs.Filter("keywords__contains", keyword)

	bb, err := qs.Count()

	if starNum != "" && pageSize != "" {

		cstarNum, err := strconv.ParseInt(starNum, 10, 64)
		if err != nil {
			return 0, nil, err
		}
		cpageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return 0, nil, err
		}
		qs = qs.Limit(cpageSize, cstarNum)
	}

	//RelatedSel  自动关联查询

	_, err = qs.RelatedSel().All(&cates)

	return bb, cates, err
}

func Tx_ChangeStatusToOk(keywords *Keywords, user *User) bool {
	flag := false
	timeeee := time.Now()

	CheckStatusssssss := keywords.CheckStatus
	if CheckStatusssssss != 2 {
		return flag
	}

	keywords.Read("Id")

	keywords.CheckStatus = CheckStatusssssss

	o := orm.NewOrm()
	o.Begin()

	_, errre := o.Update(keywords, "CheckStatus")

	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	} else {

		account := Account{User: keywords.Agent}

		err := account.Read("User")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		} else {

			account.Money = account.Money + keywords.PreRegFrozenMoney
			account.MoneyBak = account.Money
			_, err := o.Update(&account, "Money", "MoneyBak")

			if err != nil {
				beego.Error(err)
				o.Rollback()
				return flag

			} else {

				PreRegFrozenMoney := strconv.FormatFloat(keywords.PreRegFrozenMoney, 'f', 2, 64)

				memooo := user.UserName + "对" + keywords.CustomName + "进行关键词审核操作,返回冻结资金：" + PreRegFrozenMoney

				AccountDetail := AccountDetail{User: user, DetailType: 9998,
					DetailTypeName: "返回预注册冻结资金", Money: keywords.PreRegFrozenMoney,
					AccountMoney: account.Money,
					Memo:         memooo, DetailDateTime: timeeee,
				}
				_, eeerr := o.Insert(&AccountDetail)

				if eeerr != nil {
					beego.Error(eeerr)
					o.Rollback()
					return flag

				} else {

					account.Money = account.Money - keywords.Price
					account.MoneyBak = account.Money

					_, err := o.Update(&account, "Money", "MoneyBak")

					if err != nil {
						beego.Error(err)
						o.Rollback()
						return flag

					} else {
						memo2 := user.UserName + "对" + keywords.CustomName + "进行关键词审核通过操作自动正式扣款操作,扣除正式注册资金：" + PreRegFrozenMoney

						AccountDetail2 := AccountDetail{User: user,
							DetailType: 9997, DetailTypeName: "扣除申请关键词" + keywords.Keywords + "的所有资金",
							Money:          0 - account.Money,
							AccountMoney:   account.Money,
							Memo:           memo2,
							DetailDateTime: timeeee}

						_, eeerr = o.Insert(&AccountDetail2)
						if eeerr != nil {
							beego.Error(eeerr)
							o.Rollback()
							return flag

						} else {

							log := Logs{User: user, UserName: user.UserCode,
								OperateDatetime: timeeee,
								OperateInfo:     memooo,
							}

							_, eeerr = o.Insert(&log)
							if eeerr != nil {
								beego.Error(eeerr)
								o.Rollback()
								return flag

							} else {
								logs := Logs{User: user, UserName: user.UserCode,
									OperateDatetime: timeeee,
									OperateInfo:     memo2}

								_, eeerr = o.Insert(&logs)
								if eeerr != nil {
									beego.Error(eeerr)
									o.Rollback()
									return flag

								} else {

									errr := o.Commit()
									if errr != nil {
										beego.Error(errr)
									} else {
										flag = true
									}
								}
							}
						}
					}
				}
			}
		}

	}

	return flag

}

func Tx_ChangeStatusToNo(keywords *Keywords, user *User) bool {
	flag := false
	timeeee := time.Now()
	CheckStatusssssss := keywords.CheckStatus
	if CheckStatusssssss != 3 {
		return flag
	}

	keywords.Read("Id")

	keywords.CheckStatus = CheckStatusssssss

	o := orm.NewOrm()
	o.Begin()

	_, errre := o.Update(keywords, "CheckStatus")

	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	} else {

		account := Account{User: keywords.Agent}

		account.Read("User")
		account.Money = account.Money + keywords.PreRegFrozenMoney
		account.MoneyBak = account.Money
		_, err := o.Update(&account, "Money", "MoneyBak")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		} else {

			PreRegFrozenMoney := strconv.FormatFloat(keywords.PreRegFrozenMoney, 'f', 2, 64)

			memooo := user.UserName + "对" + keywords.CustomName + "进行关键词审核操作,返回冻结资金：" + PreRegFrozenMoney

			AccountDetail := AccountDetail{User: user, DetailType: 9998,
				DetailTypeName: "返回预注册冻结资金", Money: keywords.PreRegFrozenMoney,
				AccountMoney: account.Money,
				Memo:         memooo, DetailDateTime: timeeee,
			}
			_, eeerr := o.Insert(&AccountDetail)

			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			} else {

				log := Logs{User: user, UserName: user.UserCode,
					OperateDatetime: timeeee,
					OperateInfo:     memooo,
				}

				_, eeerr = o.Insert(&log)
				if eeerr != nil {
					beego.Error(eeerr)
					o.Rollback()
					return flag

				} else {

					errr := o.Commit()
					if errr != nil {
						beego.Error(errr)
					} else {
						flag = true
					}

				}
			}
		}

	}

	return flag

}

func Tx_SaveXuFei(keywords *Keywords, user *User) bool {
	flag := false
	o := orm.NewOrm()
	o.Begin()
	_, errre := o.Update(keywords, "RegPassDatetime", "ServiceYears")

	timeee := time.Now()

	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	} else {

		account := Account{User: user}

		account.Read("User")
		oney := account.Money - keywords.Price
		account.Money = oney

		account.MoneyBak = account.Money
		err := account.Update("Money", "MoneyBak")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		} else {
			ServiceYears := strconv.FormatInt(keywords.ServiceYears, 10)
			Price := strconv.FormatFloat(keywords.Price, 'f', 2, 64)

			memo := user.UserName + "对" + keywords.CustomName + "进行关键词申请操作,扣除预付款资金：" + ServiceYears + "年" + Price + "元"
			AccountDetail := AccountDetail{DetailType: 9999, User: user, DetailTypeName: "预注册冻结资金",
				DetailDateTime: timeee, Money: keywords.Price, AccountMoney: account.Money, Memo: memo}

			eeerr := AccountDetail.Insert()

			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			} else {

				log := Logs{User: user, UserName: user.UserCode,
					OperateDatetime: timeee, OperateInfo: memo}
				errrer := log.Insert()
				if errrer != nil {
					beego.Error(errrer)
					o.Rollback()
					return flag

				} else {

					errr := o.Commit()
					if errr != nil {
						beego.Error(errr)
					} else {
						flag = true
					}
				}
			}
		}

	}

	return flag

}
