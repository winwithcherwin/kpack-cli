{ pkgs ? import <nixpkgs> { }, src ? ./., subdir ? "" }:
let theSource = src; in
pkgs.buildGoModule {
  pname = "kp";
  version = "0.13.0";

  src = "${theSource}/${subdir}";

  vendorHash = "sha256-105gnhwscsgcvmvshrh6bfrvw64d8m5vzh8hg8ln8kk7mzwcjdgg";
  proxyVendor = true;
}

