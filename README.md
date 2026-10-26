# Wisp publication smoke test

A small C library used to verify compilation, GitHub provenance, build recording, artifact storage and Go module distribution through the Wisp test deployment. The exported `add` function returns the sum of its two arguments.

The workflow records releases with Wisp and serves the generated Go module. It does not publish to npm, crates.io or PyPI.
