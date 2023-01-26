# go-template

[![Test status](https://github.com/inc4/go-template/workflows/Checks/badge.svg)](https://github.com/inc4/go-template/actions?query=workflow%3A%22Checks%22)

When a new Golang repository is created, this template can be used.

Just choose this template when created through UI. Or use `--template` flag
when staring `gh repo create` command.

Directory structure is inspired by https://github.com/golang-standards/project-layout

## Getting started

- Review all files, if you haven't done it before
- Remove examples, all readmes (this one needs to be rewritten)

## Where to put the code

`/cmd` should not contain a lot of code, mostly the code is in packages.
https://github.com/golang-standards/project-layout propose to have the code
in `/pkg` and `/internal`. But it's ok to have it in `/yourpkg` if you have few
packages. Do not name packages like `pkg/api-db`, do this `pkg/api/db` instead.

`/pkg` is for packages that can be used externally. Unless you have a big
project use `/pkg` for all packages.

`/internal` is for internal packages, not intended for external usage, but
use it if you really need to fine tune what can be imported from outside.
