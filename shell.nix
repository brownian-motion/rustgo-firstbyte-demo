with import <nixpkgs> {}; {
  devEnv = stdenv.mkDerivation {
    name = "dev";
    buildInputs = [ 
      stdenv 
      go
      cargo
    ];
  };
}
