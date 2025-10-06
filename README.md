# Clerk CLI

Tool to manage Clerk pools using an API key.

## Installation

Simply install it with `go install`:

```console
go install github.com/kbence/clerkctl@latest
```

## Setup

Put your API key to the environment variable `CLERK_SECRET_KEY`:

```console
export CLERK_SECRET_KEY=<secret_key>
```

Now you can use the full functionality of `clerkctl`!

## Subcommands

### `user list`

Lists the last 10 users created. This limitation is due to Clerk Go SDK's limitation in v2. Version 3 supports limits and offsets already, but the library is not set up to be possible to use v3 as a Go package yet...

### `user delete`

Can delete a set of users either by their IDs or their email addresses (can be mixed):

```console
cleckctl user delete user_1234567890 test@test.com
```
