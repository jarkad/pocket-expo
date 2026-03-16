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
  pocket-expo = nixpkgs.buildGoModule {
    pname = "pocket-expo";
    version = "1.0.0";

    src = inputs.self;

    vendorHash = null;
  };
}
