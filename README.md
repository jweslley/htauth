# htauth

**htauth** generates encrypted passwords for basic and digest authentication.

It's similar to Apache support programs *htdigest* and *htpasswd*. However, **htauth** is much more simpler and focus only in generate encrypted passwords to standard output.

## Installing

[Download][] and put the binary somewhere in your path.

### Building from source

    git clone http://github.com/jweslley/htauth
    make

## Options

**-h string**

	Hashing encryption for passwords. Available hashing algorithms: bcrypt, sha1, plain. (default "bcrypt")

**-r string**

  The realm name to which the user name belongs. Used only to generate passwords for digest authentication.


## Usage

In order to generate an encrypted password, just execute the command below:

    htauth <username>

You will be prompted to inform your password and the encrypted password will be printed.

To save the encrypted password, redirect the output to a file:

    htauth stark > /path/to/htpasswd

Encrypted passwords for digest authentication can be generated by using the `realm` option:

    htauth -r avengers.com stark > /path/to/htdigest


## Bugs and Feedback

If you discover any bugs or have some idea, feel free to create an issue on GitHub:

    http://github.com/jweslley/htauth/issues


## License

MIT license. Copyright (c) 2018 Jonhnny Weslley <http://jonhnnyweslley.net>

See the LICENSE file provided with the source distribution for full details.


[download]: https://github.com/jweslley/htauth/releases
