package valorantgo_test

import (
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
