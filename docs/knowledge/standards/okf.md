---
type: Reference
title: Open Knowledge Format conformance
description: What OKF v0.2 requires, which parts this bundle uses, and the local rules for adding or editing concepts.
tags: [standards, okf, knowledge, markdown, agent]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
sources:
  - id: okf-spec
    resource: https://github.com/GoogleCloudPlatform/knowledge-catalog/blob/main/okf/SPEC.md
    title: Open Knowledge Format (OKF) SPEC.md, version 0.2
    author: team:google-cloud
    last_modified: 2026-08-21T00:00:00Z
---

# What OKF is

The **Open Knowledge Format** (OKF) is an open specification from Google Cloud
published in June 2026 that formalizes the *LLM-wiki* pattern: a knowledge
bundle is a directory tree of markdown files with YAML frontmatter.[^okf-spec]
There is no schema registry, no central authority, and no required tooling -
if a reader can `cat` a file, it can read OKF.

This bundle at `docs/knowledge/` is an OKF **v0.2** knowledge bundle about this
repository.

# What v0.2 requires

From the spec's conformance section, a bundle is conformant when:

1. Every non-reserved `.md` file contains a parseable YAML frontmatter block.
2. Every frontmatter block contains a non-empty `type` field.
3. Every reserved filename (`index.md`, `log.md`) follows its structure when
   present.

Everything else is optional. Consumers must not reject a bundle for missing
optional fields, unknown `type` values, unknown keys, broken links, or missing
`index.md` files.

# Conventions this bundle follows

| Rule | Detail |
|------|--------|
| Reserved filenames | `index.md` and `log.md` are never concepts; all other `.md` files are concepts. |
| Root `index.md` | The only `index.md` allowed frontmatter; it carries `okf_version: "0.2"`. |
| Subdirectory `index.md` | Plain markdown list, no frontmatter. |
| `log.md` | Date headings in `## YYYY-MM-DD` form, newest first, bold leading word per entry. |
| Links | Bundle-relative `/path/to/concept.md` form, resolved from the bundle root. |
| Timestamps | ISO 8601 with an explicit UTC offset, e.g. `2026-10-11T00:00:00Z`. |
| Actors | `human:<id>`, `process:<id>`, or `<producer>/<version>`. All hand-authored concepts here use `human:frisco`. |
| Attribution | Footnote labels match `sources[].id`, so claims join to their source by key rather than by position. |
| Trust | `generated` records authorship; `verified` records review. Both are populated because this content is hand-written and reviewed. |
| Lifecycle | `status` is `stable` for content describing shipped code, `draft` for content describing unfinished work. |

# Types used in this bundle

| `type` | Concepts |
|--------|----------|
| `Reference` | [Overview](/overview.md), [configuration](/operations/configuration.md), this page |
| `Architecture` | [Architecture](/architecture.md) |
| `API Endpoint` | [Health check](/api/health.md) |
| `Component` | [config](/components/config.md), [server](/components/server.md), [middleware](/components/middleware.md), [routes](/components/routes.md), [handlers](/components/handlers.md) |
| `Playbook` | [run-local](/operations/run-local.md), [containerization](/operations/containerization.md) |

Types are not registered centrally; consumers must tolerate unknown values.

# Adding a concept

1. Create the file in the directory that fits. `.md` files other than
   `index.md` and `log.md` are concepts.
2. Add frontmatter with at least `type`, plus `title`, `description`, `tags`,
   `status`, `generated`, and `verified`:

   ```yaml
   ---
   type: Component
   title: Short display name
   description: One sentence used by indexes and search snippets.
   tags: [example]
   status: draft
   generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
   ---
   ```

3. Write the body with structural markdown (headings, lists, tables, fenced
   code blocks) and link related concepts with bundle-relative paths.
4. Add an entry to the parent directory's `index.md`:
   `* [Title](/path/concept.md) - description`.
5. Append a `## YYYY-MM-DD` entry to `log.md`.

# Validating

There is no validator in this repository. Use the public OKF validator and
checker, or verify the three conformance rules by hand:

```bash
# Every non-reserved concept must declare a type
find docs/knowledge -name '*.md' ! -name 'index.md' ! -name 'log.md' | while read -r f; do
  head -5 "$f" | grep -q '^type:' || echo "MISSING type: $f"
done

# No index.md except the root may carry frontmatter
find docs/knowledge -mindepth 2 -name 'index.md' | while read -r f; do
  head -1 "$f" | grep -q '^---' && echo "UNEXPECTED frontmatter: $f"
done
```

# Related concepts

* [Root index](/index.md) - the bundle entry point.
* [Knowledge update log](/log.md) - change history for this bundle.

[^okf-spec]: Open Knowledge Format (OKF) SPEC.md, version 0.2
