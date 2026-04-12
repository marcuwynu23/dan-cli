# Security policy — dan-cli

## Supported versions

Security fixes are applied to the **default branch** and released **tags** of this repository. Older tags may not receive backports unless agreed with maintainers.

| Area        | Supported                         |
| ----------- | --------------------------------- |
| Latest/main | :white_check_mark:                |
| Unmaintained tags | :x: (upgrade)                |

## Reporting a vulnerability

**Do not** open a public issue for an undisclosed security vulnerability.

Please use **[GitHub private vulnerability reporting](https://github.com/marcuwynu23/dan-cli/security/advisories/new)**.

Include:

- Description, impact, and affected components (CLI flags, file I/O, dependency on `dan`/`dan-go`, etc.)
- Steps to reproduce and proof-of-concept if safe to share
- Version or commit hash

## What to expect

- **Acknowledgment:** within **48 hours** when possible  
- **Fix & disclosure:** coordinated after a patch is ready  

## Scope

This policy covers the **dan-cli** repository (command-line tool, build scripts, and bundled examples). Parser/encoder logic may also live in **[dan-go](https://github.com/marcuwynu23/dan-go)**; if the issue is purely in the library, maintainers may move the advisory to the correct repo.

## Safe harbor

Good-faith research under this policy (minimal access, no user harm, responsible disclosure) is treated as authorized.

Thank you for responsible disclosure.
