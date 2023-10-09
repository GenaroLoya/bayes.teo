package ballsBox

import (
	"fmt"
	"os"

	ft "bayes.teo/font.types"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Urn struct {
	RedBalls   int
	BlackBalls int
	ChooseProb float64
}

func Balls() {
	ft.Title.Println("Balls Box:")
	urns := map[string]Urn{
		"A": {3, 5, 0},
		"B": {2, 1, 0},
		"C": {2, 3, 0},
	}

	urnsCount := len(urns)

	for i, urn := range urns {
		urns[i] = Urn{urn.RedBalls, urn.BlackBalls, float64(1) / float64(urnsCount)}
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Urn", "Prob.", "Color", "Count", "Prob.", "Total Prob."})

	for i, urn := range urns {
		ProbBlack := float64(urn.BlackBalls) / float64(urn.RedBalls+urn.BlackBalls)
		ProbRed := float64(urn.RedBalls) / float64(urn.RedBalls+urn.BlackBalls)

		divUPBlack := urn.ChooseProb * ProbBlack
		divUPRed := urn.ChooseProb * ProbRed

		acumDownBlack := 0.0
		acumDownRed := 0.0

		for _, urn := range urns {
			acumDownBlack += urn.ChooseProb * (float64(urn.BlackBalls) / float64(urn.RedBalls+urn.BlackBalls))
			acumDownRed += urn.ChooseProb * (float64(urn.RedBalls) / float64(urn.RedBalls+urn.BlackBalls))
		}

		totalBlack := divUPBlack / acumDownBlack
		totalRed := divUPRed / acumDownRed

		t.AppendRow(table.Row{i, fmt.Sprintf("%.2f", urn.ChooseProb), "Black", urn.BlackBalls, fmt.Sprintf("%.2f", ProbBlack), fmt.Sprintf("%.2f", totalBlack)})
		t.AppendRow(table.Row{"", "", "Red", urn.RedBalls, fmt.Sprintf("%.2f", ProbRed), fmt.Sprintf("%.2f", totalRed)})
		t.AppendSeparator()
	}

	t.Render()
}
