package main

/*
# @Time    : 2020/07/17 23:33
# @Author  : ShuaiSai
# @Software: GoLand
*/

import (
	"awesomeProject/cdr/cdr"
	"awesomeProject/cdr/tools"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	t := time.Now().Add(time.Minute * -60).Format("2006-01-02 15")
	day := tools.TimeToString(time.Now().Format("2006-01-02"))
	rand.Seed(time.Now().UnixNano())

	apkfee := cdr.ApkFee{} //创建计费实例
	flag := ""

	for {
		fmt.Println("是否要生成话单，请按任意键继续，退出请按q:")
		fmt.Scanln(&flag)
		if flag == "q" {
			return
		}
		fmt.Println("请输入计费号码：")
		fmt.Scanln(&apkfee.Dn)
		fmt.Println("请输入金额：")
		fmt.Scanln(&apkfee.Money)
		fmt.Println("请输入cdr编号：")
		fmt.Scanln(&apkfee.Id)

		filepath := "D:\\陕西项目\\陕西小号计费相关\\" + "ALXH" + day + "000" + apkfee.Id + ".cdr"
		feelist := apkfee.FeeCalculate() //调用计费实例fee计算方法

		cdrslice := cdr.Cdrslice{} //创建切片实例
		cdr := cdr.NewCdr()        //使用工厂模式 返回初始值
		str := ""
		for _, fee := range feelist {
			// 生成随机数
			num := tools.RandInt64(1111111111111111, 9999999999999999)
			num1 := tools.RandInt64(10, 59)
			id := tools.TimeToString(t) + strconv.FormatInt(num1, 10) + "00" + strconv.FormatInt(num, 10)
			time := tools.TimeToString(t) + strconv.FormatInt(num1, 10) + "00"

			cdr.AddItem(id, time, apkfee.Dn, fee) //调用AddItem 为cdr设置值
			cdrslice = append(cdrslice, *cdr)
		}

		for _, feecdr := range cdrslice {
			str = feecdr.Id + feecdr.SmsType + feecdr.UserType + feecdr.FeeDn +
				feecdr.SpNum + feecdr.SpSmg + feecdr.CalledDn + feecdr.SpId + feecdr.ServiceId +
				feecdr.FeeType + feecdr.Fee + feecdr.FeeMonth + feecdr.FeeFree + feecdr.FeeInstead +
				feecdr.MoMt + feecdr.SmsStatus + feecdr.SmsFrist + feecdr.SmsNum + feecdr.FeeArea +
				feecdr.GateId + feecdr.GateId1 + feecdr.SmsGT + feecdr.StartTime + feecdr.EndTime +
				feecdr.Extend + "\n"

			tools.WriteToFile(filepath, str) //写文件
		}
	}
}
