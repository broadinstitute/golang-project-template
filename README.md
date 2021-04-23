
# golang-project-template

A Repo containing a standard layout and configuration for golang projects.

The project layout defined in this repo is a simplified version of [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
with some additional ci/cd setup added in. Some of the patterns defined in that example are overkill for our current needs, however this can easily be updated if needs change.

The basic structure provided in this repo as intended to serve as a starting point when starting a new go project and to avoid boilerplate.

## Additional Steps

1. To ensure the code coverage functionality is working, make sure github pages in the project repo.

2. To take advantage of the cloudbuild image build process, a [cloud build trigger](https://cloud.google.com/build/docs/automating-builds/create-manage-triggers) must be created to monitor the project repository.
