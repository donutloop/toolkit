# Contributing to Toolkit

## Reporting other issues

Check that [our issue database](https://github.com/donutloop/toolkit/issues)
doesn't already include that problem or suggestion before submitting an issue.
if you have ways to reproduce the issue or have additional information that may help
resolving the issue, please leave a comment. This information will help us review and fix your issue faster.

## contribution guidelines

### Before you do changes

Before contributing large or high impact changes, make the effort to coordinate
with the maintainers of the project before submitting a pull request. This
prevents you from doing extra work that may or may not be merged.

Large PRs that are just submitted without any prior communication are unlikely
to be successful.

While pull requests are the methodology for submitting changes to code, changes
are much more likely to be accepted if they are accompanied by additional
engineering work. While we don't define this explicitly, most of these goals
are accomplished through communication of the design goals and subsequent
solutions. Often times, it helps to first state the problem before presenting
solutions.

Typically, the best methods of accomplishing this are to submit an issue,
stating the problem. This issue can include a problem statement and a
checklist with requirements. If solutions are proposed, alternatives should be
listed and eliminated. Even if the criteria for elimination of a solution is
frivolous, say so.

Larger changes typically work best with design documents. These are focused on
providing context to the design at the time the feature was conceived and can
inform future documentation contributions.


### Design and cleanup proposals

You can propose new designs for existing Toolkit features. You can also design
entirely new features. We really appreciate contributors who want to refactor or
otherwise cleanup our project.

### Conventions

Fork the main branch and make changes on your fork in a feature branch:

- If it's a bug fix branch, name it XXXX-something where XXXX is the number of the issue. 
- If it's a feature branch, create an enhancement issue to announce
  your intentions, and name it XXXX-something where XXXX is the number of the issue.

Submit tests for your changes. 

Update the documentation when creating or modifying features. Test your
documentation changes for clarity, concision, and correctness, as well as a
clean documentation build.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plug-ins that do this automatically.

Pull request descriptions should be as clear as possible and include a reference
to all the issues that they address.

### Commit Messages

Commit messages must start with a capitalized and short summary (max. 50 chars)
written in the imperative, followed by an optional, more detailed explanatory
text which is separated from the summary by an empty line.

Commit messages should follow best practices, including explaining the context
of the problem and how it was solved, including in caveats or follow up changes
required. They should tell the story of the change and provide readers
understanding of what led to it.

In practice, the best approach to maintaining a nice commit message is to
leverage a `git add -p` and `git commit --amend` to formulate a solid
changeset. This allows one to piece together a change, as information becomes
available.

### Review

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Post
a comment after pushing. New commits show up in the pull request automatically,
but the reviewers are notified only when you comment.

Pull requests must be cleanly rebased on top of master.

After every commit, make sure the test suite passes. Include
documentation changes in the same pull request so that a revert would remove
all traces of the feature or fix.

Include an issue reference like `Closes #XXXX` or `Fixes #XXXX` in commits that
close an issue. Including references automatically closes the issue on a merge.

### How can I become a maintainer?

contacting the project leader at marcel.edmund.franke@gmail.com. 

## Coding Style

Unless explicitly stated, we follow all coding guidelines from the Go
community. While some of these standards may seem arbitrary, they somehow seem
to result in a solid, consistent codebase.

The rules:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Make your changes as you see fit ensuring that you create appropriate tests along with your changes. Test your changes as you go.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the sub package(s) they modify.
   * E.g. "schedule: make scheduling faster"
 * All code should pass the default levels of [`golint`](https://github.com/golang/lint).
 * Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
 * Document _all_ declarations and methods, even private ones. Declare
   expectations, caveats and anything else that may be important. If a type
   gets exported, having the comments already there will ensure it's ready.
 * All tests should run with `go test` and outside tooling should not be
   required. No, we don't need another unit testing framework. Assertion
   packages are acceptable if they provide _real_ incremental value.


## Copyright

New files that you contribute should use the standard copyright header:

```
// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License 
// license that can be found in the LICENSE file.
```

Exceptions are example and doc files 