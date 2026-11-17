# slayer-site

The landing page for [Slayer](https://github.com/janearc/slayer), the
open-source Go and Metal ray tracer by jane michelle arc. This site belongs
to jane michelle arc. Louie Sakoda maintains it, and will transfer it to
jane on request.

- `index.html` is the page. There is no build step for it.
- `assets/css/theme.css` holds every colour: Chroma's vaporwave colourway,
  solved onto the reader's contrast band by
  [libreadme](https://github.com/janearc/libreadme). Text is drawn only on
  `--ground` and `--code-ground`.
- `assets/frames/`, `assets/full/` and `assets/films/` hold pictures and
  films slayer drew, and `assets/full/` the 2K and 4K renders the contact
  sheet opens. Nothing on the page is a stand-in. Keep each file under
  10 MB.
- `catalogue/` is the manufactory's specimen book, in the site's design.
- `docs/` holds reading pages: `visual-guide.html` is copied from slayer's
  own guide at the commit it names; `ferruccio.html` is written for the
  site, from `docs/pages/`.
- Vercel deploys every push to `main`, and a preview of every pull request.

game holds what a person runs here:

    game run readable   the theme against the contrast band (CI runs it too)
    game run preview    the site at http://localhost:4174
    game run docs       docs/ rebuilt with pandoc from docs/src and docs/pages

When `readable` fails, the colour changes, not the threshold.
