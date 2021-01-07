# Project Guidelines

## General community expectations

The general community mindset for collaboration here will be:

> be meaningfully productive and easy to work with

Honestly, just be nice. This is a free product, everyone here is a volunteer,
and there are people of various backgrounds, skillsets, and values
participating. Do your part to make improving the project easy! :heart:

If you feel like an explicit list is easier to adhere to, we'd generally say to
reference [the Python Code of Conduct](https://www.python.org/psf/conduct/).
This is for no other reason than their list is pretty long and detailed, and we
feel it covers a good chunk of what we mean by "be nice".

## Code contribution expectations

In short:

Project Goals:
- consolidate installation of different kinds of Minecraft servers
- consolidate running of different kinds of Minecraft servers
- consolidate full configuration of servers and addons declaratively
- wrap operational processes of servers including observation and upgrades
- provide sensible, prod-friendly defaults for server configuration
- provide clear, understandable documentation for use
- stay as easily maintainable long-term by as few people as possible

Project *Non*-Goals:
- extend OS compatability beyond Linux to all major distributions
- handle packaging or hosting of server binaries, images, etc

Any code contributions that change general user experience, expand scope of
functionality, removes features, or generally "rocks the boat" to even a mild
degree, should be assessed against the goals and non-goals above.

If there's some contention about the reasoning or worth of a contribution, it
should be done in a linked issue -- Pull Request pages should be reserved for
discussion of actual implementation details.

## Documentation contribution expectations

Documentation is a really important part of maturing the project and incredibly
valued. Beyond emphasizing this, there's really just a short list of points to
keep in mind when contributing new documentation of any kind:

- Are there any grammatical or syntactic errors?
- Is it concise, easy to skim, with limited interpretations?
- Is this bit of information duplicated somewhere else?
- Will code changes cause this documentation to be out of date?

If a documentation contribution is a fix or update, then the criteria really
boils down to:

- Are there any grammatical or syntactic errors?
- Is this edit correct?
- Can this mistake be avoided going forward?
