package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hako/durafmt"
	log "github.com/sirupsen/logrus"
)

func requestPgcr(pgcrID int) (PGCRResponse, error) {
	url := fmt.Sprintf("https://stats.bungie.net/Platform/Destiny2/Stats/PostGameCarnageReport/%v/", pgcrID)

	pgcrHTTPClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("ERROR: Unable to create http.Request, error: %s", err.Error())
		return PGCRResponse{}, err
	}

	apiKey := os.Getenv("BUNGIE_API_KEY")

	// data.destinysets.com api key
	req.Header.Set("x-api-key", apiKey)

	res, getErr := pgcrHTTPClient.Do(req)
	if getErr != nil {
		log.Errorf("ERROR: Unable to do request, error: %v", getErr)
		return PGCRResponse{}, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Errorf("ERROR: Unable to read body, error: %v", readErr)
		return PGCRResponse{}, readErr
	}

	pgcr := PGCRResponse{}
	jsonErr := json.Unmarshal(body, &pgcr)

	if jsonErr != nil {
		log.Errorf("ERROR: Unable to parse JSON body, error: %v", jsonErr.Error())
		log.Errorf("Body: %q", string(body))
		return PGCRResponse{}, jsonErr
	}

	return pgcr, nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	standardActivitesPerSecond := 60 // should be even
	failedPgcrCount := 0
	hasEverBeenUpToDate := false

	var previousPgcr PGCRResponse
	pgcrID := 9405947263
	log.WithField("PGCR", pgcrID).Info("Starting up")

	for {
		logPgcr := log.WithField("PGCR", pgcrID)

		pgcr, err := requestPgcr(pgcrID)
		if (err != nil) || (pgcr.ErrorCode != 1) {
			failedPgcrCount += 1
		} else {
			failedPgcrCount = 0
		}

		// Handle actual exceptions seperatly from API errors
		if err != nil {
			sleepFor := time.Duration(failedPgcrCount) * time.Second
			logPgcr.WithFields(log.Fields{"err": err, "sleepDuration": sleepFor, "failCount": failedPgcrCount}).Errorf("Exception when requesting PGCR")

			time.Sleep(sleepFor)
			continue
		}

		// ErrorCode=1 means OK
		// TODO: Occasionally (about once a year) the API might go down for up to 24 hours for maintance.
		// During this time there's a higher chance that PGCR IDs will be skipped
		if pgcr.ErrorCode != 1 {
			// No need to stress so much if we get Not found - we expect that to happen
			if pgcr.ErrorStatus != "DestinyPGCRNotFound" {
				logPgcr.WithFields(log.Fields{"ErrorStatus": pgcr.ErrorStatus, "ErrorMessage": pgcr.Message, "failCount": failedPgcrCount}).Errorf("Bungie API error when requesting PGCR")
			} else {
				logPgcr.WithField("failCount", failedPgcrCount).Warnf("PGCR not found yet")
			}

			if PGCRExists(previousPgcr) {
				// Sometimes we'll overshoot trying to find the newness edge of PGCRs
				// so if we have a previous PGCR, backtrack 10 at a time.
				goBack := 10 * failedPgcrCount
				prevPgcrId := getPgcrID(previousPgcr)

				if pgcrID-goBack > prevPgcrId {
					pgcrID = pgcrID - goBack
				} else if failedPgcrCount > 10 {
					// If this failed too many time we may have just gotten stuck on a PGCR that won't
					// exist for a while (or ever), so just skip it
					logPgcr.WithField("failCount", failedPgcrCount).Warnf("Exceeded retry count, so skipping")
					pgcrID += 1
				}
			}

			time.Sleep(1 * time.Second)
			continue
		}

		//
		// PGCR exists!
		//
		endTime := getPgcrEndTime(pgcr)
		endedTimeAgo := time.Since(endTime)

		logPgcr.WithField("endedAgo", endedTimeAgo).Infof("PGCR ended %v ago", fmtDur(endedTimeAgo))

		increasePgcrIDBy := 1

		if endedTimeAgo > time.Minute*2 {
			// If we've fallen behind, then just skip a bunch by trying to guess the activities per second
			increasePgcrIDBy = int(endedTimeAgo.Seconds()) * standardActivitesPerSecond

			if PGCRExists(previousPgcr) {
				pgcrsPerSecond := getPgcrsPerSecond(pgcr, previousPgcr)
				pgcrsPerSecond = Max(pgcrsPerSecond, 10)

				increasePgcrIDBy = int(endedTimeAgo.Seconds()) * int(math.Floor(float64(pgcrsPerSecond/2)))

				// Sometimes activities get delayed and can throw out our catch up logic
				// If we're not just starting up, then cap how much we move forward by.
				if increasePgcrIDBy > 1000 && hasEverBeenUpToDate {
					increasePgcrIDBy = 1000
				}
			}

			logPgcr.WithField("increasePgcrIDBy", increasePgcrIDBy).Debugf("Catching up")
		} else {
			hasEverBeenUpToDate = true
		}

		pgcrID += increasePgcrIDBy
		previousPgcr = pgcr
	}
}

func getPgcrsPerSecond(pgcr PGCRResponse, prevPgcr PGCRResponse) int {
	if !prevPgcr.Response.Period.Before(pgcr.Response.Period) || !(getPgcrID(pgcr) > getPgcrID(prevPgcr)) {
		return 0
	}

	sincePrevPgcr := pgcr.Response.Period.Sub(prevPgcr.Response.Period)
	pgcrsSince := getPgcrID(pgcr) - getPgcrID(prevPgcr)

	pgcrsPerSecond := pgcrsSince / int(sincePrevPgcr.Seconds())

	return pgcrsPerSecond
}

func getPgcrEndTime(pgcr PGCRResponse) time.Time {
	durationSeconds := pgcr.Response.Entries[0].Values["activityDurationSeconds"].Basic.Value
	activityDuration := time.Duration(durationSeconds * 1000000000)
	endTime := pgcr.Response.Period.Add(activityDuration)

	return endTime
}

func getPgcrID(pgcr PGCRResponse) int {
	_prevPgcrId, _ := strconv.ParseUint(pgcr.Response.ActivityDetails.InstanceID, 10, 64)
	return int(_prevPgcrId)
}

func fmtDur(dur time.Duration) string {
	return durafmt.Parse(dur).LimitFirstN(2).String()
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func WritePgcrID(pgcrID string) {
	fileBytes := []byte(pgcrID)
	err := os.WriteFile("./lastPGCR.txt", fileBytes, 0644)

	if err != nil {
		log.Errorf("Failed to write lastPGCR.txt")
	}
}

func PGCRExists(pgcr PGCRResponse) bool {
	return pgcr.Response.ActivityDetails.InstanceID != ""
}
