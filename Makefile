.DEFAULT_GOAL: run

lib/libfirstbyte.a: lib/lib.h lib/firstbyte/Cargo.toml | lib/firstbyte/src 
	@cd lib/firstbyte && cargo build --release
	@cp lib/firstbyte/target/release/libfirstbyte.a lib/

.PHONY: run
run: lib/libfirstbyte.a
	go run main_static.go

.PHONY: clean
clean:
	rm -rf lib/libfirstbyte.a lib/firstbyte/target

.PHONY: test
test:
	@cd lib/firstbyte && cargo test
	go test ./...
