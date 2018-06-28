package models

import (
	"dpm/common"
	"time"
)

const (
	//简历仅自己可见
	CV_VISIBILI_TYPE_PRIVATE = 0
	//简历对所有人可见
	CV_VISIBILI_TYPE_PUBLIC = 1
	//简历用密码可见
	CV_VISIBILI_TYPE_PWD = 2

	//基本信息模块类型
	CVMTYPE_BASICINFO = "cvmtype_basicinfo"
	//头像模块类型
	CVMTYPE_HEADPORTRAIT = "cvmtype_headportrait"
	//求职意向模块类型
	CVMTYPE_JOBINTENSION = "cvmtype_job_intension"
	//项目经验模块类型
	CVMTYPE_PROJECTEXPERIENCE = "cvmtype_project_experience"
	//教育经历模块类型
	CVMTYPE_EDUCATEEXPERIENCE = "cvmtype_educate_experience"
	//自我评价模块类型
	CVMTYPE_SELFEVALUATION = "cvmtype_self_evaluation"
	//荣誉奖项模块类型
	CVMTYPE_HONORS = "cvmtype_honors"
)

type CurriculumVitae struct {
	//简历Id
	CVId string `json:"cv_id"`
	//简历名称
	CVName string `json:"cv_name,omitempty"`
	//查看密码
	CViewPwd string `json:"cview_pwd,omitempty"`
	//自定义域名
	CustomDomainName string `json:"custom_domainname,omitempty"`
	//可见类型
	CVisibiliType int `json:"cvisibili_type,omitempty"`
	//创建时间
	CVCreateTime common.JSONTime `json:"cv_createtime,omitempty"`
	//更新时间
	CVUpdateTime common.JSONTime `json:"cv_updatetime,omitempty"`
	//简历模板
	*CVTemplate `json:"cv_template,omitempty"`
	//基本信息模块
	*BasicInfoCVM `json:"basicinfo_cvm,omitempty"`
	//头像模块
	*HeadPortraitCVM `json:"headportrait_cvm,omitempty"`
	//求职意向模块
	*JobIntensionCVM `json:"jobintension_cvm,omitempty"`
	//工作经历模块
	*WorkExperienceCVM `json:"workexperience_cvm,omitempty"`
	//志愿者经历模块
	*VolunteerExperienceCVM `json:"volunteerexperience_cvm,omitempty"`
	//实习生经历模块
	*TraineeExperienceCVM `json:"traineeexperience_cvm,omitempty"`
	//教育经历模块
	*EducateExperienceCVM `json:"educateexperience_cvm,omitempty"`
	//项目经历
	*ProjectExperienceCVM `json:"projectexperience_cvm,omitempty"`
	//自我评价模块
	*SelfEvaluationCVM `json:"selfevaluation_cvm,omitempty"`
	//荣誉奖项模块
	*HonorsCVM `json:"honors_cvm,omitempty"`
	//自荐信模块
	*SelfRecomdLetterCVM `json:"selfrecomdletter_cvm,omitempty"`
	//作品展示模块
	*WorksShowCVM `json:"worksshow_cvm,omitempty"`
	//兴趣爱好模块
	*InterestCVM `json:"interest_cvm,omitempty"`
	//技能特长模块
	*SpecialityCVM `json:"speciality_cvm,omitempty"`
	//自定义描述模块
	CustomDescribeCVMs []CustomDescribeCVM `json:"customdescribe_cvms,omitempty"`
	//自定义经验模块
	CustomExperienceCVMs []CustomExperienceCVM `json:"customexperience_cvms,omitempty"`
}

//简历模块
type CVModule struct {
	//模块类型
	CVMType string `json:"cvm_type"`
	//模块名称
	CVMName string `json:"cvm_name"`
	//是否显示
	ModuleDisplay bool `json:"module_display"`
	//是否必须
	ModuleRequisite bool `json:"module_requisite"`
	//坐标位置
	ModuleCoordinate string `json:"module_coordinate"`
}

//模块标题
type TitleCVM struct {
	*CVModule
	//模块标题
	TitleName string `json:"title_name"`
	//标题显示
	TitleDisplay bool `json:"title_display"`
}

//基本信息模块
type BasicInfoCVM struct {
	*TitleCVM
	//候选人姓名
	CandidateName string `json:"candidate_name"`
	//出生年份
	YearOfBirth int `json:"yearofbirth"`
	//出生月份
	MonthOfBirth int `json:"monthofbirth"`
	//所在城市
	CityLocation string `json:"city_location"`
	//工作年限
	WorkingLife int `json:"working_life"`
	//联系电话
	PhoneNumber string `json:"phone_number"`
	//邮箱地址
	EMailAddress string `json:"email_address"`
	//简短描述
	ShortDescribe string `json:"short_describe"`
	//性别
	Sex string `json:"sex"`
	//最高学历
	HighestEducation string `json:"highest_education"`
	//民族
	Nationality string `json:"nationality"`
	//婚姻状况
	MaritalStatus string `json:"marital_status"`
	//政治面貌
	PoliticalOutlook string `json:"political_outlook"`
	//身高
	Height int `json:"height"`
	//体重
	Weight int `json:"weight"`
	//自定义字段
	CustomFields []CustomKVField `json:"custom_fields"`
}

//自定义字段
type CustomKVField struct {
	//字段名
	FieldName string `json:"field_name"`
	//字段值
	FieldValue string `json:"field_value"`
}

//头像模块
type HeadPortraitCVM struct {
	*TitleCVM
	//头像路径
	PortraitPath string `json:"portrait_path"`
	//头像风格
	PortraitStyle string `json:"portrait_style"`
}

//求职意向模块
type JobIntensionCVM struct {
	*TitleCVM
	//意向岗位
	IntentionalPost string `json:"intentional_post"`
	//职业类型：{全职，兼职，实习，不填}
	CareerType string `json:"career_type"`
	//意向城市
	IntentionalCity string `json:"intentional_city"`
	//入职时间:{随时，1周内，3个月，另议，不填}
	Hiredate string `json:"hiredate"`
	//薪资下限
	SalaryLower float64 `json:"salary_lower"`
	//薪资上限
	SalaryCap float64 `json:"salary_cap"`
	//是否面议
	SalaryPersonally bool `json:"salary_personally"`
}

//时间段类模块
type TimeRangeCVM struct {
	*TitleCVM
	//时间隐藏
	ExpTimeDisplay bool `json:"exptime_display"`
	//描述隐藏
	ExpDescDisplay bool `json:"expdesc_display"`
	//经历明细
	ExperienceItems []ExperienceItem `json:"experience_items"`
}

//时间经历
type ExperienceItem struct {
	*TitleCVM
	//开始时间
	BeginDate common.JSONTime `json:"begin_date"`
	//结束时间
	EndDate common.JSONTime `json:"end_date"`
	//补充字段名一
	ExtFieldName1st string `json:"extfield_name1st"`
	//补充字段值一
	ExtFieldVal1st string `json:"extfield_val1st"`
	//补充字段名二
	ExtFieldName2nd string `json:"extfield_name2nd"`
	//补充字段值二
	ExtFieldVal2nd string `json:"extfield_val2nd"`
	//经验描述
	ExpDesc string `json:"exp_desc"`
	//显示序号
	ViewOrder int64 `json:"view_order"`
}

//工作经历
type WorkExperienceCVM struct {
	*TimeRangeCVM
}

//志愿者经历
type VolunteerExperienceCVM struct {
	*WorkExperienceCVM
}

//实习生经历
type TraineeExperienceCVM struct {
	*WorkExperienceCVM
}

//教育经历
type EducateExperienceCVM struct {
	*TimeRangeCVM
}

//项目经历
type ProjectExperienceCVM struct {
	*TimeRangeCVM
}

//自定义经验模块
type CustomExperienceCVM struct {
	*TimeRangeCVM
}

//描述类模块
type DescribeCVM struct {
	*TitleCVM
	//描述
	CVMDescription string `json:"cvm_description"`
}

//自我评价模块
type SelfEvaluationCVM struct {
	*DescribeCVM
}

//荣誉奖项模块
type HonorsCVM struct {
	*DescribeCVM
}

//自定义描述模块
type CustomDescribeCVM struct {
	*DescribeCVM
}

//自荐信模块
type SelfRecomdLetterCVM struct {
	*DescribeCVM
}

//作品展示模块
type WorksShowCVM struct {
	*TitleCVM
	//作品
	OpuieItems map[string]interface{} `json:"opuie_items"`
}

//作品
type OpuieItem struct {
	//作品描述
	OpuieItemDesc string `json:"opuieitem_desc"`
	//作品类型:{"图片","在线"}
	OpuieItemType string `json:"opuieitem_type"`
	//访问路径
	ItemVisitPath string `json:"item_visit_path"`
}

//图片类作品
type PictureOpuieItem struct {
	*OpuieItem
	//作品标题
	PictureOpuieTitle string `json:"picture_opuie_title"`
}

//线上类作品
type OnlineOpuie struct {
	*OpuieItem
}

//标签类模块
type LabelCVM struct {
	*TitleCVM
}

//兴趣爱好模块
type InterestCVM struct {
	*LabelCVM
	InterestLabels []CommonLabel `json:"interest_labels"`
}

//技能特长模块
type SpecialityCVM struct {
	*LabelCVM
	SpecialityLabels []CommonLabel `json:"speciality_labels"`
}

//通用标签
type CommonLabel struct {
	//标签名称
	LabelName string `json:"label_name"`
	//标签类型
	LableType string `json:"lable_type"`
}

//简历默认构造函数
func NewCurriculumVitae() *CurriculumVitae {
	jtnow := common.JSONTime(time.Now())
	cv := &CurriculumVitae{
		CVisibiliType:          CV_VISIBILI_TYPE_PRIVATE,
		CVCreateTime:           jtnow,
		CVUpdateTime:           jtnow,
		BasicInfoCVM:           NewBasicInfoCVM(),
		HeadPortraitCVM:        NewHeadPortraitCVM(),
		JobIntensionCVM:        NewJobIntensionCVM(),
		ProjectExperienceCVM:   NewProjectExperienceCVM(),
		EducateExperienceCVM:   NewEducateExperienceCVM(),
		SelfEvaluationCVM:      NewSelfEvaluationCVM(),
		WorkExperienceCVM:      NewWorkExperienceCVM(),
		VolunteerExperienceCVM: NewVolunteerExperienceCVM(),
		TraineeExperienceCVM:   NewTraineeExperienceCVM(),
		HonorsCVM:              NewHonorsCVM(),
		SelfRecomdLetterCVM:    NewSelfRecomdLetterCVM(),
		WorksShowCVM:           NewWorksShowCVM(),
		InterestCVM:            NewInterestCVM(),
		SpecialityCVM:          NewSpecialityCVM(),
	}

	return cv
}

//基本信息模块默认构造函数
func NewBasicInfoCVM() *BasicInfoCVM {
	cvm := &BasicInfoCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_BASICINFO,
				CVMName:         "基本信息",
				ModuleDisplay:   true,
				ModuleRequisite: true,
			},
			TitleName:    "基本信息",
			TitleDisplay: true,
		},
	}
	return cvm
}

//头像模块默认构造函数
func NewHeadPortraitCVM() *HeadPortraitCVM {
	cvm := &HeadPortraitCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_HEADPORTRAIT,
				CVMName:         "头像",
				ModuleDisplay:   true,
				ModuleRequisite: false,
			},
			TitleName:    "头像",
			TitleDisplay: false,
		},
	}
	return cvm
}

//求职意向模块默认构造函数
func NewJobIntensionCVM() *JobIntensionCVM {
	cvm := &JobIntensionCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_JOBINTENSION,
				CVMName:         "求职意向",
				ModuleDisplay:   true,
				ModuleRequisite: true,
			},
			TitleName:    "求职意向",
			TitleDisplay: true,
		},
	}
	return cvm
}

//项目经历模块默认构造函数
func NewProjectExperienceCVM() *ProjectExperienceCVM {
	cvm := &ProjectExperienceCVM{
		TimeRangeCVM: &TimeRangeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_PROJECTEXPERIENCE,
					CVMName:         "项目经验",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "项目经验",
				TitleDisplay: true,
			},
			ExpTimeDisplay: true,
			ExpDescDisplay: true,
		},
	}
	return cvm
}

//教育经历模块默认构造函数
func NewEducateExperienceCVM() *EducateExperienceCVM {
	cvm := &EducateExperienceCVM{
		TimeRangeCVM: &TimeRangeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_EDUCATEEXPERIENCE,
					CVMName:         "教育经历",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "教育经历",
				TitleDisplay: true,
			},
			ExpTimeDisplay: true,
			ExpDescDisplay: true,
		},
	}
	return cvm
}

//自我评价模块默认构造函数
func NewSelfEvaluationCVM() *SelfEvaluationCVM {
	cvm := &SelfEvaluationCVM{
		DescribeCVM: &DescribeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_SELFEVALUATION,
					CVMName:         "自我评价",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "自我评价",
				TitleDisplay: true,
			},
		},
	}
	return cvm
}

//工作经历模块默认构造函数
func NewWorkExperienceCVM() *WorkExperienceCVM {
	cvm := &WorkExperienceCVM{
		TimeRangeCVM: &TimeRangeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_EDUCATEEXPERIENCE,
					CVMName:         "工作经历",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "工作经历",
				TitleDisplay: true,
			},
			ExpTimeDisplay: true,
			ExpDescDisplay: true,
		},
	}
	return cvm
}

//志愿者经历模块默认构造函数
func NewVolunteerExperienceCVM() *VolunteerExperienceCVM {
	cvm := &VolunteerExperienceCVM{
		WorkExperienceCVM: &WorkExperienceCVM{
			TimeRangeCVM: &TimeRangeCVM{
				TitleCVM: &TitleCVM{
					CVModule: &CVModule{
						CVMType:         CVMTYPE_EDUCATEEXPERIENCE,
						CVMName:         "志愿者经历",
						ModuleDisplay:   false,
						ModuleRequisite: false,
					},
					TitleName:    "志愿者经历",
					TitleDisplay: true,
				},
				ExpTimeDisplay: true,
				ExpDescDisplay: true,
			},
		},
	}
	return cvm
}

//实习经历模块默认构造函数
func NewTraineeExperienceCVM() *TraineeExperienceCVM {
	cvm := &TraineeExperienceCVM{
		WorkExperienceCVM: &WorkExperienceCVM{
			TimeRangeCVM: &TimeRangeCVM{
				TitleCVM: &TitleCVM{
					CVModule: &CVModule{
						CVMType:         CVMTYPE_EDUCATEEXPERIENCE,
						CVMName:         "实习经历",
						ModuleDisplay:   false,
						ModuleRequisite: false,
					},
					TitleName:    "实习经历",
					TitleDisplay: true,
				},
				ExpTimeDisplay: true,
				ExpDescDisplay: true,
			},
		},
	}
	return cvm
}

//荣誉奖项模块默认构造函数
func NewHonorsCVM() *HonorsCVM {
	cvm := &HonorsCVM{
		DescribeCVM: &DescribeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_HONORS,
					CVMName:         "荣誉奖项",
					ModuleDisplay:   false,
					ModuleRequisite: false,
				},
				TitleName:    "荣誉奖项",
				TitleDisplay: true,
			},
		},
	}
	return cvm
}

//自荐信模块默认构造函数
func NewSelfRecomdLetterCVM() *SelfRecomdLetterCVM {
	cvm := &SelfRecomdLetterCVM{
		DescribeCVM: &DescribeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_HONORS,
					CVMName:         "自荐信",
					ModuleDisplay:   false,
					ModuleRequisite: false,
				},
				TitleName:    "自荐信",
				TitleDisplay: true,
			},
		},
	}
	return cvm
}

//作品展示模块默认构造函数
func NewWorksShowCVM() *WorksShowCVM {
	cvm := &WorksShowCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_JOBINTENSION,
				CVMName:         "作品展示",
				ModuleDisplay:   false,
				ModuleRequisite: false,
			},
			TitleName:    "作品展示",
			TitleDisplay: true,
		},
	}
	return cvm
}

//兴趣爱好模块默认构造函数
func NewInterestCVM() *InterestCVM {
	cvm := &InterestCVM{
		LabelCVM: &LabelCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_JOBINTENSION,
					CVMName:         "兴趣爱好",
					ModuleDisplay:   false,
					ModuleRequisite: false,
				},
				TitleName:    "兴趣爱好",
				TitleDisplay: true,
			},
		},
	}
	return cvm
}

//技能特长模块默认构造函数
func NewSpecialityCVM() *SpecialityCVM {
	cvm := &SpecialityCVM{
		LabelCVM: &LabelCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_JOBINTENSION,
					CVMName:         "技能特长",
					ModuleDisplay:   false,
					ModuleRequisite: false,
				},
				TitleName:    "技能特长",
				TitleDisplay: true,
			},
		},
	}
	return cvm
}
