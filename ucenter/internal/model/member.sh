#! /bin/bash
goctl model mysql datasource --url="root:root@tcp(127.0.0.1:3306)/mscoin" --table="member" -c --dir .