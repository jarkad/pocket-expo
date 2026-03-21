# Changelog

## [1.1.0](https://github.com/jarkad/pocket-expo/compare/v1.0.0...v1.1.0) (2026-03-21)


### Features

* **dev:** auto-tidy go.mod in dev ([81cde13](https://github.com/jarkad/pocket-expo/commit/81cde13fc78efe4d6276f33af68ba7cc5cf53934))


### Bug Fixes

* **dev:** auto-run go generate in dev ([315d7b8](https://github.com/jarkad/pocket-expo/commit/315d7b8cfffaaff72c51af7f4d6b6056ae059abf))

## 1.0.0 (2026-03-21)


### Features

* add nix ([d951e29](https://github.com/jarkad/pocket-expo/commit/d951e29f03d001f3d28212813487e09b78bbbc66))
* add templ and sqlc as go tools ([dd3c49a](https://github.com/jarkad/pocket-expo/commit/dd3c49a0c602a986057087a8e38a0044a475a4cc))
* **db:** use dbmate for migrations ([df058f0](https://github.com/jarkad/pocket-expo/commit/df058f089b76d6e24494afbb90be82706f54843d))
* **deps:** enable dependabot golang ([9189247](https://github.com/jarkad/pocket-expo/commit/91892473d5383c7f17bedbba722d9ab34dc42d24))
* **docs:** add README.md ([433b5bb](https://github.com/jarkad/pocket-expo/commit/433b5bb69e2e0aa0c40685f5b5ce70c333082ff5))
* **gh:** add repo description ([5a0b800](https://github.com/jarkad/pocket-expo/commit/5a0b800a3ecfd093b76fd9101dd4b87fe6545992))
* **gh:** tailor ([a97d559](https://github.com/jarkad/pocket-expo/commit/a97d559a57e848c72e6210e020ca3ca037c2bb28))
* **html:** add basic error messages ([54002df](https://github.com/jarkad/pocket-expo/commit/54002df7b91c496c601e4da7c98321069b3b8ebd))
* init ([b2eec50](https://github.com/jarkad/pocket-expo/commit/b2eec50e1a25b2868c4f1377249c4e0eea893f8a))
* **nix:** add reuse pre-commit hook ([d185d8a](https://github.com/jarkad/pocket-expo/commit/d185d8a76cdd30fed2b48be9bc827163041aa3da))
* **release:** add goreleaser ([76a2358](https://github.com/jarkad/pocket-expo/commit/76a235812124cdfb6e2d1b37c492279d5266b7d0))
* **release:** add release-please ([df80ced](https://github.com/jarkad/pocket-expo/commit/df80ced6f34da22d0901d17da0fdd66553f2060c))
* **release:** upload binaries ([5d221a0](https://github.com/jarkad/pocket-expo/commit/5d221a0b82f2b9e0ba2f71811e2d1feae66170c6))


### Bug Fixes

* **build:** disable sqlc for now ([c33e972](https://github.com/jarkad/pocket-expo/commit/c33e9721febe6bac454bfe545dd419baa73dde21))
* **build:** go mod vendor ([4d29e6d](https://github.com/jarkad/pocket-expo/commit/4d29e6d550ee22dacdff46ea902033a567c17b0d))
* **legal:** add license header to .air.toml ([e64ed82](https://github.com/jarkad/pocket-expo/commit/e64ed828d3e0b83e366d120af462b915ac68f71c))
* **legal:** add license header to .go files ([c33863a](https://github.com/jarkad/pocket-expo/commit/c33863a5da7af6a452735668dce74616382312e7))
* **legal:** add license header to .templ files ([5e8d54e](https://github.com/jarkad/pocket-expo/commit/5e8d54e13c3f1944f8c52061b5160f70ac8e93a2))
* **legal:** add license header to Bootstrap assets ([762f058](https://github.com/jarkad/pocket-expo/commit/762f058a8c312e944039d1a47a014f6c32e32c25))
* **legal:** add license header to db/schema.sql ([29aa3d1](https://github.com/jarkad/pocket-expo/commit/29aa3d1610c70f409d07440bf621d9a571199e09))
* **legal:** add license header to go.mod and go.sum ([77fae54](https://github.com/jarkad/pocket-expo/commit/77fae54bfb4a22194702c6252b41093ca35f46a3))
* **legal:** add license header to HTMX assets ([f83914f](https://github.com/jarkad/pocket-expo/commit/f83914f37ac032e9ba21e7b6ba727be8f6396624))
* **legal:** add license header to tailor-generated files ([d392bb2](https://github.com/jarkad/pocket-expo/commit/d392bb25d3158545e9613b20ff47eb2f3860bcf7))
* **legal:** add license header to various .yaml configs ([f5edd84](https://github.com/jarkad/pocket-expo/commit/f5edd843146dd1b6ae8b30938acb95058f851f2a))
* **legal:** mark all go-snaps snapshots as EUPL-1.2 ([7e41562](https://github.com/jarkad/pocket-expo/commit/7e41562fe64c3ef68d5f57f39a5a3041abd84bca))
* **nix:** nix flake lock ([e2470c7](https://github.com/jarkad/pocket-expo/commit/e2470c7fda7944269820b169d989fbf369ac0a6e))
* **release:** make goreleaser honor vendor dir ([267eda6](https://github.com/jarkad/pocket-expo/commit/267eda6663f5415dbc46256d82f9e9d21cca056f))
* **release:** provide full git history to goreleaser ([d8b6492](https://github.com/jarkad/pocket-expo/commit/d8b6492c87f2139f375cbfc9c96e9ae4157f865f))
* **release:** provide github token to goreleaser ([ab9b055](https://github.com/jarkad/pocket-expo/commit/ab9b055728f413fe03fd72d749736b9a6aac5dca))
* **release:** restrict goreleaser arch ([5cf2da2](https://github.com/jarkad/pocket-expo/commit/5cf2da2bf2912b750c396e18476a9d11fad20c73))
* run sqlc from go generate ([e33507e](https://github.com/jarkad/pocket-expo/commit/e33507e1284d3a92ea8bce82b7b120e2f4c34b06))
* **svc/counter:** do not overwrite q.Get()'s error with q.Reset()'s ([0005a6b](https://github.com/jarkad/pocket-expo/commit/0005a6b628d0b7bc8d34c21eb0be0f8b843eb628))
