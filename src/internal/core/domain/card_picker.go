package domain

type CardPicker struct {
	GroupName       string        `validate:"required,min=1,max=128"`
	ExpansionName   string        `validate:"required,default_card_expansion"`
	CardName        string        `validate:"required,default_card_name=ExpansionName"`
	CollectorNumber string        `validate:"required"`
	CardCondition   CardCondition `validate:"required,card_condition"`
	Language        string        `validate:"required"`
}
