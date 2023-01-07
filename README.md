## Password

Demo application for the pwdgen Password generator.

Usage:

    -d int
          Number of digit [0-9] characters (default 4)
    -l int
          Number of lower case characters (default 4)
    -s int
          Number of special characters (default 4)
    -u int
          Number of upper case characters (default 4)

Example:

    $ go run main.go -s=2 -u=3 -l=3 -d=3
    83w#Z(0WIsc