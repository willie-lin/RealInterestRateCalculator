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

	//println(len(initialInvestment))
	for i := 0; i < phase; i++ {

		initialInvestment = append(initialInvestment, repayment)

	}

	//fmt.Println(initialInvestment)/
	// 计算实际月均收益
	irr := income(initialInvestment)

	//fmt.Println(irr)

	// 计算实际年化收益率
	actualization := math.Pow(1+irr, 12) - 1.00

	//fmt.Println("年化利率为f%:", actualization)
	fmt.Printf("年化利率为: %.6f\n", actualization)
	// 月均还款(本金+利息)

	averageMonthlyRepayment := totalMoney * irr * math.Pow(1+irr, float64(phase)) / (math.Pow(1+irr, float64(phase)) - 1)

	// 还款利息总额

	totalRepaymentInterest := totalMoney*float64(phase)*irr*math.Pow(1+irr, float64(phase))/(math.Pow(1+irr, float64(phase))-1) - totalMoney

	// 还款总额为
	totalrepayment := totalMoney * float64(phase) * irr * math.Pow(1+irr, float64(phase)) / (math.Pow(1+irr, float64(phase)) - 1)

	//fmt.Println("您每月平均还款为%", averageMonthlyRepayment)
	fmt.Printf("您每月平均还款为: %.6f\n", averageMonthlyRepayment)
	//fmt.Println("----------------------------")
	fmt.Println("----------------------------")
	//fmt.Println("您的贷款总利息为f%", totalRepaymentInterest)
	fmt.Printf("您的贷款利息为: %.6f\n", totalRepaymentInterest)
	//fmt.Println("您的总还款额为f%", totalrepayment)
	fmt.Printf("您的总还款额为: %.6f\n", totalrepayment)

	fmt.Println("************************************")

	fmt.Printf("您的贷款总额为:%.6f,"+"总分期为: %d,"+"每期还款为：%.6f,"+" 利息如下:\n", totalMoney, phase, averageMonthlyRepayment)
	fmt.Printf("实际月利息 = %.6f%% ,"+"名义年利率 =  %.6f%%,"+"实际年利率 = %.6f%%,"+"\n", irr*100, irr*12*100, actualization*100)
	fmt.Printf("每月月供为为=%.6f,"+"总利息 = %.6f,"+"\n", averageMonthlyRepayment, totalRepaymentInterest)

	fmt.Println("***************************************************")

	// 第一个月还款利息
	firstinterest := totalMoney * irr
	// 剩余利息为
	//Residualinterest := totalRepaymentInterest - firstinterest
	// 剩余本金
	remainingprincipal := totalMoney - (averageMonthlyRepayment - firstinterest)
	fmt.Printf("\n\n ----- 等额本息计算, 以 %d 个月为例 -----\n", phase)
	fmt.Printf("第 %d 个月应还利息为: %.6f , 应还本金为: %.6f, 还款总额为: %.6f，剩余欠款: %.6f\n",
		1, firstinterest, averageMonthlyRepayment-firstinterest, averageMonthlyRepayment, remainingprincipal)

	//  第 n 个月还款
	for m := 2; m < phase+1; m++ {

		curInterest := (totalMoney*irr-averageMonthlyRepayment)*math.Pow(1+irr, float64(m-1)) + averageMonthlyRepayment
		curBase := averageMonthlyRepayment - curInterest
		remainingprincipal -= curBase
		fmt.Printf("第 %d 个月应还利息为: %.6f, 应还本金为: %.6f, 还款总额为: %.6f，剩余欠款: %.6f\n ",

			m, curInterest, curBase, averageMonthlyRepayment, math.Abs(remainingprincipal))

	}

	fmt.Println("------------------------------------------------------")

	fmt.Printf("\n\n ====假如 等额本金还款，以 %d 个月为例 ======\n", phase)

	monthlyPrincipalRepayment := totalMoney / float64(phase)

	tTotalRateMoney2 := float64(0)
	remainingmoney := totalMoney

	for i := 1; i < phase+1; i++ {
		currInterest := (totalMoney - monthlyPrincipalRepayment*float64(i-1)) * irr
		tTotalRateMoney2 += currInterest
		curTotalMoney := monthlyPrincipalRepayment + currInterest
		remainingmoney = remainingmoney - monthlyPrincipalRepayment

		fmt.Printf("第 %d 个月应还利息为: %.6f, 应还本金为: %.6f, 还款总额为: %.6f，剩余欠款: %.6f\n",
			i, currInterest, monthlyPrincipalRepayment, curTotalMoney, math.Abs(remainingmoney))

	}

	fmt.Printf("\n等额本金还款，总利息 = %.6f, 比等额本息少: %.6f\n", tTotalRateMoney2, totalRepaymentInterest-tTotalRateMoney2)

	// 每月应还本金

}

func income(initialInvestment []float64) float64 {

	//fmt.Println(len(initialInvestment))
	//fmt.Println()
	//var aaa = []float64(initialInvestment)
	//fmt.Println(aaa)
	//irr, err := financial.IRR([]float64{-20000, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789, 2789})
	irr, err := financial.IRR(initialInvestment)

	//println(irr)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("IRR is:", irr)
	fmt.Printf("IRR is: %.6f\n", irr)
	//return
	return irr
}

func main() {

	var (
		money     float64
		phase     int
		Repayment float64
	)

	fmt.Println("真实利率计算器！！！")

	fmt.Println("输入借款金额：")

	fmt.Scanf("%f", &money)

	fmt.Println("输入借款分期数：")

	fmt.Scanf("%d", &phase)

	fmt.Println("输入每月还款金额：")

	fmt.Scanf("%f", &Repayment)

	fmt.Println("你的实际贷款利率为")

	fmt.Println("..............................")

	annualInterestRate(money, phase, Repayment)

}
