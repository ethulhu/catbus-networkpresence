# SPDX-FileCopyrightText: 2020 Ethel Morgan
#
# SPDX-License-Identifier: MIT

{ pkgs ? import <nixpkgs> {} }:
with pkgs;

buildGoModule rec {
  name = "catbus-networkpresence-${version}";
  version = "latest";
  goPackagePath = "go.eth.moe/catbus-networkpresence";

  modSha256 = "19hyb5h0qxfjkmissxvhcrr7xdd66iss6v1w0nmz8zag7q1qk34r";

  buildInputs = [
    arp-scan
    makeWrapper
  ];

  src = ./.;

  postInstall = ''
    wrapProgram $out/bin/arp-scan                        --set PATH ${lib.makeBinPath [ arp-scan ] }
    wrapProgram $out/bin/catbus-observer-networkpresence --set PATH ${lib.makeBinPath [ arp-scan ] }
  '';

  meta = {
    homepage = "https://ethulhu.co.uk/catbus";
    licence = stdenv.lib.licenses.mit;
  };
}
