# Task: replace day-stepper with month-grid calendar view (`/booking`)

- **Status:** in progress
- **Branch:** `dev-v0.1.7` (sequential +1 from merged `dev-v0.1.6`, PR #93)
- **Owner:** Nick (approved via Clauco's task on redma thread `calendar-view-booking-v0.1.7`, 2026-06-09T20:27:39Z)
- **Reviewer:** Clauco (contract: `.agents/pr-review-contract.md`)

## Goal

Replace the prev/next-day arrow stepper on `pages/booking/index.vue:184-209` with a month-grid date picker. Visitors see the full 30-day booking window at a glance instead of clicking through one day at a time.

## Locked design decisions (from Clauco's task brief — do not re-ask)

1. **Replace the day-stepper entirely.** No "calendar + stepper" hybrid.
2. **No availability dots on calendar cells in this pass.** Slots fetched only on day-click.

## Files

**New:**
- `composables/useDate.ts` — pure helpers (`toYYYYMMDD`, `buildMonthGrid`, `isSameDay`, `clampDate`).
- `components/booking/CalendarMonthGrid.vue` — month-grid picker, v-model on `selectedDate`.

**Modified:**
- `pages/booking/index.vue` — drop `changeDay`, `isPreviousDayDisabled`, `isNextDayDisabled`; replace date-selector block (lines 184-209) with `<CalendarMonthGrid>`. Drop inline `toYYYYMMDD` in favour of the composable import.

## Invariants from PR #93 (must survive)

- `timezoneIana` computed for POST body (raw IANA name).
- `timezone` computed for visible "Timezone: ..." label (underscores → spaces, "Unknown" fallback).
- `?topic=&summary=` deep-link prefill on `onMounted`.
- GA4 `getGaSessionInfo` clientId/sessionId on POST body.
- Real-time availability fetched client-side from `/api/booking/availability` (STRICT rule #4).
- No touches to `cloudbuild.yaml`, `backend/internal/gcal/`, `cmd/server/main.go`, Secret Manager refs, or `plugins/analytics.client.ts` cookie-consent gating.

## Pre-PR gate strategy

- **Toolchain state on workstation:** `gh` ✓ (dasiyes, repo scope); `node`/`npm` ✓; `go` ✗; `node_modules` ✗.
- **Frontend gates (`npm run lint`, `npm run generate`):** deferred to CI / owner — see "Notes / risks" in the PR body. `npm run lint` was already broken on `main` for a pre-existing eslint-config reason (Clauco flagged on PR #93 review). `node_modules` would require either `npm install` (network + a one-time setup) or owner running locally.
- **Manual interaction verification:** deferred to Clauco per task brief ("I'll capture a desktop + mobile screenshot during review... You don't need to provide screenshots yourself — verifying interaction is enough."). I'll run a self-verification pass against the source for the listed interaction points.
- **Backend:** untouched, no `go build` / `go test` needed.

## Implementation steps

1. Cut branch `dev-v0.1.7` from main.
2. **Commit 1:** `feat(booking): extract date helpers into composables/useDate.ts` — new `composables/useDate.ts` with the four pure helpers. `pages/booking/index.vue` swaps the inline `toYYYYMMDD` for the import. No behaviour change visible.
3. **Commit 2:** `feat(booking): add month-grid calendar view component` — new `components/booking/CalendarMonthGrid.vue` per spec. Not yet wired into the page.
4. **Commit 3:** `feat(booking): replace day-stepper with month-grid calendar view` — drop the prev/next-day block, mount `<CalendarMonthGrid>`, keep `formattedDate` as the "Available slots for ..." header.
5. Self-verify against the six interaction points from Clauco's brief.
6. Push, `gh pr create --body-file`, ping Clauco on the thread with the PR URL.

## Self-verification checklist (run against source before PR)

- [ ] Calendar renders current month, today highlighted as selected.
- [ ] Future-and-in-window days are clickable; past days and days beyond `maxDate` render as disabled.
- [ ] Day click emits `update:modelValue` → parent's `watch(selectedDate, fetchAvailability)` fires.
- [ ] Prev/next month nav does NOT change `selectedDate`; only `displayMonth` shifts.
- [ ] "Available slots for {{ formattedDate }}" header sits between calendar and slot grid.
- [ ] Visible "Timezone: ..." label still reads `America/New York` (or `Unknown` if Intl absent); POST body still sends raw IANA.
- [ ] `?topic=&summary=` deep-link still prefills `bookingDetails.notes` on mount.

## Out of scope (do not touch)

- Backend changes (booking availability, book, cancel, email, gcal package).
- Per-cell availability dots (needs backend batch endpoint first).
- Persisting selected date across refreshes.
- i18n of day/month names (keep `en-US`).
- Keyboard accessibility for cell navigation.
- `pages/booking-demo.vue` (leave as-is).

## Review section (to fill in after Clauco reviews)

<!-- Clauco's findings go here. -->
