# SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
#
# SPDX-License-Identifier: EUPL-1.2

{
  # std
  inputs = {
    devshell.url = "github:numtide/devshell";
    nixago.url = "github:nix-community/nixago";
    nixpkgs.url = "https://channels.nixos.org/nixos-unstable/nixexprs.tar.xz";
    std.inputs.devshell.follows = "devshell";
    std.inputs.nixago.follows = "nixago";
    std.url = "github:divnix/std";
    tailor.url = "github:wimpysworld/nix-packages";
    tailor.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs =
    inputs:
    inputs.std.growOn
      {
        inherit inputs;
        systems = [
          "x86_64-linux"
          "aarch64-linux"
        ];
        cellsFrom = ./nix;
        cellBlocks = with inputs.std.blockTypes; [
          (containers "containers")
          (devshells "shells" { ci.build = true; })
          (installables "ci-jobs" { ci.build = true; })
          (installables "packages" { ci.build = true; })
          (nixago "dotfiles")
          (runnables "operables" { ci.build = true; })
        ];
      }
      {
        devShells = inputs.std.harvest inputs.self [
          "repo"
          "shells"
        ];
        packages = inputs.std.harvest inputs.self [
          "repo"
          "packages"
        ];
        hydraJobs = inputs.std.harvest inputs.self [
          "repo"
          "ci-jobs"
        ];
      };
}
