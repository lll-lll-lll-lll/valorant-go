package valorantgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrCtxIsNill    = errors.New("context is nil")
	ErrBindType     = errors.New("bind type is not supported")
	ErrResionIsNill = errors.New("resion is empty")
)

type ValConfig struct {
	APIKey string `env:"VALORANT_API_KEY,required"`
	Resion Resion `json:"resion"`
}

func NewValConfig(resion Resion) (*ValConfig, error) {
	if resion == "" {
		return nil, ErrResionIsNill
	}
	return &ValConfig{Resion: resion}, nil
}

type Valorant struct {
	Client *http.Client
	Config *ValConfig
}

func New(client *http.Client, cfg *ValConfig) *Valorant {
	return &Valorant{
		Client: client,
		Config: cfg,
	}
}

func (v *Valorant) DoToContents(ctx context.Context) (*ContentDto, error) {
	if ctx == nil {
		return nil, ErrCtxIsNill
	}
	url := buildURL(v.Config, ContentsAboutVAL)
	req, err := buildGETReq(url, v.Config.APIKey)
	if err != nil {
		return nil, err
	}
	rawResp, err := do(ctx, v.Client, req)
	if err != nil {
		return nil, err
	}
	resp, err := bind(rawResp, &ContentDto{})
	if err != nil {
		return nil, err
	}
	cd, ok := resp.(*ContentDto)
	if !ok {
		return nil, err
	}
	return cd, nil
}

func (v *Valorant) DoToMatchIDs(ctx context.Context, queue QueueCode) (*MatchIDsRes, error) {
	if ctx == nil {
		return nil, ErrCtxIsNill
	}
	url := buildURLWithPathParam(v.Config, MatchIDs, queue.String())
	req, err := buildGETReq(url, v.Config.APIKey)
	if err != nil {
		return nil, err
	}
	rawResp, err := do(ctx, v.Client, req)
	if err != nil {
		return nil, err
	}
	resp, err := bind(rawResp, &MatchIDsRes{})
	if err != nil {
		return nil, err
	}
	midr, ok := resp.(*MatchIDsRes)
	if !ok {
		return nil, err
	}
	return midr, nil
}

func (v *Valorant) DoToMatchInfo(ctx context.Context, matchID string) (*MatchDto, error) {
	if matchID == "" {
		return nil, errors.New("matchID cannot be empty")
	}
	if ctx == nil {
		return nil, ErrCtxIsNill
	}
	url := buildURLWithPathParam(v.Config, MatchInfoByMatchID, matchID)
	req, err := buildGETReq(url, v.Config.APIKey)
	if err != nil {
		return nil, err
	}
	rawResp, err := do(ctx, v.Client, req)
	if err != nil {
		return nil, err
	}
	resp, err := bind(rawResp, &MatchDto{})
	if err != nil {
		return nil, err
	}
	md, ok := resp.(*MatchDto)
	if !ok {
		return nil, err
	}
	return md, nil
}

func (v *Valorant) DoToMatchInfoList(ctx context.Context, puuid string) (*MatchListDto, error) {
	if puuid == "" {
		return nil, errors.New("puuid cannot be empty")
	}
	if ctx == nil {
		return nil, ErrCtxIsNill
	}
	url := buildURLWithPathParam(v.Config, MatchInfoListByPUUID, puuid)
	req, err := buildGETReq(url, v.Config.APIKey)
	if err != nil {
		return nil, err
	}
	rawResp, err := do(ctx, v.Client, req)
	if err != nil {
		return nil, err
	}
	resp, err := bind(rawResp, &MatchListDto{})
	if err != nil {
		return nil, err
	}
	mld, ok := resp.(*MatchListDto)
	if !ok {
		return nil, err
	}
	return mld, nil
}

func (v *Valorant) DoToSummoner(ctx context.Context, summonerName string) (*Summoner, error) {
	if summonerName == "" {
		return nil, fmt.Errorf("summoner name cannot be empty")
	}
	url := buildURLWithPathParam(v.Config, Summoners, summonerName)
	req, err := buildGETReq(url, v.Config.APIKey)
	if err != nil {
		return nil, err
	}
	rawResp, err := do(ctx, v.Client, req)

	if err != nil {
		return nil, err
	}
	resp, err := bind(rawResp, &Summoner{})
	if err != nil {
		return nil, err
	}
	s, ok := resp.(*Summoner)
	if !ok {
		return nil, err
	}
	return s, nil
}

func do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("valorant api request failed %v", err)
	}
	return resp, nil
}

func bind(res *http.Response, target any) (any, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	switch target.(type) {
	case *ContentDto:
		var cd ContentDto
		if err := json.Unmarshal(body, &cd); err != nil {
			return nil, fmt.Errorf("bind ContentDto error %v", err)
		}
		return &cd, nil
	case *MatchIDsRes:
		var mir MatchIDsRes
		if err := json.Unmarshal(body, &mir); err != nil {
			return nil, fmt.Errorf("bind MatchIDsRes error %v", err)
		}
		return &mir, nil
	case *MatchDto:
		var md MatchDto
		if err := json.Unmarshal(body, &md); err != nil {
			return nil, fmt.Errorf("bind MatchDto error %v", err)
		}
		return &md, nil
	case *MatchListDto:
		var mld MatchListDto
		if err := json.Unmarshal(body, &mld); err != nil {
			return nil, fmt.Errorf("bind MatchListDto error %v", err)
		}
		return &mld, nil
	case *Summoner:
		var s Summoner
		if err := json.Unmarshal(body, &s); err != nil {
			return nil, fmt.Errorf("bind Summoner error %v", err)
		}
		return &s, nil
	}
	return nil, ErrBindType
}

func buildGETReq(url, valAPIKey string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Riot-Token", valAPIKey)
	req.Header.Set("Accept-Charset", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://developer.riotgames.com")
	return req, nil
}

func buildURLWithPathParam(cfg *ValConfig, path ValAPI, param string) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = fmt.Sprintf("%s.api.riotgames.com", cfg.Resion)
	u.Path = fmt.Sprint(path, param)
	return u.String()
}

func buildURL(cfg *ValConfig, path ValAPI) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = fmt.Sprintf("%s.api.riotgames.com", cfg.Resion)
	u.Path = string(path)
	return u.String()
}
