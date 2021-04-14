# go-fish

go-fish is a small tool for working with Git hooks in a project 🎣 . It is a port of [Husky](https://github.com/typicode/husky) to Go.

go-fish is intended to be used with Go projects. However, it can theoretically be used for any project since it doesn't do anything that is specific to Go projects.

## Installation

The recommended way to install go-fish is with [shed](https://github.com/getshiphub/shed).

```
shed install github.com/cszatmary/go-fish
```

## Usage

Install hooks:

```
shed run go-fish install
```

This will create a `.hooks` folder in the root of your repository.

Adding a hook:

```
shed run go-fish create pre-commit
```

This will create a `.hooks/pre-commit` script. Edit this as desired to configure the hook.

Uninstall hooks:

```
shed run go-fish uninstall
```

## How it works

go-fish uses the [core.hooksPath](https://git-scm.com/docs/git-config#Documentation/git-config.txt-corehooksPath) config option of Git to set the hooks directory to `.hooks`.
This allows Git hooks to be managed within the repository.

## License

dot is available under the [MIT License](LICENSE).

## Contributing

Contributions are welcome. Feel free to open an issue or submit a pull request.
