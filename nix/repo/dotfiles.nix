# SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
#
# SPDX-License-Identifier: EUPL-1.2

{ inputs, cell }:
let
  inherit (inputs) nixpkgs;
  inherit (inputs.std) lib;

  l = builtins // nixpkgs.lib;
in
{

  treefmt = lib.dev.mkNixago lib.cfg.treefmt {
    packages = [
      nixpkgs.nixfmt
      nixpkgs.go
      nixpkgs.taplo
    ];
    data = {
      global = {
        on-unmatched = "warn";
        excludes = [ "vendor/*" ];
      };
      formatter = {
        gofmt.command = "gofmt";
        gofmt.includes = [ "*.go" ];
        nixfmt.command = "nixfmt";
        nixfmt.includes = [ "*.nix" ];
        taplo.command = "taplo";
        taplo.includes = [ "*.toml" ];
        taplo.options = [ "fmt" ];
        templ.command = "templ";
        templ.includes = [ "*.templ" ];
        templ.options = [ "fmt" ];
      };
    };
  };

  lefthook = lib.dev.mkNixago lib.cfg.lefthook {
    packages = [
      nixpkgs.cocogitto
      nixpkgs.golangci-lint
      nixpkgs.reuse
      nixpkgs.ripsecrets
    ];
    data = {
      pre-push = {
        parallel = true;
        commands = {
          cocogitto.run = "cog check -l";
        };
      };
      pre-commit = {
        parallel = true;
        fail_on_changes = "ci"; # only when $CI=1
        commands = {
          golangci-lint = {
            run = "golangci-lint run --fix --new-from-rev=@";
            stage_fixed = true;
          };
          reuse = {
            run = "reuse lint-file '{staged_files}'";
          };
          ripsecrets = {
            run = "ripsecrets --strict-ignore '{staged_files}'";
            file_types = [ "text" ];
          };
          treefmt = {
            run = "treefmt '{staged_files}'";
            stage_fixed = true;
          };
        };
      };
    };
  };

}
