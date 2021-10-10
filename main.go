package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type matchTimeline struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}
type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}
type StatPerks struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}
type Selections struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}
type Styles struct {
	Description string       `json:"description"`
	Selections  []Selections `json:"selections"`
	Style       int          `json:"style"`
}
type Perks struct {
	StatPerks StatPerks `json:"statPerks"`
	Styles    []Styles  `json:"styles"`
}
type Participants struct {
	Assists                        int    `json:"assists"`
	BaronKills                     int    `json:"baronKills"`
	BountyLevel                    int    `json:"bountyLevel"`
	ChampExperience                int    `json:"champExperience"`
	ChampLevel                     int    `json:"champLevel"`
	ChampionID                     int    `json:"championId"`
	ChampionName                   string `json:"championName"`
	ChampionTransform              int    `json:"championTransform"`
	ConsumablesPurchased           int    `json:"consumablesPurchased"`
	DamageDealtToBuildings         int    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int    `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int    `json:"damageSelfMitigated"`
	Deaths                         int    `json:"deaths"`
	DetectorWardsPlaced            int    `json:"detectorWardsPlaced"`
	DoubleKills                    int    `json:"doubleKills"`
	DragonKills                    int    `json:"dragonKills"`
	FirstBloodAssist               bool   `json:"firstBloodAssist"`
	FirstBloodKill                 bool   `json:"firstBloodKill"`
	FirstTowerAssist               bool   `json:"firstTowerAssist"`
	FirstTowerKill                 bool   `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool   `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool   `json:"gameEndedInSurrender"`
	GoldEarned                     int    `json:"goldEarned"`
	GoldSpent                      int    `json:"goldSpent"`
	IndividualPosition             string `json:"individualPosition"`
	InhibitorKills                 int    `json:"inhibitorKills"`
	InhibitorTakedowns             int    `json:"inhibitorTakedowns"`
	InhibitorsLost                 int    `json:"inhibitorsLost"`
	Item0                          int    `json:"item0"`
	Item1                          int    `json:"item1"`
	Item2                          int    `json:"item2"`
	Item3                          int    `json:"item3"`
	Item4                          int    `json:"item4"`
	Item5                          int    `json:"item5"`
	Item6                          int    `json:"item6"`
	ItemsPurchased                 int    `json:"itemsPurchased"`
	KillingSprees                  int    `json:"killingSprees"`
	Kills                          int    `json:"kills"`
	Lane                           string `json:"lane"`
	LargestCriticalStrike          int    `json:"largestCriticalStrike"`
	LargestKillingSpree            int    `json:"largestKillingSpree"`
	LargestMultiKill               int    `json:"largestMultiKill"`
	LongestTimeSpentLiving         int    `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int    `json:"magicDamageTaken"`
	NeutralMinionsKilled           int    `json:"neutralMinionsKilled"`
	NexusKills                     int    `json:"nexusKills"`
	NexusLost                      int    `json:"nexusLost"`
	NexusTakedowns                 int    `json:"nexusTakedowns"`
	ObjectivesStolen               int    `json:"objectivesStolen"`
	ObjectivesStolenAssists        int    `json:"objectivesStolenAssists"`
	ParticipantID                  int    `json:"participantId"`
	PentaKills                     int    `json:"pentaKills"`
	Perks                          Perks  `json:"perks"`
	PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
	ProfileIcon                    int    `json:"profileIcon"`
	Puuid                          string `json:"puuid"`
	QuadraKills                    int    `json:"quadraKills"`
	RiotIDName                     string `json:"riotIdName"`
	RiotIDTagline                  string `json:"riotIdTagline"`
	Role                           string `json:"role"`
	SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int    `json:"spell1Casts"`
	Spell2Casts                    int    `json:"spell2Casts"`
	Spell3Casts                    int    `json:"spell3Casts"`
	Spell4Casts                    int    `json:"spell4Casts"`
	Summoner1Casts                 int    `json:"summoner1Casts"`
	Summoner1ID                    int    `json:"summoner1Id"`
	Summoner2Casts                 int    `json:"summoner2Casts"`
	Summoner2ID                    int    `json:"summoner2Id"`
	SummonerID                     string `json:"summonerId"`
	SummonerLevel                  int    `json:"summonerLevel"`
	SummonerName                   string `json:"summonerName"`
	TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
	TeamID                         int    `json:"teamId"`
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TripleKills                    int    `json:"tripleKills"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlaced"`
	Win                            bool   `json:"win"`
}
type Bans struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
type Baron struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Champion struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Dragon struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Inhibitor struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type RiftHerald struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Tower struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
type Objectives struct {
	Baron      Baron      `json:"baron"`
	Champion   Champion   `json:"champion"`
	Dragon     Dragon     `json:"dragon"`
	Inhibitor  Inhibitor  `json:"inhibitor"`
	RiftHerald RiftHerald `json:"riftHerald"`
	Tower      Tower      `json:"tower"`
}
type Teams struct {
	Bans       []Bans     `json:"bans"`
	Objectives Objectives `json:"objectives"`
	TeamID     int        `json:"teamId"`
	Win        bool       `json:"win"`
}
type Info struct {
	GameCreation       int64          `json:"gameCreation"`
	GameDuration       int            `json:"gameDuration"`
	GameEndTimestamp   int64          `json:"gameEndTimestamp"`
	GameID             int            `json:"gameId"`
	GameMode           string         `json:"gameMode"`
	GameName           string         `json:"gameName"`
	GameStartTimestamp int64          `json:"gameStartTimestamp"`
	GameType           string         `json:"gameType"`
	GameVersion        string         `json:"gameVersion"`
	MapID              int            `json:"mapId"`
	Participants       []Participants `json:"participants"`
	PlatformID         string         `json:"platformId"`
	QueueID            int            `json:"queueId"`
	Teams              []Teams        `json:"teams"`
	TournamentCode     string         `json:"tournamentCode"`
}

type summonerv4 struct {
	Accountid     string `json:"AccountId"`
	ProfileIconId int    `json:"ProfileIconId"`
	RevisionDate  int    `json:"RevisionDate"`
	Name          string `json:"Name"`
	Id            string `json:"Id"`
	Puuid         string `json:"Puuid"`
	SummonerLevel string `json:"SummonerLevel"`
}

func getMatchTimeline(matchId string, apiKey string) *matchTimeline {
	url := fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s", matchId, apiKey)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var match matchTimeline
	json.Unmarshal(body, &match)

	return &match
}

func getMatches(puuid string, apiKey string) []string {
	//change count param in url below to choose how many games you want, check rate limit first since all are used to call match via id after
	url := fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=50&api_key=%s", puuid, apiKey)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	stringBody := string(body)
	stringBody = strings.ReplaceAll(stringBody, "[", "")
	stringBody = strings.ReplaceAll(stringBody, "]", "")
	ascii := 34
	quoteTrim := string(ascii)
	stringBody = strings.ReplaceAll(stringBody, quoteTrim, "")
	split := strings.Split(stringBody, ",")
	return split
}

func getSummonerId(summonerName string, apiKey string) *summonerv4 {
	url := fmt.Sprintf("https://oc1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summonerName, apiKey)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	var summoner summonerv4
	json.Unmarshal(body, &summoner)
	return &summoner
}

func getChampIds() map[string]string {

	idToNameMap := make(map[string]string)
	//11.20.1 current version
	url := "https://ddragon.leagueoflegends.com/cdn/11.20.1/data/en_US/champion.json"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	data := result["data"].(map[string]interface{})

	for _, element := range data {
		ele := element.(map[string]interface{})
		id := fmt.Sprintf("%s", ele["key"])
		name := fmt.Sprintf("%v", ele["name"])
		idToNameMap[id] = name
	}

	return idToNameMap
}

func main() {

	apiKey := "" //redacted if you want to run it go and get a 1day key, they give them out for free :D rates are ok too for casual use

	ch := make(chan int)
	summoner := getSummonerId("", apiKey) //first string is ur in game name or whoev you want to search
	last3 := getMatches(summoner.Puuid, apiKey)

	var matches []*matchTimeline
	for _, match := range last3 {
		matches = append(matches, getMatchTimeline(match, apiKey))
	}
	champMaps := getChampIds()

	go func(map[string]string) {
		counter := make(map[string]int)
		for id := range ch {
			champ := champMaps[strconv.Itoa(id)]
			_, ok := champMaps[champ]
			if !ok {
				counter[champ]++
			}
		}

		for key, element := range counter {
			fmt.Println(key, " => ", element)
		}

	}(champMaps)

	for _, match := range matches {
		for _, team := range match.Info.Teams {
			for _, bans := range team.Bans {
				ch <- bans.ChampionID
			}
		}
	}
	close(ch)
  
  select{} //lol i didnt add any syncronization so if you want to see the results just let it run and nuke it manually or if you really want add a wait group or done channel
}
