# Pixr

> Create a simplistic file sharing site that allows users to quickly upload content without the need for an account.

---

~ **More will be added later** ~

Clone the repository:

```bash
git clone https://github.com/Neoplatonist/pixr.git
```

Either build the binary with `make build` or download the binary from the [releases](https://github.com/Neoplatonist/pixr/releases). Currently we only have a release for linux, but will add more binaries later. Add the binary to the cloned directory.

To use first build the frontend with

```bash
cd frontend
npm run build
```

Afterward either fill out the configs or use flags to attach the mysql database.

To start the server:

```bash
./pixr
```

For help:

```bash
./pixr -h

or

./pixr --help
```

---

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](http://opensource.org/licenses/mit-license.php)**
- Copyright 2019 © <a href="http://github.com/neoplatonist" target="_blank">Neoplatonist</a>.
