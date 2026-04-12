# Contributing to dan-cli

Thank you for helping improve the **DAN command-line tool**.

## Community

- [Code of Conduct](CODE_OF_CONDUCT.md) — required for all interaction.
- [Security policy](SECURITY.md) — report vulnerabilities privately (do not use public issues).
- [Funding](.github/FUNDING.yml) — optional support via PayPal (see repository **Sponsor** button).

## Reporting issues

Use the [bug report](.github/ISSUE_TEMPLATE/bug_report.md), [feature request](.github/ISSUE_TEMPLATE/feature_request.md), or [question](.github/ISSUE_TEMPLATE/question.md) templates when you open an issue.

## Pull requests

1. Fork the repo and create a branch from `main`.
2. The CLI depends on the Go module [`github.com/marcuwynu23/dango`](https://github.com/marcuwynu23/dan-go) (this monorepo uses `replace github.com/marcuwynu23/dango => ../dan-go` in `go.mod`).
3. Run `go build ./cmd/dan` and `go -C ../dan-go test ./...` for the **dango** module (or `make test` from this directory) before opening a PR.
4. Keep commits focused; write clear messages and describe the PR for reviewers.

## Coding style

- Match existing `gofmt` / `go fmt` formatting.
- Prefer small, reviewable changes.

## License

By contributing, you agree your contributions are licensed under the same terms as this repository (see [LICENSE](LICENSE) if present).
