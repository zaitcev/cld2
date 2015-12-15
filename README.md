# cld2

The CLD2 is a re-implementation of CLD of Project Hail.
See
  https://github.com/jgarzik/hail/tree/master/cld
  https://hail.wiki.kernel.org/index.php/CLD

Hail's CLD used Berkeley DB in order to elect the master.
It worked, but wasn't fun enough.

Systems you should examine before using CLD2 (or CLD for that matter):
 - Zookeeper
 - etcd
 - Consul
 - maybe Redis
