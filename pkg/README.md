# `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools.
