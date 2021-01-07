# Contributing

Heya :wave: we're thrilled that you're considering helping out! Hopefully this
doc can orient you on doing just that!

## How you can contribute

Given the small backing and large breadth of this project, there's actually
quite a few things that can be a huge help when taken care of:

- documentation edits and additions in README's, Wiki Guides, and Go docs
- feature requests and (detailed :eyes:) bug reports in [our
  issues](https://github.com/loksonarius/mcsm/issues)
- pull requests addressing issues, POC'ing features, cleaning up code, adding
  tests, improving code coverage, etc

Contributions both large and small in scale are welcome so long as they help
improve or maintain the quality of the project!

## Making your first contribution

Regardless of the kind of contribution you'd like to make, we suggest reviewing
our [project guidelines](GUIDELINES.md) and a few [open
issues](https://github.com/loksonarius/mcsm/issues) to get your bearings of
things around here.

After that, feel free to jump to the heading below that best matches your
intent!

### I'm just totally new to this

So you're totally new to contributing to Open Source and working with Git, or
maybe this is the first Go project you're collaborating on. If this sounds like
you, then please consider reviewing the following reference guides and docs at
your discretion. Reviewing them in their entirety may be excessive, but they
should serve as good introductions to improve your mental model for this
project!

- [Microsoft Intro to Git
  Series](https://docs.microsoft.com/en-us/learn/paths/intro-to-vc-git/): brief
  series detailing SVC with git, as well as asynchronous collaboration
- [Go by Example](https://gobyexample.com): brief snippets covering most of Go's
  features and paterns
- [A Tour of Go](https://tour.golang.org/welcome/1): guided, interactive
  examples working through Go's features
- [makeapullrequest.com](https://makeapullrequest.com) and
  [firsttimersonly.com](https://www.firsttimersonly.com): introductory guides to
  contributing to OpenSource projects in general, with some GitHub specifcs here
  and there
- [Open First-Timer
  Issues](https://github.com/loksonarius/mcsm/labels/good%20first%20issue): open
  issues in this project that are marked as "good for first timers"

### Ask for some help with something

While most projects may have a community forum or chat server for questions like
these, as of right now, this project has none. As the project grows and the need
for asynchronous topical chat increasse, we'll re-evaluate the need if/when the
time comes.

For now, please [submit your question as an
issue](https://github.com/loksonarius/mcsm/issues/new?assignees=&labels=question&template=general-question.md&title=)
on our board. At the very least, this will allow common question patterns and
needs to be identified over time, so please don't hesitate to ask!

### Submit a feature request or bug report

This is one of the easier contributions to make and can be done straight from
our issues page:

- [Submit new Bug Report](https://github.com/loksonarius/mcsm/issues/new?assignees=&labels=bug&template=bug_report.md&title=)
- [Submit new Feature Request](https://github.com/loksonarius/mcsm/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=)

Before submittin a new issue, please try searching through both [open _and_
closed Issues and Pull Requests](https://github.com/loksonarius/mcsm/issues?q=)
for mentions of your suggestion or error. If you don't find one that matches
your case, then please feel free to open an issue to make a suggestion using
either the `Feature request` template or the `Bug report` template.

### Edit or add documentation to the wiki or repo

Depending on what's being edited or added to, this can be easily accomplished
straight from the GitHub web UI or with a few extra steps with a local repo
copy.

| Type       | How to submit                                  |
| ---------- | ---------------------------------------------- |
| Wiki pages | Make edit suggestions right from the wiki page |
| Go Docs    | Fork the repository and edit locally           |
| README's   | Fork the repository and edit locally           |

The pull request made for each of these should be labeled with the
`documentation` label so that it can be prioritized for reading and review.

### Implement a bug fix or feature request

This section here details how to get from a freshly cloned repo to iterating
with automated test runs to submitting a Pull Request with code changes.

#### Dev Environment Setup

The following utilites are required for local development:

- [go 1.15+](https://golang.org)
- [Docker](https://docs.docker.com/get-docker/)
- [just](https://github.com/casey/just)
- [jq](https://stedolan.github.io/jq/)
- [(optional) direnv](https://direnv.net)

Please install them if needed, and consider reading through "Getting started"
guides for them if they're new to you.

#### Repo Layout

The following is an overview of the repo's layout:

```
.
├── CONTRIBUTING.md     # This document!
├── GUIDELINES.md   # Community, docs, and code standards
├── JUSTFILE            # Task declarations
├── build/          # Built binaries of mcsm go here
├── cmd/                # Go code focused on CLI setup
├── integration     # Integration test utils and suites
│   ├── Dockerfile      # Defines container image for running integration suites
│   ├── README.md   # Overview of our integration test setup
│   ├── servers/        # Integration suite runs stay here
│   ├── suite.sh    # Entrypoint for running and adding integration suites
│   └── suites/         # Suite definitions live here
├── internal/       # Go code that shouldn't be externally used
├── main.go             # Entrypoint for the mcsm binary -- really plain
├── pkg/            # Go code defining mcsm constructs -- publicly accessible
└── scripts/            # Scripts used by 'just' tasks
```

#### Available Tasks

To view available developer tasks, run:

```bash
just --list
```

Tasks include brief documentation explaning what each does. Important ones for
code development are:

- `build`: you'll be using this to actually compile the CLI code
- `test`: run any unit tests defined under `integration` and `pkg`
- `integration`: run integration test suites in a Docker container

#### Sample iteration loop

While some specifics vary across feature and scope, the general iteration cycle
of 'edit-build-test-commit' is pretty consistent for this project. For this
sample iteration loop, we'll be adding a new unit test case to the
`pkg/config/properties` package.

Start by editing `pkg/config/properties/util_test.go` to add a new test cause
under the `TestToPropertiesKey`.

```bash
vim pkg/config/properties/util_test.go

# make your edits using whatever editor you like
# added the following test case:
#	{
#		name: "downcases and separates words",
#		s:    "F1AbClS",
#		e:    "f1-ab-cl-s",
#	},
```

Then run the `test` task to both build and run unit tests, checking for any
compilation errors or failures in the output.

```bash
just test

# output included the following:
# --- FAIL: TestToPropertiesKey (0.00s)
#     --- FAIL: TestToPropertiesKey/downcases_and_separates_words (0.00s)
#         util_test.go:44: got -f1-ab-cl-s, expected f1-ab-cl-s
```

Now we can go edit `toPropertiesKey` in `pkg/config/properties/util.go` and
handle the case where the first letter of a key is capital.

```bash
vim pkg/config/properties/util.go

# changed lines
#   for i, c := range s {
#   	if unicode.IsUpper(c) && i > 0 {
#   		result += "-"
```

No we re-run our tests and see if that worked.

```bash
just test

# output:
# go: warning: "./internal/..." matched no packages
# no packages to test
# ?       github.com/loksonarius/mcsm/pkg/config  [no test files]
# ?       github.com/loksonarius/mcsm/pkg/config/presets  [no test files]
# ok      github.com/loksonarius/mcsm/pkg/config/properties       (cached)
# ok      github.com/loksonarius/mcsm/pkg/server  (cached)
# Tested
```

And double check our integration tests too.

```bash
just integration

# truncating output as it's pretty excessive
# Waiting for server to finish startup
# Waiting for server to finish startup
# [20:10:20] [Server thread/INFO]: Done (32.579s)! For help, type "help"
# Server startup complete -- stopping now
# Integration suite run
```

At this point we could make a git commit, clean up any code that was added, or
check our code coverage to see if we may have missed testing any branches we've
added to the code base. Regardless of the kind of feature being added or bug
being fixed, it's expected for code subissions to have passing tests and
reasonable coverage and documentation, so make sure you can verify your changes
work in some specific way! :eyes: If you need help doing this, feel free to
[open a question issue](#ask-for-some-help-with-something) detailing what you're
trying to do.

#### Submitting a Pull Request

There isn't much to do after making your changes but [open a Pull
Request](https://github.com/loksonarius/mcsm/compare) for code review! The
default template includes a few fields and formatting to get you started on
providing helpful info and context for your change.

Some important things to consider before actually opening a PR:

- be sure to review the code changes your submitting -- the diff may contain
  or be missing some code you may have not intended
- be sure to select _your_ branch as the source branch and _our_ `main` branch
  as the target for merging

There's no current expectations nor commitments around review time given this
isn't a sponsored nor staffed project, but there will be at least a "best
effort" attempt made to review and interact with open PRs until they're either
closed or merged.
