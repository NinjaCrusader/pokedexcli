{
  description = "pokedexcli - Go-based Pokedex CLI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        packages.default = pkgs.buildGoModule {
          pname = "pokedexcli";
          version = "0.1.0";

          src = ./.;

          vendorHash = null;

          ldflags = [
            "-s"
            "-w"
          ];
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/pokedexcli";
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            golangci-lint
          ];
        };
      });
}
