package strategy_test

import (
	"strategy"
	"testing"
)

func TestStrategy(t *testing.T) {
	basketballScores := []strategy.RankItem{
		{Name: "John", Score: 100},
		{Name: "Bob", Score: 50},
		{Name: "Alice", Score: 60},
		{Name: "Carol", Score: 70},
	}
	basketballScoreBoard := strategy.NewRankList(10, &strategy.InsertSort{})
	for _, v := range basketballScores {
		basketballScoreBoard.AddItem(v.Name, v.Score)
	}
	t.Logf("basketball scoreboard: %v", basketballScoreBoard.Items())

	marathonScores := []strategy.RankItem{
		{Name: "John", Score: 100},
		{Name: "Bob", Score: 50},
		{Name: "Alice", Score: 60},
		{Name: "Carol", Score: 70},
		{Name: "David", Score: 80},
		{Name: "Eric", Score: 90},
		{Name: "Frank", Score: 80},
		{Name: "George", Score: 90},
		{Name: "Harry", Score: 80},
		{Name: "Ivan", Score: 80},
		{Name: "Jerry", Score: 80},
		{Name: "Kevin", Score: 80},
		{Name: "Larry", Score: 80},
		{Name: "Mary", Score: 80},
	}
	marathonScoresBoard := strategy.NewRankList(20, &strategy.QuickSort{})
	for _, v := range marathonScores {
		marathonScoresBoard.AddItem(v.Name, v.Score)
	}
	t.Logf("marathon scoreboard: %v", marathonScoresBoard.Items())
}
