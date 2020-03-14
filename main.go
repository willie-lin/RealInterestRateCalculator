package main

import (
	"fmt"
	"github.com/orcaman/financial"
	"math"
)


// 根据你的 贷款总额、分期数、还款金额 ， 计算一次性百分比手续费，计算年化利率

//var (
//	money float64
//	phase int
//	Repayment float64
//)

func annualInterestRate(totalMoney float64, phase int, repayment float64) {

	// total_money 	借款总额
	// phase		分期数
	//repayment		每期还款总额

	// 计算内部收益率
	// 初始投资金额
	initialInvestment := []float64{-totalMoney}

	println(len(initialInvestment))
	for i := 0; i < phase; i++ {

		initialInvestment = append(initialInvestment, repayment)

	}

	fmt.Println(initialInvestment)
	// 计算月均收益

	irr := income(initialInvestment)

	fmt.Println(irr)



	// 计算年化收益率
	actualization := math.Pow(1 + irr, 12) - 1.00

	fmt.Println("年化利率为f%:", actualization)
	// 月均还款(本金+利息)

	averageMonthlyRepayment := totalMoney * irr * math.Pow(1+irr, float64(phase)) / (math.Pow(1+irr, float64(phase)) - 1)

	// 还款利息总额

	totalRepaymentInterest := totalMoney*float64(phase)*irr*math.Pow(1+irr, float64(phase))/(math.Pow(1+irr, float64(phase))-1) - totalMoney

	// 还款总额为
	totalrepayment := totalMoney*float64(phase)*irr*math.Pow(1+irr, float64(phase)) / (math.Pow(1+irr, float64(phase))-1)

	fmt.Println("您每月平均还款为f%", averageMonthlyRepayment)
	fmt.Println("----------------------------")
	fmt.Println("您的贷款利息为f%", totalRepaymentInterest)
	fmt.Println("您的总还款额为f%", totalrepayment)

}



func income(initialInvestment []float64) float64 {

	fmt.Println(len(initialInvestment))
	fmt.Println()

	var aaa = []float64(initialInvestment)

	fmt.Println(aaa)
	//irr, err := financial.IRR([]float64{-20000, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789})
	irr, err := financial.IRR(initialInvestment)

	println(irr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("IRR is: %f", irr)
	//return
	return irr
}

func main() {

	var (
		money float64
		phase int
		Repayment float64
	)

	fmt.Println("真实利率计算器！！！")

	fmt.Println("输入借款金额：")

	fmt.Scanf("%f", &money)

	fmt.Println("输入借款分期数：")

	fmt.Scanf("%d", &phase)

	fmt.Println("输入每月还款金额：")

	fmt.Scanf("%f", &Repayment)

	fmt.Println("你的实际贷款利率为" )


	fmt.Println("..............................")

	annualInterestRate(money, phase, Repayment)


}