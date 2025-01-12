# SPDX-FileCopyrightText: 2023 Philipp Jungkamp <p.jungkamp@gmx.net>
# SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
# SPDX-License-Identifier: Apache-2.0
{
  lib,
  stdenv,
  buildGoModule,
  installShellFiles,
  versionCheckHook,
  protobuf,
  protoc-gen-go,
  protoc-gen-go-grpc,
  nix-update-script,
}:
let
  version = "0.12.1";
  src = ./..;
in
buildGoModule {
  pname = "cunicu";
  inherit version src;

  vendorHash = "sha256-6if+n5Q24c05Mqp0Ci4/S3aA+PLRCvuhcjQQaIWZxwk=";

  nativeBuildInputs = [
    installShellFiles

    protobuf
    protoc-gen-go
    protoc-gen-go-grpc
  ];

  nativeInstallCheckInputs = [ versionCheckHook ];

  CGO_ENABLED = 0;

  # These packages contain networking dependent tests which fail in the sandbox
  excludedPackages = [
    "pkg/config"
    "pkg/selfupdate"
    "pkg/tty"
    "scripts"
  ];

  ldflags = [
    "-X cunicu.li/cunicu/pkg/buildinfo.Version=${version}"
    "-X cunicu.li/cunicu/pkg/buildinfo.BuiltBy=Nix"
  ];

  doInstallCheck = true;
  versionCheckProgramArg = "version";

  passthru.updateScript = nix-update-script { };

  preBuild = ''
    go generate ./...
  '';

  postInstall = lib.optionalString (stdenv.buildPlatform.canExecute stdenv.hostPlatform) ''
    $out/bin/cunicu docs --with-frontmatter
    installManPage ./docs/usage/man/*.1
    installShellCompletion --cmd cunicu \
      --bash <($out/bin/cunicu completion bash) \
      --zsh <($out/bin/cunicu completion zsh) \
      --fish <($out/bin/cunicu completion fish)
  '';

  meta = {
    description = "Zeroconf peer-to-peer mesh VPN using Wireguard and Interactive Connectivity Establishment (ICE)";
    homepage = "https://cunicu.li";
    license = lib.licenses.asl20;
    platforms = lib.platforms.linux;
    maintainers = [ lib.maintainers.stv0g ];
    mainProgram = "cunicu";
  };
}
