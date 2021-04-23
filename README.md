
# golang-project-template

A Repo containing a standard layout and configuration for golang projects.

The project layout defined in this repo is a simplified version of [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
with some additional ci/cd setup added in. Some of the patterns defined in that example are overkill for dsp-devops' current needs, however this can easily be updated if needs change.

The basic structure provided in this repo is intended to serve as a starting point when starting a new go project and to avoid boilerplate.

## Usage

Click the `use this template` button in the top right to create a new repo with the desired owner and name using the files and folder structure defined here. Make sure to check the `include all branches` option in order to include the `gh-pages` branch so that code coverage html reports work properly.

Most of the CI/CD worflows included here are intended to be generic with two exceptions.

1. [this line in the dockerfile](https://github.com/broadinstitute/golang-project-template/blob/142d0dc810fa4f3afa68e0a5d37aac03f0c3796f/Dockerfile#L13) which will need to be updated to match the actual name of any executable(s).

2. [this line in the cloudbuild.yaml](https://github.com/broadinstitute/golang-project-template/blob/52f72e50a023a91479fd83605326fec7e69188ec/cloudbuild.yaml#L24) will need to be updated with the desired name for the docker image.

## Additional Steps

1. To ensure the code coverage functionality is working, make sure github pages is enabled in the project repo.

2. To take advantage of the cloudbuild image build process, a [cloud build trigger](https://cloud.google.com/build/docs/automating-builds/create-manage-triggers) must be created to monitor the project repository and the cloud build github app must be enabled.
