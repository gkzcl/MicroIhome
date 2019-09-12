#!/usr/bin/env bash
redis-server ./conf/redis.conf

fdfs_trackerd /home/roger/go/src/MicroIhome/IhomeWeb/conf/tracker.conf restart

fdfs_storaged /home/roger/go/src/MicroIhome/IhomeWeb/conf/storage.conf restart

sudo nginx

consul agent -dev