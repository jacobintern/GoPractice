package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type YBRes struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"status_code"`
	Status     string       `json:"status"`
	Data       []MatchModel `json:"data"`
}

type MatchModel struct {
	Cate      string `json:"cate"`
	Eid       string `json:"eid"`
	Game      string `json:"game"`
	IsStarted string `json:"is_started"`
	League    string `json:"league"`
	LeagueId  int    `json:"league_id"`
	Leagueen  string `json:"leagueen"`
	Score     string `json:"score"`
	Seid      string `json:"seid"`
	StartDate string `json:"start_date"`
	Team1     string `json:"team1"`
	Team1Logo string `json:"team1_logo"`
	Team1en   string `json:"team1en"`
	Team2     string `json:"team2"`
	Team2Logo string `json:"team2_logo"`
	Team2en   string `json:"team2en"`
	MatchesID int    `json:"matches_id"`
}

type Status struct {
	DocID string `json:"_doc"`
	ID    int    `json:"_id"`
	Name  string `json:"name"`
}

type Team struct {
	DocID      string `json:"_doc"`
	ID         int    `json:"_id"`
	SID        int    `json:"_sid"`
	UID        int    `json:"uid"`
	Virtual    bool   `json:"virtual"`
	Name       string `json:"name"`
	MediumName string `json:"mediumname"`
	Abbr       string `json:"abbr"`
	Nickname   string `json:"nickname"`
	IsCountry  bool   `json:"iscountry"`
	HasLogo    bool   `json:"haslogo"`
}

type Match struct {
	DocID string   `json:"_doc"`
	ID    int      `json:"_id"`
	Dt    TimeInfo `json:"_dt"`
	Teams struct {
		Home Team `json:"home"`
		Away Team `json:"away"`
	} `json:"teams"`
}

type TimeInfo struct {
	DocID string `json:"_doc"`
	Time  string `json:"time"`
	Date  string `json:"date"`
	Tz    string `json:"tz"`
}

type Tournament struct {
	DocID   string  `json:"_doc"`
	ID      int     `json:"_id"`
	Name    string  `json:"name"`
	Matches []Match `json:"matches"`
}

type RealCategory struct {
	DocID       string       `json:"_doc"`
	ID          int          `json:"_id"`
	Name        string       `json:"name"`
	Tournaments []Tournament `json:"tournaments"`
}

type Sport struct {
	DocID          string         `json:"_doc"`
	ID             int            `json:"_id"`
	Name           string         `json:"name"`
	Realcategories []RealCategory `json:"realcategories"`
}

type Events struct {
	Event string `json:"event"`
	Dob   int    `json:"_dob"`
	Data  Data   `json:"data"`
}

type Data struct {
	Sport Sport `json:"sport"`
}

type RadarList struct {
	QueryUrl string   `json:"queryUrl"`
	Doc      []Events `json:"doc"`
}

type newModel struct {
	QueryUrl           string
	EventsBob          int
	SportID            int
	SportName          string
	RealCategoryID     int
	RealCategoryName   string
	TournamementID     int
	TournamementName   string
	Date               string
	Time               string
	MatchesID          int
	HomeTeamID         int
	HomeTeamName       string
	HomeTeamMediumName string
	AwayTeamID         int
	AwayTeamName       string
	AwayTeamMediumName string
}

func MatchList(c *gin.Context) {

	// yb list
	list := GetList()

	// thirdpard data
	radarList := GetRadar()

	// find thirdpart match id
	res := FindIn(list, radarList)

	c.JSON(http.StatusOK, res)
}

func GetList() []MatchModel {
	ybURL := "https://video.bktest666.com/video/v2/api/list/yb.txt"

	// 发起 GET 请求
	response, err := http.Get(ybURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	res := new(YBRes)
	err = json.Unmarshal(body, &res)

	if err != nil {
		fmt.Println(err)
	}

	return res.Data
}

func GetRadar() []newModel {
	radarURL := "https://widgets.fn.sportradar.com/demolmt/en/Asia:Shanghai/gismo/sport_matches/1/2024-01-08"

	// 发起 GET 请求
	response, err := http.Get(radarURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	matches := new(RadarList)
	err = json.Unmarshal(body, &matches)

	if err != nil {
		fmt.Println(err)
	}

	res := matches.ToModel()

	return res
}

func (input RadarList) ToModel() []newModel {
	res := []newModel{}

	for _, event := range input.Doc {
		for _, r := range event.Data.Sport.Realcategories {
			for _, tn := range r.Tournaments {
				for _, m := range tn.Matches {
					res = append(res, newModel{
						QueryUrl:           input.QueryUrl,
						EventsBob:          event.Dob,
						SportID:            event.Data.Sport.ID,
						SportName:          event.Data.Sport.Name,
						RealCategoryID:     r.ID,
						RealCategoryName:   r.Name,
						TournamementID:     tn.ID,
						TournamementName:   tn.Name,
						Date:               m.Dt.Date,
						Time:               m.Dt.Time,
						MatchesID:          m.ID,
						HomeTeamID:         m.Teams.Home.ID,
						HomeTeamName:       m.Teams.Home.Name,
						HomeTeamMediumName: m.Teams.Home.MediumName,
						AwayTeamID:         m.Teams.Away.ID,
						AwayTeamName:       m.Teams.Away.Name,
						AwayTeamMediumName: m.Teams.Away.MediumName,
					})
				}
			}
		}
	}

	return res
}

func FindIn(input []MatchModel, list []newModel) []MatchModel {
	res := []MatchModel{}

	for _, yb := range input {
		if yb.Leagueen == "" {
			continue
		}

		for _, m := range list {
			replacer := strings.NewReplacer(",", "", "-", "", " ", "")
			tmpA := strings.ToLower(replacer.Replace(m.RealCategoryName + m.TournamementName))
			tmpB := strings.ToLower(replacer.Replace(yb.Leagueen))

			if tmpA == "" || tmpB == "" {
				continue
			}

			if tmpA == tmpB &&
				strings.Contains(m.HomeTeamMediumName, yb.Team1en) &&
				strings.Contains(m.AwayTeamMediumName, yb.Team2en) {
				yb.MatchesID = m.MatchesID
				res = append(res, yb)
			}
		}
	}

	return res

}
