{ pkgs ? import <nixpkgs> { }, src ? ./., subdir ? "" }:
let theSource = src; in
pkgs.buildGoModule {
  pname = "kp";
  version = "0.13.0";

  src = "${theSource}/${subdir}";

  vendorHash = "sha256-2g6BWwmLkpK1c/jQJoksvK0fOh3wEzu0+P+Molme4W4=";
  proxyVendor = true;
}

