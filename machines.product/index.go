package machinesProduct

import (
	ft "bayes.teo/font.types"
)

func Machines() {
	ft.Title.Println("Machines Percentages:")
	// Definir las probabilidades iniciales
	probProbA := 0.45
	probProdB := 0.30
	probProdC := 0.25

	probFallsA := 0.03
	probFallsB := 0.04
	probFallsC := 0.05

	// a. Probabilidad de que una pieza sea defectuosa
	probFallsGeneral := probProbA*probFallsA +
		probProdB*probFallsB +
		probProdC*probFallsC

	ft.SubTitle.Println("Probabilidad defectuosa general:")
	ft.Text.Println(probFallsGeneral*100, "%")

	probProdFallA := (probProbA * probFallsA) / probFallsGeneral
	ft.SubTitle.Println("Probabilidad defectuosa producida por la máquina A:")
	ft.Text.Println(probProdFallA*100, "%")

	probProdFallB := (probProdB * probFallsB) / probFallsGeneral
	ft.SubTitle.Println("Probabilidad defectuosa producida por la máquina B:")
	ft.Text.Println(probProdFallB*100, "%")

	probProdFallC := (probProdC * probFallsC) / probFallsGeneral
	ft.SubTitle.Println("Probabilidad defectuosa producida por la máquina C:")
	ft.Text.Println(probProdFallC*100, "%")

	//Determianar la probabilidad mas alta
	fallsProdList := map[string]float64{
		"A": probProdFallA,
		"B": probProdFallB,
		"C": probProdFallC,
	}

	machineMax := ""
	probMax := 0.0
	for i, prob := range fallsProdList {

		if prob > probMax {
			probMax = prob
			machineMax = i
		}
	}
	ft.SubTitle.Println("Probabilidad más alta de defectuosos:")
	ft.Text.Println(probMax*100, "% de la máquina", machineMax)
}
