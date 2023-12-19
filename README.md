# bitbucket-cli

A [Bitbucket Enterprise](https://bitbucket.org/product/enterprise) CLI.

<!-- ## Docker container

A docker container for this project can be obtained [here](https://github.com/swisscom/bitbucket-cli/pkgs/container/bitbucket-cli).
 -->

## Options
  <!-- ?? How can USERNAME and URL both use -u?? -->

```plaintext
  -D, --debug                               # 
  -u, --username {USERNAME}                 # The username to authenticate as
  -p, --password {PASSWORD}                 # The password to authenticate with
  -t, --access-token                        # The access token to authenticate with
  -u, --url {URL}                           # 
  -c, --config {/path/to/config-file}       #
  -h, --help                                # Display this help and exit
```

Examples

```bash
bitbucket-cli [--username USERNAME] [--password PASSWORD] --url URL command args
```

## Subcommands

- `project`
  - `-k {KEY}` (required)
  - `list`
  - `clone`
    - `-o,--output-path` (default: `./`)
    - `-b,--branch` (default: `prod`)
- `repo`
  - `branch`
    - `list`
    - `compare`
  - `pr`
    - `approve`
    - `create`
    - `list`
    - `merge`
  - `security`
    - `scan`
    - `result`
- `pr`
  - `list`

### project

#### list

Lists all the repositories in a project

```bash
$ export BITBUCKET_USERNAME="my-bitbucket-username"
$ read -s BITBUCKET_PASSWORD # Type your password and then press ENTER
$ export BITBUCKET_PASSWORD
$ bitbucket-cli --url https://your-bitbucket-hostname/rest project list -k PRJKEY

project-1       https://your-bitbucket-hostname/scm/prjkey/project-1.git
project-2       https://your-bitbucket-hostname/scm/prjkey/project-2.git
project-3       https://your-bitbucket-hostname/scm/prjkey/project-3.git
```

#### clone

Clones all the repositories in a project:

```bash
$ export BITBUCKET_USERNAME="my-bitbucket-username"
$ read -r -s BITBUCKET_PASSWORD # Type your password and then press ENTER
$ export BITBUCKET_PASSWORD
$ bitbucket-cli --url https://your-bitbucket-hostname/rest project clone -k PRJKEY -o /tmp/test/
    
    head: 987a5d8c25d8adb5ba013cf1cb88cd56a189241e5048b9702f319fb6e641cf81 refs/heads/master
    head: df2b794192904e6a9265975f33510eebe680177013e86fd7002850f45389ad34 refs/heads/master
    head: 2cf20bee2c59c3b8cae6ec0820a1353ff0ca2adeecdb84ba773845cff91ab121 refs/heads/master

$ ls -la /tmp/test 
total 0
drwxr-xr-x 13 dvitali dvitali  260 Jul 21 18:09 .
drwxrwxrwt 29 root    root    1400 Jul 21 18:11 ..
drwxr-xr-x  3 dvitali dvitali  120 Jul 21 18:09 project-1
drwxr-xr-x  4 dvitali dvitali  140 Jul 21 18:09 project-2
drwxr-xr-x  3 dvitali dvitali  100 Jul 21 18:09 project-3
```

### repo

This main subcommand requires two arguments:

- `-k KEY`
- `-n NAME`

These are the identifiers for your repository.  Not including one of the two in all of the subcommands will result in an error.

### pr

This subcommand deals with PRs, please check its subcommands.

#### create

This command, subcommand of (`repo pr`) allows you to create a Pull Request.

Use it as follows:

```plaintext
bitbucket-cli repo -k "KEY" \
  -n "bitbucket-playground" \
  pr create \
  -t "Some Title" \
  -d "Some Description :thumbsup:" \
  -F "refs/heads/feature/2" -T "refs/heads/master"

Usage: bitbucket-cli repo pr create --title TITLE [--description DESCRIPTION] --from-ref FROM-REF --to-ref TO-REF [--from-key FROM-KEY] [--from-slug FROM-SLUG]

Options:
  --title TITLE, -t TITLE
                         Title of this PR
  --description DESCRIPTION, -d DESCRIPTION
                         Description of the PR
  --from-ref FROM-REF, -F FROM-REF
                         Reference of the incoming PR, e.g: refs/heads/feature-ABC-123
  --to-ref TO-REF, -T TO-REF
                         Target reference, e.g: refs/heads/master
  --from-key FROM-KEY, -K FROM-KEY
                         Project Key of the "from" repository
  --from-slug FROM-SLUG, -S FROM-SLUG
                         Repository slug of the "from" repository
  --help, -h             display this help and exit
```

#### List

Lists all the PRs for the chosen repository

```bash
$ bitbucket-cli repo -k KEY -n bitbucket-playground pr list
Some Title (ID: 2)
feature 1 (ID: 1)
```

```bash
$ bitbucket-cli repo -k KEY -n bitbucket-playground pr list -s DECLINED
feature 1 (ID: 1)
```

```plaintext
Usage: bitbucket-cli repo pr list [--state STATE]

Options:
  --state STATE, -s STATE
                         PR State, any of: ALL, OPEN, DECLINED, MERGED
  --help, -h             display this help and exit
```

### Security

#### Scan

```plain
Usage: bitbucket-cli repo security scan

Options:
  --help, -h             display this help and exit
```

##### Example

```bash
bitbucket-cli repo -k ABC -n some-repo security scan
```
