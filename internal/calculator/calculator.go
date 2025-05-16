package calculator

import "math"

func monthlyRate(yearRate float64) float64 {
	return yearRate / 12 / 100
}

func CalculateAnnuityPayment(term int, rate float64, sumCredit float64) float64 {

	monthlyRate := monthlyRate(rate)

	monthlyPayment := sumCredit * ((monthlyRate * math.Pow((1+monthlyRate), float64(term))) / (math.Pow((1+monthlyRate), float64(term)) - 1))

	return monthlyPayment

}

func CalculateDifferentPayment(term int, rate float64, sumCredit float64) float64 {

	mainDebt := sumCredit / float64(term)
	currentMonth := 1

	for term > 0 {
		monthPercentage := (sumCredit - mainDebt*(float64(currentMonth)-1)*(float64(rate/100))) / 12

		term--
	}

}
