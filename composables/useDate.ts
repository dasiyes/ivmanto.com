// composables/useDate.ts
//
// Pure date helpers for the booking page. No Vue / Nuxt imports — keeps the
// helpers trivially testable and reusable from both the page and the
// CalendarMonthGrid component. Extracted from pages/booking/index.vue as part
// of dev-v0.1.7 (month-grid calendar view).

/**
 * Format a Date as `YYYY-MM-DD` using its local time fields. Suitable for the
 * `date` query param on `/api/booking/availability`.
 */
export function toYYYYMMDD(date: Date): string {
  const year = date.getFullYear()
  const month = (date.getMonth() + 1).toString().padStart(2, '0')
  const day = date.getDate().toString().padStart(2, '0')
  return `${year}-${month}-${day}`
}

/**
 * Build the month grid for a calendar view, starting at the Sunday on-or-
 * before the 1st of `displayMonth` and ending at the Saturday on-or-after the
 * last day. The returned array contains 28, 35, or 42 cells (4–6 weeks ×
 * 7 cols): trailing all-out-of-month weeks are trimmed so single-month views
 * don't render an empty band beneath the last row. Out-of-month padding days
 * are included in the first row(s) so callers can render them as inert cells.
 */
export function buildMonthGrid(displayMonth: Date): Date[] {
  const firstOfMonth = new Date(displayMonth.getFullYear(), displayMonth.getMonth(), 1)
  const gridStart = new Date(firstOfMonth)
  gridStart.setDate(firstOfMonth.getDate() - firstOfMonth.getDay()) // Sunday = 0

  const cells: Date[] = []
  for (let i = 0; i < 42; i++) {
    const d = new Date(gridStart)
    d.setDate(gridStart.getDate() + i)
    cells.push(d)
  }
  // Trim trailing all-out-of-month weeks so the grid is 28/35/42 cells,
  // not always 42. Keeps the calendar tight against the last visible row.
  while (
    cells.length > 7
    && cells.slice(-7).every((d) => d.getMonth() !== displayMonth.getMonth())
  ) {
    cells.length -= 7
  }
  return cells
}

/**
 * True when `a` and `b` fall on the same calendar day in their local time
 * zone. Used for selected-day and today highlighting in the calendar grid.
 */
export function isSameDay(a: Date, b: Date): boolean {
  return (
    a.getFullYear() === b.getFullYear() &&
    a.getMonth() === b.getMonth() &&
    a.getDate() === b.getDate()
  )
}

/**
 * Return `date` clamped to the inclusive range [min, max]. If `date` is before
 * `min`, returns a copy of `min`; if after `max`, returns a copy of `max`;
 * otherwise returns a copy of `date` itself. None of the inputs are mutated.
 */
export function clampDate(date: Date, min: Date, max: Date): Date {
  const d = new Date(date.getFullYear(), date.getMonth(), date.getDate())
  const minD = new Date(min.getFullYear(), min.getMonth(), min.getDate())
  const maxD = new Date(max.getFullYear(), max.getMonth(), max.getDate())
  if (d.getTime() < minD.getTime()) return minD
  if (d.getTime() > maxD.getTime()) return maxD
  return d
}
