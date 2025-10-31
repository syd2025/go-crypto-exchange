package model

type Member struct {
	Id                          int64   `gorm:"columnid"`
	AliNo                       string  `gorm:"columnali_no" default:"0"`
	QrCodeUrl                   string  `gorm:"columnqr_code_url"`
	AppealSuccessTimes          int64   `gorm:"columnappeal_success_times"`
	AppealTime                  int64   `gorm:"columnappeal_times"`
	ApplicationTime             int64   `gorm:"columnapplication_time"`
	Avatar                      string  `gorm:"columnavatar"`
	Bank                        string  `gorm:"columnbank"`
	Branch                      string  `gorm:"columnbranch"`
	CardNo                      string  `gorm:"columncard_no"`
	CertifiedBussinessApplyTime int64   `gorm:"columncertified_bussiness_apply_time"`
	CertifiedBussinessCheckTime int64   `gorm:"columncertified_bussiness_check_time"`
	ChannelId                   int64   `gorm:"columnchannel_id"`
	Email                       string  `gorm:"columnemail"`
	FirstLevel                  int64   `gorm:"columnfirst_level"`
	GoogleDate                  int64   `gorm:"columngoogle_date"`
	GoogleKey                   string  `gorm:"columngoogle_key"`
	GoogleState                 int64   `gorm:"columngoogle_state"`
	IdNumber                    string  `gorm:"columnid_number"`
	InviterId                   int64   `gorm:"columninviter_id"`
	IsChannel                   int64   `gorm:"columnis_channel"`
	JyPassword                  string  `gorm:"columnjy_password"`
	LastLoginTime               int64   `gorm:"columnlast_login_time"`
	City                        string  `gorm:"columncity"`
	Country                     string  `gorm:"columncountry"`
	District                    string  `gorm:"columndistrict"`
	Province                    string  `gorm:"columnprovince"`
	LoginCount                  int64   `gorm:"columnlogin_count"`
	LoginLock                   int64   `gorm:"columnlogin_lock"`
	Margin                      string  `gorm:"columnmargin"`
	MemberLevel                 int64   `gorm:"columnmember_level"`
	MobilePhone                 string  `gorm:"columnmobile_phone"`
	Password                    string  `gorm:"columnpassword"`
	PromotionCode               string  `gorm:"columnpromotion_code"`
	PublishAdvertise            int64   `gorm:"columnpublish_advertise"`
	RealName                    string  `gorm:"columnreal_name"`
	RealNameStatus              int64   `gorm:"columnreal_name_status"`
	RegistrationTime            int64   `gorm:"columnregistration_time"`
	Salt                        string  `gorm:"columnsalt"`
	SecondLevel                 int64   `gorm:"columnsecond_level"`
	SignInAbility               int64   `gorm:"columnsign_in_ability"`
	Status                      int64   `gorm:"columnstatus"`
	ThirdLevel                  int64   `gorm:"columnthird_level"`
	Token                       string  `gorm:"columntoken"`
	TokenExpireTime             int64   `gorm:"columntoken_expire_time"`
	TransactionStatus           int64   `gorm:"columntransaction_status"`
	TransactionTime             int64   `gorm:"columntransaction_time"`
	Transactions                int64   `gorm:"columntransactions"`
	Username                    string  `gorm:"columnusername"`
	QrWeCodeUrl                 string  `gorm:"columnqr_we_code_url"`
	Wechat                      string  `gorm:"columnwechat"`
	Integration                 int64   `gorm:"columnintegration"`
	MemberGradeId               int64   `gorm:"columnmember_grade_id"`
	KycStatus                   int64   `gorm:"columnkyc_status"`
	GeneralizeTotal             int64   `gorm:"columngeneralize_total"`
	InviterParentId             int64   `gorm:"columninviter_parent_id"`
	SuperPartner                string  `gorm:"columnsuper_partner"`
	KickFee                     float64 `gorm:"columnkick_fee"`
	Power                       float64 `gorm:"columnpower"`
	TeamLevel                   int64   `gorm:"columnteam_level"`
	TeamPower                   float64 `gorm:"columnteam_power"`
	MemberLevelId               int64   `gorm:"columnmember_level_id"`
}

func (m *Member) TableName() string {
	return "member"
}

const (
	GENERAL        = 0
	REALNAME       = 1
	IDENTIFICATION = 2
)

const (
	NORMALPARTER = "0"
	SUPERPARTER  = "1"
	PSUPERPARTER = "2"
)

const (
	NORMAL   = 0
	ILLEAGAL = 1
)

func (m *Member) FillSuperPartner(partner string) {
	if partner == "" {
		m.SuperPartner = NORMALPARTER
		m.Status = NORMAL
	} else {
		if partner != NORMALPARTER {
			m.SuperPartner = partner
			m.Status = ILLEAGAL
		}
	}
}

func (m *Member) MemberLevelStr() string {
	if m.MemberLevel == GENERAL {
		return "普通会员"
	}
	if m.MemberLevel == REALNAME {
		return "实名会员"
	}
	if m.MemberLevel == IDENTIFICATION {
		return "认证会员"
	}
	return ""
}

func (m *Member) MemberRate() int32 {
	if m.SuperPartner == NORMALPARTER {
		return 0
	}
	if m.SuperPartner == SUPERPARTER {
		return 1
	}
	if m.SuperPartner == PSUPERPARTER {
		return 2
	}
	return 0
}

func NewMember() *Member {
	return &Member{}
}
