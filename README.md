# Destiny Activity Scraper

Scrapes PGCRS - activities from the Destiny 2 API - and ingest them into the database. This does not attempt to scrape _all_ PGCRs that happen, but just a subset while still staying reasonably up to date

## Context

The Destiny API publishes PGCRs (Post Game Carnage Report) for each activity/game/match played in the Destiny 2 world. When a player finishes an activity in Destiny, the game servers assign it a statically increasing ID and make it available through the API.

Because of this, you can constantly request the latest PGCRs by just fetching them in a for loop, incrementing the ID each time:

```
id = 100
while (true) {
    id += 1
    fetchActivity(id)
}
```

There are between 60-100 PGCRs generated a second, but this application only attempts to scrape a subset of them. The Bungie.net API has a per-IP rate limit of about 25 request/second.

If you request activities fast enough, eventually you will 404 as you hit an ID that hasnt been generated yet. When this happens, you just need to wait for a bit and retry a few seconds later when it should be generated.

Occasionally the game servers will delay creating the PGCR (and assigning it an ID), sometimes by a few hours. The effect of this is e.g. PGCR 100 ended at 06:00, PGCR 101 ended at 02:00, and PGCR 102 ended at 06:01. This can affect "catch up logic", so it must be handled specially.

Even rarer, but still "sometimes", will a PGCR ID just be completely skipped by the system. This also needs to be handled so the scraper does not become stuck on a single ID.
