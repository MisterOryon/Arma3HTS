# Introduction
## Thanks
First off, thank you for considering contributing to Arma3HTS.
It's people like you who keep open source projects moving forward.

## Why you should read the contributing
Following these instructions will help you communicate better, by respecting these instructions you respect the time we devote to the management and development of this open source project.
In return, We reciprocate that respect in addressing your issue, helping you finalize your pull requests and answering your questions as quickly as possible.

## The types of contributions desired
Arma3HTS is an open source project, and it is with great pleasure that we will receive your contributions.
From warmly welcoming new participants, submitting bug reports, answer other people's questions, writing to improving documentation, submitting a feature request, to working on writing code that can be integrated in Arma3HTS.

## The types of contributions that are not wanted
It is always difficult to refuse a pull request because your don't open an issue before starting work, so to avoid inconvenience for we and for you see [How to suggest a feature or enhancement](#how-to-suggest-a-feature-or-enhancement).

# Ground Rules
## For healthy community
- Communicate with others in a respectful and considerate manner and encourage new contributors
- Create issues for any changes and enhancements that you wish to make, discuss things transparently and get community feedback
- You must follow the code of [conduct](CODE_OF_CONDUCT.md)

## For code
### Code style
To avoid having to recode your functions, follow its instructions:
- Avoided if nesting, mostly loop nesting, preferred early return, try to avoid nested loops or move them into a dedicated function
- Do not add any structs if a similar structs already exists, prefer to add one or more methods in this structs
- Each method must have a single responsibility
- Format your code with [go fmt](https://go.dev/blog/gofmt) before commit your code

### Comments
Your code must be commented, don't comment every line just:

- Above of methods to explain what they do, with which variable and what it returns
- All complex instructions such as mathematical calculations or complex SQL requests
- If your structs publicly, add a comment above structs declaration to indicate its purpose and when to use it

> **Note:** Any pull request that does not have a comment, when code review you will be asked you to add comments in code to validate the pull request.

### Commits
Keep your commits as small as possible, preferably one commit per new feature/modification.
Additionally, follow:

| Commit message                                                                                                                                                                                                                          | Release type                                                                                          |
|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `fix(users): case sensitive issue for usernames`                                                                                                                                                                                        | Fix Release                                                                                           |
| `feat(users): add public/private profile option`                                                                                                                                                                                        | Feature Release                                                                                       |
| `perf(users): remove phone number`<br><br>`BREAKING CHANGE: The phone number has been removed.`<br>`The server is managed locally by the administrator of the home, the constraint of the unique telephone number does not make sense.` | Breaking Release <br /> (Note that the `BREAKING CHANGE: ` token must be in the footer of the commit) |

> **Note:** Any pull request with commits not respect its conventions will be refused.

# How to report a bug
## Security vulnerability
If you find a security vulnerability, **do NOT open** an issue send an email at **misteroryon@protonmail.com**.
To determine if you are dealing with a security issue, ask yourself these two questions:

- Can I access something that doesn't belong to me or something I shouldn't have access to?
- Can I disable something for other people?

If the answer to either of these two questions is "yes", then you are probably dealing with a security issue.

> **Note:** that even if you answer "no" to both questions, you may still face a security issue, so if you are unsure, just emailed us.

## How to file a bug report
To send a bug report, go to [issues](https://github.com/MisterOryon/Arma3HTS/issues) list on GitHub and create a new issue, choose bug report template.

# How to suggest a feature or enhancement
If you find yourself wanting a feature that doesn't exist in Arma3HTS, you're probably not alone with similar needs.
Open an issue with **enhancement** label on our [issues](https://github.com/MisterOryon/Arma3HTS/issues) list on GitHub that describes the feature you want to see, why you need it.
Try to evaluate the feature, how it might work.
Don't forget to specify if you only want to propose an idea or if you want to work on it.
Before your issue is added to the workflow, it must be approved by a maintainer.

# Your First Contribution
Not sure where to start contributing to Arma3HTS. You can start by browsing through these [issues](https://github.com/MisterOryon/Arma3HTS/issues) whit label:

- **Good first issue** — issues which should only require a few lines of code and a test or two
- **Help wanted** — issues which should be a bit more involved than beginner issues

> **Working on your first Pull Request?** <br>
> You can learn how from this _free_ series [How to Contribute to an Open Source Project on GitHub](https://kcd.im/pull-request).

## Getting Started
### How to submit a contribution
First off, assignee you on an open issue with **waiting for a developer** label or open an issue with **enhancement** label.
If you want to open an issue, see [How to suggest a feature or enhancement](#how-to-suggest-a-feature-or-enhancement), When it accepted the label **in progress** is added, and you can start working.

Make sure to follow [Ground Rules for code](#for-code) during development.

When your work is done, create a pull request and follow its steps:

- For the title reuse the same convention as for the [commits](#commits), for example `feat(remote_client): Add FTP client`
- Add the issue number in your pull request
- Tag at least one person in [maintainers](#maintainers) like reviewer

### Process for small or "obvious" fixes
If your contribution falls into one of the categories below, then you can directly create a pull request with the label **small fixes** without having to open an issue.

- Spelling / grammar fixes
- Typo correction, space and formatting changes
- Comment clean up
- Bug fixes, don't forget to adjust the unit tests in case of code change and if it exists the issue number who report the bug
- Adding logging messages or debugging output

# Code review process
Your pull request will be reviewed by one or more people present in [maintainers](#maintainers).
The response time may be different between two pull requests, generally it does not exceed 7 days.
It is possible that during the review, the reviewer asks you to add or modify certain things before validating your pull requests.
If modifications have been requested and after one month without news from you, the pull requests will be closed without being accepted.

# Community
## Maintainers
You can see a list of authors, maintainers in [README authors](README.md#authors).

> **This project is developed in my free time.** <br>
> Arma3HS is one of the projects that I am developing between 2 sandwiches during my break, It may therefore happen that I do not see your message. <br>
> If after 1 weeks you haven't response from me, don't hesitate to ping me again.
