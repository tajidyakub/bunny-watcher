# Bunny Watcher
> Watch a configured filesystem for changes and execute update to a Bunny CDN storage.

The application will watch a specific local filesystem and executes updates to a remote Bunny CDN storage when filesystem's state updated.


## Config

A configuration file must be placed in your home directory `~/.config/bunny-watcher/bunny.yaml`. Below is the example of the configuration content.

```yaml
api:
    host: storage.bunny-cdn.net
    key: a134366d-replace-with-your-key-45c0
remote:
    name: storage_name
    path: /path/inside/storage
local:
    path: ~/Projects/project-path
```

## Watcher

The application when initiated will run in the background and continously watch the state of the configured local filesystem, note that the application doesn't aware of the remote storage's state, it will overwrite any updates made through another method, e.g through the web interface.