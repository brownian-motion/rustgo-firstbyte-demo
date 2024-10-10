# Rust <-> Go FFI Demo
Demos building a static library in Rust and importing it in Go using the CGo interface.

Followed the example demonstrated by https://github.com/mediremi/rust-plus-golang .

# Compiling
## With Nix
If you use the Nix package manager, you can compile and run the demo by checking out this repo and running `nix-shell --run 'make run'`.

## Native tools
With `cargo`, `golang`, and GNU `make` installed, run `make run`.