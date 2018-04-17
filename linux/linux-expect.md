---
layout: post
title: Linux Expect
date: 2017-03-21 15:10:00
tags:
- Linux
categories: Linux
---

```bash
expect -c "
set timeout 5

spawn scp someuser@somehost:/path/to/remote/somefile.tar.gz /opt

expect {
    \"(yes/no)?\"
    {
         send \"yes\n\"
         expect \"*asswrod:\" { send \"Hello@1234\n\" }
     }
     \"*assword:\"
     {
         send \"Hello@1234\r\"
     }
}

# expect \"password:\"
# send \"Hello@1234\r\"
expect eof
"
echo "done"
```
