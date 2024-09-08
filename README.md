# Granary

A simple secret management tool. It encrypts all of your secrets inside a file so that the values can't be retrieved directly.

By default, it will use a file named `secrets` in current working directory. You can change this value by using the `-p` flag.

p.s. This is just a hobby project, not a professional tool. Use at your own risk.

## Usage

1. Create a new secret file:

    ```sh
    gran new
    ```

2. Set a secret value:

    ```sh
    gran set my-password my-secret-password
    ```

3. List all available secrets:

    ```sh
    gran list
    ```

4. Get a secret value:

    ```sh
    gran get my-password
    ```

5. Remove a secret value:

    ```sh
    gran remove my-password
    ```

6. To avoid entering the passphrase each time `gran` command is invoked, you can create a file in the current working directory which contains the passphrase itself. The name of the file should be the same as the secret file's but has the extension of `.gpass`:

    ```sh
    # for example, if you have a secret file named "secrets"
    # and the passphrase is "my-secrets-password", then you
    # should create a file named "secrets.gpass" with the
    # content of "my-secrets-password".

    # First method, create the file manually:
    gran new
    echo my-secrets-password > secrets.gpass

    # ... or the second method, do it in one command:
    gran new --create-passfile

    # now this won't prompting for passphrase again
    gran set foo bar
    ```

## Installation

0. For convenience, make sure that you have both `Go` compiler and `make` build tool installed. But if you don't want to use `make`, you can check the content of `Makefile` and run the build commands directly on your shell (still need a `Go` compiler, though).

1. If you just want to build (without installing), you can run:

    ```sh
    make
    ```

    It will create a binary named `gran` in the project directory root.

2. But if you want to install it in your system, run:

    ```sh
    make install
    ```

## How it works?

All of the secrets are encrypted inside a Granary secret file, while the key itself is derived from the entered passkey using PBKDF2.

The content of Granary secret file is a simple ASCII string with the following format:

"`header:keyString:secrets`"

1. `header` is a string with value of "`gran-secret-file`". The purpose of this segment is to check whether the file is a real Granary secret file or not.

2. `keyString` is a 2-part strings that are concatenated together into a single string, separated by a dollar sign character (`$`). The first part is a **hex**-encoded salt bytes. While the second part is the hash bytes (SHA-256) of the key (derived from the passphrase) that is used to encrypt the secrets (also **hex**-encoded).

3. `secrets` is the list of all secrets stored in JSON format. Then it is encrypted using AES with derived key and encoded in **Base64** format.

When a user try to access the file using `gran` binary, it will prompt a passphrase first.

The program then will try to check whether the given passphrase is correct or not based on the `keyString`.

If the passphrase is correct, it will proceed to decrypt the content of the `secrets`.

## License

Granary is available under the MIT license. See the `LICENSE.md` file for more info.
