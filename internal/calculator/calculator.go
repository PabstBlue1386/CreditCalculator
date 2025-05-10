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
