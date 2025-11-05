package model

type Member struct {
	Id                          string  `gorm:"column:id"`
	AliNo                       string  `gorm:"column:ali_no"`
	QrCodeUrl                   string  `gorm:"column:qr_code_url"`
	AppealTimes                 int64   `gorm:"column:appeal_times"`
	AppealSuccessTimes          int64   `gorm:"column:appeal_success_times"`
	ApplicationTime             int64   `gorm:"column:application_time"`
	Avatar                      string  `gorm:"column:avatar"`
	Bank                        string  `gorm:"column:bank"`
	Branch                      string  `gorm:"column:branch"`
	CardNo                      string  `gorm:"column:card_no"`
	CertifiedBussinessApplyTime int64   `gorm:"column:certified_bussiness_apply_time"`
	CertifiedBussinessCheckTime int64   `gorm:"column:certified_bussiness_check_time"`
	CertifiedBussinessStatus    int64   `gorm:"column:certified_bussiness_status"`
	ChannelId                   int64   `gorm:"column:channel_id"`
	Email                       string  `gorm:"column:email"`
	FirstLevel                  int64   `gorm:"column:first_level"`
	GoogleDate                  int64   `gorm:"column:google_date"`
	GoogleKey                   string  `gorm:"column:google_key"`
	GoogleState                 int32   `gorm:"column:google_state"`
	IdNumber                    string  `gorm:"column:id_number"`
	InviterId                   int64   `gorm:"column:inviter_id"`
	IsChannel                   int32   `gorm:"column:is_channel"`
	JYPassword                  string  `gorm:"column:jy_password"`
	LastLoginTime               int64   `gorm:"column:last_login_time"`
	City                        string  `gorm:"column:city"`
	Country                     string  `gorm:"column:country"`
	District                    string  `gorm:"column:district"`
	Province                    string  `gorm:"column:province"`
	LoginCount                  int32   `gorm:"column:login_count"`
	LoginLock                   int32   `gorm:"column:login_lock"`
	Margin                      string  `gorm:"column:margin"`
	MemberLevel                 int64   `gorm:"column:member_level"`
	MobilePhone                 string  `gorm:"column:mobile_phone"`
	Password                    string  `gorm:"column:password"`
	PromotionCode               string  `gorm:"column:promotion_code"`
	PublishAdvertise            int32   `gorm:"column:publish_advertise"`
	RealName                    string  `gorm:"column:real_name"`
	RealNameStatus              int32   `gorm:"column:real_name_status"`
	Salt                        string  `gorm:"column:salt"`
	SecondLevel                 int64   `gorm:"column:second_level"`
	SignInAbility               int16   `gorm:"column:sign_in_ability"`
	Status                      int16   `gorm:"column:status"`
	ThirdLevel                  int64   `gorm:"column:third_level"`
	Token                       string  `gorm:"column:token"`
	TokenExpireTime             int64   `gorm:"column:token_expire_time"`
	TransactionStatus           int16   `gorm:"column:transaction_status"`
	TransactionTime             int64   `gorm:"column:transaction_time"`
	Transaction                 int16   `gorm:"column:transaction"`
	Username                    string  `gorm:"column:username"`
	QrWeCodeUrl                 string  `gorm:"column:qr_we_code_url"`
	Wechat                      string  `gorm:"column:wechat"`
	Local                       string  `gorm:"column:local"`
	Integration                 int64   `gorm:"column:integration"`
	MemberGradeId               int64   `gorm:"column:member_grade_id"`
	KYCStatus                   int16   `gorm:"column:kyc_status"`
	GeneralizeTotal             int64   `gorm:"column:generalize_total"`
	InviterParentId             int64   `gorm:"column:inviter_parent_id"`
	SuperPartner                string  `gorm:"column:super_partner"`
	KickFee                     float64 `gorm:"column:decimal(19,2)"`
	Power                       float64 `gorm:"column:power"`
	TeamLevel                   int16   `gorm:"column:team_level"`
	TeamPower                   float64 `gorm:"column:team_power"`
	MemberLevelId               int64   `gorm:"column:member_level_id"`
}

func (m *Member) TableName() string {
	return "member"
}

func NewMember() *Member {
	return &Member{}
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
	NORMAL  = 0
	ILLEGAL = 1
)

func (m *Member) FillSuperPartner(partner string) {
	if partner == "" {
		m.SuperPartner = NORMALPARTER
		m.Status = NORMAL
	} else {
		if partner == SUPERPARTER {
			m.SuperPartner = partner
			m.Status = ILLEGAL
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
