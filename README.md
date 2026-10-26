# Wisp publication smoke test

A small C library used to verify compilation, GitHub provenance, build recording, artifact storage and Go module distribution through the Wisp test deployment. The exported `add` function returns the sum of its two arguments.

The workflow records releases with Wisp and serves the generated Go module. It does not publish to npm, crates.io or PyPI.

After the release is recorded, run the Go consumer on Linux x86_64 with Go 1.27:

```sh
cd consumer
GOPROXY=https://web-production-6bbc2.up.railway.app/api/go,https://proxy.golang.org \
GONOSUMDB='web-production-6bbc2.up.railway.app/*' go run -mod=mod .
```

The test uses the deployment's own module proxy. The consumer must print
`add(20, 22) = 42`. It does not require Wisp or its compiler toolchains locally.
