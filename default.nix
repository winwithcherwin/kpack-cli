{ pkgs ? import <nixpkgs> { }, src ? ./., subdir ? "" }:
let theSource = src; in
pkgs.buildGoModule {
  pname = "kp";
  version = "0.13.0";

  src = "${theSource}/${subdir}";

  vendorHash = "sha256-t8d1LoO14xpU6yeozuVW393Zj4PcEEHzYVD8XX9zkb8=";
  proxyVendor = true;
}

