# ldfpcfix

This little utility patches `/usr/bin/ld` to silence the annoying warning that appears when compiling with `fpc`:

```
Free Pascal Compiler version 3.0.2 [2017/07/25] for x86_64
Copyright (c) 1993-2017 by Florian Klaempfl and others
Target OS: Linux for x86-64
Compiling main.pas
Linking main
/usr/bin/ld: warning: link.res contains output sections; did you forget -T?
5 lines compiled, 0.1 sec
```

This warning has existed for years and neither ld nor fpc seems to think it's annoying enough to fix.

After running `ldfpcfix` as root, if `ld` exists in `/usr/bin/ld`, the output is:

```
Free Pascal Compiler version 3.0.2 [2017/07/25] for x86_64
Copyright (c) 1993-2017 by Florian Klaempfl and others
Target OS: Linux for x86-64
Compiling main.pas
Linking main
5 lines compiled, 0.0 sec
```

Read the sources if you need to verify exactly what `ldfpcfix` does.

It's not pretty, but it works. Don't use it if you don't like it.

### General information

Licensed under GPL2.

### Download

Download the precompiled binary for 64-bit Linux [here](https://github.com/xyproto/ldfpcfix/releases/download/v1.0/ldfpcfix-1.0.tar.gz).
