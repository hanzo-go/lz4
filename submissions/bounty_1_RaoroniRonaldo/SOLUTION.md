Para cumprir esta bounty, como Engenheiro de Software Sênior e Especialista em Bounties, preparei as alterações necessárias no `README.md` para alinhar a documentação com o fork oficial da Hanzo, removendo referências ao upstream e ao utilitário de linha de comando inexistente.

### Technical Summary

This Pull Request updates the `README.md` to correctly reflect the module's identity as `hanzo.ai/lz4/v4`. 

**Changes:**
1.  **Installation Path:** Updated `go get` command to use the vanity URL `hanzo.ai/lz4/v4`.
2.  **Badges:** 
    *   Updated **Go Reference** to point to `pkg.go.dev/hanzo.ai/lz4/v4`.
    *   Updated **CI** badge to point to the local GitHub Actions workflow (`hanzo-go/lz4`).
    *   Updated **Go Report Card** and **GitHub Tag** badges to point to the `hanzo-go/lz4` repository.
3.  **CLI Removal:** Removed the installation and usage sections for `lz4c`, as this module does not ship with the command-line tool.
4.  **Credits:** Preserved the lineage and credits to the original author (Pierre Curto) while ensuring all functional links point to the Hanzo fork.

**Verification:**
- `grep -n pierrec README.md` now only returns lines related to original authorship/credits.
- Installation via `go get hanzo.ai/lz4/v4` is the documented method.
- All badges point to the correct status of this specific repository.

---

===FILE: README.md===
<p align="center"><img src=".github/hero.svg" alt="lz4" width="880"></p>

# lz4 : LZ4 compression in pure Go

[![Go Reference](https://pkg.go.dev/badge/hanzo.ai/lz4/v4.svg)](https://pkg.go.dev/hanzo.ai/lz4/v4)
[![CI](https://github.com/hanzo-go/lz4/actions/workflows/ci.yml/badge.svg)](https://github.com/hanzo-go/lz4/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/hanzo-go/lz4)](https://goreportcard.com/report/github.com/hanzo-go/lz4)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/hanzo-go/lz4?style=social)](https://github.com/hanzo-go/lz4/tags)

## Overview

This package provides a streaming interface to [LZ4 data streams](http://fastcompression.blogspot.fr/2013/04/lz4-streaming-format-final.html) as well as low level compress and uncompress functions for LZ4 data blocks.
The implementation is based on the reference C [one](https://github.com/lz4/lz4).

## Install

Assuming you have the go toolchain installed:

```
go get hanzo.ai/lz4/v4
```

## Usage

The package is fully compatible with the standard library `io.Reader` and `io.Writer` interfaces.

```go
import "hanzo.ai/lz4/v4"
```

Refer to the [Go Reference](https://pkg.go.dev/hanzo.ai/lz4/v4) for detailed API documentation and examples.

## Contributors

The lz4 project was originally created by [Pierre Curto](https://github.com/pierrec). This fork is maintained by [Hanzo](https://hanzo.ai).

Thanks to all the [contributors](https://github.com/hanzo-go/lz4/graphs/contributors) who have helped with the project.

## License

[BSD-3-Clause](LICENSE)
===END_FILE===

---

### Passos para Submissão (Instruções para você):

1.  **Fork e Clone:** Certifique-se de ter feito o fork de `hanzo-go/lz4`.
2.  **Branch:** Crie uma branch chamada `fix/readme-links-and-cleanup`.
3.  **Commit:** Aplique as mudanças acima no arquivo `README.md`.
4.  **Pull Request:**
    *   **Título:** `docs: update README to point to hanzo.ai vanity path and remove CLI`
    *   **Descrição:**
        ```markdown
        Fixes #1

        This PR updates the README to align with the hanzo.ai/lz4/v4 module path:
        - Updated installation instructions to use the vanity URL.
        - Updated all badges (Go Reference, CI, Report Card, Tags) to point to this repository.
        - Removed documentation for the `lz4c` CLI tool which is not included in this fork.
        - Maintained original credits to Pierre Curto.
        ```
5.  **Hanzo Leaderboard:** Não esqueça de vincular sua conta GitHub em [hanzo.ai/leaderboard](https://hanzo.ai/leaderboard) para que os pontos da bounty sejam creditados após o merge.