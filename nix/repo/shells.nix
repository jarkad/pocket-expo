# SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
#
# SPDX-License-Identifier: EUPL-1.2

{ inputs, cell }:
let
  inherit (inputs) nixpkgs;
  inherit (inputs.std.lib.dev) mkShell;
  inherit (inputs.std) std;

  l = builtins // nixpkgs.lib;
in
nixpkgs.lib.mapAttrs (_: mkShell) {
  default = {
    name = "pocket-expo";
    imports = [
      std.devshellProfiles.default
    ];
    packages = [
      inputs.tailor.packages.tailor
      nixpkgs.air
      nixpkgs.gcc
      nixpkgs.go
      nixpkgs.golangci-lint-langserver
      nixpkgs.gopls
      nixpkgs.templ
    ];
    commands = [
      { package = nixpkgs.gh; }
      { package = nixpkgs.just; }
    ];
    nixago = [
      cell.dotfiles.lefthook
      cell.dotfiles.treefmt
    ];
  };
}
