{ pkgs ? import <nixpkgs> { }, src ? ./., subdir ? "" }:
let theSource = src; in
pkgs.buildGoModule {
  pname = "kp";
  version = "0.13.0";

  src = "${theSource}/${subdir}";

  vendorHash = "";
  proxyVendor = true;
}

