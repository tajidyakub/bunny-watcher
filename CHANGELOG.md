# CHANGELOG

## v0.2.0

Prepares Bunny API Client as an internal package.

- [ ] Init Bunny API Http Client.
- [ ] Authenticate api key.
- [ ] List existing storage.
- [ ] Validates the configured api host and key is correct.
- [ ] Validates the existence of configured remote storage and path.



## v0.1.0

Prepares the main logic of main function, ensure the config file exists and initialize a default config file if it is not yet exists.

- [x] Defines Config struct to read and write the `yaml` file.
- [x] `initConfig(path string) err errors` function in the main package to initialize a default config file in `~/.config/bunny-watcher/bunny.yaml`.
- [x] Check config file existence during start up and exit if not yet exists after creating the default one.
- [ ] Validates local path defined exists.
