package structual

import "time"

const (
	// TEAMA A队
	TEAMA = iota
	// TEAMB B队
	TEAMB
)

// Team 队伍
type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

// Player 选手
type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

// HistoricalData 历史战绩
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

// Match 比赛记录
type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAMB:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

// TeamFlyweightFactory 结构
type TeamFlyweightFactory struct {
	createdTeams map[int]*Team
}

// NewTeamFactory 工厂
func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

// GetTeam 接口实现
func (t *TeamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.createdTeams[teamName] != nil {
		return t.createdTeams[teamName]
	}

	team := getTeamFactory(teamName)
	t.createdTeams[teamName] = &team
	return t.createdTeams[teamName]
}

// GetNumberOfObjects 接口实现
func (t *TeamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
