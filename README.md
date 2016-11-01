# rsc

[![Circle CI](https://circleci.com/gh/nextrevision/rsc.svg?style=svg)](https://circleci.com/gh/nextrevision/rsc)

RSC is a [Runscope](https://runscope.com) command line client used to manage resources in a Runscope account.

## Usage

```
Runscope Client (rsc) provides a CLI for interacting
                with the Runscope service.

Usage:
  rsc [command]

Available Commands:
  delete      delete a resource from Runscope
  import      import tests from a path of configs and templates
  list        show a listing of resources
  show        show a showing of resources
  version     print the version number of rsc

Flags:
      --debug          enable debug output
      --token string   runscope authentication token (default "9b3bb9bc-01d5-4cf8-80d6-49183c6235b0")
      --verbose        enable verbose output

Use "rsc [command] --help" for more information about a command.
```

## Authenticating to Runscope

In order to use rsc, you need to obtain an access token for your account (you can find details on how to obtain an access token [here](https://www.runscope.com/docs/api/authentication#apps)). Once you have your access token, you can set it either by the `RSC_TOKEN` environment variable or by passing it as a flag with `--token`.

### Environment Variable

```
$ export RSC_TOKEN=xxxxxxxxxx-xxxx-xxxx-xxxxxxxxxxx
$ rsc ...
```

### CLI Flag

```
$ rsc --token xxxxxxxxxx-xxxx-xxxx-xxxxxxxxxxx ...
```

## Examples

All examples assume that the Runscope API token has been exported like the following:

```
$ export RSC_TOKEN=xxxxxxxxxx-xxxx-xxxx-xxxxxxxxxxx
```

### Listing Buckets

Returns a listing of all buckets for an account:

```
$ rsc list buckets
  NAME              TEAM        DEFAULT
  MyFirstBucket     Team1       true
  MySecondBucket    Team1       false
  ProdBucket        Team2       true
```

### Listing Tests

Returns a listing of all tests for a given bucket:

```
$ rsc list tests -b MyFirstBucket
  NAME                         CREATED BY      LAST RUN                    LAST STATUS   DESCRIPTION
  TestFoo                      Bob             2016-08-22T23:11:02-04:00   success       A decent test for Foo
  TestBar                      Jerry           2016-06-13T23:02:12-09:00   error         Tests that Bar is compliant with oth...
  TestBanana                   Bob             2016-08-22T23:11:02-04:00   success
```

### Importing Buckets and Tests

Imports buckets and tests based on JSON config files and templates (see examples):

```
$ rsc import --debug --path examples
DEBU[0000] Creating bucket: MyBucket
INFO[0000] Created bucket: MyBucket
DEBU[0001] Found bucket by name: MyBucket
DEBU[0001] Creating test: MyFirstTest
INFO[0002] Created test: MyFirstTest
DEBU[0002] Found bucket by name: MyBucket
DEBU[0002] Creating test: MySecondTest
INFO[0003] Created test: MySecondTest
```

#### Dry Runs

You can validate configs by running import with the flag `--dry-run`. This will perform every action except making the change to the bucket or test.

```
$ rsc import --debug --dry-run --path examples
INFO[0000] Loading configs from examples...
INFO[0000] Found 1 configs...
INFO[0000] Loading templates from examples...
INFO[0000] Found 1 templates...
INFO[0000] Importing 1 buckets...
INFO[0001] Would have created bucket: MyBucket
INFO[0001] Importing 2 tests...
INFO[0002] Would have created test: MyFirstTest
INFO[0002] Would have created test: MySecondTest
```

#### Template Functions

You can create templates for your tests (see `examples` directory). Inside of these templates, you can access information about other tests via template functions.

* `<% TriggerURL bucketname testname %>`: returns the trigger URL for a given test
* `<% TriggerURL bucketname testname %>`: returns the batch URL for a given test

## Developing

```
go get github.com/Masterminds/glide
make test
```
