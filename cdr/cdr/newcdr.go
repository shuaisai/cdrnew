package cdr

import (
	"github.com/shopspring/decimal"
	"strconv"
)

type CdrNew struct {
	Id         string
	SmsType    string
	UserType   string
	FeeDn      string
	SpNum      string
	SpSmg      string
	CalledDn   string
	SpId       string
	ServiceId  string
	FeeType    string
	Fee        string
	FeeMonth   string
	FeeFree    string
	FeeInstead string
	MoMt       string
	SmsStatus  string
	SmsFrist   string
	SmsNum     string
	FeeArea    string
	GateId     string
	GateId1    string
	SmsGT      string
	StartTime  string
	EndTime    string
	Extend     string
}

type ApkFee struct {
	Dn    string
	Money float64
	Id    string
}

type Cdrslice []CdrNew

func (cdr *CdrNew) AddItem(id, time, dn, fee string) {
	cdr.Id = id
	cdr.FeeDn = "86" + dn
	cdr.CalledDn = "86" + dn
	cdr.Fee = fee
	cdr.StartTime = time
	cdr.EndTime = time
}

func (apkfee ApkFee) FeeCalculate() []string {
	var feelist []string
	//intmoney := int(apkfee.Money * 100)
	intmoney := decimal.NewFromFloat(apkfee.Money).Mul(decimal.NewFromFloat(100)).IntPart()
	loop := true
	for loop {
		if intmoney > 999999 {
			intmoney -= 999999
			feelist = append(feelist, "999999")
		} else {
			endmoney := strconv.FormatInt(intmoney, 10)
			endmoney1 := ""
			switch len(endmoney) {
			case 6:
				endmoney1 = endmoney
			case 5:
				endmoney1 = "0" + endmoney
			case 4:
				endmoney1 = "00" + endmoney
			case 3:
				endmoney1 = "000" + endmoney
			case 2:
				endmoney1 = "0000" + endmoney
			case 1:
				endmoney1 = "00000" + endmoney
			}
			feelist = append(feelist, endmoney1)
			loop = false
		}
	}
	return feelist
}

func NewCdr() *CdrNew {
	return &CdrNew{
		//Id:         "",
		SmsType:  "01",
		UserType: "0",
		//feeDn:      "",
		SpNum: "12345",
		SpSmg: "00001",
		//calledDn:   "",
		SpId:      "123456789012",
		ServiceId: "1234567890",
		FeeType:   "2",
		//fee:        "",
		FeeMonth:   "0     ",
		FeeFree:    "0     ",
		FeeInstead: "0",
		MoMt:       "2",
		SmsStatus:  "0   ",
		SmsFrist:   "5",
		SmsNum:     "1 ",
		FeeArea:    "0771",
		GateId:     "07715",
		GateId1:    "00001",
		SmsGT:      "8613010591500",
		//startTime:  "",
		//endTime:    "",
		Extend: "              ",
	}
}
