# CDKTF-HARVESTER
This repo contains the code of a cdktf application to create and manage resources in harvester via terraform-cdk

## Development
If need to do some development, the dev environment and build environment is within the flake.nix inside of this repo. <br>
```(bash)
# running nix develop to get inside of the shell with all the dependencies needed for the project
nix develop
```
## Build
In order to run the build for the project use the following command:
```(bash)
nix build
```
The compiled go binary of the application will be in the result/bin directory from the root of the repo
