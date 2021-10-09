package main

import "time"

type PGCRResponse struct {
	Response struct {
		Period             time.Time `json:"period"`
		StartingPhaseIndex int       `json:"startingPhaseIndex"`
		ActivityDetails    struct {
			ReferenceID          int    `json:"referenceId"`
			DirectorActivityHash int    `json:"directorActivityHash"`
			InstanceID           string `json:"instanceId"`
			Mode                 int    `json:"mode"`
			Modes                []int  `json:"modes"`
			IsPrivate            bool   `json:"isPrivate"`
			MembershipType       int    `json:"membershipType"`
		} `json:"activityDetails"`
		Entries []struct {
			Standing int `json:"standing"`
			Score    struct {
				Basic struct {
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
				} `json:"basic"`
			} `json:"score"`
			Player struct {
				DestinyUserInfo struct {
					IconPath                  string `json:"iconPath"`
					CrossSaveOverride         int    `json:"crossSaveOverride"`
					ApplicableMembershipTypes []int  `json:"applicableMembershipTypes"`
					IsPublic                  bool   `json:"isPublic"`
					MembershipType            int    `json:"membershipType"`
					MembershipID              string `json:"membershipId"`
					DisplayName               string `json:"displayName"`
				} `json:"destinyUserInfo"`
				CharacterClass string `json:"characterClass"`
				ClassHash      int64  `json:"classHash"`
				RaceHash       int64  `json:"raceHash"`
				GenderHash     int64  `json:"genderHash"`
				CharacterLevel int    `json:"characterLevel"`
				LightLevel     int    `json:"lightLevel"`
				EmblemHash     int64  `json:"emblemHash"`
			} `json:"player"`
			CharacterID string `json:"characterId"`
			Values      map[string]struct {
				Basic struct {
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
				} `json:"basic"`
			} `json:"values"`
			// Values      struct {
			// 	Assists struct {
			//
			// 	} `json:"assists"`
			// 	Completed struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"completed"`
			// 	Deaths struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"deaths"`
			// 	Kills struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"kills"`
			// 	OpponentsDefeated struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"opponentsDefeated"`
			// 	Efficiency struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"efficiency"`
			// 	KillsDeathsRatio struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"killsDeathsRatio"`
			// 	KillsDeathsAssists struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"killsDeathsAssists"`
			// 	Score struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"score"`
			// 	ActivityDurationSeconds struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"activityDurationSeconds"`
			// 	CompletionReason struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"completionReason"`
			// 	FireteamID struct {
			// 		Basic struct {
			// 			Value        int64  `json:"value"`
			// 			DisplayValue string `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"fireteamId"`
			// 	StartSeconds struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"startSeconds"`
			// 	TimePlayedSeconds struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"timePlayedSeconds"`
			// 	PlayerCount struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"playerCount"`
			// 	TeamScore struct {
			// 		Basic struct {
			// 			Value        float64 `json:"value"`
			// 			DisplayValue string  `json:"displayValue"`
			// 		} `json:"basic"`
			// 	} `json:"teamScore"`
			// } `json:"values"`
			// Extended struct {
			// 	Weapons []struct {
			// 		ReferenceID int64 `json:"referenceId"`
			// 		Values      struct {
			// 			UniqueWeaponKills struct {
			// 				Basic struct {
			// 					Value        float64 `json:"value"`
			// 					DisplayValue string  `json:"displayValue"`
			// 				} `json:"basic"`
			// 			} `json:"uniqueWeaponKills"`
			// 			UniqueWeaponPrecisionKills struct {
			// 				Basic struct {
			// 					Value        float64 `json:"value"`
			// 					DisplayValue string  `json:"displayValue"`
			// 				} `json:"basic"`
			// 			} `json:"uniqueWeaponPrecisionKills"`
			// 			UniqueWeaponKillsPrecisionKills struct {
			// 				Basic struct {
			// 					Value        float64 `json:"value"`
			// 					DisplayValue string  `json:"displayValue"`
			// 				} `json:"basic"`
			// 			} `json:"uniqueWeaponKillsPrecisionKills"`
			// 		} `json:"values"`
			// 	} `json:"weapons"`
			// 	Values struct {
			// 		PrecisionKills struct {
			// 			Basic struct {
			// 				Value        float64 `json:"value"`
			// 				DisplayValue string  `json:"displayValue"`
			// 			} `json:"basic"`
			// 		} `json:"precisionKills"`
			// 		WeaponKillsGrenade struct {
			// 			Basic struct {
			// 				Value        float64 `json:"value"`
			// 				DisplayValue string  `json:"displayValue"`
			// 			} `json:"basic"`
			// 		} `json:"weaponKillsGrenade"`
			// 		WeaponKillsMelee struct {
			// 			Basic struct {
			// 				Value        float64 `json:"value"`
			// 				DisplayValue string  `json:"displayValue"`
			// 			} `json:"basic"`
			// 		} `json:"weaponKillsMelee"`
			// 		WeaponKillsSuper struct {
			// 			Basic struct {
			// 				Value        float64 `json:"value"`
			// 				DisplayValue string  `json:"displayValue"`
			// 			} `json:"basic"`
			// 		} `json:"weaponKillsSuper"`
			// 		WeaponKillsAbility struct {
			// 			Basic struct {
			// 				Value        float64 `json:"value"`
			// 				DisplayValue string  `json:"displayValue"`
			// 			} `json:"basic"`
			// 		} `json:"weaponKillsAbility"`
			// 	} `json:"values"`
			// } `json:"extended"`
		} `json:"entries"`
		Teams []interface{} `json:"teams"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
