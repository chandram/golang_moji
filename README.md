# GoLang Using Service Weaver

<p align="center">Write, Deploy & Manage Distributed Apps</p>

<hr/>

<p>Learn GoLang using Service Weaver, based on <a href="https://github.com/ServiceWeaver/workshops">ServiceWeaver Workshop</a></p>

1. Hello, World!
2. A Search Component & add Logging facility
3. Unit Testing your Search Component
4. Implement HTTP Server
5. Run Multiple Process


```
$> mkdir emojis/
$> cd emojis/
$> go mod init emojis
$> go mod tidy
$> weaver generate .
$> go run .
```
# Hello, World!

# A Search Component & Logging
```
$> weaver generate .
$> go run .
```
# Unit Testing your Search Component
```
$> go test
```
# Implement  HTTP Server

How to serve HTTP traffic via Listeners
Create a ([TOML]()) config file `config.toml` with the following contents to
configure the address of the [listener](https://serviceweaver.dev/docs.html#step-by-step-tutorial-listeners).

```toml
[single]
listeners.emojis = {address = "localhost:9001"}
```
```
$> weaver generate .
$> SERVICEWEAVER_CONFIG=config.toml go run .
```

Curl or visit url [localhost:9001/search?q=pig]()

Run
```
 weaver single status
 
 $ ~/go/bin/weaver single status
╭──────────────────────────────────────────────────╮
│ DEPLOYMENTS                                      │
├─────┬──────────────────────────────────────┬─────┤
│ APP │ DEPLOYMENT                           │ AGE │
├─────┼──────────────────────────────────────┼─────┤
│ .   │ 0bb023b0-afe1-41fa-8ad8-59f890002576 │ 46s │
╰─────┴──────────────────────────────────────┴─────╯
╭───────────────────────────────────────────────────╮
│ COMPONENTS                                        │
├─────┬────────────┬─────────────────┬──────────────┤
│ APP │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS │
├─────┼────────────┼─────────────────┼──────────────┤
│ .   │ 0bb023b0   │ emojis.Searcher │ 27193        │
│ .   │ 0bb023b0   │ weaver.Main     │ 27193        │
│ .   │ 0bb023b0   │ main            │ 27193        │
╰─────┴────────────┴─────────────────┴──────────────╯
╭──────────────────────────────────────────────╮
│ LISTENERS                                    │
├─────┬────────────┬──────────┬────────────────┤
│ APP │ DEPLOYMENT │ LISTENER │ ADDRESS        │
├─────┼────────────┼──────────┼────────────────┤
│ .   │ 0bb023b0   │ emojis   │ 127.0.0.1:9000 │
╰─────┴────────────┴──────────┴────────────────╯

```

# Run Multiple Process

Service Weaver application in a single process with `go run`, 
inorder to run multiple processes add to TOML file

```
....
[multi]
listeners.hello = {address = "localhost:12345"}
```

You can see the status with

```

$ ~/go/bin/weaver multi status
╭─────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                         │
├────────┬──────────────────────────────────────┬─────┤
│ APP    │ DEPLOYMENT                           │ AGE │
├────────┼──────────────────────────────────────┼─────┤
│ emojis │ aec589ad-d020-4f03-9846-f6a0e4e27890 │ 29s │
╰────────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────╮
│ COMPONENTS                                           │
├────────┬────────────┬─────────────────┬──────────────┤
│ APP    │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS │
├────────┼────────────┼─────────────────┼──────────────┤
│ emojis │ aec589ad   │ emojis.Searcher │ 27736, 27737 │
│ emojis │ aec589ad   │ weaver.Main     │ 27712, 27727 │
╰────────┴────────────┴─────────────────┴──────────────╯
╭─────────────────────────────────────────────────╮
│ LISTENERS                                       │
├────────┬────────────┬──────────┬────────────────┤
│ APP    │ DEPLOYMENT │ LISTENER │ ADDRESS        │
├────────┼────────────┼──────────┼────────────────┤
│ emojis │ aec589ad   │ emojis   │ 127.0.0.1:9000 │
╰────────┴────────────┴──────────┴────────────────╯

```

# Cache

```
$ ~/go/bin/weaver multi metrics cache
╭────────────────────────────────────────────────────────────────────────╮
│ // Number of Search cache hits                                         │
│ search_cache_hits: COUNTER                                             │
├───────────────────┬────────────────────┬───────────────────────┬───────┤
│ serviceweaver_app │ serviceweaver_node │ serviceweaver_version │ Value │
├───────────────────┼────────────────────┼───────────────────────┼───────┤
│ emojis            │ 143eae9d           │ 195d0cb0              │ 0     │
│ emojis            │ 3f38a335           │ 195d0cb0              │ 0     │
│ emojis            │ 5d11815a           │ 195d0cb0              │ 0     │
│ emojis            │ 7aab7eb2           │ 195d0cb0              │ 1     │
│ emojis            │ 99d0291f           │ 195d0cb0              │ 0     │
│ emojis            │ fab9091a           │ 195d0cb0              │ 0     │
╰───────────────────┴────────────────────┴───────────────────────┴───────╯
╭────────────────────────────────────────────────────────────────────────╮
│ // Number of Search cache misses                                       │
│ search_cache_misses: COUNTER                                           │
├───────────────────┬────────────────────┬───────────────────────┬───────┤
│ serviceweaver_app │ serviceweaver_node │ serviceweaver_version │ Value │
├───────────────────┼────────────────────┼───────────────────────┼───────┤
│ emojis            │ 143eae9d           │ 195d0cb0              │ 0     │
│ emojis            │ 3f38a335           │ 195d0cb0              │ 1     │
│ emojis            │ 5d11815a           │ 195d0cb0              │ 0     │
│ emojis            │ 7aab7eb2           │ 195d0cb0              │ 0     │
│ emojis            │ 99d0291f           │ 195d0cb0              │ 0     │
│ emojis            │ fab9091a           │ 195d0cb0              │ 0     │
╰───────────────────┴────────────────────┴───────────────────────┴───────╯

```


# Useful links

* [Service Weaver Installation](https://serviceweaver.dev/docs.html#installation): https://serviceweaver.dev/docs.html#installation
* [Single & Multiple Components](https://serviceweaver.dev/docs.html#installation): https://serviceweaver.dev/docs.html#step-by-step-tutorial-components
* [How to weaver_Run](https://serviceweaver.dev/docs.html#installation): https://pkg.go.dev/github.com/ServiceWeaver/weaver#Run