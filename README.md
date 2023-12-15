# Arma3HTS
Hey Arma3 administrators, if you need a small software to synchronize your mods on your server with an Arma3 perset this is made for you!

Arma3HTS is a CLI written with go, which allows you thanks to an Arma3 perset and !workshop folder to connect to your remote server and to synchronize its mods.

## Getting Started
These instructions are intended for end users so they can get a copy of the project working on their machine, for developers see [contributing](#Contributing).


### Prerequisites
- You need [Go](https://go.dev/doc/install) **1.21.x**

### Installing
#### Compile Arma3HTS yourself (recommended)
First clone the repository and move to the new folder:
```shell
git clone git@github.com:MisterOryon/Arma3HTS.git
cd Arma3HTS
```

Install the project dependencies:
```shell
go mod download
```

Compile the application:
```shell
go build -o build/Arma3HTS.exe
```

### FAQ
#### What commands are available?
You can get a detailed list of available commands using the -h argument, if you use this argument on a command you will get a help page.
```shell
Arma3HTS.exe -h
```
```shell
Arma3HTS.exe check -h
```

#### What connection modes are supported?
Currently only **sftp** is supported however you can always open an issue to suggest or offer feature enrichment.

#### My "nameModsXXX / xxx" mod mark as not found even it is present on my local machine?
This is normal by default when Steam downloads a mod it replaces the characters “/” with “-“,
To resolve this problem create a copy of your preset and replace the character “/” with “-“ in the name of your mod.

#### "unable to read known_hosts file"?
This error is due to the fact that you do not have any known_hosts files on your machine,
to resolve this you simply need to add an empty known_hosts file.

open PowerShell (windows):
```shell
cd $env:USERPROFILE
mkdir .ssh
cd .\.ssh\
ni "known_hosts"
```

open Bash (linux):
```shell
cd ~
mkdir .ssh
cd .ssh
touch "known_hosts"
```

## Built With

- [sftp](https://github.com/pkg/sftp) - A sftp client in Go
- [cobra](https://github.com/spf13/cobra) - A Framework for Modern CLI Apps in Go
- [uilive](https://github.com/gosuri/uilive) - A library for updating terminal output in realtime
- [crypto](https://golang.org/x/crypto)
- [net](https://golang.org/x/net)


## Contributing
- Read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct
- Read [CONTRIBUTING.md](CONTRIBUTING.md) for details on the process for submitting pull requests to us

## Versioning
For the versions available, see the [tags on this repository](https://github.com/MisterOryon/Arma3HTS/tags).

## Authors
- **[MisterOryon](https://github.com/MisterOryon)** - _Initial work and maintainer_

See also the list of [contributors](https://github.com/MisterOryon/Arma3HTS/contributors) who participated in this project.

## License
This project is licensed under the AGPL-3.0 License - see the [LICENSE.txt](LICENSE.txt) file for details.

## Acknowledgments
- [@PurpleBooth](https://github.com/PurpleBooth) for README template
- [Contributor Covenant](https://www.contributor-covenant.org/) for code of conduct
- [zaro67](https://discord.com/) for testing this tool on your server
