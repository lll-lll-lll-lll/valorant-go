package valorantgo

import "time"

type MatchIDsRes struct {
	CurrentTime time.Time `json:"currentTime"`
	IDs         []string  `json:"matchIds"`
}

// API /val/content/v1/contents
type ContentDto struct {
	Version      string           `json:"version"`
	Characters   []ContentItemDto `json:"characters"`
	Maps         []ContentItemDto `json:"maps"`
	Chromas      []ContentItemDto `json:"chromas"`
	Skins        []ContentItemDto `json:"skins"`
	SkinLevels   []ContentItemDto `json:"skinLevels"`
	Equips       []ContentItemDto `json:"equips"`
	GameModes    []ContentItemDto `json:"gameModes"`
	Sprays       []ContentItemDto `json:"sprays"`
	SprayLevels  []ContentItemDto `json:"sprayLevels"`
	Charms       []ContentItemDto `json:"charms"`
	CharmLevels  []ContentItemDto `json:"charmLevels"`
	PlayerCards  []ContentItemDto `json:"playerCards"`
	PlayerTitles []ContentItemDto `json:"playerTitles"`
	Acts         []ActDto         `json:"acts"`
}

type ContentItemDto struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDto `json:"localizedNames"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
	// This field is only included for maps and game modes. These values are used in the match response.
	AssetPath string `json:"assetPath"`
}

type LocalizedNamesDto struct {
	ArAE string `json:"ar-AE"`
	DeDE string `json:"de-DE"`
	EnGB string `json:"en-GB"`
	EnUS string `json:"en-US"`
	EsES string `json:"es-ES"`
	EsMX string `json:"es-MX"`
	FrFR string `json:"fr-FR"`
	IdID string `json:"id-ID"`
	ItIT string `json:"it-IT"`
	JaJP string `json:"ja-JP"`
	KoKR string `json:"ko-KR"`
	PlPL string `json:"pl-PL"`
	PtBR string `json:"pt-BR"`
	RuRU string `json:"ru-RU"`
	ThTH string `json:"th-TH"`
	TrTR string `json:"tr-TR"`
	ViVN string `json:"vi-VN"`
	ZhCN string `json:"zh-CN"`
	ZhTW string `json:"zh-TW"`
}

type ActDto struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDto `json:"localizedNames"`
	ID             string            `json:"id"`
	IsActive       bool              `json:"isActive"`
}

// API: /val/match/v1/matches/{matchId}

type MatchDto struct {
	MatchInfo    MatchInfoDto     `json:"matchInfo"`
	Players      []PlayerDto      `json:"players"`
	Coaches      []CoachDto       `json:"coaches"`
	Teams        []TeamDto        `json:"teams"`
	RoundResults []RoundResultDto `json:"roundResults"`
}

type MatchInfoDto struct {
	MatchID            string `json:"matchId"`
	MapID              string `json:"mapId"`
	GameLengthMillis   int    `json:"gameLengthMillis"`
	GameStartMillis    int64  `json:"gameStartMillis"`
	ProvisioningFlowID string `json:"provisioningFlowId"`
	IsCompleted        bool   `json:"isCompleted"`
	CustomGameName     string `json:"customGameName"`
	QueueID            string `json:"queueId"`
	GameMode           string `json:"gameMode"`
	IsRanked           bool   `json:"isRanked"`
	SeasonID           string `json:"seasonId"`
}

type PlayerDto struct {
	Puuid           string         `json:"puuid"`
	GameName        string         `json:"gameName"`
	TagLine         string         `json:"tagLine"`
	TeamId          string         `json:"teamId"`
	PartyId         string         `json:"partyId"`
	CharacterId     string         `json:"characterId"`
	Stats           PlayerStatsDto `json:"stats"`
	CompetitiveTier int            `json:"competitiveTier"`
	PlayerCard      string         `json:"playerCard"`
	PlayerTitle     string         `json:"playerTitle"`
}

type PlayerStatsDto struct {
	Score          int             `json:"score"`
	RoundsPlayed   int             `json:"roundsPlayed"`
	Kills          int             `json:"kills"`
	Deaths         int             `json:"deaths"`
	Assists        int             `json:"assists"`
	PlaytimeMillis int             `json:"playtimeMillis"`
	AbilityCasts   AbilityCastsDto `json:"abilityCasts"`
}

type AbilityCastsDto struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

type CoachDto struct {
	PUUID  string `json:"puuid"`
	TeamID string `json:"teamId"`
}
type TeamDto struct {
	TeamId       string `json:"teamId"`
	Won          bool   `json:"won"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

type RoundResultDto struct {
	RoundNum              int                   `json:"roundNum"`
	RoundResult           string                `json:"roundResult"`
	RoundCeremony         string                `json:"roundCeremony"`
	WinningTeam           string                `json:"winningTeam"`
	BombPlanter           string                `json:"bombPlanter"`
	BombDefuser           string                `json:"bombDefuser"`
	PlantRoundTime        int                   `json:"plantRoundTime"`
	PlantPlayerLocations  []PlayerLocationsDto  `json:"plantPlayerLocations"`
	PlantLocation         LocationDto           `json:"plantLocation"`
	PlantSite             string                `json:"plantSite"`
	DefuseRoundTime       int                   `json:"defuseRoundTime"`
	DefusePlayerLocations []PlayerLocationsDto  `json:"defusePlayerLocations"`
	DefuseLocation        LocationDto           `json:"defuseLocation"`
	PlayerStats           []PlayerRoundStatsDto `json:"playerStats"`
	RoundResultCode       string                `json:"roundResultCode"`
}

type PlayerRoundStatsDto struct {
	PUUID   string      `json:"puuid"`
	Kills   []KillDto   `json:"kills"`
	Damage  []DamageDto `json:"damage"`
	Score   int         `json:"score"`
	Economy EconomyDto  `json:"economy"`
	Ability AbilityDto  `json:"ability"`
}

type PlayerLocationsDto struct {
	PUUID       string      `json:"puuid"`
	ViewRadians float64     `json:"viewRadians"`
	Location    LocationDto `json:"location"`
}

type LocationDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type FinishingDamageDto struct {
	DamageType          string `json:"damageType"`
	DamageItem          string `json:"damageItem"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
}

type KillDto struct {
	TimeSinceGameStartMillis  int                  `json:"timeSinceGameStartMillis"`
	TimeSinceRoundStartMillis int                  `json:"timeSinceRoundStartMillis"`
	Killer                    string               `json:"killer"`
	Victim                    string               `json:"victim"`
	VictimLocation            LocationDto          `json:"victimLocation"`
	Assistants                []string             `json:"assistants"`
	PlayerLocations           []PlayerLocationsDto `json:"playerLocations"`
	FinishingDamage           FinishingDamageDto   `json:"finishingDamage"`
}

type AbilityDto struct {
	GrenadeEffects  string `json:"grenadeEffects"`
	Ability1Effects string `json:"ability1Effects"`
	Ability2Effects string `json:"ability2Effects"`
	UltimateEffects string `json:"ultimateEffects"`
}

type DamageDto struct {
	Receiver  string `json:"receiver"`
	Damage    int    `json:"damage"`
	Legshots  int    `json:"legshots"`
	Bodyshots int    `json:"bodyshots"`
	Headshots int    `json:"headshots"`
}

type EconomyDto struct {
	LoadoutValue int    `json:"loadoutValue"`
	Weapon       string `json:"weapon"`
	Armor        string `json:"armor"`
	Remaining    int    `json:"remaining"`
	Spent        int    `json:"spent"`
}

// API: /val/match/v1/matchlists/by-puuid/%s

type MatchListDto struct {
	PUUID   string              `json:"puuid"`
	History []MatchListEntryDto `json:"history"`
}
type MatchListEntryDto struct {
	MatchID             string        `json:"matchid"`
	GameStartTimeMillis time.Duration `json:"gameStartTimeMillis"`
	TeamID              string        `json:"teamid"`
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}
