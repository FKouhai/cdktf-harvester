{
  description = "Nix flake to build cdktf-harvester and dev environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      gomod2nix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ gomod2nix.overlays.default ];
          config.allowUnfree = true;
        };
        cdktf-harvester = pkgs.buildGoModule {
          name = "cdktf-harvester";
          src = ./cdktf-app;
          vendorHash = "sha256-aqfUIXSo9hQCw1hVekjqnlIxwqfIN32qbtUUKQjyq14=";
          proxyVendor = true;
          doCheck = false;
          postInstall = ''
            mv $out/bin/stack $out/bin/cdktf-harvester
          '';
        };
      in
      with pkgs;
      {
        defaultPackage = cdktf-harvester;
        devShells.default = mkShell {
          buildInputs = [
            nodePackages.cdktf-cli
            go
            nodejs_20
            terraform
          ];
        };
      }
    );
}
