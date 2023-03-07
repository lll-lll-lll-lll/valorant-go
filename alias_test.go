package valorantgo_test

import (
	"fmt"
	"testing"

	vg "github.com/lll-lll-lll-lll/valorant-go"
)

func Test_Code(t *testing.T) {
	t.Parallel()
	testcase := map[string]struct {
		queueCode vg.QueueCode
		want      string
	}{
		"competitive":    {queueCode: vg.Competitive, want: "competitive"},
		"unrated":        {queueCode: vg.Unrated, want: "unrated"},
		"spikerush":      {queueCode: vg.Spikerush, want: "spikerush"},
		"tournamentmode": {queueCode: vg.TournamentMode, want: "tournamentmode"},
		"deathmatch":     {queueCode: vg.DeathMatch, want: "deathmatch"},
		"onefa":          {queueCode: vg.OneFa, want: "onefa"},
		"ggteam":         {queueCode: vg.GGTeam, want: "ggteam"},
		"default":        {queueCode: 100, want: "competitive"},
	}
	for name, tt := range testcase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if tt.queueCode.String() != tt.want {
				t.Errorf("got %v, want %v", name, tt.want)
			}
		})
	}
}

func Test_URL_Queue(t *testing.T) {
	t.Parallel()
	testcase := map[string]struct {
		input      vg.ValAPI
		inputQueue string
		want       string
	}{
		"MatchInfoByMatchID": {
			input:      vg.MatchInfoByMatchID,
			inputQueue: "1",
			want:       "/val/match/v1/matches/1",
		},
		"empty": {
			input: "",
			want:  "",
		},
	}
	for name, tt := range testcase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := fmt.Sprint(tt.input, tt.inputQueue)
			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}

}
