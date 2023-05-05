# go-mono
Monorepo sample Project for golang using bazel

- first install bazel [install link](https://bazel.build/install/ubuntu )
- install gazel 
- run the gazel with bazel, This will autogenerate the BUILD.bazel files for all of the packages.

note : You'll need a `BIULD` file in each directory with Go code <br>
note: in `BUILD` file change this line to your repo `# gazelle:prefix github.com/amirhosseinmoayedi/go`<br>
install Gazelle: 
```bash
go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
```
re-run this command to upgrade Gazelle whenever you upgrade rules_go in your repository.

or install just in the project: 
```bash
go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle
```

run to create `.bazel` files:
```bash
bazel run //:gazelle
```

to cleare the cache of bazel
```bash
bazel clean --expunge
```

after each update in `go.mod`, will create `go_third_party.bzl` which is for dependency management of project:
```bash
 gazelle update-repos --from_file=go.mod -to_macro=go_third_party.bzl%go_deps
```

to build service:
```bash
bazel build //services/name_of_service
```
to run service:
```bash
bazel run //services/name_of_service

bazel run //services/name_of_service
```

create docker image
```bash
bazel run //services/name_of_service:image
```