# Contributing to DNSKit

Thank you for taking the time to contribute! :tada:

The following guidelines help keep the project maintainable and welcoming.

## Getting Started

1. Fork the repository and create your branch from `main`.
2. If you've added code, run `go fmt` and `go test ./...` to make sure your changes build and pass tests.
3. Ensure your code follows standard Go conventions and adds tests where appropriate.
4. Use [Conventional Commits](https://www.conventionalcommits.org/) for your commit messages.
5. Open a pull request and describe your changes in detail.

## Code Style

- Go code should be formatted with `go fmt`.
- Run `golangci-lint run` if available to catch common issues.
- Keep functions and files focused; prefer small, well-named units of code.

## Documentation

- Update documentation or comments whenever behavior changes.
- New features should be accompanied by relevant README or doc updates.

## Reporting Bugs and Requesting Features

- Search existing issues before opening a new one to avoid duplicates.
- When filing a bug, include clear steps to reproduce and expected behavior.
- Feature requests should explain the problem being solved and, if possible, propose an API or CLI design.

## Security

For sensitive security reports, please follow the instructions in [SECURITY.md](SECURITY.md) and email [kontakt@nicokempe.de](mailto:kontakt@nicokempe.de) directly instead of opening an issue.

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](../LICENSE).

Thanks again for helping make DNSKit better!
