jfc
---

Jesus Fucking Christ

A simple thing to just fucking retry a command

```shell
# Try three times, wait three seconds
> jfc -d 3s -t 3 -- ls /aasdasdasd
ls: cannot access '/aasdasdasd': No such file or directory
ls: cannot access '/aasdasdasd': No such file or directory
ls: cannot access '/aasdasdasd': No such file or directory
> jfc -d 3s -t 3 -- ls /
[ .. ]
# etc
```

LICENSE
-------

Slightly modified BSD 3-clause
