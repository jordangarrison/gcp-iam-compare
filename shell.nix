let
  unstable = import
    (fetchTarball "https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz")
    { };
in
{ nixpkgs ? import <nixpkgs> { } }:
with nixpkgs;
mkShell {
  name = "gcp-iam-compare";

  buildInputs = with pkgs; [
    unstable.go_1_18
  ];

  VAULT_ADDR = "https://vault.stag.flokubernetes.com";
}
