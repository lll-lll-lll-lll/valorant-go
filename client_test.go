package valorantgo_test

import (
	"fmt"
	"testing"

	vg "github.com/lll-lll-lll-lll/valorant-go"
)

func Test_URL(t *testing.T) {
	t.Parallel()
	testcase := map[string]struct {
		input vg.ValAPI
		want  string
	}{
		"ContentsAboutVAL_URL": {
			input: vg.ContentsAboutVAL,
			want:  "/val/content/v1/contents",
		},
		"MatchInfoByMatchID": {
			input: vg.MatchInfoByMatchID,
			want:  "/val/match/v1/matches/",
		},
	}
	for name, tt := range testcase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tt.input
			if tt.want != string(got) {
				t.Errorf("want %v, got %v", tt.want, got)
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
