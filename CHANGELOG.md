# Changelog

## [0.1.0](https://github.com/jarkad/pocket-expo/compare/v0.1.0...v0.1.0) (2026-03-21)


### Features

* add nix ([2018f79](https://github.com/jarkad/pocket-expo/commit/2018f79433d2bb0bb2bb6a4807e619edee80ffaf))
* add templ and sqlc as go tools ([9834516](https://github.com/jarkad/pocket-expo/commit/983451646f3bff6c12f59f3d67065daae8bcd74c))
* **db:** use dbmate for migrations ([838c2bf](https://github.com/jarkad/pocket-expo/commit/838c2bf5c369f65596b876e69d5172a66ce8219b))
* **deps:** enable dependabot golang ([9189247](https://github.com/jarkad/pocket-expo/commit/91892473d5383c7f17bedbba722d9ab34dc42d24))
* **docs:** add README.md ([7bf824f](https://github.com/jarkad/pocket-expo/commit/7bf824f2ef7e11eccafb87db3b846a8b1a730d25))
* **gh:** add repo description ([5a0b800](https://github.com/jarkad/pocket-expo/commit/5a0b800a3ecfd093b76fd9101dd4b87fe6545992))
* **gh:** tailor ([a97d559](https://github.com/jarkad/pocket-expo/commit/a97d559a57e848c72e6210e020ca3ca037c2bb28))
* **html:** add basic error messages ([ba36667](https://github.com/jarkad/pocket-expo/commit/ba366674d3d72123438ac2b492bf04b847fce513))
* init ([b2eec50](https://github.com/jarkad/pocket-expo/commit/b2eec50e1a25b2868c4f1377249c4e0eea893f8a))
* **nix:** add reuse pre-commit hook ([fe95d85](https://github.com/jarkad/pocket-expo/commit/fe95d85897422aa5491420e40bf2fe9a2a4a5c10))
* **release:** add goreleaser ([76a2358](https://github.com/jarkad/pocket-expo/commit/76a235812124cdfb6e2d1b37c492279d5266b7d0))
* **release:** add release-please ([df80ced](https://github.com/jarkad/pocket-expo/commit/df80ced6f34da22d0901d17da0fdd66553f2060c))
* **release:** upload binaries ([0717c91](https://github.com/jarkad/pocket-expo/commit/0717c91629d378c5e1ba1ccc1666537ca6c0f2bf))


### Bug Fixes

* **build:** disable sqlc for now ([ec98927](https://github.com/jarkad/pocket-expo/commit/ec989273d386a2998fdb3d9266e60eae74cc9ed8))
* **build:** go mod vendor ([300c57c](https://github.com/jarkad/pocket-expo/commit/300c57c52c3ab85b9e6e367f902b5e797dfff131))
* **legal:** add license header to .air.toml ([750b5ae](https://github.com/jarkad/pocket-expo/commit/750b5ae647246280d6c73650d1c3eb32e6f3ba5a))
* **legal:** add license header to .go files ([75dc8a1](https://github.com/jarkad/pocket-expo/commit/75dc8a1a4d45d77acc9dd35d698647d99ced8625))
* **legal:** add license header to .templ files ([78318aa](https://github.com/jarkad/pocket-expo/commit/78318aa72e39549585bb6eb9782b07ed627712c5))
* **legal:** add license header to Bootstrap assets ([6cac39c](https://github.com/jarkad/pocket-expo/commit/6cac39c06dc2a67b53a8f546effbfb8602b48ea0))
* **legal:** add license header to db/schema.sql ([6241e02](https://github.com/jarkad/pocket-expo/commit/6241e02a45924e3079d086c74a5670b365ecffdf))
* **legal:** add license header to go.mod and go.sum ([cd9b0fe](https://github.com/jarkad/pocket-expo/commit/cd9b0fe41b6663b3c9345246902369c7a3e924e5))
* **legal:** add license header to HTMX assets ([06ca850](https://github.com/jarkad/pocket-expo/commit/06ca850a3ea68cf7d1e77d499513e16c9d522c76))
* **legal:** add license header to tailor-generated files ([ce808f6](https://github.com/jarkad/pocket-expo/commit/ce808f654846d0e65a3a56fa38f6749f14ffec6c))
* **legal:** add license header to various .yaml configs ([bb4da6b](https://github.com/jarkad/pocket-expo/commit/bb4da6bc99a1a0dde0148a9113e13d9ddadbc006))
* **legal:** mark all go-snaps snapshots as EUPL-1.2 ([644d630](https://github.com/jarkad/pocket-expo/commit/644d630b427dcae3613e04a990cee250e4eead53))
* **nix:** nix flake lock ([2d43fd1](https://github.com/jarkad/pocket-expo/commit/2d43fd10f83cce0da0f786ea9e5f80f171e5cef0))
* **release:** make goreleaser honor vendor dir ([48ea9c2](https://github.com/jarkad/pocket-expo/commit/48ea9c29c8dd9fd613dc80c144debf4764421bb3))
* **release:** provide full git history to goreleaser ([a62b853](https://github.com/jarkad/pocket-expo/commit/a62b853474bc60ff28198f8832867f6f5af8a86b))
* **release:** provide github token to goreleaser ([8813acf](https://github.com/jarkad/pocket-expo/commit/8813acf3bb4f6c8b6922ffe33ec916c04533a20e))
* **release:** restrict goreleaser arch ([5cf2da2](https://github.com/jarkad/pocket-expo/commit/5cf2da2bf2912b750c396e18476a9d11fad20c73))
* run sqlc from go generate ([138608d](https://github.com/jarkad/pocket-expo/commit/138608d6587c93ea0f60b13543ee056ef43801b6))
* **svc/counter:** do not overwrite q.Get()'s error with q.Reset()'s ([e7f7142](https://github.com/jarkad/pocket-expo/commit/e7f7142105cfd074f8b99eb5a7e189ce49700f8f))


### Miscellaneous Chores

* **release:** release as v0.1.0 ([62a23e8](https://github.com/jarkad/pocket-expo/commit/62a23e8ba0792c1d32cd47fe58e470b8adcdc623))

## 0.1.0 (2026-03-21)


### Features

* add nix ([2018f79](https://github.com/jarkad/pocket-expo/commit/2018f79433d2bb0bb2bb6a4807e619edee80ffaf))
* add templ and sqlc as go tools ([9834516](https://github.com/jarkad/pocket-expo/commit/983451646f3bff6c12f59f3d67065daae8bcd74c))
* **db:** use dbmate for migrations ([838c2bf](https://github.com/jarkad/pocket-expo/commit/838c2bf5c369f65596b876e69d5172a66ce8219b))
* **deps:** enable dependabot golang ([9189247](https://github.com/jarkad/pocket-expo/commit/91892473d5383c7f17bedbba722d9ab34dc42d24))
* **docs:** add README.md ([7bf824f](https://github.com/jarkad/pocket-expo/commit/7bf824f2ef7e11eccafb87db3b846a8b1a730d25))
* **gh:** add repo description ([5a0b800](https://github.com/jarkad/pocket-expo/commit/5a0b800a3ecfd093b76fd9101dd4b87fe6545992))
* **gh:** tailor ([a97d559](https://github.com/jarkad/pocket-expo/commit/a97d559a57e848c72e6210e020ca3ca037c2bb28))
* **html:** add basic error messages ([ba36667](https://github.com/jarkad/pocket-expo/commit/ba366674d3d72123438ac2b492bf04b847fce513))
* init ([b2eec50](https://github.com/jarkad/pocket-expo/commit/b2eec50e1a25b2868c4f1377249c4e0eea893f8a))
* **nix:** add reuse pre-commit hook ([fe95d85](https://github.com/jarkad/pocket-expo/commit/fe95d85897422aa5491420e40bf2fe9a2a4a5c10))
* **release:** add goreleaser ([76a2358](https://github.com/jarkad/pocket-expo/commit/76a235812124cdfb6e2d1b37c492279d5266b7d0))
* **release:** add release-please ([df80ced](https://github.com/jarkad/pocket-expo/commit/df80ced6f34da22d0901d17da0fdd66553f2060c))
* **release:** upload binaries ([0717c91](https://github.com/jarkad/pocket-expo/commit/0717c91629d378c5e1ba1ccc1666537ca6c0f2bf))


### Bug Fixes

* **build:** disable sqlc for now ([ec98927](https://github.com/jarkad/pocket-expo/commit/ec989273d386a2998fdb3d9266e60eae74cc9ed8))
* **build:** go mod vendor ([300c57c](https://github.com/jarkad/pocket-expo/commit/300c57c52c3ab85b9e6e367f902b5e797dfff131))
* **legal:** add license header to .air.toml ([750b5ae](https://github.com/jarkad/pocket-expo/commit/750b5ae647246280d6c73650d1c3eb32e6f3ba5a))
* **legal:** add license header to .go files ([75dc8a1](https://github.com/jarkad/pocket-expo/commit/75dc8a1a4d45d77acc9dd35d698647d99ced8625))
* **legal:** add license header to .templ files ([78318aa](https://github.com/jarkad/pocket-expo/commit/78318aa72e39549585bb6eb9782b07ed627712c5))
* **legal:** add license header to Bootstrap assets ([6cac39c](https://github.com/jarkad/pocket-expo/commit/6cac39c06dc2a67b53a8f546effbfb8602b48ea0))
* **legal:** add license header to db/schema.sql ([6241e02](https://github.com/jarkad/pocket-expo/commit/6241e02a45924e3079d086c74a5670b365ecffdf))
* **legal:** add license header to go.mod and go.sum ([cd9b0fe](https://github.com/jarkad/pocket-expo/commit/cd9b0fe41b6663b3c9345246902369c7a3e924e5))
* **legal:** add license header to HTMX assets ([06ca850](https://github.com/jarkad/pocket-expo/commit/06ca850a3ea68cf7d1e77d499513e16c9d522c76))
* **legal:** add license header to tailor-generated files ([ce808f6](https://github.com/jarkad/pocket-expo/commit/ce808f654846d0e65a3a56fa38f6749f14ffec6c))
* **legal:** add license header to various .yaml configs ([bb4da6b](https://github.com/jarkad/pocket-expo/commit/bb4da6bc99a1a0dde0148a9113e13d9ddadbc006))
* **legal:** mark all go-snaps snapshots as EUPL-1.2 ([644d630](https://github.com/jarkad/pocket-expo/commit/644d630b427dcae3613e04a990cee250e4eead53))
* **nix:** nix flake lock ([2d43fd1](https://github.com/jarkad/pocket-expo/commit/2d43fd10f83cce0da0f786ea9e5f80f171e5cef0))
* **release:** make goreleaser honor vendor dir ([48ea9c2](https://github.com/jarkad/pocket-expo/commit/48ea9c29c8dd9fd613dc80c144debf4764421bb3))
* **release:** provide full git history to goreleaser ([a62b853](https://github.com/jarkad/pocket-expo/commit/a62b853474bc60ff28198f8832867f6f5af8a86b))
* **release:** provide github token to goreleaser ([8813acf](https://github.com/jarkad/pocket-expo/commit/8813acf3bb4f6c8b6922ffe33ec916c04533a20e))
* **release:** restrict goreleaser arch ([5cf2da2](https://github.com/jarkad/pocket-expo/commit/5cf2da2bf2912b750c396e18476a9d11fad20c73))
* run sqlc from go generate ([138608d](https://github.com/jarkad/pocket-expo/commit/138608d6587c93ea0f60b13543ee056ef43801b6))
* **svc/counter:** do not overwrite q.Get()'s error with q.Reset()'s ([e7f7142](https://github.com/jarkad/pocket-expo/commit/e7f7142105cfd074f8b99eb5a7e189ce49700f8f))


### Miscellaneous Chores

* **release:** release as v0.1.0 ([62a23e8](https://github.com/jarkad/pocket-expo/commit/62a23e8ba0792c1d32cd47fe58e470b8adcdc623))
