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
      nixpkgs.gcc
      nixpkgs.golangci-lint-langserver
      nixpkgs.gopls
    ];
    commands = [
      { package = nixpkgs.dbmate; }
      { package = nixpkgs.gh; }
      { package = nixpkgs.go; }
      { package = nixpkgs.just; }
      { package = nixpkgs.sqlc; }
      { package = nixpkgs.templ; }
    ];
    nixago = [
      cell.dotfiles.lefthook
      cell.dotfiles.treefmt
    ];
  };
}
