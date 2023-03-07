package valorantgo

type QueueCode int

// Queue Type for MatchIDs API
const (
	Competitive QueueCode = iota + 1
	Unrated
	Spikerush
	TournamentMode
	DeathMatch
	OneFa
	GGTeam
)

// String default is "competitive"
func (qc QueueCode) String() string {
	switch qc {
	case Competitive:
		return "competitive"
	case Unrated:
		return "unrated"
	case Spikerush:
		return "spikerush"
	case TournamentMode:
		return "tournamentmode"
	case DeathMatch:
		return "deathmatch"
	case OneFa:
		return "onefa"
	case GGTeam:
		return "ggteam"
	default:
		return "competitive"
	}
}

type ValAPI string

// API Path
const (
	// Returns a list of match IDs completed within 12 hours
	MatchIDs             ValAPI = "/val/match/v1/recent-matches/by-queue/" // path param == queue
	ContentsAboutVAL     ValAPI = "/val/content/v1/contents"
	MatchInfoByMatchID   ValAPI = "/val/match/v1/matches/"              // path param == match_id
	MatchInfoListByPUUID ValAPI = "/val/match/v1/matchlists/by-puuid/"  // path param == puuid
	Summoners            ValAPI = "/lol/summoner/v4/summoners/by-name/" // path param == summoner name
)

type Resion string

const (
	JP1  Resion = "jp1"
	Asia Resion = "asia"
	BR1  Resion = "br1"
	EUN1 Resion = "eun1"
	EUW1 Resion = "euw1"
	KR   Resion = "kr"
	LA1  Resion = "la1"
	LA2  Resion = "la2"
	NA1  Resion = "na1"
	OC1  Resion = "oc1"
	PH2  Resion = "ph2"
	RU   Resion = "ru"
	SG2  Resion = "sg2"
	TH2  Resion = "th2"
	TR1  Resion = "tr1"
	TW2  Resion = "tw2"
	VN2  Resion = "vn2"
)
