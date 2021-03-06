# What is Crust?

[![Build Status](https://drone.crust.tech/api/badges/crusttech/crust/status.svg)](https://drone.crust.tech/crusttech/crust)
[![Go Report Card](https://goreportcard.com/badge/github.com/crusttech/crust)](https://goreportcard.com/report/github.com/crusttech/crust)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcrusttech%2Fcrust.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcrusttech%2Fcrust?ref=badge_shield)

Crust brings your user ecosystem and essential applications together on one platform, unifying them via CRM, Team Messaging and Advanced Identity and Access Management.

**Crust Messaging** is a secure, high performance, open source Slack alternative that allows your teams to collaborate more efficiently, as well as communicate safely with other organisations or customers.

**Crust CRM** is the highly flexible, scalable and open source Salesforce alternative, that enables you to sell faster and interact with leads, clients and team members easier then ever before. Seamless integration with Crust Messaging and Crust Identity and Access Management make it the most complete and flexible self-hosted CRM platform on the market.

**Crust Unify** manages user experience for Crust applications, such as Compose and Messaging, as well as providing an integrated interface for third party or other bespoke applications. 100% responsive and with an intuitive design, Crust Unify increases productivity and ease of access to all IT resources.

## Contributing

### Setup

Copy `.env.example` to `.env` and make proper modifications for your local environment.

An access to a (local) instance of MySQL must be available.
Configure access to your database with `SYSTEM_DB_DSN`, `MESSAGING_DB_DSN` and `COMPOSE_DB_DSN`.

The database will be populated with migrations at the start of each service. You don't need to pre-populate the database, just make sure that your permissions include CREATE and ALTER capabilities.

### Running in local environment for development

Everything should be set and ready to run with `make realize`. This utilizes realize tool that monitors codebase for changes and restarts api http server for every file change. It is not 100% so it needs help (manual restart) in certain cases (new files added, changes in non .go files etc..)

### Making changes

Please refer to each project's style guidelines and guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.

 1. **Fork** the repo on GitHub
 2. **Clone** the project to your own machine
 3. **Commit** changes to your own branch
 4. **Push** your work back up to your fork
 5. Submit a **Pull request** so that we can review your changes

NOTE: Be sure to merge the latest master from "upstream" before making a pull request!


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcrusttech%2Fcrust.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcrusttech%2Fcrust?ref=badge_large)