package match

type stats struct {
	ParticipantID                   int  `json:"participantId"`
	Win                             bool `json:"win"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	Kills                           int  `json:"kills"`
	Deaths                          int  `json:"deaths"`
	Assists                         int  `json:"assists"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	KillingSprees                   int  `json:"killingSprees"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	DoubleKills                     int  `json:"doubleKills"`
	TripleKills                     int  `json:"tripleKills"`
	QuadraKills                     int  `json:"quadraKills"`
	PentaKills                      int  `json:"pentaKills"`
	UnrealKills                     int  `json:"unrealKills"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	VisionScore                     int  `json:"visionScore"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	TurretKills                     int  `json:"turretKills"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	ChampLevel                      int  `json:"champLevel"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	WardsKilled                     int  `json:"wardsKilled"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
}

type timeline struct {
	ParticipantID      int `json:"participantId"`
	CreepsPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"creepsPerMinDeltas"`
	XpPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"xpPerMinDeltas"`
	GoldPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"goldPerMinDeltas"`
	CsDiffPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"csDiffPerMinDeltas"`
	XpDiffPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"xpDiffPerMinDeltas"`
	DamageTakenPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"damageTakenPerMinDeltas"`
	DamageTakenDiffPerMinDeltas struct {
		TenEnd    float64 `json:"10-end,omitempty"`
		ThirtyEnd float64 `json:"30-end,omitempty"`
		Twenty30  float64 `json:"20-30,omitempty"`
		Ten20     float64 `json:"10-20"`
		Zero10    float64 `json:"0-10"`
	} `json:"damageTakenDiffPerMinDeltas"`
}
