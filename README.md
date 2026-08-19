# bigo-action

Report asymptotic complexity regressions introduced by a pull request.

Runs [bigo](https://github.com/RomanAgaltsev/bigo) on the head and base of a
PR, diffs the two reports, and posts a single comment that updates in place.

## Usage

```yaml
- uses: actions/checkout@v4
  with:
    fetch-depth: 0 # required: the base commit must be reachable
- uses: actions/setup-go@v5
  with:
    go-version-file: go.mod
- uses: RomanAgaltsev/bigo-action@v1
  with:
    fail-on: none # none (default) | break | regression
```

## Inputs

| Input | Default | Meaning |
|---|---|---|
| `working-directory` | `.` | Directory of the Go module to analyze. |
| `packages` | `./...` | Package patterns to analyze. |
| `fail-on` | `none` | `none` reports only; `break` fails on a broken budget or a new function already over budget; `regression` also fails on a proven asymptotic regression in unbudgeted code. |
| `comment` | `true` | Post/update a single PR comment. |
| `bigo-version` | `latest` | Version of bigo to install, e.g. `v1.48.2`. |
| `github-token` | `github.token` | Token used to post the comment. |

## Output

`findings` — the rendered markdown body.

## What it will not do

**Nothing fails the job on a new unverifiable bound.** Losing visibility is
worth telling you about, but it is not a defect, and failing on it would only
pressure you into avoiding code bigo cannot yet see.

**Report-only is the default on purpose.** bigo's analysis surface is pre-stable
across minors, and a tool that breaks your build by surprise is a tool you
uninstall. Turn `fail-on` up once you trust it on your codebase.

**A base that does not build is not your PR's fault.** If the base commit fails
to build, or predates your adoption of bigo, the Action says so and reports the
head side only.

## What bigo can and cannot bound

bigo bounds roughly a third of hand-written functions in large real codebases
and reports the rest as ⊤ (unverifiable). ⊤ is a deliberate answer, not a
failure: the project's prime directive is that it must never emit a wrong bound.
See [what bigo does not count](https://github.com/RomanAgaltsev/bigo#what-bigo-does-not-count-yet).

## License

MIT — see [LICENSE](LICENSE).
