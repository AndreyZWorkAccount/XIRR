package xirr

import (
	"time"
	"math"
	"NumericMethods/utility"
)

//NPV
func NetPresentValue(irrValue float64, payments []IPayment, firstPaymentDate *time.Time, daysInYear uint16 ) float64{
	if irrValue <= IrrMinValue{
		irrValue = IrrDefaultValue
	}
	var npv = 0.0
	for _,payment := range payments{
		npv += netPresentValueForSinglePeriod(irrValue,payment, firstPaymentDate, daysInYear)
	}
	return npv
}
func netPresentValueForSinglePeriod( irrValue float64, payment IPayment, startDate *time.Time, daysInYear uint16 )  float64 {
	var diffInDays = utility.DiffInDays(payment.Date(), startDate)
	periodNumber := diffInDays/float64(daysInYear)
	return payment.Amount()*math.Pow(1+irrValue, periodNumber)
}

//d(NPV)/dx
func NetPresentValueDerivative(irrValue float64, payments []IPayment, firstPaymentDate *time.Time, daysInYear uint16 ) float64{
	if irrValue <= IrrMinValue{
		irrValue = IrrDefaultValue
	}
	var npv = 0.0
	for _,payment := range payments{
		npv += netPresentValueDerivativeForSinglePeriod(irrValue,payment, firstPaymentDate, daysInYear)
	}
	return npv
}
func netPresentValueDerivativeForSinglePeriod(irrValue float64, payment IPayment, startDate *time.Time, daysInYear uint16) float64{
	diffInDays := utility.DiffInDays(payment.Date(), startDate)
	daysInYearF := float64(daysInYear)
	return  payment.Amount() * (1.0 / daysInYearF ) * diffInDays * math.Pow(1.0 + irrValue, (diffInDays/daysInYearF) - 1.0 );
}





