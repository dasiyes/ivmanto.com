# SEO Manual Fixes TODO
> Source: Semrush Site Audit — ivmanto.com — Generated March 19, 2026

---

## Errors (must fix)

### [ ] 2 structured data items are invalid
- **Amount:** 2
- **Issue:** Structured data items contain fields that do not meet Google's guidelines. Crawlers cannot properly understand the markup, risking loss of rich snippets and favorable rankings.
- **Fix:** Validate pages using the [Rich Results Test](https://search.google.com/test/rich-results) tool. Review against [schema.org](https://schema.org) and Google documentation. Correct any fields that fail validation.

---

## Warnings (should fix)

### [ ] 63 pages have low text-HTML ratio
- **Amount:** 63
- **Issue:** Text-to-HTML ratio is 10% or less. Search engines favor content-rich pages; more code than text also slows page load and crawling.
- **Fix:** Move embedded scripts and styles to separate files. Review page HTML structure and reduce inline code so text content outweighs markup.

### [ ] 14 pages have a low word count
- **Amount:** 14
- **Issue:** Fewer than 200 words on the page. Low word count is a negative quality signal to search engines.
- **Fix:** Expand on-page content to include more than 200 meaningful words per page.

### [ ] 6 pages have too much text within the title tags
- **Amount:** 6
- **Issue:** Title tags exceed 70 characters and will be truncated in search results, reducing click-through rates.
- **Fix:** Rewrite page titles to be 70 characters or fewer.

---

## Notices (good to fix)

### [ ] 19 links have non-descriptive anchor text
- **Amount:** 19
- **Issue:** Anchor text like "click here" or "right here" gives no context to search engines about the linked page, hurting indexing and rankings.
- **Fix:** Replace generic anchor text with succinct, descriptive text that reflects the content of the linked page.

### [ ] 8 links have no anchor text
- **Amount:** 8
- **Issue:** Links with empty, naked (URL-only), or special-character-only anchors make it hard for crawlers to understand the target page.
- **Fix:** Add short, descriptive anchor text to all links that currently have none.

### [ ] 2 pages are blocked from crawling
- **Amount:** 2
- **Issue:** Pages blocked via `robots.txt` or `noindex` meta tag will never appear in search results.
- **Fix:** Audit `robots.txt` and page-level `noindex` tags. Ensure no pages with valuable content are blocked by mistake.

### [ ] 2 orphaned pages in sitemaps
- **Amount:** 2
- **Issue:** Pages in `sitemap.xml` that have no internal links waste crawl budget and signal poor site structure.
- **Fix:** For each orphaned page: remove it from the sitemap if it's no longer needed, add internal links to it if the content is valuable, or leave it only if it serves a specific standalone need.

### [ ] 2 pages require content optimization
- **Amount:** 2
- **Issue:** Pages need on-page clarity improvements for better user engagement, snippet eligibility, and topical authority.
- **Fix:** Review the specific advice in the Semrush error details for each page and optimise accordingly.

### [ ] llms.txt not found
- **Amount:** 1
- **Issue:** No `llms.txt` file present. This file helps AI search engines understand the site's content.
- **Fix:** Create a `llms.txt` file in the website root directory with a short summary, guidance, and links to key resources. See [llmstxt.org](https://llmstxt.org) for format examples. Note: `public/llms.txt` already exists in the repo — verify it is being served correctly at the root.

### [ ] 1 page has only one incoming internal link
- **Amount:** 1
- **Issue:** A page with only one internal link has very little chance of ranking or receiving visits.
- **Fix:** Identify the under-linked page and add relevant internal links to it from other pages on the site.
