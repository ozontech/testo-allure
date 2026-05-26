# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-05-26

### Added

- An option `WithMaxAttachmentSize` to limit a size of large attachments by trimming them.

## [1.0.3] - 2026-05-26

### Fixed

- Fixed incorrect timings for tests when other plugins would call `Parallel()` in `BeforeEach` hook.

## [1.0.2] - 2026-05-21

### Added

- Support for sub-suites.

## [1.0.1] - 2026-05-20

### Fixed

- Fixed data race when running parallel tests.
- Correctly trim `#00`-like suffixes from full name.

## [1.0.0] - 2026-05-13

### Added

- Initial stable version.

[1.1.0]: https://github.com/ozontech/testo-allure/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/ozontech/testo-allure/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/ozontech/testo-allure/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/ozontech/testo-allure/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/ozontech/testo-allure/releases/tag/v1.0.0
