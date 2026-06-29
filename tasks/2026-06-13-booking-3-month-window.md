# Task: extend booking window to 3 months on `/booking`

- **Status:** in progress
- **Branch:** `dev-v0.1.8` (sequential +1 from merged `dev-v0.1.7`, PR #94)
- **Owner:** Nick (kickoff in chat, 2026-06-13 ~02:00; not yet Clauco-approved — night-shift work)
- **Reviewer:** Clauco (contract: `.agents/pr-review-contract.md`)

## Goal

The booking page (`/booking`) currently caps the navigable date range at
`today + 30 days` via `const bookingWindowDays = 30` in
`pages/booking/index.vue`. The AfB Consultancy Google Calendar has AfB events
populated through end of July but no further; the website's month-grid
calendar (introduced in PR #94) allows navigation up to today+30d and then
the next-month button is disabled. Result: visitors can see at most the
first ~3 weeks of July and the page then appears "empty" — the slots are
real, but the UI can't reach them.

Fix: bump `bookingWindowDays` to 90 (≈3 months) and align the AfB scheduler
to populate events that cover a 90-day rolling window.

## Locked design decisions (do not re-ask)

1. **90 days, not 120.** "3 months" interpreted as ~13 weeks, not a full
   quarter. 90 days fits `today → today + 90d` cleanly.
2. **Use the existing `bookingWindowDays` constant, not a config flag.** The
   page already uses it for `maxBookingDate`. No new prop, no new env var.
3. **Bump scheduler `WEEKS_AHEAD` to 12, not introduce a calendar-length
   constant.** The scheduler's create window stays anchored at
   `now + 4 weeks` (matches existing semantics — slots visible ≥ 4 weeks
   out) and just extends 4 → 12 weeks.
4. **Add identity-based dedup to the scheduler** while we're in there.
   Without it, the next Monday cron run would see "the day has 2 AfB
   events at e.g. 09:30 and 14:00", pick a non-overlapping slot at 11:00,
   and create a third. With the new per-day count guard (`already >= 2 →
   skip`), the scheduler is idempotent across runs and won't double-fill
   days that are already at capacity.
5. **Backfill now, don't wait for next Monday.** Run the scheduler
   manually after the source change so the calendar is full when the
   morning comes.
6. **No backend changes.** The availability handler
   (`backend/internal/booking/handler.go:handleGetAvailability`) just
   queries the calendar for the requested date; no implicit date cap.

## Files

**Modified:**
- `pages/booking/index.vue` — `bookingWindowDays` 30 → 90, plus a 4-line
  comment cross-referencing the scheduler.
- `~/.hermes/profiles/ivmo/scripts/afb_scheduler.py` — `WEEKS_AHEAD` 4 →
  12, identity-based dedup (per-day AfB count guard) in the main loop.
- `~/.hermes/profiles/ivmo/skills/ivmo/calendar-scheduler-automation/SKILL.md`
  — update the AfB-hardcoded-constants table (WEEKS_AHEAD 4 → 12) and
  note the new identity-based dedup behaviour.

**Not touched:**
- `components/booking/CalendarMonthGrid.vue` — the month-grid component
  already respects the parent's `:max-date` prop, so it scales for free
  to 90 days. Verified by reading lines 21-49 and 82-87.
- `backend/internal/booking/handler.go` — N/A.
- `composables/useDate.ts` — N/A.
- `cloudbuild.yaml`, `backend/internal/gcal/`, `cmd/server/main.go`
  startup, Secret Manager refs, `plugins/analytics.client.ts` consent
  gating — not in scope.

## Pre-PR gate strategy

- **`npm run generate`** — ran locally, **PASS**. 122 routes prerendered,
  exit 0, no errors. (Deferred tooling-gap from PR #94 no longer applies
  here — `node_modules/` is present on this workstation as of the
  night-shift.)
- **`npm run lint`** — DEFERRED. Still broken on `main` for the
  pre-existing eslint-config reason (flat `eslint.config.ts` vs pinned
  ESLint 8.57) flagged in PR #93 review. Not introduced by this PR.
  Tracked separately.
- **Manual page verification** — deferred to Clauco per the pattern from
  PR #94. The diff is one constant + 4 lines of comment; interaction
  surface is unchanged.
- **Backend** — untouched, no `go build` / `go test` needed.

## Backfill run

After the source change, ran:

```
python3 ~/.hermes/profiles/ivmo/scripts/afb_scheduler.py --weeks-ahead 12
```

Result: **Created 90 AfB events, 0 errors, 15 days skipped** (the 15
working days July 11 → July 31 that already had 2 AfB events each from
the previous run).

Post-run inspection (`afb_inspect.py --from 2026-06-13 --to 2026-10-15`):
**70 working days, June 15 → October 2, every one with exactly 2 AfB
events, 0 duplicates, 0 overlaps, 0 non-AfB conflicts.** The 90-day
window (today + 90 = Sept 11) is fully populated; the scheduler leaves
~3 weeks of headroom (→ Oct 2) for the next Monday cron to refill from.

## Self-verification checklist (run against source)

- [ ] `bookingWindowDays = 90` in `pages/booking/index.vue` (line 21 area).
- [ ] `WEEKS_AHEAD = 12` in `~/.hermes/profiles/ivmo/scripts/afb_scheduler.py`
      (line 45).
- [ ] Identity-based dedup: `existing_afb_per_day` precomputed and
      consulted before slot generation in the scheduler main loop.
- [ ] `today + 90` actually allows navigating to Sept 11. Verified by
      reading the `maxBookingDate` line in the page and the
      `isNextMonthDisabled` gate in the month-grid component
      (`CalendarMonthGrid.vue:82-87`).
- [ ] Weekend cells in the month-grid still render as clickable (existing
      behaviour). Days with no AfB events will show "There are no
      available slots for this day" (existing copy, unchanged). Out of
      scope to grey out weekends — would be a small follow-up if the
      owner wants it.
- [ ] `pages/booking-demo.vue` (the iframe-embed variant) is left
      untouched. It has its own data path and is not part of this fix.

## Out of scope (re-stated, not touched)

- Per-cell availability dots on the month grid (still deferred per the
  PR #94 brief).
- Persisting the selected date across refreshes.
- i18n of day/month names.
- `npm run lint` eslint-config fix (separate ticket).
- Weekend-cell grey-out.
