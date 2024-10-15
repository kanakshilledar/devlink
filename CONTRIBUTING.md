# Contributing to DevLink

Thank you for considering contributing to **DevLink**! We welcome contributions
from everyone and appreciate your interest in improving this project. Please
follow the guidelines below to ensure a smooth contribution process.

> [!NOTE]
> By contributing to **DevLink**, you agree that your contributions will be
> licensed under the GNU General Public License v3.0 (GPL-3.0).

## Table of Contents

- [Getting Started](#getting-started)
- [Contributing Guidelines](#contributing-guidelines)
- [Submitting Changes](#submitting-changes)
- [Reporting Issues](#reporting-issues)
- [Acknowledgements](#acknowledgements)

## Getting Started

* Fork the repository.
* Clone the repository via `git ssh`.
* Create a branch and make some changes
* Commit and start a PR.

## Contributing Guidelines

* Follow the project structure: Ensure that your code follows the existing
  project structure.
* Write clear commit messages: Use descriptive commit messages that clearly
  convey the purpose of your changes. Refer the commit log to get a brief idea
  of the commit messages.
* Add the DCO in your commit messages
  The DCO is an attestation attached to every contribution made by every
  developer. In the commit message of the contribution, the developer simply
  adds a `Signed-off-by` statement and thereby agrees to the DCO. When a
  developer submits a patch, it is a commitment that the contributor has the
  right to submit the patch per the license. The DCO agreement is shown below
  and at http://developercertificate.org/.
  ```text
  Developer's Certificate of Origin 1.1

  By making a contribution to this project, I certify that:
  
  (a) The contribution was created in whole or in part by me and I
  have the right to submit it under the open source license
  indicated in the file; or
  
  (b) The contribution is based upon previous work that, to the
  best of my knowledge, is covered under an appropriate open
  source license and I have the right under that license to
  submit that work with modifications, whether created in whole
  or in part by me, under the same open source license (unless
  I am permitted to submit under a different license), as
  Indicated in the file; or
  
  (c) The contribution was provided directly to me by some other
  person who certified (a), (b) or (c) and I have not modified
  it.
  
  (d) I understand and agree that this project and the contribution
  are public and that a record of the contribution (including
  all personal information I submit with it, including my
  sign-off) is maintained indefinitely and may be redistributed
  consistent with this project or the open source license(s)
  involved.
  ```
  The “sign-off” in the DCO is a “Signed-off-by:” line in each commit’s log
  message. The Signed-off-by: line must be in the following format:
  ```text
  Signed-off-by: Your Name <your.email@example.com>
  ```
  For your commits, replace:
    * `Your Name` with your legal name (pseudonyms, hacker handles, and the
      names of groups are not allowed).
    * `your.email@example.com` with the same email address you are using to
      author
      the commit.
* Document your code: Add comments and documentation where necessary to explain
  your code. Add documentation in `godoc` format so that godoc will
  automatically pick up from the code.
* Adhere to coding standards: Ensure your code follows the coding standards and
  style guide of the project. Use tools like `go fmt` for Go projects to
  maintain consistent formatting. Make sure the files have line endings with
  `LF` not `CRLF`.

## Submitting Changes

* To submit the changes push your code to GitHub
  ```shell
  $ git push origin my-feature-branch
  ```
* Start a PR by providing a clear description of the changes. If it fixes some
  issue mention that issue as well.

## Reporting Issues

If you encounter any bugs or have feature requests, please report them by
opening an issue in the repository. Make sure to provide a detailed description
and any relevant information to help us understand and address the issue.

## Acknowledgements

Thank you for contributing to DevLink! Your efforts help make this project
better for everyone. We appreciate your time and input.