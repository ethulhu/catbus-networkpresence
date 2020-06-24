# SPDX-FileCopyrightText: 2020 Ethel Morgan
#
# SPDX-License-Identifier: MIT

{ pkgs ? import <nixpkgs> {} }:
with pkgs;

buildGoModule rec {
  name = "catbus-networkpresence-${version}";
  version = "latest";
  goPackagePath = "go.eth.moe/catbus-networkpresence";

  modSha256 = "1zs9i676vd2wi4v59kwj5cqm373hk41kn8jim3jkrx2pl1cd2hi4";

  buildInputs = [
    arp-scan
    makeWrapper
  ];

  src = ./.;

  postInstall = ''
    wrapProgram $out/bin/arp-scan --set PATH ${lib.makeBinPath [ arp-scan ] }
  '';

  meta = {
    homepage = "https://ethulhu.co.uk/catbus";
    licence = stdenv.lib.licenses.mit;
  };
}
